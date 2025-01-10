package utils

import (
	"strings"

	"github.com/jinzhu/inflection"
)

const serviceSuffix = "Service"

type EntityName struct {
	Raw      string
	Singular string
	Plural   string
}

func (n *EntityName) IsPlural() bool {
	return n.Raw == n.Plural
}

func ParseEntityNameFromServiceName(srvName string) EntityName {
	return createEntityName(strings.TrimSuffix(srvName, serviceSuffix))
}

func createEntityName(rawName string) EntityName {
	return EntityName{
		Raw:      rawName,
		Singular: inflection.Singular(rawName),
		Plural:   inflection.Plural(rawName),
	}
}

func ParseEntityNameFromListRequestName(msgName string) EntityName {
	return createEntityName(parseSingleValue(isListRequestRegex, msgName))
}

func ParseEntityNameFromListResponseName(msgName string) EntityName {
	return createEntityName(parseSingleValue(isListResponseRegex, msgName))
}
