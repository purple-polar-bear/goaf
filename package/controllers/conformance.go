package apifcontrollers

// The conformance controller returns the conformance classes of the API
// TODO: build this list - at least partly - automatically, based on
// configuration
type ConformanceController interface {
  ConformanceClasses() []string
}

type conformanceController struct {
}

func NewConformanceController() ConformanceController {
  return &conformanceController{
  }
}

func (controller *conformanceController) ConformanceClasses() []string {
  return []string{
    "http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/core",
    "http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/oas30",
    "http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/html",
    "http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/geojson",
  }
}

func (controller *conformanceController) Handle(context templates.Context, render templates.RenderConformanceFunc) {
}
