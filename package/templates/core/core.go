package coretemplates

import(
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

// Interface definition for features renderer
type RenderCoreType interface {
  RenderLandingpage(context *models.Webcontext, landingpage *viewmodels.Landingpage)
  RenderConformance(context *models.Webcontext, conformanceClasses *viewmodels.Conformanceclasses)
  RenderAPI(context *models.Webcontext, api interface{})
}
