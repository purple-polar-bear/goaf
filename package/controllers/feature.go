package apifcontrollers

import(
  "net/http"

  "oaf-server/package/features"
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"

  // "github.com/go-spatial/geom/encoding/geojson"
)

type FeatureController struct {

}

func (controller *FeatureController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(w http.ResponseWriter, r *http.Request, routeParameters models.MatchedRouteParameters) {
    // featuresRoute := app.Templates("features", "")

    featureService, ok := app.GetService("features").(features.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    id := routeParameters.Get("item_id")
    feature := featureService.Feature(id)
    resource := &viewmodels.Feature{
      Feature: feature,
      Links: []viewmodels.Link{},
    }
    renderer.RenderItem(models.NewWebcontext(w, r), resource)
  }
}
