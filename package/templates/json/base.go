package jsontemplates

import(
  "encoding/json"
  "net/http"
  "oaf-server/package/models"
)

func RenderPage(context templates.Context, resource interface{}) {
  writer := context.Writer
  encodedContent, err = json.Marshal(landingpage)
  if err != nil {
    jsonError(writer, "JSON MARSHALLER", err.Error(), http.StatusInternalServerError)
    return
  }

  writer.Header().Set("Content-Type", "application/json")
  writer.WriteHeader(http.StatusOK)
  _, _ = writer.Write(encodedContent)
}
