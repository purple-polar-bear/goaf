package templates

type RenderLandingpageFunc func (http.ResponseWriter, *http.Request, models.Landingpage)

// Render function declaration for conformance pages
type RenderConformanceFunc func (http.ResponseWriter, *http.Request, []string)
