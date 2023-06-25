package dot_test

import (
	"github.com/mowshon/dot"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Part struct {
	Slug  string
	Count int
	User  map[string]string
}

type Info struct {
	Title string
	Pages map[string]float64
	Pipe  chan Part
}

type Key uint

const (
	FirstKey Key = iota
	SecondKey
)

type Data struct {
	More Info
	A    map[string]Info
	B    map[string]string
	C    map[string]map[string]int
	D    map[string]map[string]Info
	E    []int
	F    []Info
	G    [3]int
	H    chan Info
	I    chan int
	J    [3]Info
	K    map[uint64]string
	L    map[Key]string
	M    map[[2]int]string
}

func TestInvalidType(t *testing.T) {
	data := Data{}
	obj, err := dot.New(&data)
	assert.Nil(t, err)

	if err := obj.Insert("G.0", 15.5); assert.Error(t, err) {
		assert.ErrorContains(t, err, "a int type array cannot contain a float64 type value in path G.0")
	}

	if err := obj.Insert("I", "abc"); assert.Error(t, err) {
		assert.ErrorContains(t, err, "channel of type int cannot contain a value of type string in path I")
	}

	if err := obj.Insert("F.-1", 100); assert.Error(t, err) {
		assert.ErrorContains(t, err, "a slice of type dot_test.Info cannot contain a value of type int in path F.-1")
	}

	if err := obj.Insert("B.key", 55); assert.Error(t, err) {
		assert.ErrorContains(t, err, "the map value is of type string and cannot contain a value of type int in path B.key")
	}

	if err := obj.Insert("More", -1); assert.Error(t, err) {
		assert.ErrorContains(t, err, "type dot_test.Info cannot contain a value of type int in path More")
	}
}

func TestNewFailure(t *testing.T) {
	data := Data{}
	if obj, err := dot.New(data); assert.Error(t, err) {
		assert.ErrorContains(t, err, "expected a pointer")
		assert.Nil(t, obj)
	}
}

func TestUnknownPlaceholder(t *testing.T) {
	data := Data{}
	obj, err := dot.New(&data)
	assert.Nil(t, err)

	// Placeholder
	obj.Replace("ArrayAsKey", [2]int{0, 0})

	if err := obj.Insert("M.ArrayAsKey", "array-as-key"); assert.Nil(t, err) {
		assert.Exactly(t, "array-as-key", data.M[[2]int{0, 0}])
	}

	if err := obj.Insert("M.XYZ", "array-as-key"); assert.Error(t, err) {
		assert.ErrorContains(t, err, `unknown placeholder of type [2]int as map key in path "M.XYZ"`)
	}
}

func TestSpecificMapKeys(t *testing.T) {
	a := make(map[int32]string)
	if obj, err := dot.New(&a); assert.Nil(t, err) {
		if err := obj.Insert("5", "int-string"); assert.Nil(t, err) {
			assert.Exactly(t, "int-string", a[5])
		}
	}

	b := make(map[float32]string)
	if obj, err := dot.New(&b); assert.Nil(t, err) {
		if err := obj.Insert("5,2", "float32-string"); assert.Nil(t, err) {
			assert.Exactly(t, "float32-string", b[5.2])
		}
	}

	c := make(map[uint8]string)
	if obj, err := dot.New(&c); assert.Nil(t, err) {
		if err := obj.Insert("3", "uint8-string"); assert.Nil(t, err) {
			assert.Exactly(t, "uint8-string", c[3])
		}

		if err := obj.Insert("-3", "uint8-string"); assert.Error(t, err) {
			assert.ErrorContains(t, err, `the map key has an invalid key-value "-3" in path "-3" of type uint8`)
		}
	}

	d := make(map[bool]string)
	if obj, err := dot.New(&d); assert.Nil(t, err) {
		if err := obj.Insert("true", "bool-string"); assert.Nil(t, err) {
			assert.Exactly(t, "bool-string", d[true])
		}

		if err := obj.Insert("xyz", "bool-string"); assert.Error(t, err) {
			assert.ErrorContains(t, err, `the map key has an invalid key-value "xyz" in path "xyz" of type bool`)
		}
	}

	e := make(map[complex64]string)
	var index complex64 = complex(3, -5)
	if obj, err := dot.New(&e); assert.Nil(t, err) {
		if err := obj.Insert("(3-5i)", "complex64-string"); assert.Nil(t, err) {
			assert.Exactly(t, "complex64-string", e[index])
		}
	}
}

func TestMapPlaceholder(t *testing.T) {
	data := Data{}
	if obj, err := dot.New(&data); assert.Nil(t, err) {
		obj.Replace("FirstKey", FirstKey)
		obj.Replace("SecondKey", SecondKey)
		obj.Replace("ThirdKey", "key")

		if err := obj.Insert("L.FirstKey", "FirstValue"); assert.Nil(t, err) {
			assert.Exactly(t, "FirstValue", data.L[FirstKey])
		}

		if err := obj.Insert("L.SecondKey", "SecondValue"); assert.Nil(t, err) {
			assert.Exactly(t, "SecondValue", data.L[SecondKey])
		}

		if err := obj.Insert("L.InvalidKey", "SecondValue"); assert.Error(t, err) {
			assert.ErrorContains(
				t, err,
				`the map key has an invalid key-value "InvalidKey" in path "L.InvalidKey" of type dot_test.Key`,
			)
		}

		if err := obj.Insert("L.ThirdKey", "SecondValue"); assert.Error(t, err) {
			assert.ErrorContains(
				t, err,
				`the map key type is dot_test.Key you cannot use the placeholder of type string in path "L.ThirdKey"`,
			)
		}
	}
}

func TestInsertInPrimitiveDataTypes(t *testing.T) {
	var a string
	if obj, err := dot.New(&a); assert.Nil(t, err) {
		if err := obj.Insert("", "value"); assert.Nil(t, err) {
			assert.Exactly(t, "value", a)
		}
	}

	var info map[string]string
	if obj, err := dot.New(&info); assert.Nil(t, err) {
		if err := obj.Insert("key", "value"); assert.Nil(t, err) {
			assert.Exactly(t, "value", info["key"])
		}
	}

	var list []string
	if obj, err := dot.New(&list); assert.Nil(t, err) {
		if err := obj.Insert("-1", "value"); assert.Nil(t, err) {
			assert.Exactly(t, "value", list[0])
		}
	}
}

func TestInChannel(t *testing.T) {
	data := Data{}
	obj, err := dot.New(&data)
	assert.Nil(t, err)

	if err := obj.Insert("I", 99); assert.Nil(t, err) {
		info := <-data.I
		assert.Exactly(t, 99, info)
	}

	if err := obj.Insert("I.Field", 99); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: I.Field")
	}

	insert := "Title from Channel"
	if err := obj.Insert("H.Title", insert); assert.Nil(t, err) {
		info := <-data.H
		assert.Exactly(t, insert, info.Title)
	}

	if err := obj.Insert("H.Title.Field", 1); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: H.Title.Field")
	}

	if err := obj.Insert("H.Pipe.User.role", "admin"); assert.Nil(t, err) {
		content := <-data.H
		pipe := <-content.Pipe

		assert.Exactly(t, "admin", pipe.User["role"])
	}

	if err := obj.Insert("H.Pipe.User.role.Field", "admin"); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: H.Pipe.User.role.Field")
	}
}

func TestInMap(t *testing.T) {
	data := Data{}
	obj, err := dot.New(&data)
	assert.Nil(t, err)

	insert := "Title from map"
	if err := obj.Insert("A.first.Title", insert); assert.Nil(t, err) {
		assert.Exactly(t, insert, data.A["first"].Title)
	}

	if err := obj.Insert("C.first.second", 55); assert.Nil(t, err) {
		assert.Exactly(t, 55, data.C["first"]["second"])
	}

	if err := obj.Insert("C.first.second", 100); assert.Nil(t, err) {
		assert.Exactly(t, 100, data.C["first"]["second"])
	}

	if err := obj.Insert("C.first.second.third", 55); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: C.first.second.third")
	}

	if err := obj.Insert("D.first.second.Title", "Title from Map"); assert.Nil(t, err) {
		assert.Exactly(t, "Title from Map", data.D["first"]["second"].Title)
	}

	if err := obj.Insert("D.first.second.Field", 55); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: D.first.second.Field")
	}

	if err := obj.Insert("A.first.Title1", insert); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: A.first.Title1")
	}
}

func TestInStruct(t *testing.T) {
	data := Data{}
	obj, err := dot.New(&data)
	assert.Nil(t, err)

	insert := "Title from More"
	if err := obj.Insert("More.Title", insert); assert.Nil(t, err) {
		assert.Exactly(t, insert, data.More.Title)
	}

	if err := obj.Insert("More.Pages.total", float64(100)); assert.Nil(t, err) {
		assert.Exactly(t, float64(100), data.More.Pages["total"])
	}

	if err := obj.Insert("More.Title.Unknown", insert); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: More.Title.Unknown")
	}
}

func TestArray(t *testing.T) {
	data := Data{}
	obj, err := dot.New(&data)
	assert.Nil(t, err)

	if err := obj.Insert("G.1", 5); assert.Nil(t, err) {
		assert.Exactly(t, data.G, [3]int{0, 5, 0})
	}

	if err := obj.Insert("G", [3]int{1, 2, 3}); assert.Nil(t, err) {
		assert.Exactly(t, data.G, [3]int{1, 2, 3})
	}

	if err := obj.Insert("G.0", 0); assert.Nil(t, err) {
		assert.Exactly(t, data.G, [3]int{0, 2, 3})
	}

	if err := obj.Insert("G.0.Field", 0); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: G.0.Field")
	}

	if err := obj.Insert("G.4", 0); assert.Error(t, err) {
		assert.ErrorContains(t, err, "index 4 out of range in path G.4 of type [3]int")
	}

	if err := obj.Insert("G.-1", 0); assert.Error(t, err) {
		assert.ErrorContains(t, err, "index -1 out of range in path G.-1 of type [3]int")
	}

	insert := "Array Title"
	if err := obj.Insert("J.1.Title", insert); assert.Nil(t, err) {
		assert.Exactly(t, data.J, [3]Info{
			{}, {Title: insert}, {},
		})
	}

	if err := obj.Insert("G.a", 0); assert.Error(t, err) {
		assert.ErrorContains(t, err, `invalid value "a" as an array index`)
	}
}

func TestSlice(t *testing.T) {
	data := Data{}
	obj, err := dot.New(&data)
	assert.Nil(t, err)

	if err := obj.Insert("E", []int{1, 2}); assert.Nil(t, err) {
		assert.Exactly(t, data.E, []int{1, 2})
	}

	if err := obj.Insert("E.-1", 3); assert.Nil(t, err) {
		assert.Exactly(t, data.E, []int{1, 2, 3})
	}

	if err := obj.Insert("E.0", 0); assert.Nil(t, err) {
		assert.Exactly(t, data.E, []int{0, 2, 3})
	}

	if err := obj.Insert("E.0.Field", 4); assert.Error(t, err) {
		assert.ErrorContains(t, err, "unknown path: E.0.Field")
	}

	if err := obj.Insert("E.4", 4); assert.Error(t, err) {
		assert.ErrorContains(t, err, "index 4 out of range in path E.4")
	}

	if err := obj.Insert("F", []Info{{Title: "F-First"}}); assert.Nil(t, err) {
		assert.Exactly(t, data.F[0].Title, "F-First")
	}

	if err := obj.Insert("F.-1.Field", "Value"); assert.Error(t, err) {
		assert.ErrorContains(t, err, "F.-1.Field")
	}

	if err := obj.Insert("F.-1.Title", "F-Second"); assert.Nil(t, err) {
		assert.Len(t, data.F, 2)
		assert.Exactly(t, data.F[0].Title, "F-First")
		assert.Exactly(t, data.F[1].Title, "F-Second")
	}

	if err := obj.Insert("F.0.Title", "F-New"); assert.Nil(t, err) {
		assert.Len(t, data.F, 2)
		assert.Exactly(t, data.F[0].Title, "F-New")
	}

	if err := obj.Insert("E.a", 0); assert.Error(t, err) {
		assert.ErrorContains(t, err, `invalid value "a" as a slice index`)
	}
}
