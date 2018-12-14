package spec

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type Spec struct {
	Resources map[string]*SpecFile
}

func ReadSpec(folder string) (*Spec, error) {
	resourceFiles, err := ioutil.ReadDir(folder)

	if err != nil {
		log.Println("Failed to enumerate resources folder : ", err)

		return nil, err
	}

	resources := map[string]*SpecFile{}

	for _, file := range resourceFiles {
		fileName := filepath.Base(file.Name())
		resourceName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

		resource, err := parseSpecFile(filepath.Join(folder, file.Name()))

		if err != nil {
			log.Println("Failed to parse resource : ", err)

			return nil, err
		}

		resources[resourceName] = resource
	}

	spec := Spec{
		Resources: resources,
	}

	return &spec, nil
}