package jsontemplates

import(
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

func RenderConformance(context *models.Webcontext, conformanceClasses *viewmodels.Conformanceclasses) {
  RenderPage(context, conformanceClasses)
}
