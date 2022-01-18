package coremodels

type BaseController interface {
  HandleFunc(Application, interface{}) ControllerFunc
}
