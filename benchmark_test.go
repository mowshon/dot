package dot_test

import (
	"github.com/mowshon/dot"
	"testing"
)

type ComplexStruct struct {
	NestedStruct struct {
		Field1 string
		Field2 int
		Field3 []string
		Field4 map[string]string
	}
	NestedMap map[string]map[string]string
}

func BenchmarkDotInsert(b *testing.B) {
	data := ComplexStruct{}
	obj, _ := dot.New(&data)

	b.Run("Insert Struct Field", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj.Insert("NestedStruct.Field1", "My Value")
		}
	})

	b.Run("Insert Map Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj.Insert("NestedMap.key1.subkey1", "test")
		}
	})

	b.Run("Insert Slice Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj.Insert("NestedStruct.Field3.0", "test")
		}
	})
}

func BenchmarkDotReplace(b *testing.B) {
	data := ComplexStruct{}
	obj, _ := dot.New(&data)

	b.Run("Replace Struct Field", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj.Insert("NestedStruct.Field1", "My Value")
			obj.Insert("NestedStruct.Field1", "New Value")
		}
	})

	b.Run("Replace Map Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj.Insert("NestedMap.key1.subkey1", "test")
			obj.Insert("NestedMap.key1.subkey1", "New Value")
		}
	})

	b.Run("Replace Slice Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj.Insert("NestedStruct.Field3.0", "test")
			obj.Insert("NestedStruct.Field3.0", "New Value")
		}
	})
}
