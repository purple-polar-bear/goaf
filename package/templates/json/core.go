package jsontemplates

import(
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
  "oaf-server/package/templates/core"
)

type CoreRenderer struct {
}

func NewCoreRenderer() coretemplates.RenderCoreType {
  return &CoreRenderer{}
}

func (renderer *CoreRenderer) RenderLandingpage(context *models.Webcontext, landingpage *viewmodels.Landingpage) {
  RenderPage(context, landingpage)
}

func (renderer *CoreRenderer) RenderConformance(context *models.Webcontext, conformanceClasses *viewmodels.Conformanceclasses) {
  RenderPage(context, conformanceClasses)
}

func (renderer *CoreRenderer) RenderAPI(context *models.Webcontext, api interface{}) {
  RenderPage(context, api)
}
