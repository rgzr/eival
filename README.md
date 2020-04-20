# eival

*eival* is a Go utility package for converting your `interface{}` to your desired basic type through automatic type assertion against all its compatible types and conversion when necessary.

For example, retrieving a `float64`:
* from a `string` will use `strconv.ParseFloat()`
* from a `int` will convert it directly
* from a `bool` will return `1.0` or `0.0`
* from a `date` will return `strconv`
* from a `map[string]interface{}` will fail

# Usage

```golang
package main

import (
    ev "github.com/rgzr/eival"
    "time"
    "fmt"
)

func main() {
    var m = map[string]interface{}{
        "a": time.Now(),
        "b": "123",
        "c": "myNumber",
        "d": "myOtherNumber",

    }

    // The StringZ(), IntZ()... functions return the value on success.
    fmt.Println(ev.StringZ(123)+"!!") // prints "123!!"
    fmt.Println(ev.StringZ(m["a"])) // prints "2009-11-10T23:00:00Z"
    fmt.Println(ev.IntZ(m["b"])) // prints "123"
    // They return the zero value when conversion is not possible.
    fmt.Println(ev.IntZ(m["c"])) // prints 0
    
    // The String(), Int()... functions return the value and a boolean to check if the conversion failed
    number, ok := ev.Int(m["d"])
    if !ok {
        fmt.Println("not yet") // will print "not yet"
    }
    fmt.Println(number) // will print the zero value

    // The MapZI(), SliceZI()... functions (only available to a few types, like slices or maps) return an initialized, zero-length object if the conversion fails or the underlying object was nil.
    var myNilMap map[string]interface{}
    fmt.Println(ev.MapZI(myNilMap) == nil) // will print false 
}
```

# Notes

Greatly inspired by github.com/jaracil/ei but taking a different approach.

For a full list of the possible conversions, take a look at the source code, as it's almost self-explanatory.