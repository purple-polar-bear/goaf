package featurescontrollers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"oaf-server/package/core/services"
	"oaf-server/package/core/models"
	coreviewmodels "oaf-server/package/core/viewmodels"
	"oaf-server/package/features/models"
	"oaf-server/package/features/services"
	"oaf-server/package/features/templates/core"
	"oaf-server/package/features/viewmodels"
)

type FeaturesController struct {
}

func (controller *FeaturesController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    featuresRoute := app.Templates("features", "")
		r.ParseForm()
		urlValues := r.Form

    featureService, ok := app.GetService("features").(featureservices.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

		coreservice, ok := app.GetService("core").(coreservices.CoreService)
    if !ok {
      panic("Cannot find coreservice")
    }

    encoding := coreservice.ContentTypeUrlEncoder()
    featureParams := buildFeatureParams(app, routeParameters, urlValues)
    features := featureService.Features(r, featureParams)
    links := BuildFeaturesLinks(handler, app, encoding, featuresRoute, featureParams, features, featureParams)

    resource := viewmodels.NewFeatureCollection()
    items := features.Items()
    itemLength := len(items)
    resource.Features = make([]interface{}, itemLength)
    for index, item := range items {
      resource.Features[index] = item
    }

    resource.Links = links
    resource.NumberReturned = itemLength
    renderer.RenderItems(coremodels.NewWebcontext(w, r), resource)
  }
}

func buildFeatureParams(app coremodels.Application, routeParameters coremodels.MatchedRouteParameters, urlValues url.Values) *featuremodels.FeaturesParams {
  params := featuremodels.NewFeaturesParams()
  params.CollectionId = routeParameters.Get("collection_id")
	params.Offset = ConvertStringToIntegerWithDefault(urlValues.Get("offset"), 0)
	params.Limit = ConvertStringToIntegerWithDefault(urlValues.Get("limit"), 100)
  return params
}

func BuildFeaturesLinks(handler coremodels.Handler, app coremodels.Application, encoding *coremodels.ContentTypeUrlEncoding, templates []coremodels.Handler, params *featuremodels.FeaturesParams, features featuremodels.Features, featureParams *featuremodels.FeaturesParams) []*coreviewmodels.Link {
	baseUrl := app.Config().FullUri()
	hrefParams := make(map[string]string)
  hrefParams["collection_id"] = featureParams.CollectionId

  result := []*coreviewmodels.Link{}
	// current link
  for _, template := range templates {
		baseHref := template.Href(baseUrl, hrefParams, encoding)
    link := &coreviewmodels.Link{
      Title: template.Title(),
      Rel: template.Rel(handler.Type()),
      Type: template.Type(),
      Href: BuildFeaturesUrl(baseHref, featureParams.Limit, featureParams.Offset),
    }

		result = append(result, link)
	}

	// next link (if applicable)
	if features.HasNext() {
		for _, template := range templates {
			baseHref := template.Href(baseUrl, hrefParams, encoding)
	    link := &coreviewmodels.Link{
	      Title: template.Title(),
	      Rel: template.Rel(handler.Type()),
	      Type: template.Type(),
	      Href: BuildFeaturesUrl(baseHref, features.NextLimit(), features.NextOffset()),
	    }

			result = append(result, link)
		}
	}

  return result
}

//
// Utility functions
//

func ConvertStringToIntegerWithDefault(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return result
}

func BuildFeaturesUrl(baseUrl string, limit int, offset int) string {
	hasQuery := strings.Contains(baseUrl, "?")
	delimiter := "?"
	if hasQuery {
		delimiter = "&"
	}
	return baseUrl + delimiter + "offset=" + strconv.Itoa(limit) + "&limit=" + strconv.Itoa(offset)
}
