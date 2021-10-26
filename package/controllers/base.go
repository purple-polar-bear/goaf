package apifcontrollers

import(
  "oaf-server/package/models"
)

type BaseController interface {
  HandleFunc(models.Application, interface{}) models.ControllerFunc
}
