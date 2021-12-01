package apifcontrollers

import(
  "net/http"

  "oaf-server/package/features"
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"
)

type CollectionsController struct {

}

func (controller *CollectionsController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(w http.ResponseWriter, r *http.Request, routeParameters models.MatchedRouteParameters) {
    // Given some provider

    collections := []*viewmodels.Collection{}
    links := AddCollectionsLinkList(app)
    collectionRoute := app.Templates("featurecollection", "")
    collectionItemsRoute := append(collectionRoute, app.Templates("features", "")...)

    featureService, ok := app.GetService("features").(features.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    for _, collection := range featureService.Collections() {
      collections = append(collections, BuildCollection(app, collection, collectionItemsRoute))
      links = AddCollectionLinkList(links, app, collection, collectionRoute)
    }

    resource := &viewmodels.Collections{
      Collections: collections,
      Links: links,
    }
    renderer.RenderCollections(models.NewWebcontext(w, r), resource)
  }
}

func BuildCollection(app models.Application, collection features.Collection, templates []models.Handler) *viewmodels.Collection {
  return &viewmodels.Collection{
    Id: collection.Id(),
    Title: collection.Title(),
    Description: collection.Description(),
    Links: AddCollectionLinkList([]*viewmodels.Link{}, app, collection, templates),
  }
}

func AddCollectionsLinkList(app models.Application) []*viewmodels.Link {
  links := []*viewmodels.Link{}
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  templates := app.Templates("featurecollections", "")
  for _, template := range templates {
    link := &viewmodels.Link{
      Title: template.Title(),
      Rel: template.Rel(),
      Type: template.Type(),
      Href: template.Href(baseUrl, params),
    }

    links = append(links, link)
  }

  return links
}

func AddCollectionLinkList(links []*viewmodels.Link, app models.Application, collection features.Collection, templates []models.Handler) []*viewmodels.Link {
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  params["collection_id"] = collection.Id()

  for _, template := range templates {
    link := &viewmodels.Link{
      Title: template.Title(),
      Rel: template.Rel(),
      Type: template.Type(),
      Href: template.Href(baseUrl, params),
    }

    links = append(links, link)
  }

  return links
}
