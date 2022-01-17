package apifcontrollers

import(
  "net/http"

  "oaf-server/package/core"
  "oaf-server/package/features"
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"
)

type CollectionsController struct {

}

func (controller *CollectionsController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(handler models.Handler, w http.ResponseWriter, r *http.Request, routeParameters models.MatchedRouteParameters) {
    featureService, ok := app.GetService("features").(features.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    coreservice, ok := app.GetService("core").(apifcore.CoreService)
    if !ok {
      panic("Cannot find coreservice")
    }

    encoding := coreservice.ContentTypeUrlEncoder()
    collections := []*viewmodels.Collection{}
    links := AddCollectionsLinkList(handler, app, encoding)
    collectionRoute := app.Templates("featurecollection", "")
    collectionItemsRoute := append(collectionRoute, app.Templates("features", "")...)

    for _, collection := range featureService.Collections() {
      collections = append(collections, BuildCollection(handler, app, encoding, collection, collectionItemsRoute))
      links = AddCollectionLinkList(handler, links, app, encoding, collection, collectionRoute)
    }

    resource := &viewmodels.Collections{
      Collections: collections,
      Links: links,
    }
    renderer.RenderCollections(models.NewWebcontext(w, r), resource)
  }
}

func BuildCollection(handler models.Handler, app models.Application, encoding *models.ContentTypeUrlEncoding, collection features.Collection, templates []models.Handler) *viewmodels.Collection {
  return &viewmodels.Collection{
    Id: collection.Id(),
    Title: collection.Title(),
    Description: collection.Description(),
    Links: AddCollectionLinkList(handler, []*viewmodels.Link{}, app, encoding, collection, templates),
  }
}

func AddCollectionsLinkList(handler models.Handler, app models.Application, encoding *models.ContentTypeUrlEncoding) []*viewmodels.Link {
  links := []*viewmodels.Link{}
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  templates := app.Templates("featurecollections", "")
  for _, template := range templates {
    link := &viewmodels.Link{
      Title: template.Title(),
      Rel: template.Rel(handler.Type()),
      Type: template.Type(),
      Href: template.Href(baseUrl, params, encoding),
    }

    links = append(links, link)
  }

  return links
}

func AddCollectionLinkList(handler models.Handler, links []*viewmodels.Link, app models.Application, encoding *models.ContentTypeUrlEncoding, collection features.Collection, templates []models.Handler) []*viewmodels.Link {
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  params["collection_id"] = collection.Id()

  for _, template := range templates {
    link := &viewmodels.Link{
      Title: template.Title(),
      Rel: template.Rel(handler.Type()),
      Type: template.Type(),
      Href: template.Href(baseUrl, params, encoding),
    }

    links = append(links, link)
  }

  return links
}
