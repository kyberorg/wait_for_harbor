package config

import "strings"

const (
	HostnameMissing        = "Harbor Hostname cannot be empty"
	RobotNameMissing       = "Harbor Robot or Username cannot be empty"
	TokenMissing           = "Harbor Token or Password cannot be empty"
	ImageProjectMissing    = "Image Project cannot be empty"
	ImageRepositoryMissing = "Image Repository cannot be empty"
	ImageTagMissing        = "Image Tag cannot be empty"
	ImageShaMissing        = "Image SHA cannot be empty. Nothing to search for."
	AllGood                = "all good"
)

func (conf applicationConfiguration) ValidateParams() (bool, string) {
	//hostname
	if isEmptyString(conf.HarborParams.Hostname) {
		return false, HostnameMissing
	}
	//schema
	if &conf.HarborParams.Schema == nil || strings.ToLower(conf.HarborParams.Schema) != Http {
		conf.HarborParams.Schema = defaultHarborSchema
	}
	//robot
	if isEmptyString(conf.HarborParams.Robot) {
		return false, RobotNameMissing
	}
	//token
	if isEmptyString(conf.HarborParams.Token) {
		return false, TokenMissing
	}

	//image project
	if isEmptyString(conf.ImageParams.Project) {
		return false, ImageProjectMissing
	}

	//image repo
	if isEmptyString(conf.ImageParams.Repo) {
		return false, ImageRepositoryMissing
	}

	//image tag
	if isEmptyString(conf.ImageParams.Tag) {
		return false, ImageTagMissing
	}

	//image SHA
	if isEmptyString(conf.ImageParams.Sha) {
		return false, ImageShaMissing
	}
	if !strings.HasPrefix(conf.ImageParams.Sha, Sha256Prefix) {
		conf.ImageParams.Sha = Sha256Prefix + ":" + conf.ImageParams.Sha
	}

	//interval
	if &conf.Interval == nil {
		conf.Interval = defaultInterval
	}
	//timeout
	if &conf.Timeout == nil {
		conf.Timeout = defaultTimeout
	}

	return true, AllGood
}

func isEmptyString(str string) bool {
	return len(strings.TrimSpace(str)) <= 0
}
