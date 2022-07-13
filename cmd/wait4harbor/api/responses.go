package api

import "time"

type ErrorResponse struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

type SuccessResponse struct {
	Accessories   interface{} `json:"accessories"`
	AdditionLinks struct {
		BuildHistory struct {
			Absolute bool   `json:"absolute"`
			Href     string `json:"href"`
		} `json:"build_history"`
		Vulnerabilities struct {
			Absolute bool   `json:"absolute"`
			Href     string `json:"href"`
		} `json:"vulnerabilities"`
	} `json:"addition_links"`
	Digest     string `json:"digest"`
	ExtraAttrs struct {
		Architecture string `json:"architecture"`
		Author       string `json:"author"`
		Config       struct {
			Entrypoint   []string `json:"Entrypoint"`
			Env          []string `json:"Env"`
			ExposedPorts struct {
				Eight080TCP struct {
				} `json:"8080/tcp"`
			} `json:"ExposedPorts"`
			Labels struct {
				Maintainer string `json:"maintainer"`
			} `json:"Labels"`
			User       string `json:"User"`
			WorkingDir string `json:"WorkingDir"`
		} `json:"config"`
		Created time.Time `json:"created"`
		Os      string    `json:"os"`
	} `json:"extra_attrs"`
	Icon              string      `json:"icon"`
	ID                int         `json:"id"`
	Labels            interface{} `json:"labels"`
	ManifestMediaType string      `json:"manifest_media_type"`
	MediaType         string      `json:"media_type"`
	ProjectID         int         `json:"project_id"`
	PullTime          time.Time   `json:"pull_time"`
	PushTime          time.Time   `json:"push_time"`
	References        interface{} `json:"references"`
	RepositoryID      int         `json:"repository_id"`
	Size              int         `json:"size"`
	Tags              []struct {
		ArtifactID   int       `json:"artifact_id"`
		ID           int       `json:"id"`
		Immutable    bool      `json:"immutable"`
		Name         string    `json:"name"`
		PullTime     time.Time `json:"pull_time"`
		PushTime     time.Time `json:"push_time"`
		RepositoryID int       `json:"repository_id"`
		Signed       bool      `json:"signed"`
	} `json:"tags"`
	Type string `json:"type"`
}
