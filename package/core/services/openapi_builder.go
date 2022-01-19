package coreservices

import(
  "fmt"

  "oaf-server/package/core/models"

  "github.com/getkin/kin-openapi/openapi3"
)

type OpenAPIBuilder interface {
  OpenAPI() *openapi3.T
  Build([]interface{}) *openapi3.T
  AddComponentParameter(parameterName string, parameter *openapi3.Parameter)
}

type openAPIBuilder struct {
  // OpenAPI loader object
  Loader *openapi3.Loader

  // Base template used for quick building of the final template
  BaseTemplate *openapi3.T

  // Server configuration
  Config coremodels.Serverconfig

  // Result template
  T *openapi3.T
}

func NewOpenAPIBuilder(basetemplate *openapi3.T, loader *openapi3.Loader, config coremodels.Serverconfig) OpenAPIBuilder {
  return &openAPIBuilder{
    BaseTemplate: basetemplate,
    Loader: loader,
    Config: config,
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
  // Change the paths
  replacedPaths := make(map[string]*openapi3.PathItem)
  for path, item := range builder.T.Paths {
    replacedPaths[builder.Config.Mountingpath() + path] = item
  }
  builder.T.Paths = replacedPaths

  // Iterate through all the services in order to add specifications
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
