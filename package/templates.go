package apif

import(
  "oaf-server/package/models"
  "oaf-server/package/templates/json"
  "oaf-server/package/templates/html"
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

// Shortcut functions to add JSON responses to all endpoints
func AddFeaturesJSONTemplates(engine Engine) {
  renderer := jsontemplates.NewFeatureRenderer()
  engine.AddTemplate("featurecollections", "data collections in json format", "application/json", "data", renderer)
  engine.AddTemplate("featurecollection", "data collection in json format", "application/json", "data", renderer)
  engine.AddTemplate("features", "data items in json format", "application/json", "data", renderer)
  engine.AddTemplate("feature", "data item in json format", "application/json", "data", renderer)
}

func AddFeaturesHTMLTemplates(engine Engine) {
  renderer := htmltemplates.NewFeatureRenderer()
  engine.AddTemplate("featurecollections", "data collections in html format", "text/html", "data", renderer)
  engine.AddTemplate("featurecollection", "data collection in html format", "text/html", "data", renderer)
  engine.AddTemplate("features", "data items in html format", "text/html", "data", renderer)
  engine.AddTemplate("feature", "data item in html format", "text/html", "data", renderer)
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
