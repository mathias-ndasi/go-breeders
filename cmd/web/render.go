package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) renderTemplate(responseWriter http.ResponseWriter, templateName string, data *templateData) {
	var rawTemplate *template.Template

	// If we are using the template cache, try to get the template from our map, stored in the receiver
	if app.config.useCache {
		if cachedTemplate, ok := app.templateMap[templateName]; ok {
            rawTemplate = cachedTemplate
        }
	}

	if rawTemplate == nil {
		newTemplate, error := app.buildTemplateFromDisk(templateName)
		if error!= nil {
            log.Println("Error building template: ", error);
			return
        }

		log.Println("Building template from disk: ", templateName)
		rawTemplate = newTemplate;
	}

	if data == nil {
		data = &templateData{}
	}

	if error := rawTemplate.ExecuteTemplate(responseWriter, templateName, data); error != nil {
		log.Println("Error executing template: ", error);
		http.Error(responseWriter, error.Error(), http.StatusInternalServerError)
        return
	}
}

func (app *application) buildTemplateFromDisk(templateName string) (*template.Template, error) {
	templateFiles := []string{
        "./templates/base.layout.gohtml",
        "./templates/partials/header.partial.gohtml",
        "./templates/partials/footer.partial.gohtml",
		fmt.Sprintf("./templates/%s", templateName),
    }

	rawTemplate, error := template.ParseFiles(templateFiles...);
	if error!= nil {
        return nil, error
    }

	app.templateMap[templateName] = rawTemplate;

    return rawTemplate, nil
}
