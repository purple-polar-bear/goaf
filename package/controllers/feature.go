package apifcontrollers

import(
  "net/http"

  "oaf-server/package/models"
)

type FeatureController struct {

}

func (controller *FeatureController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
  }
}
