package featurescontrollers

import(
  "net/http"

  "oaf-server/package/core/services"
  "oaf-server/package/core/models"
  coreviewmodels "oaf-server/package/core/viewmodels"
  "oaf-server/package/features/models"
  "oaf-server/package/features/services"
  "oaf-server/package/features/templates/core"

  // "github.com/go-spatial/geom/encoding/geojson"
)

type FeatureController struct {

}

func (controller *FeatureController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    templates := app.Templates("feature", "")

    featureService, ok := app.GetService("features").(featureservices.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    coreservice, ok := app.GetService("core").(coreservices.CoreService)
    if !ok {
      panic("Cannot find coreservice")
    }

    encoding := coreservice.ContentTypeUrlEncoder()
    collectionId := routeParameters.Get("collection_id")
    featureId := routeParameters.Get("item_id")
    feature := featureService.Feature(collectionId, featureId)
    baseUrl := app.Config().FullUri()
    hrefParams := make(map[string]string)
    hrefParams["collection_id"] = collectionId
    hrefParams["item_id"] = featureId
    links := []*coreviewmodels.Link{}
    // current link
    for _, template := range templates {
  		baseHref := template.Href(baseUrl, hrefParams, encoding)
      link := &coreviewmodels.Link{
        Title: template.Title(),
        Rel: template.Rel(handler.Type()),
        Type: template.Type(),
        Href: baseHref,
      }

  		links = append(links, link)
  	}

    resource := &featuremodels.Feature{
      Feature: feature,
      Links: links,
    }
    renderer.RenderItem(coremodels.NewWebcontext(w, r), resource)
  }
}
