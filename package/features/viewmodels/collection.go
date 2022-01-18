package viewmodels

import(
  "oaf-server/package/core/viewmodels"
)

type Collection struct {
  /* indicator about the type of the items in the collection (the default value is 'feature').	*/
	ItemType string `json:"itemType,omitempty"`
	Links []*viewmodels.Link `json:"links"`
	/* human readable title of the collection	*/
	Title string `json:"title,omitempty"`
	/* the list of coordinate reference systems supported by the service	*/
	Crs []string `json:"crs,omitempty"`
	/* a description of the features in the collection	*/
	Description string `json:"description,omitempty"`
	/* The extent of the features in the collection. In the Core only spatial and temporal
    extents are specified. Extensions may add additional members to represent other
    extents, for example, thermal or pressure ranges.	*/
	Extent *Extent `json:"extent,omitempty"`
	/* identifier of the collection used, for example, in URIs	*/
	Id string `json:"id"`
}
