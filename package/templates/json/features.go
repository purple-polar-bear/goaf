package jsontemplates

import(
  "oaf-server/package/core/models"
  "oaf-server/package/core/templates/json"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"
)

type FeatureRenderer struct {
}

func NewFeatureRenderer() coretemplates.RenderFeaturesType {
  return &FeatureRenderer{}
}

func (renderer *FeatureRenderer) RenderCollections(context *coremodels.Webcontext, collections *viewmodels.Collections) {
  jsontemplates.RenderPage(context, collections)
}

func (renderer *FeatureRenderer) RenderCollection(context *coremodels.Webcontext, collection *viewmodels.Collection) {
  jsontemplates.RenderPage(context, collection)
}

func (renderer *FeatureRenderer) RenderItems(context *coremodels.Webcontext, items *viewmodels.Features) {
  jsontemplates.RenderPage(context, items)
}

func (renderer *FeatureRenderer) RenderItem(context *coremodels.Webcontext, item *viewmodels.Feature) {
  jsontemplates.RenderPage(context, item)
}
