package coremodels

// The ContentTypeUrlEncoding is the encoding of the content type (or format)
// through the url parameters. This option is always optional.
//
// Example url: '/api?format=json'
// will be resolved by a ContentTypeUrlEncoding with the following properties
// ParameterName = "format"
// Encodings["json"] = "application/json"
//
// The OverrideHeader determines if the url encoding must be parsed (if possible)
// before the Accept headers of the request are parsed, or after. Default is true.
type ContentTypeUrlEncoding struct {
  ParameterName string
  OverrideHeader bool
  Encodings map[string][]string
  ReverseEncodings map[string]string
}

// Creates a new instance of ContentTypeUrlEncoding with the given parameterName
func NewContentTypeUrlEncoding(parameterName string) *ContentTypeUrlEncoding {
  return &ContentTypeUrlEncoding{
    ParameterName: parameterName,
    OverrideHeader: true,
    Encodings: make(map[string][]string),
    ReverseEncodings: make(map[string]string),
  }
}

func (encoder *ContentTypeUrlEncoding) AddContentType(parameterValue string, contentType string) {
  encoder.Encodings[parameterValue] = append(encoder.Encodings[parameterValue], contentType)
  encoder.ReverseEncodings[contentType] = parameterValue
}
