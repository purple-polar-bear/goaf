package viewmodels

type Landingpage struct {
  Title string `json:"title"`
  Description string `json:"description"`
  Links []*Link `json:"links"`
  License string
  LicenseName string
}
