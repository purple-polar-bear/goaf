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
    templates := app.Templates("feature", "")

    featureService, ok := app.GetService("features").(features.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    collectionId := routeParameters.Get("collection_id")
    featureId := routeParameters.Get("item_id")
    feature := featureService.Feature(collectionId, featureId)
    baseUrl := app.Config().FullUri()
    hrefParams := make(map[string]string)
    hrefParams["collection_id"] = collectionId
    hrefParams["item_id"] = featureId
    links := []*viewmodels.Link{}
    // current link
    for _, template := range templates {
  		baseHref := template.Href(baseUrl, hrefParams)
      link := &viewmodels.Link{
        Title: template.Title(),
        Rel: template.Rel(),
        Type: template.Type(),
        Href: baseHref,
      }

  		links = append(links, link)
  	}

    resource := &viewmodels.Feature{
      Feature: feature,
      Links: links,
    }
    renderer.RenderItem(models.NewWebcontext(w, r), resource)
  }
}
