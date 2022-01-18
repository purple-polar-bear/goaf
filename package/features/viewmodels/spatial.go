package viewmodels

type Spatial struct {
	/* One or more bounding boxes that describe the spatial extent of the dataset.
    In the Core only a single bounding box is supported. Extensions may support
    additional areas. If multiple areas are provided, the union of the bounding
    boxes describes the spatial extent.	*/
	Bbox [][]float64 `json:"bbox,omitempty"`
	/* Coordinate reference system of the coordinates in the spatial extent
    (property `bbox`). The default reference system is WGS 84 longitude/latitude.
    In the Core this is the only supported coordinate reference system.
    Extensions may support additional coordinate reference systems and add
    additional enum values.	*/
	Crs string `json:"crs,omitempty"`
}
