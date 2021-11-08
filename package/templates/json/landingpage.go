package jsontemplates

import(
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

func RenderLandingpage(context *models.Webcontext, landingpage *viewmodels.Landingpage) {
  RenderPage(context, landingpage)
}
