package apifcore

import(
  "fmt"

  "github.com/getkin/kin-openapi/openapi3"
)

type OpenAPIBuilder interface {
  OpenAPI() *openapi3.T
  Build([]interface{}) *openapi3.T
  AddComponentParameter(parameterName string, parameter *openapi3.Parameter)
}

type openAPIBuilder struct {
  Loader *openapi3.Loader
  BaseTemplate *openapi3.T
  T *openapi3.T
}

func NewOpenAPIBuilder(basetemplate *openapi3.T, loader *openapi3.Loader) OpenAPIBuilder {
  return &openAPIBuilder{
    BaseTemplate: basetemplate,
    Loader: loader,
  }
}

func (builder *openAPIBuilder) OpenAPI() *openapi3.T {
  return builder.T
}

func (builder *openAPIBuilder) AddComponentParameter(parameterName string, parameter *openapi3.Parameter) {
  parameterRef := &openapi3.ParameterRef{Value: parameter}
  builder.T.Components.Parameters[parameterName] = parameterRef
}

func (builder *openAPIBuilder) Build(services []interface{}) *openapi3.T {
  builder.duplicateTemplate()
  for _, service := range services {
    openapiservice := service.(OpenAPIService)
    openapiservice.BuildOpenAPISpecification(builder)
  }

  // Validating the Swagger file
  err := builder.T.Validate(builder.Loader.Context)
  if err != nil {
    // panic(fmt.Sprintf("Could not validate Swagger definition: %s", err))
  }

  return builder.T
}

//
// Internal API
//

func (builder *openAPIBuilder) duplicateTemplate() {
  api := builder.BaseTemplate
  builder.T = &openapi3.T{}
  data, err := api.MarshalJSON()
  if err != nil {
    panic(fmt.Sprintf("Error in base OpenAPI file: %s", err))
  }

  builder.T.UnmarshalJSON(data)
}
