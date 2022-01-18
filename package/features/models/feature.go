package featuremodels

import (
	"encoding/json"

  "oaf-server/package/core/viewmodels"
)

type Feature struct {
	// overwrite ID in geojson.Feature so strings are also allowed as id
	Feature interface{}
	// Added Links in de document
	Links []*viewmodels.Link `json:"links,omitempty"`
}

func (c *Feature) MarshalJSON() ([]byte, error) {
	p, err := json.Marshal(c.Feature)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(p, &data); err != nil {
		return nil, err
	}

	data["Links"] = c.Links

	return json.Marshal(data)
}
