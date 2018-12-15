package main

import "fmt"

type studentRecord struct {
	id          int
	name        string
	grade       int
	extraFields map[string]interface{}
}

func (sr *studentRecord) assignToField(key string, value interface{}) {
	if sr.extraFields == nil {
		sr.extraFields = make(map[string]interface{})
	}
	sr.extraFields[key] = value
}

func (sr studentRecord) getValue(key string) (interface{}, bool) {
	v, ok := sr.extraFields[key]
	return v, ok
}

type keyValue struct {
	key   string
	value interface{}
}

func main() {
	sr := studentRecord{1, "Bob", 75, nil}

	fieldName := "someValue"
	value, ok := sr.getValue(fieldName)
	fmt.Printf("Found value for field %s: %t\tValue: %v\n", fieldName, ok, value)

	keyValuePairs := []keyValue{
		keyValue{"studentRecordTitle", "Title"},
		keyValue{"studentRecordYear", 1987},
		keyValue{"studentRecordAudit", true},
	}

	for _, pair := range keyValuePairs {
		key, value := pair.key, pair.value
		sr.assignToField(key, value)
		v, ok := sr.getValue(key)
		fmt.Printf("Found value for field %s: %t\tValue: %v\n", key, ok, v)
	}
}
