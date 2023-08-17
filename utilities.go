package gorae

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func convertWordDefinitionToJSON(wordDefinition *WordDefinition) (string, error) {
	wordDefinitionJson, err := json.Marshal(wordDefinition)
	if err != nil {
		return "", err
	}

	return string(wordDefinitionJson), nil
}

func sanitizeStrNum(strNum string) (int, error) {
	strNum = strings.Trim(strNum, " .")

	num, err := strconv.Atoi(strNum)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	return num, nil
}
