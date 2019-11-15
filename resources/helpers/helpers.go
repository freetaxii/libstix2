package helpers

import (
	"errors"
	"strings"
)

//StringInSlice - check if slice contains string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ValidSlice - check if slice contains value from vocabs
func ValidSlice(methodName string, list, vocabsList []string) (bool, error) {
	for _, val := range list {
		if !StringInSlice(val, vocabsList) {
			return false, errors.New("the " + methodName + " property should be one of list: " + strings.Join(vocabsList, ", "))
		}
	}

	return true, nil
}

//AddToList - append comma-delimited values or array of values to the first argument or
func AddToList(list []string, data interface{}) ([]string, error) {
	arr := []string{}

	switch data.(type) {
	case string:
		arr = strings.Split(data.(string), ",")
	case []string:
		arr = data.([]string)
	default:
		return list, errors.New("Wrong data param type")
	}

	for _, val := range arr {
		list = append(list, strings.TrimSpace(val))
	}

	return list, nil
}
