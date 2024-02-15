package main

import (
	"fmt"
	"time"

	"github.com/bregydoc/gtranslate"
	"github.com/fogleman/gg"
	"github.com/notneet/NNbot/helper"
)

func main() {
	// Example: generate and save image with the translated text
	textToGenerate := "こんにちは" // Japanese greeting "Konnichiwa"
	generateAndSaveImage(textToGenerate)

	// Example: Translate from English to Bahasa Indonesia
	textEn := "Good Morning!"
	result, err := tl(textEn, gtranslate.TranslationParams{From: "en", To: "ja"})
	helper.PanicIfError(err)

	fmt.Printf("result: %v\n", result)
}

/*

 */

func tl(rawText string, opts gtranslate.TranslationParams) (string, error) {
	translated, err := gtranslate.TranslateWithParams(rawText, opts)
	if err != nil {
		return "", err
	}

	return translated, nil
}

func generateAndSaveImage(text string) {
	const W = 600
	const H = 400

	// Create a context with a white background
	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Customize drawing
	dc.SetRGB(0, 0, 0)

	// Replace this path with the path to the downloaded font file
	fontPath := "./resource/font/NotoSansJP-Medium.ttf"
	err := dc.LoadFontFace(fontPath, 100)
	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	dc.DrawStringAnchored(text, W/2, H/2, 0.5, 0.5)

	// Save the image to a file
	fileName := fmt.Sprintf("generated_image_%d.png", time.Now().UnixNano())
	err = dc.SavePNG(fileName)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Image saved as", fileName)
}
