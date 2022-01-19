package coreservices

import(
  "oaf-server/package/core/models"
  "github.com/getkin/kin-openapi/openapi3"
)

type CoreService interface {
  OpenAPI() *openapi3.T
  SetContact(*ContactInfo)
  SetLicense(*LicenseInfo)
  SetContentTypeUrlEncoder(*coremodels.ContentTypeUrlEncoding)
  ContentTypeUrlEncoder() *coremodels.ContentTypeUrlEncoding

  // Method to rebuild the OpenAPI specification. It is recommended to call
  // this method, in order to easy error handling.
  RebuildOpenAPI([]interface{})
}

type OpenAPIService interface {
  BuildOpenAPISpecification(OpenAPIBuilder)
}

type ContactInfo struct {
  Name string
  Url string
}

type LicenseInfo struct {
  Name string
  Url string
}

type coreService struct {
  ContactInfo *ContactInfo
  LicenseInfo *LicenseInfo
  ContentTypeUrlEncoding *coremodels.ContentTypeUrlEncoding

  loader *openapi3.Loader
  basicOpenAPI *openapi3.T
  openAPI *openapi3.T
  config coremodels.Serverconfig
}

func NewCoreService() CoreService {
  loader := openapi3.NewLoader()
  loader.IsExternalRefsAllowed = true

  serviceSpecPath := "./package/core/oaf.yml"
  openapi, err := loader.LoadFromFile(serviceSpecPath)
  if err != nil {
    panic("Cannot Loadswagger from file :" + serviceSpecPath)
  }

  return &coreService{
    loader: loader,
    basicOpenAPI: openapi,
    openAPI: openapi,
  }
}

func (service *coreService) SetContact(info *ContactInfo) {
  service.ContactInfo = info
}

func (service *coreService) SetLicense(info *LicenseInfo) {
  service.LicenseInfo = info
}

func (service *coreService) SetContentTypeUrlEncoder(contentTypeUrlEncoding *coremodels.ContentTypeUrlEncoding) {
  service.ContentTypeUrlEncoding = contentTypeUrlEncoding
}

func (service *coreService) ContentTypeUrlEncoder() *coremodels.ContentTypeUrlEncoding {
  return service.ContentTypeUrlEncoding
}

func (service *coreService) OpenAPI() *openapi3.T {
  return service.openAPI
}

func (service *coreService) RebuildOpenAPI(services []interface{})() {
  builder := NewOpenAPIBuilder(service.basicOpenAPI, service.loader, service.config)
  service.openAPI = builder.Build(services)
}

func (service *coreService) BuildOpenAPISpecification(builder OpenAPIBuilder) {
  target := builder.OpenAPI()
  if service.ContactInfo != nil {
    target.Info.Contact.Name = service.ContactInfo.Name
    target.Info.Contact.URL = service.ContactInfo.Url
  }

  if service.LicenseInfo != nil {
    target.Info.License.Name = service.LicenseInfo.Name
    target.Info.License.URL = service.LicenseInfo.Url
  }

  // ContentType encoding via url parameter
  encoder := service.ContentTypeUrlEncoding
  if encoder != nil {
    parameter := openapi3.NewQueryParameter(encoder.ParameterName).
      WithDescription("The optional format parameter determines the outputformat, default json.").
      WithRequired(false).
      WithSchema(openapi3.NewStringSchema())
    parameter.Style = openapi3.SerializationForm
    explode := false
    parameter.Explode = &explode
    parameterRef := &openapi3.ParameterRef{Value: parameter}
    target.Components.Parameters[encoder.ParameterName] = parameterRef
    pathParameterRef := &openapi3.ParameterRef{Ref: "#/components/parameters/" + encoder.ParameterName, Value: parameter}
    for _, path := range target.Paths {
      for _, operation := range path.Operations() {
        operation.Parameters = append(operation.Parameters, pathParameterRef)
      }
    }
  }

  // add server
  server := &openapi3.Server{
    URL: service.config.FullHost(),
    Description: "Production server",
  }
  target.Servers = append(target.Servers, server)
}

// Implement ConfigurableService
func (service *coreService) SetConfig(config coremodels.Serverconfig) {
  service.config = config
}


func featureServiceBuildOpenAPI(builder OpenAPIBuilder) {
  parameterName := "featureId"
  parameter := openapi3.NewPathParameter(parameterName).
    WithDescription("local identifier of a feature").
    WithRequired(true).
    WithSchema(openapi3.NewStringSchema())
  parameter.Style = openapi3.SerializationForm
  builder.AddComponentParameter(parameterName, parameter)

  parameterName = "collectionId"
  parameter = openapi3.NewPathParameter(parameterName).
    WithDescription("local identifier of a collection").
    WithRequired(true).
    WithSchema(openapi3.NewStringSchema())
  parameter.Style = openapi3.SerializationForm
  builder.AddComponentParameter(parameterName, parameter)
}
