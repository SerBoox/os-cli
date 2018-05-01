package controllers

import (
	"errors"

	"github.com/serboox/os-cli/configs"
	"github.com/serboox/os-cli/models"
	"github.com/serboox/os-cli/views"
)

//CreateInstance create instance in openstack
func CreateInstance(
	cliArgs *configs.CliArgs,
) error {

	resAuthTokens := models.ResAuthTokens{}
	headers, err := resAuthTokens.Post(cliArgs)
	if err != nil {
		return err
	}

	token := headers.Get(configs.XSubjectTokenKey)
	if token == "" {
		return errors.New("Token not be empty")
	}

	novaURL, err := resAuthTokens.FindEndpointURL(configs.Nova)
	if err != nil {
		return err
	}

	resServers := models.ResServers{}
	_, err = resServers.Post(cliArgs, token, novaURL+"/servers")
	if err != nil {
		return err
	}

	views.ShowNewServer(resServers)

	return nil
}
