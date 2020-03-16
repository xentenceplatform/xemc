package main

import (
	"flag"
	"fmt"
	"github.com/xentenceplatform/xemc/models"
)

var modelYamlPath string

func displayApplicationHeader() (int, error) {
	return fmt.Println("xemc: Xentence Model Compiler")
}

func init() {
	configureModelPathFlagArgument()
}

func main() {
	displayApplicationHeader()
	parseArguments()
	processModelFileIfShould(hasValidModelPath())
}

func configureModelPathFlagArgument() {
	configureStringFlagArgument(
		&modelYamlPath,
		"model",
		"m",
		"",
		"Path to the YAML model file to compile")
}

func hasValidModelPath() bool {
	var validModelPath bool

	if modelYamlPath == "" {
		fmt.Println(getModelPathMissingString())
		validModelPath = false
	} else {
		validModelPath = true
	}
	return validModelPath
}

func getModelPathMissingString() string {
	return "No model was specified, specify a model with -m or -model"
}

func configureStringFlagArgument(storagePointer *string, longHandflag string, shortHandFlag string, defaultValue string, usage string) {
	flag.StringVar(storagePointer, longHandflag, defaultValue, usage)
	flag.StringVar(storagePointer, shortHandFlag, defaultValue, usage)
}

func parseArguments() {
	flag.Parse()
}

func processModelFileIfShould(shouldAttempt bool) {
	if shouldAttempt {
		outputString := models.ReadModel(modelYamlPath)

		fmt.Println(outputString)
	}
}
