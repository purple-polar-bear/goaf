package jsontemplates

import(
  "encoding/json"
  "fmt"
  "net/http"
  "oaf-server/package/core/models"
)

func RenderPage(context *coremodels.Webcontext, resource interface{}) {
  writer := context.W
  encodedContent, err := json.Marshal(resource)
  if err != nil {
    jsonError(writer, "JSON MARSHALLER", err.Error(), http.StatusInternalServerError)
    return
  }

  writer.Header().Set("Content-Type", "application/json")
  writer.WriteHeader(http.StatusOK)
  _, _ = writer.Write(encodedContent)
}

func jsonError(w http.ResponseWriter, code string, msg string, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	result, err := json.Marshal(&coremodels.Exception{
		Code:        code,
		Description: msg,
	})

	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("problem marshaling error: %v", msg)))
	} else {
		_, _ = w.Write(result)
	}
}
