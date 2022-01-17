package apifcontrollers

import(
  "net/http"
  "oaf-server/package/core"
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"
)

type LandingpageController struct {
}

func (controller *LandingpageController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderCoreType)

  return func(handler models.Handler, w http.ResponseWriter, r *http.Request, routeParameters models.MatchedRouteParameters) {
    links := controller.buildLinks(handler, app)
    config := app.Config()

    resource := &viewmodels.Landingpage{
      Title: config.Title(),
      Description: config.Description(),
      Links: links,
    }

    renderer.RenderLandingpage(models.NewWebcontext(w, r), resource)
  }
}

func (controller *LandingpageController) buildLinks(handler models.Handler, app models.Application) []*viewmodels.Link {
  result := []*viewmodels.Link{}
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  coreservice, ok := app.GetService("core").(apifcore.CoreService)
  if !ok {
    panic("Cannot find coreservice")
  }

  encoding := coreservice.ContentTypeUrlEncoder()
  for _, route := range app.Routes() {
    if !route.LandingpageVisible() {
      continue
    }

    for _, item := range route.Handlers() {
      link := &viewmodels.Link{
        Title: item.Title(),
        Rel: item.Rel(handler.Type()),
        Type: item.Type(),
        Href: item.Href(baseUrl, params, encoding),
      }

      result = append(result, link)
    }
  }

  return result
}
