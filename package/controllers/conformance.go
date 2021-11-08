package apifcontrollers

import(
  "net/http"

  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"
)

// The conformance controller returns the conformance classes of the API
// TODO: build this list - at least partly - automatically, based on
// configuration
type ConformanceController interface {
  models.BaseController
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

func (controller *conformanceController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderConformanceType)
  return func(w http.ResponseWriter, r *http.Request) {
    resource := &viewmodels.Conformanceclasses{
      ConformsTo: controller.ConformanceClasses(),
    }

    renderer.RenderConformance(models.NewWebcontext(w, r), resource)
  }
}
