package jsontemplates

import(
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"
)

type FeatureRenderer struct {
}

func NewFeatureRenderer() coretemplates.RenderFeaturesType {
  return &FeatureRenderer{}
}

func (renderer *FeatureRenderer) RenderCollections(context *models.Webcontext, collections *viewmodels.Collections) {
  RenderPage(context, collections)
}

func (renderer *FeatureRenderer) RenderCollection(context *models.Webcontext, collection *viewmodels.Collection) {
  RenderPage(context, collection)
}

func (renderer *FeatureRenderer) RenderItems(context *models.Webcontext, items *viewmodels.Features) {
  RenderPage(context, items)
}

func (renderer *FeatureRenderer) RenderItem(context *models.Webcontext, item *viewmodels.Feature) {
  RenderPage(context, item)
}
