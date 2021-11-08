package models

type BaseController interface {
  HandleFunc(Application, interface{}) ControllerFunc
}
