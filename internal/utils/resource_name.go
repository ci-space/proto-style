package utils

import (
	"strings"

	"github.com/jinzhu/inflection"
)

const serviceSuffix = "Service"

type ResourceName struct {
	Raw      string
	Singular string
	Plural   string
}

func (n *ResourceName) IsPlural() bool {
	return n.Raw == n.Plural
}

func ParseResourceNameFromServiceName(srvName string) ResourceName {
	return createResourceName(strings.TrimSuffix(srvName, serviceSuffix))
}

func createResourceName(rawName string) ResourceName {
	return ResourceName{
		Raw:      rawName,
		Singular: inflection.Singular(rawName),
		Plural:   inflection.Plural(rawName),
	}
}

func ParseResourceNameFromListRequestName(msgName string) ResourceName {
	return createResourceName(parseSingleValue(isListRequestRegex, msgName))
}

func ParseResourceNameFromListResponseName(msgName string) ResourceName {
	return createResourceName(parseSingleValue(isListResponseRegex, msgName))
}
