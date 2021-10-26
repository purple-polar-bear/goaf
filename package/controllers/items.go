package apifcontrollers

import(
  "net/http"

  "oaf-server/package/models"
)

type ItemsController struct {

}

func (controller *ItemsController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
  }
}
