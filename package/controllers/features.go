package apifcontrollers

import(
  "net/http"

  "oaf-server/package/models"
)

type FeaturesController struct {

}

func (controller *FeaturesController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
  }
}
