package goconfig

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	cfgTag        = "cfg"
	cfgDefaultTag = "cfg-default"
)

// Load retrieves the value for struct keys from Provider
// cfg tag defines the value key to lookup
// cfg-default tag defines the default value if value not found by provider
// if strict parameter is true then the process stops if loader cannot find a key and
// no default value provided
func Load(model interface{}, provider Provider, strict bool) error {
	return load(model, provider, strict)
}

func load(model interface{}, provider Provider, strict bool) error {
	t := reflect.TypeOf(model).Elem()
	v := reflect.ValueOf(model).Elem()

	for i := 0; i < t.NumField(); i++ {
		typeField := t.Field(i)
		valueField := v.Field(i)

		if typeField.Type.Kind() == reflect.Struct {
			value := valueField.Addr().Interface()
			err := load(value, provider, strict)
			if err != nil {
				return fmt.Errorf("Unable to set value for %s: %w", typeField.Name, err)
			}
			continue
		}

		key, ok := typeField.Tag.Lookup(cfgTag)
		if !ok {
			continue
		}

		value, err := provider.Lookup(key)
		if err != nil {
			defaultValue, ok := typeField.Tag.Lookup(cfgDefaultTag)
			if !ok {
				if strict {
					return fmt.Errorf("Unable to load value for field %s: %w", typeField.Name, err)
				}
				continue
			}

			value = defaultValue
		}

		valueField = v.FieldByIndex(typeField.Index)
		err = assignValue(&valueField, typeField.Type, value)
		if err != nil {
			return fmt.Errorf("Unable to set value for field %s: %w", typeField.Name, err)
		}
	}

	return nil
}

func assignValue(reflectValue *reflect.Value, fieldType reflect.Type, value string) error {
	switch fieldType.Kind() {
	case reflect.String:
		reflectValue.SetString(value)
	case reflect.Bool:
		value, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		reflectValue.SetBool(value)
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		reflectValue.SetInt(value)
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		value, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		reflectValue.SetUint(value)
	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		reflectValue.SetFloat(value)
	case reflect.Slice:
		parts := strings.Split(value, ",")
		if len(parts) == 0 || len(parts) == 1 && parts[0] == "" {
			break
		}

		sliceType := fieldType.Elem()
		slice := reflect.MakeSlice(fieldType, len(parts), len(parts))
		for i, v := range parts {
			valueRef := slice.Index(i)
			err := assignValue(&valueRef, sliceType, v)
			if err != nil {
				return fmt.Errorf("Unable to set index %d : %w", i, err)
			}
		}

		reflectValue.Set(slice)

	default:
		return fmt.Errorf("Could not set value for %s type", fieldType.Name())
	}
	return nil
}
