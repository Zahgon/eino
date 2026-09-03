package compose

import (
	"reflect"
)

type FieldMapping struct {
	fromNodeKey string
	from        string
	to          string

	customExtractor func(input any) (any, error)
}

func (m *FieldMapping) String() string { _ = "STUB: not implemented"; return "" }

func FromField(from string) *FieldMapping { _ = "STUB: not implemented"; return nil }

func ToField(to string, opts ...FieldMappingOption) *FieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func MapFields(from, to string) *FieldMapping { _ = "STUB: not implemented"; return nil }

func (m *FieldMapping) FromNodeKey() string { _ = "STUB: not implemented"; return "" }

func (m *FieldMapping) FromPath() FieldPath { _ = "STUB: not implemented"; return *new(FieldPath) }

func (m *FieldMapping) ToPath() FieldPath { _ = "STUB: not implemented"; return *new(FieldPath) }

func (m *FieldMapping) Equals(o *FieldMapping) bool { _ = "STUB: not implemented"; return false }

type FieldPath []string

func (fp *FieldPath) join() string { _ = "STUB: not implemented"; return "" }

func splitFieldPath(path string) FieldPath { _ = "STUB: not implemented"; return *new(FieldPath) }

const pathSeparator = "\x1F"

func FromFieldPath(fromFieldPath FieldPath) *FieldMapping { _ = "STUB: not implemented"; return nil }

func ToFieldPath(toFieldPath FieldPath, opts ...FieldMappingOption) *FieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func MapFieldPaths(fromFieldPath, toFieldPath FieldPath) *FieldMapping {
	_ = "STUB: not implemented"
	return nil
}

type FieldMappingOption func(*FieldMapping)

func WithCustomExtractor(extractor func(input any) (any, error)) FieldMappingOption {
	_ = "STUB: not implemented"
	return *new(FieldMappingOption)
}

func (m *FieldMapping) targetPath() FieldPath { _ = "STUB: not implemented"; return *new(FieldPath) }

func buildFieldMappingConverter[I any]() func(input any) (any, error) {
	_ = "STUB: not implemented"
	return nil
}

func buildStreamFieldMappingConverter[I any]() func(input streamReader) streamReader {
	_ = "STUB: not implemented"
	return nil
}

func convertTo(mappings map[string]any, typ reflect.Type) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func assignOne(destValue reflect.Value, taken any, to string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func instantiateIfNeeded(field reflect.Value) { _ = "STUB: not implemented"; return }

func newInstanceByType(typ reflect.Type) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func checkAndExtractFromField(fromField string, input reflect.Value) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

type errMapKeyNotFound struct {
	mapKey string
}

func (e *errMapKeyNotFound) Error() string { _ = "STUB: not implemented"; return "" }

type errInterfaceNotValidForFieldMapping struct {
	interfaceType reflect.Type
	actualType    reflect.Type
}

func (e *errInterfaceNotValidForFieldMapping) Error() string { _ = "STUB: not implemented"; return "" }

func checkAndExtractFromMapKey(fromMapKey string, input reflect.Value) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func checkAndExtractFieldType(paths []string, typ reflect.Type) (extracted reflect.Type, remainingPaths FieldPath, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), *new(FieldPath), nil
}

var strType = reflect.TypeOf("")

func fieldMap(mappings []*FieldMapping, allowMapKeyNotFound bool, uncheckedSourcePaths map[string]FieldPath) func(any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil
}

func streamFieldMap(mappings []*FieldMapping, uncheckedSourcePaths map[string]FieldPath) func(streamReader) streamReader {
	_ = "STUB: not implemented"
	return nil
}

func takeOne(inputValue reflect.Value, inputType reflect.Type, from string) (taken any, takenType reflect.Type, err error) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Type), nil
}

func isFromAll(mappings []*FieldMapping) bool { _ = "STUB: not implemented"; return false }

func fromFields(mappings []*FieldMapping) bool { _ = "STUB: not implemented"; return false }

func isToAll(mappings []*FieldMapping) bool { _ = "STUB: not implemented"; return false }

func validateStructOrMap(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func validateFieldMapping(predecessorType reflect.Type, successorType reflect.Type, mappings []*FieldMapping) (

	typeHandler *handlerPair,

	uncheckedSourcePath map[string]FieldPath,
	err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
