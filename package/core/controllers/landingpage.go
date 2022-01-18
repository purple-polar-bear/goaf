package corecontrollers

import(
  "net/http"
  "oaf-server/package/core/services"
  "oaf-server/package/core/models"
  "oaf-server/package/core/viewmodels"
  "oaf-server/package/core/templates"
)

type LandingpageController struct {
}

func (controller *LandingpageController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderCoreType)

  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    links := controller.buildLinks(handler, app)
    config := app.Config()

    resource := &viewmodels.Landingpage{
      Title: config.Title(),
      Description: config.Description(),
      Links: links,
    }

    renderer.RenderLandingpage(coremodels.NewWebcontext(w, r), resource)
  }
}

func (controller *LandingpageController) buildLinks(handler coremodels.Handler, app coremodels.Application) []*viewmodels.Link {
  result := []*viewmodels.Link{}
  baseUrl := app.Config().FullUri()
  params := make(map[string]string)
  coreservice, ok := app.GetService("core").(coreservices.CoreService)
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
