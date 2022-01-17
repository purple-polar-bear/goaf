package models

type ConfigurableService interface {
  SetConfig(Serverconfig)
}
