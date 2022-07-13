package config

import (
	"github.com/kyberorg/wait4harbor/cmd/wait4harbor/api"
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
)

const Http = "http"
const Https = "https"
const Sha256Prefix = "sha256"

const (
	defaultHarborSchema = Https
	defaultInterval     = 2 * time.Second
	defaultTimeout      = 1 * time.Minute
)

var (
	harborHostname  = kingpin.Flag("harbor.hostname", "Harbor Endpoint Address").Required().String()
	harborSchema    = kingpin.Flag("harbor.schema", "HTTP or HTTPS").Default(defaultHarborSchema).String()
	harborRobotName = kingpin.Flag("harbor.robot", "Full name for Robot or Username").Required().String()
	harborToken     = kingpin.Flag("harbor.token", "Token or password").Required().String()

	imageProject = kingpin.Flag("image.project", "Harbor Project which contains image").Required().
			String()
	imageRepo = kingpin.Flag("image.repo", "Harbor Repository").Required().String()
	imageTag  = kingpin.Flag("image.tag", "Image Tag").Required().String()
	imageSha  = kingpin.Flag("image.sha", "Image SHA256 to search for").Required().String()

	interval = kingpin.Flag("interval", "Time between checks").Duration()
	timeout  = kingpin.Flag("timeout", "").Duration()
)

//internal vars
var (
	appConfig *applicationConfiguration
)

//applicationConfiguration application configuration values
type applicationConfiguration struct {
	HarborParams *api.HarborParams
	ImageParams  *api.ImageParams
	Interval     time.Duration
	Timeout      time.Duration
}

//GetAppConfig returns application configuration object
func GetAppConfig() *applicationConfiguration {
	return appConfig
}

func init() {
	//parse flags
	kingpin.Parse()

	appConfig = &applicationConfiguration{
		HarborParams: &api.HarborParams{
			Hostname: *harborHostname,
			Schema:   *harborSchema,
			Robot:    *harborRobotName,
			Token:    *harborToken,
		},
		ImageParams: &api.ImageParams{
			Project: *imageProject,
			Repo:    *imageRepo,
			Tag:     *imageTag,
			Sha:     *imageSha,
		},
		Interval: *interval,
		Timeout:  *timeout,
	}

}
