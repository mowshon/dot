# Connect.the.Dots, Control.your.Data!
`dot` is a versatile and powerful Go package aimed to simplify and streamline data manipulations in complex structures. This robust package offers you the ability to insert values into struct fields, arrays, slices, maps, and even channels using a dot-separated path, providing an intuitive and concise way to navigate and alter nested data structures in Go. From nested struct fields to dynamic collections, and from map keys of different types to channels, `dot` delivers an incredible amount of flexibility and convenience to your Go programming.

## Features
* **Dot-separated Path Insertion**: Insert values into struct fields, arrays, slices, maps, and channels using a dot-separated path like `Field1.Field2.Field3`.
* **Nested Struct Support**: Capable of handling complex, deeply nested Go structs.
* **Flexible Value Type**: Accepts any Go value types, allowing for wide-ranging manipulations.
* **Support for Collection Types**: This package not only works with struct fields but also supports insertion into `slices`, `arrays`, `maps`, and even `channels`, making it a truly versatile tool for manipulating data structures in Go.
* **Map Key Placeholders**: Enables the use of placeholders for `map keys`, providing a flexible way to interact with Go maps.
* **Array and Slice Indexing**: Replace values in an array or a slice by specifying an index in the path. This feature makes it easy to modify specific elements without having to iterate over the entire data structure.
* **Slice Appending**: Easily add a new value at the end of a slice by specifying `-1` in the path. This intuitive feature simplifies dynamic data modification, enhancing the versatility of your Go programs.
* **Smart Value Replacing**: Whether you're working with basic types or complex data structures, `dot` offers intelligent replacing that keeps your data intact and your code clean.

Whether you are aiming to replace a specific element in a slice or add a new entry to a map, `dot` gets you there effortlessly. Backed by smart placeholder handling, the package ensures smooth operation with map keys of varied types, giving you the power to modify data structures precisely the way you want.

## Table of Contents
- [Getting Started](#getting-started)
- [Constructor Usage](#constructor-usage)
- [Insertion into Struct Fields](#insertion-into-structure-fields)
- [Working with Maps](#working-with-maps)
- [Insertion and Replacement in Slices](#insertion-and-replacement-in-slices)
- [Working with Arrays](#working-with-arrays)
- [Insertion into Channels](#insertion-into-channels)
- [Working with Different Map Keys: Placeholders](#working-with-different-map-keys-placeholders)
- [Pros, Cons and Use Cases](#pros-cons-and-use-cases)
- [Benchmark Results](#benchmark-results)

## Getting Started

Getting started with `dot` is straightforward. To use `dot` in your Go application, you will first need to download and install it. This can be done with the `go get` command, as shown below:

```
go get github.com/mowshon/dot
```

After installing the package, you can import it into your Go files with the following line of code:

```golang
import "github.com/mowshon/dot"
```

You are now ready to use `dot` to effortlessly manipulate complex data structures in Go. The upcoming sections of this documentation will guide you through the various methods and features of `dot`, each complete with illustrative examples and detailed explanations.

Stay tuned to discover the power and elegance of `dot`, your tool for precise and convenient data manipulations in Go!

## Constructor Usage

One of the first steps when working with dot is creating an instance using your target structure. This can be done using **dot**'s constructor method, `New`.

### Passing a Variable to the Constructor

`dot.New` is the constructor method that creates and returns a new `dot` instance. The method **requires one argument**: a pointer to the variable that you want to manipulate. This variable is typically a struct, but it can also be a slice, array, map, or channel.

Here is a basic example of how to create a `dot` instance:

```golang
type MyStruct struct {
    Field1 struct {
        Field2 string
    }
}

func main() {
    data := MyStruct{}
    obj, err := dot.New(&data)
    
    if err != nil {
        // handle error
    }

    // You can now use obj to manipulate data...
}
```

In the example above, we first define a struct named `MyStruct` that we want to manipulate. Inside the `main` function, we create an instance of `MyStruct`, `data`. We then pass a pointer to `data` to `dot.New`, which returns a `dot` instance (`obj`) and an error (`err`).

If `dot.New` encounters an error while creating the `dot` instance, it will return a non-nil `err`. Always check `err` to ensure that the `dot` instance was created successfully.

After successful creation, you can use the `dot` instance (`obj` in the example) to perform various operations on `data`, which will be detailed in the following sections.

With the `dot` instance in your hands, you're ready to dive into the functionalities provided by `dot`. The next sections will explore the `Insert` method and its versatility in modifying different data types and structures in Go.

## Insertion into Structure Fields

Once you've instantiated a `dot` object using the constructor, you can start leveraging its capabilities to manipulate data structures. One of the most powerful features `dot` offers is the `Insert` method. It allows you to insert values into struct fields using a dot-separated path.

### Insert Method Syntax

The `Insert` method requires two arguments:

* The first argument is a string that represents **the path to the field** where you want to insert the value. The path should be constructed as a dot-separated sequence of field names starting from the topmost struct field name.
* The second argument is the **value** you want to insert. This can be any Go data type.

Here's an example of how you can insert a value into a struct field using the `Insert` method:

```golang
type MyStruct struct {
    Field1 struct {
        Field2 string
    }
}

func main() {
    data := MyStruct{}
    obj, err := dot.New(&data)

    if err != nil {
        // handle error
    }

    err = obj.Insert("Field1.Field2", "My Value")

    if err != nil {
        // handle error
    }

    fmt.Println(data.Field1.Field2) // Prints: My Value
}
```

In this example, we're inserting the string "My Value" into the field `Field2` which is nested inside `Field1` of our struct `MyStruct`. The `Insert` method traverses the dot-separated path `"Field1.Field2"` to reach the desired location and then inserts the provided value.

If the `Insert` method encounters an error while trying to insert the value, it will return a non-nil `err`. Always check `err` to ensure that the operation was successful.

With `dot`, even the most deeply nested struct fields are just a dot-separated path away! The upcoming sections will provide insights on how to work with more complex data structures using `dot`. Stay tuned!

## Working with Maps

`dot` not only allows you to handle regular structure fields but also supports complex data types like maps. This section will guide you on how to use the `Insert` method to manipulate maps in your structures.

### Inserting a Value into a Map Field

A map is a built-in data type in Go that associates values with keys. To insert a value into a map field, you can use the `Insert` method in the same way as for struct fields. The only difference is that the path will now include a map key.

Here's an example:

```golang
type MyStruct struct {
    MyMap map[string]int
}

func main() {
    data := MyStruct{}
    obj, err := dot.New(&data)

    if err != nil {
        // handle error
    }

    err = obj.Insert("MyMap.year", 2023)

    if err != nil {
        // handle error
    }

    fmt.Println(data.MyMap["year"]) // Prints: 2023
}
```

In this example, we're inserting the integer 2023 into the map `MyMap` with the key `"year"`. The `Insert` method finds the map using the first part of the path ("MyMap") and then inserts the value at the map key specified in the path.

Again, don't forget to check the `err` returned by the `Insert` method to ensure that the operation was successful.

Being able to manipulate maps so conveniently brings you a step closer to mastering the management of complex data structures with `dot`. Next, we'll explore how to use `dot` to work with slices and arrays.

## Insertion and Replacement in Slices

`dot` is a versatile tool not only for handling struct fields and maps but also for working with dynamic collections such as slices. This section will explain how you can use the `Insert` method for modifying slices in your structures.

### Inserting a Value at the End of a Slice

To insert a new value at the end of a slice, you can use the `Insert` method and specify `-1` in the path. Here's how it works:

```golang
type Data struct {
    Title string
}

type MyStruct struct {
    Field3 []Data
}

func main() {
    data := MyStruct{}
    obj, err := dot.New(&data)

    if err != nil {
        // handle error
    }

    err = obj.Insert("Field3.-1", Data{Title: "new Title"})

    if err != nil {
        // handle error
    }

    fmt.Println(data.Field3[0].Title) // Prints: new Title
}
```

In the above example, `Data{Title: "new Title"}` is added to the end of the slice `Field3`. The `-1` in the path `"Field3.-1"` indicates that the new value should be appended to the slice.

### Replacing a Value by Slice Index

To replace a value at a specific index in a slice, you can use the `Insert` method and specify the index in the path. Here's an example:

```golang
type Data struct {
    Title string
}

type MyStruct struct {
    Field3 []Data
}

func main() {
    data := MyStruct{Field3: []Data{{Title: "old Title"}}}
    obj, err := dot.New(&data)

    if err != nil {
        // handle error
    }

    err = obj.Insert("Field3.0.Title", "replace Title")

    if err != nil {
        // handle error
    }

    fmt.Println(data.Field3[0].Title) // Prints: replace Title
}
```

In this example, the string `"replace Title"` replaces the string at the specified index in the slice `Field3`. The index `0` in the path `"Field3.0.Title"` determines the position of the slice element that will be modified.

With `dot`, it's easy to modify slices in your structures. Next, we'll see how to use `dot` to work with arrays.

## Working with Arrays

Arrays in Go are fixed-length sequences of items of the same type. `dot` provides the functionality to modify array elements using its `Insert` method. However, as arrays have a fixed length, you cannot append new elements to them, so the `-1` value in the path is not applicable.

Here's an example of how to replace a value in an array using `dot`:

```golang
type MyStruct struct {
    Field5 [3]int
}

func main() {
    data := MyStruct{}
    obj, _ := dot.New(&data)

    err := obj.Insert("Field5.1", 2023)

    if err != nil {
        // handle error
    }

    fmt.Println(data.Field5[1]) // Prints: 2023
}
```

In this example, `Field5` is an array of integers of size 3. The `Insert` method replaces the integer at index 1 with 2023. The index `1` in the path `"Field5.1"` determines the position of the array element that will be modified.

Using `dot`, it's straightforward to modify arrays in your structures, giving you an extra tool in your belt for handling complex data structures in Go.

## Insertion into Channels

Channels are a powerful feature in Go that provide a way for two goroutines to communicate with each other and synchronize their execution. `dot` gives you the ability to insert values into channels with its `Insert` method.

### Inserting a Value into a Channel

The Insert method can be used to send a value into a channel field in your structures. The following example demonstrates how to do this:

```golang
type MyStruct struct {
    FieldChannel chan string
}

func main() {
    data := MyStruct{}
    obj, err := dot.New(&data)

    if err != nil {
        // handle error
    }

    go func() {
        err = obj.Insert("FieldChannel", "value for channel")

        if err != nil {
            // handle error
        }
    }()

    message := <-data.FieldChannel
    fmt.Println(message) // Prints: value for channel
}
```

In the above example, `"value for channel"` is sent to the `FieldChannel` channel. The `Insert` method takes care of the necessary operations to send the value into the channel. This operation is equivalent to the Go code:

```golang
data.FieldChannel <- "value for channel"
```

This concludes the basic usage of the `Insert` method in different contexts. Next, we'll discuss an advanced feature of `dot` that allows you to deal with map keys of different types using **placeholders**.

## Working with Different Map Keys: Placeholders

Working with maps where keys are not strings can be tricky. But don't worry, `dot` has you covered. `dot` provides the ability to define placeholders for map keys of different types.

### Defining and Replacing Placeholders

A placeholder in `dot` is a representation of a map key of any type. To define a placeholder, you can use the `Replace` method.

Here's how you can define and use placeholders in `dot`:

```golang
type Key string
const UniqueID Key = "Some"

type MyStruct struct {
    Field4 map[Key]string
}

func main() {
    data := MyStruct{}
    obj, _ := dot.New(&data)

    // Define placeholder
    obj.Replace("First", UniqueID)

    // Use placeholder in path
    err := obj.Insert("Field4.First", "value for map")

    if err != nil {
        // handle error
    }

    fmt.Println(data.Field4[UniqueID]) // Prints: value for map
}
```

In this example, `First` is defined as a placeholder for `UniqueID` using the `Replace` method. Once the placeholder is defined, you can use it in the path string passed to the `Insert` method.

When the `Insert` method encounters a placeholder in the path, it replaces the placeholder with its corresponding map key before inserting the value. This operation is equivalent to the Go code:

```golang
data.Field4[UniqueID] = "value for map"
```

This concludes our discussion on placeholders in `dot`, rounding out your understanding of `dot`'s powerful features for dealing with complex data structures.

## Pros, Cons and Use Cases

While the `dot` package provides great flexibility and convenience when working with complex data structures in Go, there are some considerations and potential disadvantages to keep in mind:

* **Performance Impact**: The use of reflection, which `dot` relies on to access and manipulate data structures, can be slower compared to direct field access and manipulation. This might not be an issue for smaller applications, but for high-performance applications where every millisecond counts, this could be a potential drawback.
* **Error Handling**: The package returns an error if the path is not found or if there is a type mismatch. While this is good for catching mistakes, it does mean that you need to handle these errors properly, which could add additional complexity to your code.
* **Loss of Type Safety**: One of the benefits of Go is its strong type system which can catch many errors at compile time. However, with the `dot` package, since you're specifying paths as strings and using reflection, some errors can only be caught at runtime, which may lead to increased debugging time.

Despite these potential drawbacks, the `dot` package can be **incredibly useful** in various situations:

* **Data Transformations**: If you frequently need to transform complex data structures (e.g., nested structs, maps, arrays), the `dot` package can simplify this process significantly.
* **Configuration Management**: If your Go application works with complex configuration data stored in nested structures, `dot` can make it easier to retrieve and update configuration values.
* **JSON Path-like Operations**: If you need to perform JSONPath-like operations but on Go data structures, `dot` provides a similar functionality.

As always, the decision to use a package like `dot` should be based on your specific use case, taking into account the trade-offs between readability, maintainability, performance, and development speed.

## Benchmark Results

We conducted a benchmark comparison between the `dot` package and the native Golang style for data insertion into different data types in a 5-level nested structure. Here are the results:

```
BenchmarkDotInsert-12       	  320488	      3536 ns/op
BenchmarkNativeInsert-12    	442920729	         2.661 ns/op
```

### Interpretation

* `BenchmarkDotInsert-12` is the function that tests the performance of our `dot` package. The benchmark was able to perform `320,488` iterations of the function in the default time (1 second), with each operation taking `3,536` nanoseconds (ns) or `3.536` microseconds (µs).
* `BenchmarkNativeInsert-12` is the function that tests the performance of the equivalent native Golang operations. The benchmark was able to perform a staggering `442,920,729` iterations in the default time, with each operation taking only `2.661` nanoseconds (ns).

## Conclusion

Based on the benchmark results, it's evident that the native Golang operations are faster than the operations performed by the `dot` package. However, speed isn't everything when it comes to code. While the native Golang style may be faster, it can also be more verbose and complex, especially when dealing with deeply nested data structures.

On the other hand, the `dot` package, although a bit slower, provides a cleaner, simpler syntax that is easier to read and maintain. It also handles errors robustly, which can be a major advantage in many scenarios. Therefore, it's recommended to use the `dot` package when you prioritize code simplicity and maintainability over execution speed.
