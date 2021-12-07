package htmltemplates

import(
  "html/template"
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

// Transforms a renderlandingpage function into a renderlandingpage object
func NewFeatureRenderer() *FeatureRenderer {
  templates := NewTemplate([]string{
    "collections.html",
    "collection.html",
    "features.html",
    "feature.html",
  })

  return &FeatureRenderer{
    Templates: templates,
  }
}

// Internal
type FeatureRenderer struct {
  Templates *template.Template
}

func (renderer *FeatureRenderer) RenderCollections(context *models.Webcontext, collections *viewmodels.Collections) {
  renderer.Templates.ExecuteTemplate(context.W, "collections.html", collections)
}

func (renderer *FeatureRenderer) RenderCollection(context *models.Webcontext, collection *viewmodels.Collection) {
  renderer.Templates.ExecuteTemplate(context.W, "collection.html", collection)
}

func (renderer *FeatureRenderer) RenderItems(context *models.Webcontext, items *viewmodels.Features) {
  renderer.Templates.ExecuteTemplate(context.W, "features.html", items)
}

func (renderer *FeatureRenderer) RenderItem(context *models.Webcontext, item *viewmodels.Feature) {
  renderer.Templates.ExecuteTemplate(context.W, "feature.html", item)
}
