package featureservices

import (
  "oaf-server/package/core/services"
  "github.com/getkin/kin-openapi/openapi3"
)

type FeatureServiceOpenAPIBuilder interface {
  AddBBox()
  AddDatetime()
  AddLimit()
  AddOffset()
}

type featureServiceOpenAPIBuilder struct {
  Builder coreservices.OpenAPIBuilder
}

func NewFeatureServiceOpenAPIBuilder(builder coreservices.OpenAPIBuilder) FeatureServiceOpenAPIBuilder {
  return &featureServiceOpenAPIBuilder{
    Builder: builder,
  }
}

func (service *featureServiceOpenAPIBuilder) AddBBox() {
  schema := openapi3.NewOneOfSchema(
    openapi3.NewFloat64Schema().WithMaxItems(4).WithMinItems(4),
    openapi3.NewFloat64Schema().WithMaxItems(6).WithMinItems(6),
  )
  schema.Type = "array"
  parameter := openapi3.NewQueryParameter("bbox").
    WithDescription("Only features that have a geometry that intersects the bounding box are selected.\nThe bounding box is provided as four or six numbers, depending on whether the\ncoordinate reference system includes a vertical axis (height or depth):\n\n* Lower left corner, coordinate axis 1\n* Lower left corner, coordinate axis 2\n* Minimum value, coordinate axis 3 (optional)\n* Upper right corner, coordinate axis 1\n* Upper right corner, coordinate axis 2\n* Maximum value, coordinate axis 3 (optional)\n\nThe coordinate reference system of the values is WGS 84 longitude/latitude\n(http://www.opengis.net/def/crs/OGC/1.3/CRS84) unless a different coordinate\nreference system is specified in the parameter `bbox-crs`.\n\nFor WGS 84 longitude/latitude the values are in most cases the sequence of\nminimum longitude, minimum latitude, maximum longitude and maximum latitude.\nHowever, in cases where the box spans the antimeridian the first value\n(west-most box edge) is larger than the third value (east-most box edge).\n\nIf the vertical axis is included, the third and the sixth number are\nthe bottom and the top of the 3-dimensional bounding box.\n\nIf a feature has multiple spatial geometry properties, it is the decision of the\nserver whether only a single spatial geometry property is used to determine\nthe extent or all relevant geometries.").
    WithRequired(false).
    WithSchema(schema)
  parameter.Style = openapi3.SerializationForm
  explode := false
  parameter.Explode = &explode
  service.Builder.AddComponentParameter("bbox", parameter)
}

func (service *featureServiceOpenAPIBuilder) AddDatetime() {
  parameter := openapi3.NewQueryParameter("datetime").
    WithDescription("Either a date-time or an interval, open or closed. Date and time expressions\nadhere to RFC 3339. Open intervals are expressed using double-dots.\n\nExamples:\n\n* A date-time: \"2018-02-12T23:20:50Z\"\n* A closed interval: \"2018-02-12T00:00:00Z/2018-03-18T12:31:12Z\"\n* Open intervals: \"2018-02-12T00:00:00Z/..\" or \"../2018-03-18T12:31:12Z\"\n\nOnly features that have a temporal property that intersects the value of\n`datetime` are selected.\n\nIf a feature has multiple temporal properties, it is the decision of the\nserver whether only a single temporal property is used to determine\nthe extent or all relevant temporal properties.").
    WithRequired(false).
    WithSchema(openapi3.NewStringSchema())
  parameter.Style = openapi3.SerializationForm
  explode := false
  parameter.Explode = &explode
  service.Builder.AddComponentParameter("datetime", parameter)
}

func (builder *featureServiceOpenAPIBuilder) AddLimit() {

}

func (builder *featureServiceOpenAPIBuilder) AddOffset() {

}
