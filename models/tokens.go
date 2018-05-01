package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/serboox/os-cli/configs"
	"github.com/serboox/os-cli/exceptions"
)

// ReqAuthTokens structure declared request JSON
type ReqAuthTokens struct {
	ReqAuth ReqAuthTokensAuth `json:"auth"`
}

// ReqAuthTokensAuth structure declared request JSON part
type ReqAuthTokensAuth struct {
	IDEntity ReqAuthTokensIDEntity `json:"identity"`
}

// ReqAuthTokensIDEntity structure declared request JSON part
type ReqAuthTokensIDEntity struct {
	Methods  []string              `json:"methods"`
	Password ReqAuthTokensPassword `json:"password"`
}

// ReqAuthTokensPassword structure declared request JSON part
type ReqAuthTokensPassword struct {
	User ReqAuthTokensUser `json:"user"`
}

// ReqAuthTokensUser structure declared request JSON part
type ReqAuthTokensUser struct {
	Name     string              `json:"name"`
	Domain   ReqAuthTokensDomain `json:"domain"`
	Password string              `json:"password"`
}

// ReqAuthTokensDomain structure declared request JSON part
type ReqAuthTokensDomain struct {
	Name string `json:"name"`
}

// ResAuthTokens structure declared response JSON
type ResAuthTokens struct {
	Token ResAuthTokensToken `json:"token"`
}

// ResAuthTokensToken structure declared response JSON part
type ResAuthTokensToken struct {
	IsDomani  bool                   `json:"is_domain"`
	Methods   []string               `json:"methods"`
	Roles     []ResAuthTokensRoles   `json:"roles"`
	ExpiresAt string                 `json:"expires_at"`
	Project   ResAuthTokensProject   `json:"project"`
	Catalog   []ResAuthTokensCatalog `json:"catalog"`
}

// ResAuthTokensRoles structure declared response JSON part
type ResAuthTokensRoles struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ResAuthTokensProject structure declared response JSON part
type ResAuthTokensProject struct {
	ID     string              `json:"id"`
	Name   string              `json:"name"`
	Domain ResAuthTokensDomain `json:"domain"`
}

// ResAuthTokensDomain structure declared response JSON part
type ResAuthTokensDomain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ResAuthTokensCatalog structure declared response JSON part
type ResAuthTokensCatalog struct {
	ID        string                   `json:"id"`
	Type      string                   `json:"type"`
	Name      string                   `json:"name"`
	Endpoints []ResAuthTokensEndpoints `json:"endpoints"`
}

// ResAuthTokensEndpoints structure declared response JSON part
type ResAuthTokensEndpoints struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Interface string `json:"interface"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
}

//Post Password authentication with unscoped authorization
func (res *ResAuthTokens) Post(cliArgs *configs.CliArgs) (
	headers http.Header, err error,
) {
	methodName := "ResAuthTokens.Post"

	req := ReqAuthTokens{
		ReqAuth: ReqAuthTokensAuth{
			IDEntity: ReqAuthTokensIDEntity{
				Methods: []string{"password"},
				Password: ReqAuthTokensPassword{
					User: ReqAuthTokensUser{
						Name: cliArgs.Login,
						Domain: ReqAuthTokensDomain{
							Name: "Default",
						},
						Password: cliArgs.Password,
					},
				},
			},
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
		urlMethod:  "POST",
		url:        cliArgs.AuthHost,
		headers: map[string]string{
			"Content-Type": "application/json",
		},
		res:       res,
		newReader: newReader,
	}

	resp, err := ctx.Send()
	return resp.Header, err
}

//FindEndpointURL find endpoint in result set
func (res *ResAuthTokens) FindEndpointURL(end configs.Endpoint) (
	string, error,
) {
	for _, catalog := range res.Token.Catalog {
		if catalog.Name == configs.GetEntrypoint(end) {
			for _, point := range catalog.Endpoints {
				if point.Interface == "public" {
					return point.URL, nil
				}
			}
		}
	}
	return "", errors.New("Endpoint not be found in result set")
}
