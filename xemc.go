package main

import (
	"flag"
	"fmt"
	"github.com/xentenceplatform/xemc/models"
)

var modelYamlPath string

func init() {
	flag.StringVar(&modelYamlPath, "model", "", "Path to the YAML model file to compile")
	flag.StringVar(&modelYamlPath, "m", "", "Path to the YAML model file to compile")
}

func main() {
	fmt.Println("xemc: Xentence Model Compiler")

	flag.Parse()

	if modelYamlPath == "" {
		fmt.Println("No model was specified, specify a model with -m or -model")
	} else {
		fmt.Println(models.ReadModel(modelYamlPath))
	}
}
