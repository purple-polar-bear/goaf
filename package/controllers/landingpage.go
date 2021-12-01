package apifcontrollers

import(
  "net/http"
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"
)

type LandingpageController struct {
}

func (controller *LandingpageController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderLandingpageType)

  return func(w http.ResponseWriter, r *http.Request, routeParameters models.MatchedRouteParameters) {
    links := controller.buildLinks(app)
    config := app.Config()

    resource := &viewmodels.Landingpage{
      Title: config.Title(),
      Description: config.Description(),
      Links: links,
    }

    renderer.RenderLandingpage(models.NewWebcontext(w, r), resource)
  }
}

func (controller *LandingpageController) buildLinks(app models.Application) []*viewmodels.Link {
  result := []*viewmodels.Link{}
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  for _, route := range app.Routes() {
    if !route.LandingpageVisible() {
      continue
    }

    for _, item := range route.Handlers() {
      link := &viewmodels.Link{
        Title: item.Title(),
        Rel: item.Rel(),
        Type: item.Type(),
        Href: item.Href(baseUrl, params),
      }

      result = append(result, link)
    }
  }

  return result
}
