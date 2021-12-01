package viewmodels

import (
	"github.com/go-spatial/geom/encoding/geojson"
)

type Feature struct {
	// overwrite ID in geojson.Feature so strings are also allowed as id
	ID interface{} `json:"id,omitempty"`
	geojson.Feature
	// Added Links in de document
	Links []Link `json:"links,omitempty"`

	Foo interface{}
}

func (c *Feature) MarshalJSON() ([]byte, error) {
	nestedResult := c.MarshalJSON()
	
	return json.Marshal(map[string]interface{}{
	}
}
