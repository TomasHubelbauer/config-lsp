package docvalues

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CustomValueContext interface {
	GetIsContext() bool
}

type EmptyValueContext struct{}

func (EmptyValueContext) GetIsContext() bool {
	return true
}

var EmptyValueContextInstance = EmptyValueContext{}

type CustomValue struct {
	FetchValue func(context CustomValueContext) DeprecatedValue
}

func (v CustomValue) GetTypeDescription() []string {
	return v.FetchValue(EmptyValueContextInstance).GetTypeDescription()
}

func (v CustomValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	return v.FetchValue(EmptyValueContextInstance).DeprecatedCheckIsValid(value)
}

func (v CustomValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return v.FetchValue(EmptyValueContextInstance).DeprecatedFetchCompletions(line, cursor)
}

func (v CustomValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return v.FetchValue(EmptyValueContextInstance).DeprecatedFetchHoverInfo(line, cursor)
}