package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/serboox/os-cli/src/exceptions"
)

// Response common interface
type Response interface{}

type sendDataCtx struct {
	methodName string
	urlMethod  string
	url        string
	res        Response
	headers    map[string]string
	newReader  *strings.Reader
}

func (ctx *sendDataCtx) Send() (res *http.Response, err error) {
	newReq, err := http.NewRequest(
		ctx.urlMethod, //"POST"
		ctx.url,
		ctx.newReader,
	)
	if err != nil {
		return nil, &exceptions.HTTPError{
			Method:  ctx.methodName,
			Message: err.Error(),
		}
	}

	for key, value := range ctx.headers {
		newReq.Header.Add(key, value)
	}

	res, err = http.DefaultClient.Do(newReq)
	if err != nil {
		return nil, &exceptions.JSONError{
			Method:  ctx.methodName,
			Message: err.Error(),
		}
	}

	defer res.Body.Close()

	resBytes, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(resBytes, ctx.res)
	if err != nil {
		return nil, &exceptions.JSONError{
			Method:  ctx.methodName,
			Message: err.Error(),
			JSON:    string(resBytes),
		}
	}

	return res, nil
}
