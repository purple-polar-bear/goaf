package core

import (
	"oaf-server/package/core/viewmodels"
	"oaf-server/package/features/models"

	"github.com/go-spatial/geom/encoding/geojson"
)

type FeatureCollection struct {
	NumberReturned int64           `json:"numberReturned,omitempty"`
	TimeStamp      string          `json:"timeStamp,omitempty"`
	Type           string          `json:"type"`
	Features       []*Feature      `json:"features"`
	Links          []*viewmodels.Link `json:"links,omitempty"`
	NumberMatched  int64           `json:"numberMatched,omitempty"`
	Crs            string          `json:"crs,omitempty"`
	Offset         int64           `json:"-"`
	Next           bool
	RequestParams  *featuremodels.FeaturesParams
}

type Feature struct {
	// overwrite ID in geojson.Feature so strings are also allowed as id
	ID interface{} `json:"id,omitempty"`
	geojson.Feature
	// Added Links in de document
	Links []*viewmodels.Link `json:"links,omitempty"`
}

func (features *FeatureCollection) HasNext() bool {
	return features.Next
}

func (features *FeatureCollection) NextLimit() int {
	if !features.HasNext() {
		return 0
	}

	return features.RequestParams.Limit
}

func (features *FeatureCollection) NextOffset() int {
	return features.RequestParams.Offset + features.RequestParams.Limit
}

func (features *FeatureCollection) Items() []*Feature {
	return features.Features
}
