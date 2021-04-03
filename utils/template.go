package utils

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func populateTemplate() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const sharedPath = "templates/shared"
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(sharedPath + "/_layout.html"))
	template.Must(layout.ParseFiles(sharedPath+"/_header.html", sharedPath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

func ServeTemplate(data interface{}, filename string, w http.ResponseWriter) {
	file := fmt.Sprintf("%s.html", filename)
	tmpl := populateTemplate()
	t := tmpl[file]
	t.Execute(w, data)
}
