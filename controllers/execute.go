package controllers

import (
	"github.com/ledongthuc/pdf"
	"mangtas/features"
	"net/http"
)

func Execute(http.ResponseWriter, *http.Request) {
	pdf.DebugOn = true

	// Read local pdf file
	content, err := features.ReadPdf("static/test.pdf")
	if err != nil {
		panic(err)
	}

	wordsCountedByFrequency := features.WordCount(content)

	features.PrintAnswer(wordsCountedByFrequency)
}
