package goldi

import (
	"fmt"
	"strings"
	"unicode"
)

type ConfiguredType struct {
	*TypeConfigurator
	EmbeddedType TypeFactory
}

func NewConfiguredType(embeddedType TypeFactory, configuratorTypeID, configuratorMethod string) *ConfiguredType {
	if embeddedType == nil {
		panic(fmt.Errorf("refusing to create a new ConfiguredType with nil as embedded type"))
	}

	configuratorTypeID = strings.TrimSpace(configuratorTypeID)
	configuratorMethod = strings.TrimSpace(configuratorMethod)

	if configuratorTypeID == "" || configuratorMethod == "" {
		panic(fmt.Errorf("can not create a new ConfiguredType with empty configurator type or method (%q, %q)", configuratorTypeID, configuratorMethod))
	}

	if unicode.IsLower(rune(configuratorMethod[0])) {
		panic(fmt.Errorf("can not create a new ConfiguredType with unexproted configurator method %q", configuratorMethod))
	}

	return &ConfiguredType{
		TypeConfigurator: NewTypeConfigurator(configuratorTypeID, configuratorMethod),
		EmbeddedType:     embeddedType,
	}
}

func (t *ConfiguredType) Arguments() []interface{} {
	return append(t.EmbeddedType.Arguments(), "@"+t.ConfiguratorTypeID)
}

func (t *ConfiguredType) Generate(parameterResolver *ParameterResolver) interface{} {
	embedded := t.EmbeddedType.Generate(parameterResolver)
	err := t.Configure(embedded, parameterResolver.Container)
	if err != nil {
		panic(err)
	}

	return embedded
}