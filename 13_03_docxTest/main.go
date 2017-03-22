package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	// docx это zip с документами xml
	reader, err := zip.OpenReader("1.docx")
	if err != nil {
		log.Fatalln(err)
	}

	// поиск в архиве
	files := reader.File
	var file *zip.File
	for _, f := range files {
		// поиск файла с текстом
		if f.Name == "word/document.xml" {
			file = f
		}
	}
	// не найден документ с текстом
	if file == nil {
		log.Println("document.xml file not found")
	}

	// открытие файла с текстом
	dataFile, _ := file.Open()
	defer dataFile.Close()

	b, err := ioutil.ReadAll(dataFile)
	if err != nil {
		log.Fatalln(err)
	}

	type Body struct {
		Paragraph []string `xml:"p>r>t"`
	}

	type Document struct {
		XMLName xml.Name `xml:"document"`
		Body    Body     `xml:"body"`
	}
	var v Document

	err = xml.Unmarshal(b, &v)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Value: %v\n", v.Body.Paragraph)
}
