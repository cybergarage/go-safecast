# go-safecast

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-safecast)
[![test](https://github.com/cybergarage/go-safecast/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-safecast/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-safecast.svg)](https://pkg.go.dev/github.com/cybergarage/go-safecast)

`go-safecast` is a utility function package for safe casting between primitive data types of Go. The following shows an unsafe cast example that causes cast overflow problems, but the unsafe casting is not detected at runtime and compile time.

```
package main

import (
    "fmt"
    "math"
)

func main() {
    from := math.MaxInt8 + 1
    to := int8(from)
    fmt.Printf("%v != %v", to, from)

    // Output:
    // -128 != 128
}
```

`go-safecast` provides utility functions for safe casting, and raises an error if the casting is unsafe as the following.

```
package main

import (
    "github.com/cybergarage/go-safecast/safecast"
)

func main() {
    from := math.MaxInt8 + 1
    var to int8
    err := safecast.FromInt(from, &to)
    if err := nil {
        fmt.Println(err.Error())
    }

    // Output:
    // cast error : out of range 128 > *int8
}
```

Currently, `go-safecast` provides the following safe casting functions.Please see [![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-safecast.svg)](https://pkg.go.dev/github.com/cybergarage/go-safecast) to know the functions and examples in more detail.

|Function                                    |To                                                                             |
|--------------------------------------------|-------------------------------------------------------------------------------|
|func FromInt(from int64, to any) error      | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromInt8(from int8, to any) error      | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromInt16(from int16, to any) error    | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromInt32(from int32, to any) error    | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromInt64(from int64, to any) error    | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromUint(from uint, to any) error      | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromUint8(from uint8, to any) error    | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromUint16(from uint16, to any) error  | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromUint32(from uint32, to any) error  | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromUint64(from uint64, to any) error  | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string, *bool |
|func FromFloat32(from float32, to any) error| *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string |
|func FromFloat64(from float64, to any) error| *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *string |
|func FromString(from string, to any) error  | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float64, *float32, *bool, *string *[]byte |
|func FromBool(from bool, to any) error      | *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *bool, *string |

