package dot_test

import (
	"github.com/mowshon/dot"
	"testing"
)

type Inner struct {
	Field string
}

type Nested struct {
	MapField    map[string]Inner
	SliceField  []Inner
	StructField Inner
	ArrayField  [1]Inner
}

func BenchmarkDotInsert(b *testing.B) {
	data := Nested{}
	data.MapField = make(map[string]Inner)
	data.SliceField = make([]Inner, 1)
	obj, _ := dot.New(&data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		obj.Insert("MapField.key.Field", "My Value")
		obj.Insert("SliceField.0.Field", "My Value")
		obj.Insert("StructField.Field", "My Value")
		obj.Insert("ArrayField.0.Field", "My Value")
	}
}

func BenchmarkNativeInsert(b *testing.B) {
	data := Nested{}
	data.MapField = make(map[string]Inner)
	data.SliceField = make([]Inner, 1)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		add, _ := data.MapField["key"]
		add.Field = "My Value"

		data.SliceField[0].Field = "My Value"
		data.StructField.Field = "My Value"
		data.ArrayField[0].Field = "My Value"
	}
}
