package models

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
)

func ReadModel(filepath string) string {
	success, content := readFileIntoByteArray(filepath)
	yamlMap := unmarshalYamlIntoMapIfShould(content, success)

	return processYamlMap(yamlMap)
}

func unmarshalYamlIntoMapIfShould(content []byte, shouldAttempt bool) map[interface{}]interface{} {
	var result map[interface{}]interface{}

	if shouldAttempt {
		result = unmarshalYamlIntoMap(content)
	}

	return result
}

func unmarshalYamlIntoMap(content []byte) map[interface{}]interface{} {
	contentMap := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(content), &contentMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return contentMap
}

func readFileIntoByteArray(filepath string) (bool, []byte) {
	success := false
	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	} else {
		success = true
	}
	return success, content
}

func processYamlMap(contentMap map[interface{}]interface{}) string {
	packageName := getAsString(contentMap["package"])
	modelsOutputString := processModels(contentMap["models"], packageName)

	result := convertYamlToOutputString(packageName, modelsOutputString)

	return result
}

func convertYamlToOutputString(packageName string, modelsOutputString string) string {
	return fmt.Sprintf("package: %s\n\n", packageName) + modelsOutputString
}

func processModels(models interface{}, packageName string) string {
	var result string

	for _, model := range models.([]interface{}) {
		result += processModel(model.(map[interface{}]interface{}))
	}

	return result
}

func processModel(model map[interface{}]interface{}) string {

	attributesOutputString := processAttributes(model["attributes"])

	result := convertModelToOutputString(
		getAsString(model["name"]),
		getAsString(model["description"]),
		attributesOutputString)

	// TODO Extend this to load the model into a struct that can be passed to the writer
	//  said struct should be passed up the chain as a second return result

	return result
}

func convertModelToOutputString(name string, desc string, attributes string) string {
	result := fmt.Sprintf("%s\n// %s\n", name, desc)
	result += attributes
	result += "\n\n"
	return result
}

func processAttributes(attributes interface{}) string {
	var result string

	for _, attribute := range attributes.([]interface{}) {
		result += processAttribute(attribute.(map[interface{}]interface{}))
	}

	return result
}

func processAttribute(attribute map[interface{}]interface{}) string {

	name := getAsString(attribute["name"])
	typ := getAsString(attribute["type"])
	defaultValue := convertValue(typ, getAsString(attribute["default"]))
	typ = convertType(typ)
	desc := getAsString(attribute["description"])

	result := convertAttributeToOutputString(name, typ, defaultValue, desc)

	// TODO Extend this to load the attribute into a struct that can be passed to the writer
	//  said struct should be passed up the chain as a second return result


	return result
}

func convertAttributeToOutputString(name string, typ string, defaultValue string, desc string) string {
	result := fmt.Sprintf("%s %s := %s\t// %s\n",
		name,
		typ,
		defaultValue,
		desc)
	return result
}

func getAsString(input interface{}) string {
	worked, result := getYamlStringAsStringIfShould(input, "", true)
	worked, result = getYamlIntAsStringIfShould(input, result, !worked)
	worked, result = getYamlFloatAsStringIfShould(input, result, !worked)
	worked, result = getYamlBooleanAsStringIfShould(input, result, !worked)
	worked, result = getYamlArrayAsStringIfShould(input, result, !worked)

	return result
}

func getYamlArrayAsStringIfShould(input interface{}, defaultValue string, shouldAttempt bool) (bool, string) {
	var succeeded bool
	var result string

	succeeded = false

	if shouldAttempt {
		result, succeeded = getYamlArrayAsString(input)

		if !succeeded {
			result = defaultValue
		}

	} else {
		result = defaultValue
	}

	return succeeded, result
}

func getYamlBooleanAsStringIfShould(input interface{}, defaultValue string, shouldAttempt bool) (bool, string) {
	var succeeded bool
	var result string

	succeeded = false

	if shouldAttempt {
		result, succeeded = getYamlBooleanAsString(input)

		if !succeeded {
			result = defaultValue
		}

	} else {
		result = defaultValue
	}

	return succeeded, result
}

func getYamlFloatAsStringIfShould(input interface{}, defaultValue string, shouldAttempt bool) (bool, string) {
	var succeeded bool
	var result string

	succeeded = false

	if shouldAttempt {
		result, succeeded = getYamlFloatAsString(input)

		if !succeeded {
			result = defaultValue
		}

	} else {
		result = defaultValue
	}

	return succeeded, result
}

func getYamlIntAsStringIfShould(input interface{}, defaultValue string, shouldAttempt bool) (bool, string) {
	var succeeded bool
	var result string

	succeeded = false

	if shouldAttempt {
		result, succeeded = getYamlIntAsString(input)

		if !succeeded {
			result = defaultValue
		}

	} else {
		result = defaultValue
	}

	return succeeded, result
}

func getYamlStringAsStringIfShould(input interface{}, defaultValue string, shouldAttempt bool) (bool, string) {
	var succeeded bool
	var result string

	succeeded = false

	if shouldAttempt {
		result, succeeded = getYamlStringAsString(input)

		if !succeeded {
			result = defaultValue
		}

	} else {
		result = defaultValue
	}

	return succeeded, result
}

func getYamlStringAsString(input interface{}) (string, bool) {
	result, worked := input.(string)

	return result, worked
}

func getYamlIntAsString(input interface{}) (string, bool) {
	var result string

	valInt, worked := input.(int)

	if worked {
		result = strconv.Itoa(valInt)
	}
	return result, worked
}

func getYamlFloatAsString(input interface{}) (string, bool) {
	var result string

	valFloat, worked := input.(float64)

	if worked {
		result = strconv.FormatFloat(valFloat, 'f', 50, 64)
	}
	return result, worked
}

func getYamlBooleanAsString(input interface{}) (string, bool) {
	var result string

	valBool, worked := input.(bool)

	if worked {
		result = strconv.FormatBool(valBool)
	}
	return result, worked
}

func getYamlArrayAsString(input interface{}) (string, bool) {
	var result string

	valArr, worked := input.([]interface{})

	if worked {
		result = "[" + valArr[0].(string) + "]"
	}

	return result, worked
}

func convertValue(instanceType string, value string) string {
	result := value

	if result == "" {
		result = "nil"
	}

	return result
}

func convertType(instanceType string) string {
	return instanceType
}
