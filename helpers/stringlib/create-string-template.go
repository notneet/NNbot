package stringlib

import (
	"text/template"

	"github.com/notneet/NNbot/helpers"
)

func CreateStringTemplate(templateName string, text string) *template.Template {
	textMsg, err := template.New(templateName).Parse(text)
	helpers.PanicIfError(err)

	return textMsg
}
