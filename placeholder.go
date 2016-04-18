package placeholder

import (
  "fmt"
  "strings"
  "strconv"
  "path/filepath"
  "html/template"
  "net/http"
)

type Placeholder struct {
  Width  int
  Height int
  Fill   string
  Stroke string
}

var funcMap = template.FuncMap {
  "minus": func (a, b int) int {
    return a - b
  },
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi There")
}

// Handler for URL paths /svg/WIDTH/HEIGHT/[FILL/STROKE]
func svg(w http.ResponseWriter, r *http.Request) {
  url := strings.Split(r.URL.Path, "/")
  width, _ := strconv.Atoi(url[2])
  height, _ := strconv.Atoi(url[3])
  fill := "#CCCCCC"
  stroke := "#222222"
  if (len(url) > 4) {
    fill = url[4]
    if (len(url) > 5) {
      stroke = url[5]
    }
  }

  values := &Placeholder{Height: height,Width: width,Fill: fill,Stroke: stroke}
  fmt.Printf("%+v\n", values)
  t, err := template.New("placeholder.svg").ParseFiles(filepath.Join("templates", "placeholder.svg"))

  if err != nil {
    fmt.Printf("%v", err)
  }

  // Set the content type to image/svg
  w.Header().Set("Content-Type", "image/svg+xml")
  rendererr := t.Execute(w, values)

  if rendererr != nil {
    fmt.Printf("%v", rendererr)
  }
}

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/svg/", svg)
  http.ListenAndServe(":8080", nil)
}
