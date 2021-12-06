package coretemplates

import(
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

// Interface definition for features renderer
type RenderCoreType interface {
  RenderCollections(context *models.Webcontext, collections *viewmodels.Collections)
  RenderCollection(context *models.Webcontext, collection *viewmodels.Collection)
}
