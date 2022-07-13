package api

const (
	HarborApiPath           = "/api/v2.0"
	HarborPingEndpoint      = "/ping"
	HarborProjectEndpoint   = "/projects"
	HarborRepositorySubPath = "/repositories"
	HarborArtifactsSubPath  = "/artifacts"
)

type HarborParams struct {
	Hostname string
	Schema   string
	Robot    string
	Token    string
}

type ImageParams struct {
	Project string
	Repo    string
	Tag     string
	Sha     string
}

type HarborApi struct {
	apiUrl     string
	harborInfo *HarborParams
	imgInfo    *ImageParams
}

func GetHarborApi(harborParams *HarborParams, imageParams *ImageParams) *HarborApi {
	hapi := &HarborApi{}
	hapi.apiUrl = harborParams.Schema + "://" + harborParams.Hostname + HarborApiPath
	hapi.imgInfo = imageParams
	hapi.harborInfo = harborParams
	return hapi
}

func (hapi HarborApi) GetPingEndpoint() string {
	return hapi.apiUrl + HarborPingEndpoint
}

func (hapi HarborApi) GetLookupEndpoint() string {
	return hapi.apiUrl + HarborProjectEndpoint + "/" + hapi.imgInfo.Project +
		HarborRepositorySubPath + "/" + hapi.imgInfo.Repo +
		HarborArtifactsSubPath + "/" + hapi.imgInfo.Tag
}

func (hapi HarborApi) PrintFullImagePath() string {
	return hapi.harborInfo.Hostname + "/" + hapi.imgInfo.Project + "/" + hapi.imgInfo.Repo + ":" + hapi.imgInfo.Tag +
		"@" + hapi.imgInfo.Sha
}
