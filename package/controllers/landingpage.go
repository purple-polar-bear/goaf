package apifcontrollers

import(
  "net/http"
)

type LandingPage struct {

}

func (controller *LandingPage) Handle(w http.ResponseWriter, request *http.Request) {
  w.Write([]byte("Test"))
}
