package jsontemplates

import(
  "oaf-server/package/models"
)

func RenderLandingpage(context *models.Webcontext, landingpage *models.Landingpage) {
  RenderPage(context, landingpage)
}
