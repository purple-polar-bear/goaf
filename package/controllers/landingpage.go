package apifcontrollers

import(
  "net/http"
  "oaf-server/package/models"
  "oaf-server/package/templates/core"
)

type LandingpageController struct {
}

func (controller *LandingpageController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderLandingpageType)

  return func(w http.ResponseWriter, r *http.Request) {
    links := controller.buildLinks(app)

    resource := &models.Landingpage{
      Title: app.Title(),
      Description: app.Description(),
      Links: links,
    }

    renderer.RenderLandingpage(nil, resource)
  }
}

func (controller *LandingpageController) buildLinks(app models.Application) []*models.Link {
  result := []*models.Link{}
  for _, item := range app.Templates("", "") {
    rel := item.CalculateRelation()
    if rel == "" {
      continue
    }

    link := &models.Link{
      Rel: rel,
      Title: item.Title,
      Type: item.Type,
      Href: item.Href(),
    }
    result = append(result, link)
  }

  return result
}
