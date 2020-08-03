package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// SetTextPrinter ...
func SetTextPrinter(langCode string) (p *message.Printer) {
	switch langCode {
	case "ru":
		p = message.NewPrinter(language.Russian)
	case "uz":
		p = message.NewPrinter(language.Uzbek)
	default:
		p = message.NewPrinter(language.English)
	}
	return p
}
