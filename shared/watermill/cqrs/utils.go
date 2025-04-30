package cqrs

import (
	"regexp"
	"strings"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/dreadster3/pawcare/shared/events"
)

func StructNameDotLower(v interface{}) string {
	if event, ok := v.(events.IEvent); ok {
		return event.EventName()
	}

	structName := cqrs.StructName(v)

	regex := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	structName = regex.ReplaceAllString(structName, "${1}.${2}")

	return strings.ToLower(structName)
}
