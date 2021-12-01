package apifcontrollers

import(
  "net/http"

  "oaf-server/package/features"
  "oaf-server/package/models"
  "oaf-server/package/templates/core"
)

type FeatureController struct {

}

func (controller *FeatureController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(w http.ResponseWriter, r *http.Request, routeParameters models.MatchedRouteParameters) {
    featuresRoute := app.Templates("features", "")

    featureService, ok := app.GetService("features").(features.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    id := routeParameters.Get("item_id")
    resource := featureService.Feature(id)
    renderer.RenderItem(models.NewWebcontext(w, r), resource)
  }
}
