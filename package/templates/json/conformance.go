package jsontemplates

import(
  "oaf-server/package/models"
)

func RenderConformance(context *models.Webcontext, conformanceClasses *models.Conformanceclasses) {
  RenderPage(context, conformanceClasses)
}
