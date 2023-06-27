package dot

import (
	"fmt"
	"github.com/goccy/go-reflect"
	"strconv"
)

// inArray inserts the data into the array at the specified path
func (d *Dot) inArray(innerObj reflect.Value, currentPath string, parts []string) error {
	// Getting the array index
	index, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf(`invalid value "%s" as an array index`, parts[0])
	}

	// Getting the array index
	if innerObj.Len() <= index || index < 0 {
		return fmt.Errorf(
			"index %d out of range in path %s of type %s",
			index, currentPath, innerObj.Type(),
		)
	}

	// Create a variable that matches the type of value
	value := reflect.New(innerObj.Type().Elem()).Elem()

	// Recursively insert the value in the path indicated
	if err = d.insert(value, currentPath, parts[1:], Array); err != nil {
		return err
	}

	// Insert a value at the specified array index
	innerObj.Index(index).Set(value)

	return nil
}
