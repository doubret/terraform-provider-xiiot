package main

import (
	"flag"
	"github.com/doubret/terraform-provider-xiiot/generator/spec"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var inputFolder = flag.String("i", "", "input folder to read specs from")

func tf_schema_type(fieldType string) string {
	fieldType = strings.TrimSuffix(fieldType, "[]")

	isEnum := spec.IsEnum(fieldType)

	fieldType = strings.TrimPrefix(fieldType, "(")
	fieldType = strings.TrimSuffix(fieldType, ")")

	base := "TypeString"

	switch fieldType {
	case "ip", "ip_mask", "string":
		base = "TypeString"
	case "double", "int":
		base = "TypeInt"
	case "bool":
		base = "TypeBool"
	}

	if isEnum {
		base = "TypeString"
	}

	return base
}

func main() {
	flag.Parse()

	if inputFolder == nil {
		log.Print("No input spec folder specified")

		return
	} else {
		model, err := spec.ReadSpec(*inputFolder)

		if err != nil {
			log.Println("Failed to read spec : ", err)
		} else {
			funcMap := template.FuncMap{
				"title":          strings.Title,
				"is_array":       spec.IsArray,
				"go_type":        spec.GoType,
				"go_base_type":   spec.GoBaseType,
				"name":           spec.Name,
				"tf_schema_type": tf_schema_type,
				"is_in":          spec.IsIn,
			}

			templates := template.Must(template.New("").Funcs(funcMap).ParseFiles(
				"templates/common.tmpl",
				"templates/resource.tmpl",
				"templates/provider.tmpl",
			))

			for key, value := range model.Resources {
				context := struct {
					Name   string
					Schema *spec.SpecFile
				}{
					key,
					value,
				}

				writer, err := os.Create(filepath.Join("xiiot", "resources", key+".go"))

				if err != nil {
					log.Println("Failed to create file : ", err)
				}

				err = templates.ExecuteTemplate(writer, "resource.tmpl", context)

				if err != nil {
					log.Println("Failed to execute template : ", err)
				}
			}

			writer, err := os.Create(filepath.Join("xiiot", "provider.go"))
			err = templates.ExecuteTemplate(writer, "provider.tmpl", model)
			if err != nil {
				log.Println("Failed to execute provider template : ", err)
			}

			return
		}
	}
}
