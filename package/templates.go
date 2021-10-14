package apif

type Templates struct {
  ContentTypeQueryParameter string
  Templates []*Template
}

type Template struct {
  Route Route
  Type string
  Title string
  HandleFunc ControllerFunc
}

// Shortcut functions to add JSON responses to all endpoints
func AddDefaultJSONTemplates(engine Engine) {
  engine.AddConformanceTemplate("conformance capabilities in json format", "application/json", jsontemplates.RenderConformance)
}

// Add a conformance template to the engine
func (engine *engine) AddConformanceTemplate(title string, contentType string, renderFunc RenderConformanceFunc) {
  handler := func(w http.ResponseWriter, r *http.Request) {
    // TODO: add try..catch for 500 errors
    result := engine.conformanceController.Handle(w, r)
    renderFunc(w, r, result)
  }

  newTemplate := &Template{
    Route: engine.GetRoute("conformance"),
    Type: "application/json",
    Title: "Conformance capabilities in json format",
    HandleFunc: handler,
  }
  engine.Templates = append(engine.Templates, newTemplate)
}
