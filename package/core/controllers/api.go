package corecontrollers

import(
  "net/http"

  "oaf-server/package/core/models"
  "oaf-server/package/core/services"
  "oaf-server/package/core/templates"
)

type APIController struct {
}

func (controller *APIController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderCoreType)

  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    apiService, ok := app.GetService("core").(coreservices.CoreService)
    if !ok {
      panic("Cannot find featureservice")
    }

    resource := apiService.OpenAPI()
    renderer.RenderAPI(coremodels.NewWebcontext(w, r), resource)
  }
}
