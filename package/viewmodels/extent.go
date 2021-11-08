package viewmodels

type Extent struct {
	/* The spatial extent of the features in the collection.	*/
	Spatial *Spatial `json:"spatial,omitempty"`
	/* The temporal extent of the features in the collection.	*/
	Temporal *Temporal `json:"temporal,omitempty"`
}
