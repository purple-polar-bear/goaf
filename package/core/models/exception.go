package coremodels

type Exception struct {
  Code string `json:"code"`
	Description string `json:"description,omitempty"`
}
