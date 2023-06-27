package dot

import (
	"github.com/goccy/go-reflect"
)

// inChannel inserts the value into the channel on the specified pat
func (d *Dot) inChannel(innerObj reflect.Value, currentPath string, parts []string) error {
	// Create a channel if it does not exist
	if innerObj.IsNil() {
		innerObj.Set(reflect.MakeChan(innerObj.Type(), 1))
	}

	// Create a variable that matches the type of value
	value := reflect.New(innerObj.Type().Elem()).Elem()
	if err := d.insert(value, currentPath, parts[1:], Channel); err != nil {
		return err
	}

	// Insert a value in the channel
	reflect.Select([]reflect.SelectCase{
		{
			Dir:  reflect.SelectSend,
			Chan: innerObj,
			Send: value,
		},
	})

	return nil
}
