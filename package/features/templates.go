package apifeatures

import(
  "oaf-server/package/core"
  "oaf-server/package/features/templates/json"
  "oaf-server/package/features/templates/html"
)

// Shortcut functions to add JSON responses to all endpoints
func AddFeaturesJSONTemplates(engine apicore.Engine) {
  renderer := jsontemplates.NewFeatureRenderer()
  engine.AddTemplate("featurecollections", "data collections in json format", "application/json", "data", renderer)
  engine.AddTemplate("featurecollection", "data collection in json format", "application/json", "data", renderer)
  engine.AddTemplate("features", "data items in json format", "application/json", "data", renderer)
  engine.AddTemplate("feature", "data item in json format", "application/json", "data", renderer)
}

func AddFeaturesHTMLTemplates(engine apicore.Engine) {
  renderer := htmltemplates.NewFeatureRenderer()
  engine.AddTemplate("featurecollections", "data collections in html format", "text/html", "data", renderer)
  engine.AddTemplate("featurecollection", "data collection in html format", "text/html", "data", renderer)
  engine.AddTemplate("features", "data items in html format", "text/html", "data", renderer)
  engine.AddTemplate("feature", "data item in html format", "text/html", "data", renderer)
}
