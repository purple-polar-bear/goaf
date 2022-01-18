package corecontrollers

import(
  "net/http"

  "oaf-server/package/core/models"
  "oaf-server/package/core/viewmodels"
  "oaf-server/package/core/templates"
)

// The conformance controller returns the conformance classes of the API
// TODO: build this list - at least partly - automatically, based on
// configuration
type ConformanceController interface {
  coremodels.BaseController
  ConformanceClasses() []string
}

type conformanceController struct {
  conformanceClasses []string
}

func NewConformanceController() ConformanceController {
  return &conformanceController{
    conformanceClasses: defaultConformanceClasses(),
  }
}

func (controller *conformanceController) ConformanceClasses() []string {
  return controller.conformanceClasses
}

func defaultConformanceClasses() []string {
  return []string{
    "http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/core",
    "http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/oas30",
    "http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/html",
    "http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/geojson",
  }
}

func (controller *conformanceController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderCoreType)
  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    resource := &viewmodels.Conformanceclasses{
      ConformsTo: controller.ConformanceClasses(),
    }

    renderer.RenderConformance(coremodels.NewWebcontext(w, r), resource)
  }
}
