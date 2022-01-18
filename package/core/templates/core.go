package coretemplates

import(
  "oaf-server/package/core/models"
  "oaf-server/package/core/viewmodels"
)

// Interface definition for features renderer
type RenderCoreType interface {
  RenderLandingpage(context *coremodels.Webcontext, landingpage *viewmodels.Landingpage)
  RenderConformance(context *coremodels.Webcontext, conformanceClasses *viewmodels.Conformanceclasses)
  RenderAPI(context *coremodels.Webcontext, api interface{})
}
