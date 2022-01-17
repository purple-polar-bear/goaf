package htmltemplates

import(
  "html/template"
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

// Transforms a renderlandingpage function into a renderlandingpage object
func NewCoreRenderer() *CoreRenderer {
  templates := NewTemplate([]string{
    "conformance.html",
    "landingpage.html",
    "api.html",
  })

  return &CoreRenderer{
    Templates: templates,
  }
}

// Internal
type CoreRenderer struct {
  Templates *template.Template
}

func (renderer *CoreRenderer) RenderLandingpage(context *models.Webcontext, landingpageClasses *viewmodels.Landingpage) {
  renderer.Templates.ExecuteTemplate(context.W, "landingpage.html", landingpageClasses)
}

func (renderer *CoreRenderer) RenderConformance(context *models.Webcontext, conformanceClasses *viewmodels.Conformanceclasses) {
  renderer.Templates.ExecuteTemplate(context.W, "conformance.html", conformanceClasses)
}

func (renderer *CoreRenderer) RenderAPI(context *models.Webcontext, api interface{}) {
  renderer.Templates.ExecuteTemplate(context.W, "api.html", api)
}
