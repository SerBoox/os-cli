package models

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/serboox/os-cli/src/configs"
	"github.com/serboox/os-cli/src/exceptions"
)

// ReqServers structure declared request JSON
type ReqServers struct {
	Server ReqServersItem `json:"server"`
}

// ReqServersItem structure declared request JSON part
type ReqServersItem struct {
	Name             string                     `json:"name"`
	ImageRef         string                     `json:"imageRef"`
	FlavorRef        string                     `json:"flavorRef"`
	AvailabilityZone string                     `json:"availability_zone"`
	DiskConfig       string                     `json:"OS-DCF:diskConfig"`
	SecurityGroups   []ReqServersSecurityGroups `json:"security_groups"`
}

// ReqServersSecurityGroups structure declared request JSON part
type ReqServersSecurityGroups struct {
	Name string `json:"name"`
}

// ResServers structure declared response JSON
type ResServers struct {
	Server ResServersItem `json:"server"`
}

// ResServersItem structure declared response JSON part
type ResServersItem struct {
	ID             string                     `json:"id"`
	SecurityGroups []ResServersSecurityGroups `json:"security_groups"`
	DiscConfig     string                     `json:"OS-DCF:diskConfig"`
	Links          []ResServersLinks          `json:"links"`
	AdminPass      string                     `json:"adminPass"`
}

// ResServersSecurityGroups structure declared response JSON part
type ResServersSecurityGroups struct {
	Name string `json:"name"`
}

// ResServersLinks structure declared response JSON part
type ResServersLinks struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// Post Create instance
func (res *ResServers) Post(cliArgs *configs.CliArgs, token, url string) (
	resp *http.Response, err error,
) {
	methodName := "ResServers.Post"

	securityGroups := make([]ReqServersSecurityGroups, 0, 1)
	securityGroups = append(
		securityGroups,
		ReqServersSecurityGroups{
			Name: "default",
		},
	)

	req := ReqServers{
		Server: ReqServersItem{
			Name:             cliArgs.InstName,
			ImageRef:         cliArgs.ImageRef,
			FlavorRef:        strconv.FormatInt(cliArgs.FlavorRef, 10),
			AvailabilityZone: configs.DefaultAvailabilityZone,
			DiskConfig:       "AUTO",
			SecurityGroups:   securityGroups,
		},
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, &exceptions.JSONError{
			Method:  methodName,
			Message: err.Error(),
			JSON:    string(reqBytes),
		}
	}

	newReader := strings.NewReader(string(reqBytes))

	ctx := sendDataCtx{
		methodName: methodName,
		urlMethod:  http.MethodPost,
		url:        url,
		headers: map[string]string{
			"Content-Type": "application/json",
			configs.XAuth:  token,
		},
		res:       res,
		newReader: newReader,
	}

	return ctx.Send()
}
