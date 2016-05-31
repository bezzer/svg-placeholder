package main

import (
  "fmt"
  "log"
  "strconv"
  "strings"
  "regexp"
  "html/template"
  "net/http"
  "github.com/NYTimes/gziphandler"
)

// Placeholder Holds the values for building an SVG
type Placeholder struct {
  Width  int
  Height int
  BorderWidth int
  BorderHeight int
  StrokeWidth int
  Fill string
  FillEnd string
  StrokeColor string
  TextColor string
  Message string
  ShowText bool
}

// SVG placeholder template (string formated to remove newlines and spaces)
const svgTemplate = "<svg width=\"{{.Width}}\" height=\"{{.Height}}\" xmlns=\"http://www.w3.org/2000/svg\">" +
  "{{if .FillEnd}}" +
    "<linearGradient id=\"lg\">"+
      "<stop offset=\"0%\" stop-color=\"#{{.Fill}}\"/>" +
      "<stop offset=\"100%\" stop-color=\"#{{.FillEnd}}\"/>" +
    "</linearGradient>" +
  "{{end}}" +
  "<rect x=\"{{.StrokeWidth}}\" y=\"{{.StrokeWidth}}\" width=\"{{.BorderWidth}}\" height=\"{{.BorderHeight}}\" style=\"fill:{{if .FillEnd}}url(#lg){{else}}#{{.Fill}}{{end}};stroke:#{{.StrokeColor}};stroke-width:{{.StrokeWidth}}\"/>" +
  "<text x=\"50%\" y=\"50%\" font-size=\"18\" text-anchor=\"middle\" alignment-baseline=\"middle\" font-family=\"monospace, sans-serif\" fill=\"#{{.TextColor}}\">"+
    "{{if .ShowText}}" +
      "{{if .Message}}" +
        "{{.Message}}" +
      "{{else}}" +
        "{{.Width}}&#215;{{.Height}}" +
      "{{end}}" +
    "{{end}}" +
  "</text>" +
"</svg>"

// Default strokewidth
const strokeWidth = 2

// Pre-parse the template
var templates = template.Must(template.New("svg").Parse(svgTemplate))

// Patern matcher for SVG URLs
var svgPatern = regexp.MustCompile(`\/(\d+)(?:x(\d+))?(?:\/([\da-f]{6}|[\da-f]{3})(?:-([\da-f]{6}|[\da-f]{3}))?)?(?:\/([\da-f]{6}|[\da-f]{3})(?:-([\da-f]{6}|[\da-f]{3}))?(?:\/(.+))?)?`)

// Handler for URL paths /svg/WIDTH/HEIGHT/[FILL/STROKE]
func svg(w http.ResponseWriter, r *http.Request) {
  var showText bool
  var width, height int
  var fill, fillEnd, strokeColor, textColor, message string
  
  // Lowercase the path to simpify the regex
  path := strings.ToLower(r.URL.Path)
  
  fill = "DEDEDE"
  textColor = "555555"
  strokeColor = "555555"
  
  if svgPatern.MatchString(path) {
    // Output SVG
    matches := svgPatern.FindStringSubmatch(path)
    
    // Width must always be defined
    width, _ = strconv.Atoi(matches[1])
    // Height defaults to width (square) if not defined
    height = width
    if len(matches[2]) > 0 {
      height, _ = strconv.Atoi(matches[2])
    }
    // Determine whether to show text based on width/height
    showText = width >= 75 && height >= 40
    // Fill colour
    if len(matches[3]) > 0 {
      fill = matches[3]
    }
    // Fill end colour (for gradients) 
    if len(matches[4]) > 0 {
      fillEnd = matches[4]
    }
    // Stroke colour
    if len(matches[5]) > 0 {
      textColor = matches[5]
    }
    // Border color
    if len(matches[6]) > 0 {
      strokeColor = matches[6]
    } else {
      strokeColor = textColor;
    }
    // Text
    if len(matches[7]) > 0 {
      message = matches[7]
    }
  } else {
    // Show error image
    width = 300
    height = 100
    message = "Unsupported"
  }
  
  values := &Placeholder{
    Height: height,
    Width: width,
    Fill: fill,
    FillEnd: fillEnd,
    StrokeColor: strokeColor,
    TextColor: textColor,
    StrokeWidth: strokeWidth,
    Message: message,
    ShowText: showText,
    BorderWidth: width - strokeWidth * 2,
    BorderHeight: height - strokeWidth * 2}

  // Set the content type to image/svg
  w.Header().Set("Content-Type", "image/svg+xml")
  w.Header().Set("Cache-Control", "max-age=31536000")
  w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%dx%d.svg", width, height))
  // Construct the output
  rendererr := templates.ExecuteTemplate(w, "svg", values)

  if rendererr != nil {
    log.Printf("Error rendering template: %v", rendererr)
  } else {
    log.Printf("SVG Placeholder of width %d, height %d, fill %s and stroke %s generated. Message: %s", width, height, fill, strokeColor, message)
  }
}

func main() {
  fileHandler := http.FileServer(http.Dir("www"))
  http.Handle("/", gziphandler.GzipHandler(fileHandler))
  http.Handle("/svg/", gziphandler.GzipHandler(http.StripPrefix("/svg", http.HandlerFunc(svg))))
  http.ListenAndServe(":5000", nil)
}
