package controllers

import (
	"errors"
	"net/http"

	"github.com/serboox/os-cli/src/configs"
	"github.com/serboox/os-cli/src/models"
	"github.com/serboox/os-cli/src/views"
)

//CreateInstance create instance in openstack
func CreateInstance(
	cliArgs *configs.CliArgs,
) (err error) {
	var (
		headers http.Header
		token   string
		novaURL string
	)

	resAuthTokens := new(models.ResAuthTokens)

	if headers, err = resAuthTokens.Post(cliArgs); err != nil {
		return err
	}

	if token := headers.Get(configs.XSubjectTokenKey); token == "" {
		return errors.New("token not be empty")
	}

	if novaURL, err = resAuthTokens.FindEndpointURL(
		configs.Nova,
	); err != nil {
		return err
	}

	resServers := new(models.ResServers)
	if _, err = resServers.Post(
		cliArgs,
		token,
		novaURL+"/servers",
	); err != nil {
		return err
	}

	views.ShowNewServer(resServers)

	return nil
}
