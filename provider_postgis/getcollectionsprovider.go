package provider_postgis

import (
	"fmt"
	"net/http"
	cg "wfs3_server/codegen"
	pc "wfs3_server/provider_common"
)

type GetCollectionsProvider struct {
	data cg.Collections
}

func (provider *PostgisProvider) NewGetCollectionsProvider(r *http.Request) (cg.Provider, error) {

	path := r.URL.Path // collections
	ct := r.Header.Get("Content-Type")

	p := &GetCollectionsProvider{}

	csInfo := cg.Collections{Links: []cg.Link{}, Collections: []cg.Collection{}}
	// create Links
	hrefBase := fmt.Sprintf("%s%s", provider.CommonProvider.ServiceEndpoint, path) // /collections
	links, _ := pc.CreateLinks("collections ", hrefBase, "self", ct)
	csInfo.Links = append(csInfo.Links, links...)
	for _, cn := range provider.PostGis.Layers {
		clinks, _ := pc.CreateLinks("collection "+cn.Identifier, fmt.Sprintf("%s/%s", hrefBase, cn.Identifier), "item", ct)
		csInfo.Links = append(csInfo.Links, clinks...)
	}

	for _, cn := range provider.PostGis.Layers {

		cInfo := cg.Collection{
			Id:          cn.Identifier,
			Title:       cn.Identifier,
			Description: cn.Description,
			Crs:         []string{},
			Links:       []cg.Link{},
		}

		chrefBase := fmt.Sprintf("%s/%s", hrefBase, cn.Identifier)

		clinks, _ := pc.CreateLinks("collection "+cn.Identifier, chrefBase, "self", ct)
		cInfo.Links = append(cInfo.Links, clinks...)

		cihrefBase := fmt.Sprintf("%s/items", chrefBase)
		ilinks, _ := pc.CreateLinks("items "+cn.Identifier, cihrefBase, "item", ct)
		cInfo.Links = append(cInfo.Links, ilinks...)
		csInfo.Collections = append(csInfo.Collections, cInfo)
	}

	p.data = csInfo

	return p, nil
}

func (provider *GetCollectionsProvider) Provide() (interface{}, error) {
	return provider.data, nil
}

func (provider *GetCollectionsProvider) String() string {
	return "getcollections"
}

func (provider *GetCollectionsProvider) SrsId() string {
	return "n.a."
}
