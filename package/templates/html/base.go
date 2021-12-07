package htmltemplates

import(
  "html/template"
  "strings"
)

func NewTemplate(pages []string) *template.Template {
  path := "package/templates/html/templates/"
  pagesWithPath := []string{}
  for _, page := range pages {
    pagesWithPath = append(pagesWithPath, path + page)
  }

  return template.Must(template.New("templates").Funcs(
   template.FuncMap{
     "isOdd":       func(i int) bool { return i%2 != 0 },
     "titleize":  func(title string) string { return strings.Title(title) },
   },
  ).ParseFiles(pagesWithPath...))
}
