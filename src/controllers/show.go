package controllers

import (
	"errors"
	"net/http"

	"github.com/serboox/os-cli/src/configs"
	"github.com/serboox/os-cli/src/models"
	"github.com/serboox/os-cli/src/views"
)

// ShowInstances show instance list
func ShowInstances(cliArgs *configs.CliArgs) (err error) {
	var (
		headers http.Header
		token   string
		novaURL string
	)

	resAuthTokens := new(models.ResAuthTokens)

	if headers, err = resAuthTokens.Post(cliArgs); err != nil {
		return
	}

	if token = headers.Get(configs.XSubjectTokenKey); token == "" {
		return errors.New("token not be empty")
	}

	if novaURL, err = resAuthTokens.FindEndpointURL(
		configs.Nova,
	); err != nil {
		return
	}

	resServersDetail := models.ResServersDetail{}
	if _, err = resServersDetail.Get(
		token,
		novaURL+"/servers/detail",
	); err != nil {
		return
	}

	if err = views.ShowServersDetail(resServersDetail); err != nil {
		return
	}

	return
}
