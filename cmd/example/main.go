package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"github.com/notneet/NNbot/pkg/common"
)

func main() {
	// Example: generate and save image with the translated text
	textToGenerate := "こんにちは" // Japanese greeting "Konnichiwa"
	GenerateAndSaveImage(textToGenerate)

	// Example: Translate from English to Bahasa Indonesia
	textEn := "Good Morning!"
	result, err := TL(textEn, gtranslate.TranslationParams{From: "en", To: "id"})
	common.PanicIfError(err)

	fmt.Printf("result: %v\n", result)
}
