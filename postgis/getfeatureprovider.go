package postgis

import (
	"errors"
	"fmt"
	"net/http"
	"oaf-server/codegen"
	"oaf-server/provider"
)

type GetFeatureProvider struct {
	data  *Feature
	srsid string
}

func (pp *PostgisProvider) NewGetFeatureProvider(r *http.Request) (codegen.Provider, error) {

	collectionId, featureId, _ := codegen.ParametersForGetFeature(r)

	featureIdParam := featureId
	bboxParam := pp.PostGis.BBox

	p := &GetFeatureProvider{srsid: fmt.Sprintf("EPSG:%d", pp.PostGis.SrsId)}

	path := r.URL.Path
	ct := r.Header.Get("Content-Type")

	for _, cn := range pp.PostGis.Layers {
		// maybe convert to map, but not thread safe!
		if cn.Identifier != collectionId {
			continue
		}

		pathItem := pp.ApiProcessed.Paths.Find("/collections/pand/items/{featureId}")
		if pathItem == nil {
			return p, errors.New("Invalid path :" + path)
		}

		for k := range r.URL.Query() {
			if notfound := pathItem.Get.Parameters.GetByInAndName("query", k) == nil; notfound {
				return p, errors.New("Invalid query parameter :" + k)
			}
		}

		whereMap := make(map[string]string)
		fcGeoJSON, err := pp.PostGis.GetFeatures(r.Context(), pp.PostGis.db, cn, whereMap, 0, 1, featureIdParam, bboxParam)

		if err != nil {
			return nil, err
		}

		if len(fcGeoJSON.Features) >= 1 {
			feature := fcGeoJSON.Features[0]

			hrefBase := fmt.Sprintf("%s%s", pp.CommonProvider.ServiceEndpoint, path) // /collections
			links, _ := provider.CreateLinks("feature", hrefBase, "self", ct)
			feature.Links = links

			p.data = feature

		} else {
			return p, fmt.Errorf("Feature with id: %s not found", string(featureId))
		}

		return p, nil
	}

	return p, errors.New("Cannot find layer : " + collectionId)
}

func (gfp *GetFeatureProvider) Provide() (interface{}, error) {
	return gfp.data, nil
}

func (gfp *GetFeatureProvider) String() string {
	return "getfeature"
}

func (gfp *GetFeatureProvider) SrsId() string {
	return gfp.srsid
}
