package utils

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/imdario/mergo"
)

// MergeObjects merges the fields of two different structs. The structs can be of
// different types, but must have the same json fields.
// The `dst` must be a pointer.
func MergeObjects(dst interface{}, src interface{}) error {

	if dst != nil && reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return errors.New("dst must be a pointer")
	}

	j, err := json.Marshal(src)
	if err != nil {
		return err
	}

	newSrc := reflect.New(reflect.TypeOf(dst).Elem()).Interface()
	err = json.Unmarshal(j, &newSrc)
	if err != nil {
		return err
	}

	if err := mergo.Merge(dst, newSrc, mergo.WithOverride); err != nil {
		return err
	}

	return nil
}
