package coretemplates

import(
  "oaf-server/package/models"
)

// Function definition for landingpage classes
type RenderLandingpageFunc func (*models.Webcontext, *models.Landingpage)

// Interface definition for landingpage classes
type RenderLandingpageType interface {
  RenderLandingpage(context *models.Webcontext, landingpage *models.Landingpage)
}

// Transforms a renderlandingpage function into a renderlandingpage object
func NewRenderLandingpageType(fun RenderLandingpageFunc) RenderLandingpageType {
  return &renderLandingpageType{
    renderLandingpageFunc: fun,
  }
}

// Internal

type renderLandingpageType struct {
  renderLandingpageFunc RenderLandingpageFunc
}

func (object *renderLandingpageType) RenderLandingpage(context *models.Webcontext, landingpageClasses *models.Landingpage) {
  object.renderLandingpageFunc(context, landingpageClasses)
}
