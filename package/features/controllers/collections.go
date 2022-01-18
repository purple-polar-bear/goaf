package featurescontrollers

import(
  "net/http"

  "oaf-server/package/core/services"
  "oaf-server/package/core/models"
  coreviewmodels "oaf-server/package/core/viewmodels"
  "oaf-server/package/features/models"
  "oaf-server/package/features/services"
  "oaf-server/package/features/viewmodels"
  "oaf-server/package/features/templates/core"
)

type CollectionsController struct {

}

func (controller *CollectionsController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    featureService, ok := app.GetService("features").(featureservices.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    coreservice, ok := app.GetService("core").(coreservices.CoreService)
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
    renderer.RenderCollections(coremodels.NewWebcontext(w, r), resource)
  }
}

func BuildCollection(handler coremodels.Handler, app coremodels.Application, encoding *coremodels.ContentTypeUrlEncoding, collection featuremodels.Collection, templates []coremodels.Handler) *viewmodels.Collection {
  return &viewmodels.Collection{
    Id: collection.Id(),
    Title: collection.Title(),
    Description: collection.Description(),
    Links: AddCollectionLinkList(handler, []*coreviewmodels.Link{}, app, encoding, collection, templates),
  }
}

func AddCollectionsLinkList(handler coremodels.Handler, app coremodels.Application, encoding *coremodels.ContentTypeUrlEncoding) []*coreviewmodels.Link {
  links := []*coreviewmodels.Link{}
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  templates := app.Templates("featurecollections", "")
  for _, template := range templates {
    link := &coreviewmodels.Link{
      Title: template.Title(),
      Rel: template.Rel(handler.Type()),
      Type: template.Type(),
      Href: template.Href(baseUrl, params, encoding),
    }

    links = append(links, link)
  }

  return links
}

func AddCollectionLinkList(handler coremodels.Handler, links []*coreviewmodels.Link, app coremodels.Application, encoding *coremodels.ContentTypeUrlEncoding, collection featuremodels.Collection, templates []coremodels.Handler) []*coreviewmodels.Link {
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  params["collection_id"] = collection.Id()

  for _, template := range templates {
    link := &coreviewmodels.Link{
      Title: template.Title(),
      Rel: template.Rel(handler.Type()),
      Type: template.Type(),
      Href: template.Href(baseUrl, params, encoding),
    }

    links = append(links, link)
  }

  return links
}
