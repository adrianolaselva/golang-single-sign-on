package common

import (
	"net/http"
	"strconv"
	"strings"
)

type HTTPRequestCommon interface {
	GetPaginateParameters() (*QueryParameter, error)
}

type HTTPRequestCommonImpl struct {
	request *http.Request
}

func NewHTTPRequestCommon(request *http.Request) *HTTPRequestCommonImpl {
	return &HTTPRequestCommonImpl{request}
}

func (h *HTTPRequestCommonImpl) GetPaginateParameters() (*QueryParameter, error) {
	queryParameter := QueryParameter{}
	queryParameter.Limit = 10
	queryParameter.Page = 1

	queryStrings := h.request.URL.Query()

	parameters := make(map[string]interface{})
	for key := range queryStrings {
		value := strings.Join(queryStrings[key], "")
		switch key {
			case "page":
				queryParameter.Page, _ = strconv.Atoi(value)
				h.request.URL.Query().Del(key)
				continue
			case "limit":
				queryParameter.Limit, _ = strconv.Atoi(value)
				h.request.URL.Query().Del(key)
				continue
			default:
				switch value {
					case "true":
						parameters[key], _ = strconv.ParseBool(value)
						continue
					case "false":
						parameters[key], _ = strconv.ParseBool(value)
						continue
				}
				parameters[key] = value
		}
	}

	queryParameter.Filters = parameters

	return &queryParameter, nil
}

type QueryParameter struct {
	Limit int
	Page int
	Filters map[string]interface{}
	OrderBy string
	OrderDir string
}