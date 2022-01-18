package viewmodels

type Temporal struct {
	/* Coordinate reference system of the coordinates in the temporal extent
    (property `interval`). The default reference system is the Gregorian calendar.
    In the Core this is the only supported temporal coordinate reference system.
    Extensions may support additional temporal coordinate reference systems and add
    additional enum values.	*/
	Trs string `json:"trs,omitempty"`
	/* One or more time intervals that describe the temporal extent of the dataset.
    The value `null` is supported and indicates an open time interval.
    In the Core only a single time interval is supported. Extensions may support
    multiple intervals. If multiple intervals are provided, the union of the
    intervals describes the temporal extent.	*/
	Interval [][]string `json:"interval,omitempty"`
}
