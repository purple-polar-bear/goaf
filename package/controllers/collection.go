package apifcontrollers

import(
  "net/http"

  "oaf-server/package/models"
)

type CollectionController struct {

}

func (controller *CollectionController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
  }
}
