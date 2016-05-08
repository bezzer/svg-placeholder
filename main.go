package main

import (
  "fmt"
  "strconv"
  "regexp"
  "path/filepath"
  "html/template"
  "net/http"
)

// Placeholder Holds the values for building an SVG
type Placeholder struct {
  Width  int
  Height int
  BorderWidth int
  BorderHeight int
  StrokeWidth int
  Fill   string
  Stroke string
  Message string
}

// SVG placeholder template
const svgTemplate = `<svg width="{{.Width}}" height="{{.Height}}" xmlns="http://www.w3.org/2000/svg">
  <rect x="{{.StrokeWidth}}" y="{{.StrokeWidth}}" width="{{.BorderWidth}}" height="{{.BorderHeight}}" style="fill: #{{.Fill}}; stroke: #{{.Stroke}}; stroke-width: {{.StrokeWidth}}"/>
  <text x="50%" y="50%" font-size="24" text-anchor="middle" alignment-baseline="middle" font-family="sans-serif" fill="{{.Stroke}}">
    {{if .Message}}
      {{.Message}}
    {{else}}
      {{.Width}}x{{.Height}}
    {{end}}
  </text>
</svg>`

const strokeWidth = 2

// Pre-parse the template
var templates = template.Must(template.New("svg").Parse(svgTemplate))

// Patern matcher for SVG URLs
var svgPatern = regexp.MustCompile("\\/(\\d+)(?:\\/(\\d+))?(?:\\/([0-9a-fA-F]{6}))?(?:\\/([0-9a-fA-F]{6}))?")

// Default handler
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi There")
}

// Handler for URL paths /svg/WIDTH/HEIGHT/[FILL/STROKE]
func svg(w http.ResponseWriter, r *http.Request) {
  
  var width, height int
  var fill, stroke, message string
  fill = "CCCCCC"
  stroke = "333333"
  
  if svgPatern.MatchString(r.URL.Path) {
    // Output SVG
    matches := svgPatern.FindStringSubmatch(r.URL.Path)
    fmt.Printf("%v Matches: %+v\n", len(matches), matches)
    
    // Width must always be defined
    width, _ = strconv.Atoi(matches[1])
    // Height defaults to width (square) if not defined
    height = width
    if len(matches[2]) > 0 {
      height, _ = strconv.Atoi(matches[2])
    }
    // Fill colour
    if len(matches[3]) > 0 {
      fill = matches[3]
    }
    // Stroke colour
    if len(matches[4]) > 0 {
      stroke = matches[4]
    } 
  } else {
    // Show error image
    width = 300
    height = 100
    message = "Unsupported request"
  }
  
  values := &Placeholder{
    Height: height,
    Width: width,
    Fill: fill,
    Stroke: stroke,
    StrokeWidth: strokeWidth,
    Message: message,
    BorderWidth: width - strokeWidth * 2,
    BorderHeight: height - strokeWidth * 2}

  // Set the content type to image/svg
  w.Header().Set("Content-Type", "image/svg+xml")
  w.Header().Set("Cache-Control", "max-age=31536000")
  
  // Construct the output
  rendererr := templates.ExecuteTemplate(w, "svg", values)

  if rendererr != nil {
    fmt.Printf("Error rendering template: %v", rendererr)
  }
}

func main() {
  var iconsHandler = http.StripPrefix("/images/icons/", http.FileServer(http.Dir(filepath.Join("images", "icons"))))
  http.Handle("/favicon.ico", iconsHandler)
  http.Handle("/images/icons/", iconsHandler)
  http.HandleFunc("/", handler)
  http.Handle("/svg/", http.StripPrefix("/svg", http.HandlerFunc(svg)))
  http.ListenAndServe(":5000", nil)
}
