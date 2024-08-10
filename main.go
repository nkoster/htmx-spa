package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Helper function to load file content into a template.HTML
func loadFileContent(filePath string) template.HTML {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return template.HTML("Error loading file")
	}
	return template.HTML(content)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Load index.html from the "ui" directory
		t, _ := template.ParseFiles(filepath.Join("ui", "index.html"))
		content := loadFileContent(filepath.Join("templates", "home.html")) // Load the content of home.html
		if err := t.Execute(w, map[string]template.HTML{"Content": content}); err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/page_a.html", func(w http.ResponseWriter, r *http.Request) {
		// Load index.html from the "ui" directory and content from the "templates" directory
		t, _ := template.ParseFiles(filepath.Join("ui", "index.html"))
		// Load the content of page_a.html
		content := loadFileContent(filepath.Join("templates", "page_a.html"))
		if err := t.Execute(w, map[string]template.HTML{"Content": content}); err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/page_b.html", func(w http.ResponseWriter, r *http.Request) {
		// Load index.html from the "ui" directory and content from the "templates" directory
		t, _ := template.ParseFiles(filepath.Join("ui", "index.html"))
		// Load the content of page_b.html
		content := loadFileContent(filepath.Join("templates", "page_b.html"))
		if err := t.Execute(w, map[string]template.HTML{"Content": content}); err != nil {
			fmt.Println(err)
		}
	})

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
