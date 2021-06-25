package postgis

import (
	"net/http"
	"oaf-server/codegen"

	"github.com/getkin/kin-openapi/openapi3"
)

type GetApiProvider struct {
	data        *openapi3.T
	contenttype string
}

func (pp *PostgisProvider) NewGetApiProvider(r *http.Request) (codegen.Provider, error) {
	p := &GetApiProvider{}
	p.contenttype = r.Header.Get("Content-Type")

	p.data = pp.ApiProcessed
	return p, nil
}

func (gap *GetApiProvider) Provide() (interface{}, error) {
	return gap.data, nil
}

func (gap *GetApiProvider) ContentType() string {
	return gap.contenttype
}

func (gap *GetApiProvider) String() string {
	return "api"
}

func (gap *GetApiProvider) SrsId() string {
	return "n.a"
}
