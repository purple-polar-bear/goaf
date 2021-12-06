package htmltemplates

import(
  "html/template"
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

// Transforms a renderlandingpage function into a renderlandingpage object
func NewCoreRenderer() *CoreRenderer {
  path := "package/templates/html/templates/"
  templates := template.Must(template.New("templates").Funcs(
    template.FuncMap{
      "isOdd":       func(i int) bool { return i%2 != 0 },
    },
  ).ParseFiles(
    path + "conformance.html",
    path + "landingpage.html",
  ))
  return &CoreRenderer{
    Templates: templates,
  }
}

// Internal
type CoreRenderer struct {
  Templates *template.Template
}

func (renderer *CoreRenderer) RenderLandingpage(context *models.Webcontext, landingpageClasses *viewmodels.Landingpage) {
  writer := context.W
  // buf := new(bytes.Buffer)
  renderer.Templates.ExecuteTemplate(writer, "landingpage.html", landingpageClasses)
  /*
  encodedContent = buf.Bytes()
  w.Header().Set("Content-Type", "text/html")
  w.WriteHeader(http.StatusOK)
  _, _ = w.Write(encodedContent)
  */
}

func (renderer *CoreRenderer) RenderConformance(context *models.Webcontext, conformanceClasses *viewmodels.Conformanceclasses) {
  writer := context.W
  renderer.Templates.ExecuteTemplate(writer, "conformance.html", conformanceClasses)
}
