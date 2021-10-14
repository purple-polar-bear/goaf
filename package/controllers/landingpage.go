package apifcontrollers

import(
  "net/http"
  "oaf-server/package/models"
  "oaf-server/package/templates"
)

type LandingPageController struct {
}

func (controller *LandingPage) Handle(w http.ResponseWriter, r *http.Request, app models.Engine) *models.Landingpage{
  links := controller.buildLinks(app)

  resource := &models.Landingpage{
    Title: app.Title(),
    Description: app.Description(),
    Links: links,
  }

  return resource
}

func (controller *LandingPage) buildLinks(app models.Engine) []models.Link {
  result := []models.Link
  for _, item := app.Templates() {
    rel := determineRelation()
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
