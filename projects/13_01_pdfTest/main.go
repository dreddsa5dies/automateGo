// чтение и запись PDF
package main

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/ledongthuc/pdf"
)

func main() {
	// чтение PDF
	content, err := readPdf("basic.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	// создание и запись PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, content)
	err = pdf.OutputFileAndClose("hello.pdf")

	return
}

func readPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	var textBuilder bytes.Buffer
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		t := p.Content().Text
		for i := range t {
			textBuilder.WriteString(t[i].S)
		}
	}
	return textBuilder.String(), nil
}
