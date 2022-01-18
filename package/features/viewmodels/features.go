package viewmodels

import(
	"time"

  "oaf-server/package/core/viewmodels"
)

type Features struct {
	NumberReturned int            		`json:"numberReturned,omitempty"`
	TimeStamp      string         		`json:"timeStamp,omitempty"`
	Type           string         		`json:"type"`
	Features       []interface{}  		`json:"features"`
	Links          []*viewmodels.Link	`json:"links,omitempty"`
	NumberMatched  int            		`json:"numberMatched,omitempty"`
	Crs            string         		`json:"crs,omitempty"`
	Offset         int            		`json:"-"`
}

const CRS_GPS = "http://www.opengis.net/def/crs/EPSG/0/4326"

// Default CRS is the GPS system
func NewFeatureCollection() *Features {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc).String()

  return &Features{
    Type: "FeatureCollection",
    Crs: CRS_GPS,
		TimeStamp: now,
		Features: []interface{}{},
		Links: []*viewmodels.Link{},
  }
}
