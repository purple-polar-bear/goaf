package coretemplates

import(
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

// Interface definition for features renderer
type RenderFeaturesType interface {
  RenderCollections(context *models.Webcontext, collections *viewmodels.Collections)
  RenderCollection(context *models.Webcontext, collection *viewmodels.Collection)
  RenderItems(context *models.Webcontext, items *viewmodels.Features)
  RenderItem(context *models.Webcontext, item *viewmodels.Feature)
}
