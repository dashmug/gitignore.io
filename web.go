package main

import (
    "path"
    "net/http"
    "html/template"
)

type HomeData struct {
  TemplateCount int
}

type DocsData struct {
  TemplateCount int
}

// Home
func HomeHandler(w http.ResponseWriter, r *http.Request) {
  fp := path.Join("templates", "index.html")

  homeData := HomeData{Metrics.templateCount}

  tmpl, err := template.ParseFiles(fp)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }

  if err := tmpl.Execute(w, homeData); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func DocumentationHandler(w http.ResponseWriter, r *http.Request) {
  fp := path.Join("templates", "docs.html")
  docsData := DocsData{Metrics.templateCount}

  tmpl, err := template.ParseFiles(fp)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }

  if err := tmpl.Execute(w, docsData); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
  fp := path.Join("", "README.md")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }

  if err := tmpl.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
