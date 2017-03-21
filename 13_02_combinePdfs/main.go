package main

import (
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
	pdfL "github.com/ledongthuc/pdf"
)

func main() {
	// в какой папке исполняемый файл
	pwdDir, _ := os.Getwd()

	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// запись ошибок и инфы в файл
	log.SetOutput(fLog)

	//открытие папки
	dir, err := os.Open(pwdDir)
	if err != nil {
		log.Fatalln(err)
	}
	defer dir.Close()

	// список файлов
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalln(err)
	}

	// создание и запись PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "combinePdfs")

	for _, fi := range fileInfos {
		// проверка на .pdf
		if strings.HasSuffix(fi.Name(), ".pdf") {
			r, err := pdfL.Open(fi.Name())
			if err != nil {
				log.Fatalln(err)
			}
			totalPage := r.NumPage()

			for pageIndex := 2; pageIndex <= totalPage; pageIndex++ {
				p := r.Page(pageIndex)
				if p.V.IsNull() {
					continue
				}
				// объединение PDF
				pdf.AddPage()
				pdf.RawWriteStr(p.V.String())
			}
		}
	}

	err = pdf.OutputFileAndClose("save.pdf")
}
