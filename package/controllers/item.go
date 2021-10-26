package apifcontrollers

import(
  "net/http"

  "oaf-server/package/models"
)

type ItemController struct {

}

func (controller *ItemController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
  }
}
