package apif

import(
  "oaf-server/package/models"
  "oaf-server/package/templates/core"
  "oaf-server/package/templates/json"
)

type Templates struct {
  ContentTypeQueryParameter string
  Templates []*Template
}

type Template struct {
  Route Route
  Type string
  Title string
  HandleFunc models.ControllerFunc
}

// Shortcut functions to add JSON responses to all endpoints
func AddDefaultJSONTemplates(engine Engine) {
  engine.AddTemplate("conformance", "conformance capabilities in json format", "application/json", coretemplates.NewRenderConformanceType(jsontemplates.RenderConformance))
}

// Add a template to the engine
func (engine *engine) AddTemplate(name string, title string, contenttype string, renderer interface{}) {
  handler := engine.Controller(name).HandleFunc(engine, renderer)
  route := engine.router.Route(name)
  route.AddRoute(contenttype, handler)
}
