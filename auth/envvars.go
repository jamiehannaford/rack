package auth

import (
	"os"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
)

func envvars() (gophercloud.AuthOptions, string, error) {
	ao, err := rackspace.AuthOptionsFromEnv()
	region := os.Getenv("RS_REGION_NAME")
	return ao, region, err
}
