package coretemplates

import(
  "oaf-server/package/models"
  "oaf-server/package/viewmodels"
)

// Function definition for landingpage classes
type RenderLandingpageFunc func (*models.Webcontext, *viewmodels.Landingpage)

// Interface definition for landingpage classes
type RenderLandingpageType interface {
  RenderLandingpage(context *models.Webcontext, landingpage *viewmodels.Landingpage)
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

func (object *renderLandingpageType) RenderLandingpage(context *models.Webcontext, landingpageClasses *viewmodels.Landingpage) {
  object.renderLandingpageFunc(context, landingpageClasses)
}
