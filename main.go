package main

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type DataModel struct {
	Company string
	Contact string
	Country string
}

func main() {

	var templ *template.Template
	var err error

	// use Go's default HTML template generation tools to generate your HTML
	if templ, err = template.ParseFiles("./templates/sample.html"); err != nil {
		panic(err)
	}

	templateData := struct {
		Title       string
		Description string
		DataModels  []DataModel
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		DataModels: []DataModel{
			DataModel{
				Company: "Jhon Lewis",
				Contact: "Maria Anders",
				Country: "Germany",
			},
			DataModel{
				Company: "Lo",
				Contact: "Maria S",
				Country: "TH",
			},
			DataModel{
				Company: "Mark",
				Contact: "Maria S",
				Country: "US",
			},
		},
	}

	// apply the parsed HTML template data and keep the result in a Buffer
	var body bytes.Buffer
	if err = templ.Execute(&body, templateData); err != nil {
		panic(err)
	}

	// initalize a wkhtmltopdf generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	// enable this if the HTML file contains local references such as images, CSS, etc.
	page.EnableLocalFileAccess.Set(true)

	// add the page to your generator
	pdfg.AddPage(page)

	// manipulate page attributes as needed
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// magic
	err = pdfg.Create()
	if err != nil {
		panic(err)
	}

	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		panic(err)
	}

}
