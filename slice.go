package dot

import (
	"fmt"
	"reflect"
	"strconv"
)

// inSlice inserts a value into a slice along the specified path
func (d *Dot) inSlice(innerObj reflect.Value, currentPath string, parts []string) error {
	// Getting the slice index
	index, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf(`invalid value "%s" as a slice index`, parts[0])
	}

	// Initialising a new slice
	if innerObj.IsNil() {
		innerObj.Set(reflect.MakeSlice(innerObj.Type(), 0, 0))
	}

	// If the index is greater than -1, it means that the value
	// at the specified index must be replaced
	if index > -1 {
		if innerObj.Len() <= index {
			return fmt.Errorf("index %d out of range in path %s", index, currentPath)
		}

		// Accessing the value of this index
		innerObj = innerObj.Index(index)

		// Insert a value along the specified path
		return d.insert(innerObj, currentPath, parts[1:], Slice)
	}

	// Index -1 means that the value should be inserted at the end of the slice
	return d.appendValue(innerObj, currentPath, parts[1:])
}

func (d *Dot) appendValue(innerObj reflect.Value, currentPath string, remainingParts []string) error {
	value := reflect.New(innerObj.Type().Elem()).Elem()

	if err := d.insert(value, currentPath, remainingParts, Slice); err != nil {
		return err
	}

	innerObj.Set(reflect.Append(innerObj, value))

	return nil
}
