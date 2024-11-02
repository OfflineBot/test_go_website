
package main

import (
    "html/template"
    "log"
    "net/http"
)


func mainHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/main.html")
    if err != nil {
        http.Error(w, "Could not load template", http.StatusInternalServerError)
        log.Println("Error loading template:", err)
        return
    }
    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Failed to render template", http.StatusInternalServerError)
        log.Println("Error executing template:", err)
    }
}


func main() {
    http.HandleFunc("/", mainHandler)
    log.Println("Server started at http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
