package coretemplates

import(
  "oaf-server/package/core/models"
  "oaf-server/package/viewmodels"
)

// Interface definition for features renderer
type RenderFeaturesType interface {
  RenderCollections(context *coremodels.Webcontext, collections *viewmodels.Collections)
  RenderCollection(context *coremodels.Webcontext, collection *viewmodels.Collection)
  RenderItems(context *coremodels.Webcontext, items *viewmodels.Features)
  RenderItem(context *coremodels.Webcontext, item *viewmodels.Feature)
}
