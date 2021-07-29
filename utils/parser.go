package utils

import (
	"io/ioutil"
	"os"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/yaml"
)

var log = ctrl.Log.WithName("utils")

// ParseYAML parses the YAML path and assigns the decoded values
// into the second parameter
func ParseYAML(path string, result interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		log.Error(err, "Unable to open the file from the given path")
		return err
	}
	defer file.Close()
	dataInByte, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error(err, "Unable to read from the reader")
	}
	err = yaml.Unmarshal(dataInByte, result)
	if err != nil {
		log.Error(err, "Unable to unmarshal bytes into result interface")
		return err
	}
	return nil
}
