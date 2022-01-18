package jsontemplates

import(
  "oaf-server/package/core/models"
  "oaf-server/package/core/viewmodels"
  "oaf-server/package/core/templates"
)

type CoreRenderer struct {
}

func NewCoreRenderer() coretemplates.RenderCoreType {
  return &CoreRenderer{}
}

func (renderer *CoreRenderer) RenderLandingpage(context *coremodels.Webcontext, landingpage *viewmodels.Landingpage) {
  RenderPage(context, landingpage)
}

func (renderer *CoreRenderer) RenderConformance(context *coremodels.Webcontext, conformanceClasses *viewmodels.Conformanceclasses) {
  RenderPage(context, conformanceClasses)
}

func (renderer *CoreRenderer) RenderAPI(context *coremodels.Webcontext, api interface{}) {
  RenderPage(context, api)
}
