package apifcontrollers

import(
  "net/http"

  "oaf-server/package/features"
  "oaf-server/package/models"
  "oaf-server/package/templates/core"
  "oaf-server/package/viewmodels"
)

type FeaturesController struct {

}

func (controller *FeaturesController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(w http.ResponseWriter, r *http.Request) {
    featuresRoute := app.Templates("features", "")

    featureService, ok := app.GetService("features").(features.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    featureParams := buildFeatureParams(app)
    features := featureService.Features(featureParams)
    links := BuildFeaturesLinks(featuresRoute, featureParams, features)

    resource := viewmodels.NewFeatureCollection()
    items := features.Items()
    itemLength := len(items)
    resource.Features = make([]interface{}, itemLength)
    for index, item := range items {
      resource.Features[index] = item
    }
    
    resource.Links = links
    resource.NumberReturned = itemLength
    renderer.RenderItems(models.NewWebcontext(w, r), resource)
  }
}

func buildFeatureParams(app models.Application) *features.FeaturesParams {
  return features.NewFeaturesParams()
}

func BuildFeaturesLinks(templates []models.Handler, params *features.FeaturesParams, items features.Features) []*viewmodels.Link {
  result := []*viewmodels.Link{}

  return result
}
