package main

import (
	"fmt"
	"github.com/kyberorg/wait4harbor/cmd/wait4harbor/api"
	"github.com/kyberorg/wait4harbor/cmd/wait4harbor/config"
	"log"
	"time"
)

const DoneFlag = "doneFlag"
const NotFound = "notFound"
const Failed = "failed"

var hapi *api.HarborApi
var resty *api.Resty
var sha string

func main() {
	paramsValid, errorMessage := config.GetAppConfig().ValidateParams()
	if !paramsValid {
		log.Fatalln(errorMessage)
	}

	harborParams := config.GetAppConfig().HarborParams
	imageParams := config.GetAppConfig().ImageParams

	hapi = api.GetHarborApi(harborParams, imageParams)
	resty = api.GetResty(hapi)
	sha = imageParams.Sha

	pingSuccess := resty.Ping()
	if pingSuccess {
		fmt.Println("Harbor API alive. Starting checks...")
	} else {
		log.Fatalf("Harbor API unreachable. Endpoint: %s\n", hapi.GetPingEndpoint())
	}

	timeout := time.Now().Add(config.GetAppConfig().Timeout)
	for {
		checkResult := doCheck()
		if checkResult == DoneFlag {
			fmt.Println("Great success")
			break
		} else if time.Now().After(timeout) {
			fmt.Println("Timeout reached")
			break
		}
		doPause()
	}
}

func doCheck() string {
	fmt.Printf("Looking for %s \n", hapi.PrintFullImagePath())
	found, errMessage := resty.LookupForSha(sha)
	if found {
		log.Println("Tag+SHA match found. Exiting")
		return DoneFlag
	}
	if errMessage != nil {
		log.Printf("Request failed. Error Message: %s \n", errMessage)
		return Failed
	}
	log.Println("digest mismatch")
	return NotFound
}

func doPause() {
	fmt.Printf("Pausing for %f seconds.\n", config.GetAppConfig().Interval.Seconds())
	time.Sleep(config.GetAppConfig().Interval)
}
