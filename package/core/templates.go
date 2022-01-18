package apicore

import(
  "oaf-server/package/core/models"
  "oaf-server/package/core/templates/json"
  "oaf-server/package/core/templates/html"
)

type Templates struct {
  ContentTypeQueryParameter string
  Templates []*Template
}

type Template struct {
  Route Route
  Type string
  Title string
  HandleFunc coremodels.ControllerFunc
}

func AddBaseJSONTemplates(engine Engine) {
  renderer := jsontemplates.NewCoreRenderer()
  engine.AddTemplate("landingpage", "this landing page in json format", "application/json", "self", renderer)
  engine.AddTemplate("conformance", "conformance capabilities in json format", "application/json", "conformance", renderer)
  engine.AddTemplate("api", "API capabilities in json format", "application/vnd.oai.openapi+json;version=3.0", "service-desc", renderer)
}

// Shortcut functions to add HTML responses to base endpoints
func AddBaseHTMLTemplates(engine Engine) {
  renderer := htmltemplates.NewCoreRenderer()
  engine.AddTemplate("landingpage", "alternate landing page in html format", "text/html", "self", renderer)
  engine.AddTemplate("conformance", "conformance capabilities in html format", "text/html", "conformance", renderer)
  engine.AddTemplate("api", "API capabilities in html format", "text/html", "service-doc", renderer)
}

// Add a template to the engine
func (engine *engine) AddTemplate(name string, title string, contenttype string, rel string, renderer interface{}) {
  controller := engine.Controller(name)
  if controller == nil {
    panic("Cannot find controller: " + name)
  }

  handler := &Handler{
    title: title,
    rel: rel,
    controllerFunc: controller.HandleFunc(engine, renderer),
  }

  route := engine.router.Route(name)
  route.AddHandler(contenttype, handler)
}
