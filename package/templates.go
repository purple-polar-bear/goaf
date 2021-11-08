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

func AddBaseJSONTemplates(engine Engine) {
  engine.AddTemplate("landingpage", "this landing page in json format", "application/json", "self", coretemplates.NewRenderLandingpageType(jsontemplates.RenderLandingpage))
  engine.AddTemplate("conformance", "conformance capabilities in json format", "application/json", "conformance", coretemplates.NewRenderConformanceType(jsontemplates.RenderConformance))
}

// Shortcut functions to add JSON responses to all endpoints
func AddFeaturesJSONTemplates(engine Engine) {
  renderer := jsontemplates.NewFeatureRenderer()
  engine.AddTemplate("featurecollections", "data collections in json format", "application/json", "data", renderer)
  engine.AddTemplate("featurecollection", "data collection in json format", "application/json", "data", renderer)
  engine.AddTemplate("features", "data items in json format", "application/json", "data", renderer)
  engine.AddTemplate("feature", "data item in json format", "application/json", "data", renderer)
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
