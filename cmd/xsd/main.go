package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

func main() {
	ExampleXSD()
}

func ExampleXSD() {
	xsdfile := filepath.Join("test", "xmldsig-core-schema.xsd")
	f, err := os.Open(xsdfile)
	if err != nil {
		log.Printf("failed to open file: %s", err)
		return
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("failed to read file: %s", err)
		return
	}

	s, err := xsd.Parse(buf)
	if err != nil {
		log.Printf("failed to parse XSD: %s", err)
		return
	}
	defer s.Free()

	d, err := libxml2.ParseString(`<foo></foo>`)
	if err != nil {
		log.Printf("failed to parse XML: %s", err)
		return
	}
	defer d.Free()

	if err := s.Validate(d); err != nil {
		for _, e := range err.(xsd.SchemaValidationError).Errors() {
			log.Printf("error: %s", e.Error())
		}
		return
	}

	log.Printf("validation successful!")
}
