package apifcontrollers

import(
  "net/http"

  "oaf-server/package/models"
)

type CollectionsController struct {

}

func (controller *CollectionsController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
  }
}
