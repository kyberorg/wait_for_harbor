package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Resty struct {
	hapi   *HarborApi
	client *resty.Client
	req    *resty.Request
}

func GetResty(hapi *HarborApi) *Resty {
	r := &Resty{
		hapi:   hapi,
		client: resty.New(),
	}
	r.req = r.client.R()
	r.req.SetHeader("accept", "application/json")
	r.req.SetBasicAuth(hapi.harborInfo.Robot, hapi.harborInfo.Token)
	return r
}

func (r Resty) Ping() bool {
	resp, err := r.client.R().Get(r.hapi.GetPingEndpoint())
	if err != nil {
		fmt.Printf("Ping request failed: %s \n", err)
		return false
	}
	return resp.IsSuccess()
}

func (r Resty) LookupForSha(sha string) (bool, error) {
	resp, err := r.client.R().Get(r.hapi.GetLookupEndpoint())
	if err != nil {
		fmt.Printf("Ping request failed: %s \n", err)
		return false, err
	}

	switch resp.StatusCode() {
	case 200:
		var data SuccessResponse
		err := json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return false, errors.New("failed to parse JSON")
		}
		realDigest := data.Digest
		if realDigest != sha {
			fmt.Printf("Digest mismatch. \n Excepted: %s, \n Real: %s \n", sha, realDigest)
		}
		return realDigest == sha, nil
	case 403:
		fmt.Println("Got 403. Raw resp: " + string(resp.Body()))
		return false, errors.New("no such project found")
	case 404:
		fmt.Println("Got 404. Raw resp: " + string(resp.Body()))
		var errJson ErrorResponse
		err := json.Unmarshal(resp.Body(), &errJson)
		if err != nil {
			return false, err
		}
		return false, errors.New(errJson.Errors[0].Message)
	default:
		fmt.Printf("Got unknown error. Status %d. Raw resp: %s", resp.StatusCode(), string(resp.Body()))
		return false, errors.New("got Unknown Error")
	}
}
