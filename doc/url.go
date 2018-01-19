package doc

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/rumyantseva/test2doc/parse"
)

type URL struct {
	rawURL            *url.URL
	ParameterizedPath string
	Parameters        []Parameter
}

func NewURL(req *http.Request) *URL {
	u := &URL{
		rawURL: req.URL,
	}
	u.ParameterizedPath, u.Parameters = paramPath(req)
	return u
}

func paramPath(req *http.Request) (string, []Parameter) {
	uri, err := url.QueryUnescape(req.URL.Path)
	if err != nil {
		// fall back to unescaped uri
		uri = req.URL.Path
	}

	vars := (*parse.Extractor)(req)
	params := []Parameter{}

	for k, v := range vars {
		var helper Parameter
		if p, ok := h.parameters[uri]; ok {
			helper = p
		}

		uri = strings.Replace(uri, "/"+v, "/{"+k+"}", 1)
		params = append(params, MakeParameter(k, v, helper))
	}

	var queryKeys []string
	queryParams := req.URL.Query()

	for k, vs := range queryParams {
		queryKeys = append(queryKeys, k)

		// just take first value
		params = append(params, MakeParameter(k, vs[0], Parameter{}))
	}

	var queryKeysStr string
	if len(queryKeys) > 0 {
		queryKeysStr = "{?" + strings.Join(queryKeys, ",") + "}"
	}

	uri = uri + queryKeysStr

	return uri, params
}
