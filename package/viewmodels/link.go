package viewmodels

type Link struct {
  Href string `json:"href"`
	Hreflang string `json:"hreflang,omitempty"`
	Length int64 `json:"length,omitempty"`
	Rel string `json:"rel,omitempty"`
	Title string `json:"title,omitempty"`
	Type string `json:"type,omitempty"`
}
