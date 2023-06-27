package dot

import (
	"errors"
	"fmt"
	"github.com/goccy/go-reflect"
	"strconv"
	"strings"
)

// parseType checks if the key does not match one of the primitive types
func parseType(kind reflect.Kind, value string) (any, error) {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.ParseInt(value, 10, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.ParseUint(value, 10, 64)
	case reflect.Float32, reflect.Float64:
		// This is a forced measure, as the path to the value is divided by a dot
		newValue := strings.ReplaceAll(value, ",", ".")
		return strconv.ParseFloat(newValue, 64)
	case reflect.Complex64, reflect.Complex128:
		return strconv.ParseComplex(value, 128)
	case reflect.Bool:
		return strconv.ParseBool(value)
	case reflect.String, reflect.Interface:
		return value, nil
	default:
		return nil, errors.New("unknown type")
	}
}

// keyValue prepares the value of the map key itself
func (d *Dot) keyValue(key reflect.Value, value string, currentPath string) (any, error) {
	// Trying to find the appropriate key in placeholders
	if placeholder, ok := d.Placeholders[value]; ok {
		refPlaceholder := reflect.ValueOf(placeholder)
		if refPlaceholder.Type() != key.Type() {
			return nil, fmt.Errorf(
				`the map key type is %s you cannot use the placeholder of type %s in path "%s"`,
				key.Type(), refPlaceholder.Type(), currentPath,
			)
		}

		return placeholder, nil
	}

	// Trying to catch a primitive key type
	result, err := parseType(key.Kind(), value)

	if err != nil {
		// It was not possible to convert the value into a primitive map key type
		if strings.Contains(err.Error(), "strconv") {
			return nil, fmt.Errorf(
				`the map key has an invalid key-value "%s" in path "%s" of type %s`,
				value, currentPath, key.Type(),
			)
		}

		// If no primitive type was found, it means there is no key placeholder
		return nil, fmt.Errorf(
			`unknown placeholder of type %s as map key in path "%s"`,
			key.Type(), currentPath,
		)
	}

	return result, nil
}

// prepareKey creates a key of the appropriate type for the map
func (d *Dot) prepareKey(innerObj reflect.Value, value string, currentPath string) (reflect.Value, error) {
	thisMap := reflect.TypeOf(innerObj.Interface())
	key := reflect.New(thisMap.Key()).Elem()

	// Prepare the contents of the key to match its type
	content, err := d.keyValue(key, value, currentPath)
	if err != nil {
		return reflect.Value{}, err
	}

	// Convert the type so that it exactly matches the required type
	keyValue := reflect.ValueOf(content)
	if keyValue.Type().ConvertibleTo(key.Type()) {
		keyValue = keyValue.Convert(key.Type())
	}

	// Inserting the key value
	key.Set(keyValue)

	return key, nil
}

// inMap inserts a value into the map along the indicated path
func (d *Dot) inMap(innerObj reflect.Value, currentPath string, parts []string) error {
	// Create a map if one does not exist
	if innerObj.IsNil() {
		innerObj.Set(reflect.MakeMap(innerObj.Type()))
	}

	// The map key can be of any type. As with the value,
	// we must first create a key of the correct type.
	key, err := d.prepareKey(innerObj, parts[0], currentPath)
	if err != nil {
		return err
	}

	// Initialise the value of the corresponding map type
	value := reflect.New(innerObj.Type().Elem()).Elem()

	// Insert the value recursively into the variable we just created
	if err := d.insert(value, currentPath, parts[1:], Map); err != nil {
		return err
	}

	// Having a key of the appropriate type, specify its value
	innerObj.SetMapIndex(key, value)

	return nil
}
