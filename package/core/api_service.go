package apifcore

import(
  "fmt"

  "github.com/getkin/kin-openapi/openapi3"
)

type CoreService interface {
  OpenAPI() *openapi3.T
  SetContact(*ContactInfo)
  SetLicense(*LicenseInfo)
  SetContentTypeUrlEncoder(*ContentTypeUrlEncoding)
  ContentTypeUrlEncoder() *ContentTypeUrlEncoding

  // Method to rebuild the OpenAPI specification. It is recommended to call
  // this method, in order to easy error handling.
  RebuildOpenAPI()
}

type ContactInfo struct {
  Name string
  Url string
}

type LicenseInfo struct {
  Name string
  Url string
}

type ContentTypeUrlEncoding struct {
  ParameterName string
  OverrideHeader bool
  Encodings map[string]string
}

type coreService struct {
  ContactInfo *ContactInfo
  LicenseInfo *LicenseInfo
  ContentTypeUrlEncoding *ContentTypeUrlEncoding

  loader *openapi3.Loader
  basicOpenAPI *openapi3.T
  openAPI *openapi3.T
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

func (service *coreService) SetContentTypeUrlEncoder(contentTypeUrlEncoding *ContentTypeUrlEncoding) {
  service.ContentTypeUrlEncoding = contentTypeUrlEncoding
}

func (service *coreService) ContentTypeUrlEncoder() *ContentTypeUrlEncoding {
  return service.ContentTypeUrlEncoding
}

func (service *coreService) OpenAPI() *openapi3.T {
  return service.openAPI
}

func (service *coreService) RebuildOpenAPI()() {
  api := service.basicOpenAPI
  copy := &openapi3.T{}
  data, err := api.MarshalJSON()
  if err != nil {
    panic(fmt.Sprintf("Error in base OpenAPI file: %s", err))
  }

  copy.UnmarshalJSON(data)
  if service.ContactInfo != nil {
    copy.Info.Contact.Name = service.ContactInfo.Name
    copy.Info.Contact.URL = service.ContactInfo.Url
  }

  if service.LicenseInfo != nil {
    copy.Info.License.Name = service.LicenseInfo.Name
    copy.Info.License.URL = service.LicenseInfo.Url
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
    copy.Components.Parameters[encoder.ParameterName] = parameterRef
    pathParameterRef := &openapi3.ParameterRef{Ref: "#/components/parameters/" + encoder.ParameterName, Value: parameter}
    for _, path := range copy.Paths {
      for _, operation := range path.Operations() {
        operation.Parameters = append(operation.Parameters, pathParameterRef)
      }
    }
  }

  // Add stuff from other services
  featureServiceBuildOpenAPI(copy)

  // Validating the Swagger file
  // err = copy.Validate(service.loader.Context)
  if err != nil {
    panic(fmt.Sprintf("Could not validate Swagger definition: %s", err))
  }

  service.openAPI = copy
}

func featureServiceBuildOpenAPI(openapi *openapi3.T) {
  parameterName := "featureId"
  parameter := openapi3.NewPathParameter(parameterName).
    WithDescription("local identifier of a feature").
    WithRequired(true).
    WithSchema(openapi3.NewStringSchema())
  AddComponentParameter(openapi, parameterName, parameter)

  parameterName = "collectionId"
  parameter = openapi3.NewPathParameter(parameterName).
    WithDescription("local identifier of a collection").
    WithRequired(true).
    WithSchema(openapi3.NewStringSchema())
  AddComponentParameter(openapi, parameterName, parameter)
}

func AddComponentParameter(openapi *openapi3.T, parameterName string, parameter *openapi3.Parameter) {
  parameterRef := &openapi3.ParameterRef{Value: parameter}
  openapi.Components.Parameters[parameterName] = parameterRef
}

func NewContentTypeUrlEncoding(parameterName string) *ContentTypeUrlEncoding {
  return &ContentTypeUrlEncoding{
    ParameterName: parameterName,
    OverrideHeader: true,
    Encodings: make(map[string]string),
  }
}

func (encoder *ContentTypeUrlEncoding) AddContentType(parameterValue string, contentType string) {
  encoder.Encodings[parameterValue] = contentType
}
