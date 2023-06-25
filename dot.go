package dot

import (
	"fmt"
	"reflect"
	"strings"
)

// Repetitive error message
var errUnknownPath = "unknown path: %s"

// Scenario is a type to define a scenario for
// further action depending on the type
type Scenario uint

// Basic scenarios depending on type
const (
	Array Scenario = iota
	Channel
	Map
	Slice
	Var
)

// Each scenario has its own error message when the type does not match
var errMsg = map[Scenario]string{
	Array:   "a %s type array cannot contain a %s type value in path %s",
	Channel: "channel of type %s cannot contain a value of type %s in path %s",
	Slice:   "a slice of type %s cannot contain a value of type %s in path %s",
	Map:     "the map value is of type %s and cannot contain a value of type %s in path %s",
	Var:     "type %s cannot contain a value of type %s in path %s",
}

// Dot provides the necessary methods for manipulating
// the value from the data types provided
type Dot struct {
	Object       reflect.Value  // Object is the reflection interface of the object provided for manipulation
	Content      any            // Content is the value to be inserted in the specified path
	Placeholders map[string]any // Placeholders contains substitutes by name for specific map key types
}

// New initialises a new structure with the necessary data for value manipulation
// The object provided must be a pointer
func New(obj any) (*Dot, error) {
	innerObj := reflect.ValueOf(obj)
	if innerObj.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("expected a pointer")
	}

	return &Dot{
		Object:       innerObj.Elem(),
		Placeholders: make(map[string]any),
	}, nil
}

// Replace creates a new placeholder by replacing the map key with a specific value
//
// e.g. map path: Insert("Data.mapKey", 1), Replace("mapKey", SpecificConstant)
// it is equivalent to: Data[SpecificConstant] = 1
func (d *Dot) Replace(key string, value any) {
	d.Placeholders[key] = value
}

// Insert receives a specific path separated by dots and the value to be inserted into that path
func (d *Dot) Insert(path string, content any) error {
	// Save the content in the structure
	d.Content = content

	// Separate the received path by a point
	parts := strings.Split(path, ".")
	if path == "" {
		parts = []string{}
	}

	return d.insert(d.Object, "", parts, Var)
}

// insert is called recursively to insert a value into the specified path
func (d *Dot) insert(innerObj reflect.Value, previousPath string, parts []string, source Scenario) error {
	// Preparing the current path on the current segment
	currentPath := preparePath(previousPath, parts)
	for index, fieldName := range parts {
		// Preparing a value insertion path
		currentPath = preparePath(previousPath, parts[:index+1])

		// Removing segments already traversed from the path
		remainingParts := parts[index:]

		// Determine the type of current path segment
		switch innerObj.Kind() {
		case reflect.Map:
			err := d.inMap(innerObj, currentPath, remainingParts)

			if err != nil {
				return err
			}

			return nil
		case reflect.Slice:
			err := d.inSlice(innerObj, currentPath, remainingParts)

			if err != nil {
				return err
			}

			return nil
		case reflect.Array:
			err := d.inArray(innerObj, currentPath, remainingParts)
			if err != nil {
				return err
			}

			return nil
		case reflect.Struct:
			innerObj = innerObj.FieldByName(fieldName)

			// The value is inserted into the channel immediately,
			// so we call the method to insert the value into the channel
			if innerObj.Kind() == reflect.Chan {
				return d.inChannel(innerObj, currentPath, remainingParts)
			}
		default:
			// If it is logical to already insert a value in the specified path,
			// but the path has not yet ended, it means that the path is specified incorrectly
			if len(remainingParts) != 0 {
				return fmt.Errorf(errUnknownPath, currentPath)
			}
		}

		if innerObj.Kind() == reflect.Invalid {
			return fmt.Errorf(errUnknownPath, currentPath)
		}
	}

	return set(innerObj, currentPath, d.Content, source)
}

// set is the final step for inserting a value on the specified path
func set(innerObj reflect.Value, currentPath string, content any, source Scenario) error {
	value := reflect.ValueOf(content)

	// Checking for type matching
	if innerObj.Type() != value.Type() {
		return fmt.Errorf(
			errMsg[source], innerObj.Type(), value.Type(), currentPath,
		)
	}

	innerObj.Set(value)
	return nil
}
