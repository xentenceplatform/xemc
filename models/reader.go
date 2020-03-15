package models

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
)

func ReadModel(filepath string) string {
	var result string
	content, error := ioutil.ReadFile(filepath)

	if error != nil {
		log.Fatal(error)
		result = ""
	} else {
		contentMap := make(map[interface{}]interface{})

		error = yaml.Unmarshal([]byte(content), &contentMap)
		if error != nil {
			log.Fatalf("error: %v", error)
		}

		result = processModels(contentMap)
	}

	return result
}

func processModels(contentMap map[interface{}]interface{}) string {
	result := fmt.Sprintf("package: %s\n\n", contentMap["package"])

	for _, model := range contentMap["models"].([]interface{}) {
		result += processModel(model.(map[interface{}]interface{}))
	}

	return result
}

func processModel(model map[interface{}]interface{}) string {
	result := fmt.Sprintf("%s\n// %s\n", model["name"], model["description"])

	for _, model := range model["attributes"].([]interface{}) {
		result += processAttribute(model.(map[interface{}]interface{}))
	}

	result += "\n\n"

	return result
}

func processAttribute(attribute map[interface{}]interface{}) string {

	typ := getAsString(attribute["type"])

	result := fmt.Sprintf("%s %s := %s\t// %s\n",
		attribute["name"],
		convertType(typ),
		convertValue(typ, getAsString(attribute["default"])),
		attribute["description"])

	return result
}

func getAsString(input interface{}) string {
	result, worked := input.(string)

	if !worked {
		valInt, worked := input.(int)

		if !worked {
			valFloat, worked := input.(float64)

			if !worked {
				valBool, worked := input.(bool)

				if !worked {
					valArr, worked := input.([]interface{})

					if worked {
						result = "[" + valArr[0].(string) + "]"
					}
				} else {
					result = strconv.FormatBool(valBool)
				}

			} else {
				result = strconv.FormatFloat(valFloat, 'f', 50, 64)
			}
		} else {
			result = strconv.Itoa(valInt)
		}
	}
	return result
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
