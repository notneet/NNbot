package main

import "github.com/bregydoc/gtranslate"

func TL(rawText string, opts gtranslate.TranslationParams) (string, error) {
	translated, err := gtranslate.TranslateWithParams(rawText, opts)
	if err != nil {
		return "", err
	}

	return translated, nil
}
