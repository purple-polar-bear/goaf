package apifcontrollers

import(
  "net/http"

  "oaf-server/package/models"
  "oaf-server/package/core"
  coretemplates "oaf-server/package/templates/core"
)

type APIController struct {
}

func (controller *APIController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderCoreType)

  return func(w http.ResponseWriter, r *http.Request, routeParameters models.MatchedRouteParameters) {
    apiService, ok := app.GetService("core").(apifcore.CoreService)
    if !ok {
      panic("Cannot find featureservice")
    }

    resource := apiService.OpenAPI()
    renderer.RenderAPI(models.NewWebcontext(w, r), resource)
  }
}
