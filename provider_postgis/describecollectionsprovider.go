package provider_postgis

import (
	"encoding/json"
	"fmt"
	"net/http"
	. "wfs3_server/codegen"
	pc "wfs3_server/provider_common"
)

type DescribeCollectionsProvider struct {
	data Content
}

func (provider *PostgisProvider) NewDescribeCollectionsProvider(r *http.Request) (Provider, error) {

	path := r.URL.Path // collections
	ct := r.Header.Get("Content-Type")

	if ct == "" {
		ct = JSONContentType
	}

	p := &DescribeCollectionsProvider{}

	csInfo := Content{Links: []Link{}, Collections: []CollectionInfo{}}
	// create Links
	hrefBase := fmt.Sprintf("%s%s", provider.serviceEndpoint, path) // /collections
	links, _ := pc.CreateLinks(hrefBase, "self", ct)
	csInfo.Links = append(csInfo.Links, links...)
	for _, cn := range provider.PostGis.Layers {
		clinks, _ := pc.CreateLinks(fmt.Sprintf("%s/%s", hrefBase, cn.Identifier), "item", ct)
		csInfo.Links = append(csInfo.Links, clinks...)
	}

	for _, cn := range provider.PostGis.Layers {

		cInfo := CollectionInfo{
			Name:        cn.Identifier,
			Title:       cn.Identifier,
			Description: cn.Description,
			Crs:         []string{},
			Links:       []Link{},
		}

		chrefBase := fmt.Sprintf("%s/%s", hrefBase, cn.Identifier)

		clinks, _ := pc.CreateLinks(chrefBase, "self", ct)
		cInfo.Links = append(cInfo.Links, clinks...)

		cihrefBase := fmt.Sprintf("%s/items", chrefBase)
		ilinks, _ := pc.CreateLinks(cihrefBase, "item", ct)
		cInfo.Links = append(cInfo.Links, ilinks...)
		csInfo.Collections = append(csInfo.Collections, cInfo)
	}

	p.data = csInfo

	return p, nil
}

func (provider *DescribeCollectionsProvider) Provide() (interface{}, error) {
	return provider.data, nil
}

func (provider *DescribeCollectionsProvider) MarshalJSON(interface{}) ([]byte, error) {
	return json.Marshal(provider.data)
}

func (provider *DescribeCollectionsProvider) MarshalHTML(interface{}) ([]byte, error) {
	return json.Marshal(provider.data)
}