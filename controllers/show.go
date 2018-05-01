package controllers

import (
	"errors"

	"github.com/serboox/os-cli/configs"
	"github.com/serboox/os-cli/models"
	"github.com/serboox/os-cli/views"
)

//ShowInstances show instance list
func ShowInstances(
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

	resServersDetail := models.ResServersDetail{}
	_, err = resServersDetail.Get(token, novaURL+"/servers/detail")
	if err != nil {
		return err
	}

	views.ShowServersDetail(resServersDetail)

	return nil
}
