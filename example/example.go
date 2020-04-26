package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/rgzr/eival"
)

const (
	INT64 = iota
	INT
	INT32
	INT16
	INT8
	UINT64
	UINT
	UINT32
	UINT16
	UINT8
	BYTE
	FLOAT32
	FLOAT64
	BOOL
	STRING
	TIME
	BYTES
	SLICE
	MAP
)

var names = map[int]string{
	INT64:   "INT64",
	INT:     "INT",
	INT32:   "INT32",
	INT16:   "INT16",
	INT8:    "INT8",
	UINT64:  "UINT64",
	UINT:    "UINT",
	UINT32:  "UINT32",
	UINT16:  "UINT16",
	UINT8:   "UINT8",
	BYTE:    "BYTE",
	FLOAT32: "FLOAT32",
	FLOAT64: "FLOAT64",
	BOOL:    "BOOL",
	STRING:  "STRING",
	TIME:    "TIME",
	BYTES:   "BYTES",
	SLICE:   "SLICE",
	MAP:     "MAP",
}

var types = map[int]string{
	INT64:   "int64",
	INT:     "int",
	INT32:   "int32",
	INT16:   "int16",
	INT8:    "int8",
	UINT64:  "uint64",
	UINT:    "uint",
	UINT32:  "uint32",
	UINT16:  "uint16",
	UINT8:   "uint8",
	BYTE:    "byte",
	FLOAT32: "float32",
	FLOAT64: "float64",
	BOOL:    "bool",
	STRING:  "string",
	TIME:    "time.Time",
	BYTES:   "[]byte",
	SLICE:   "[]interface{}",
	MAP:     "map[string]interface{}",
}

type Result struct {
	V  interface{}
	Ok bool
}

func R(v interface{}, ok bool) *Result {
	return &Result{V: v, Ok: ok}
}

type S []interface{}
type M map[string]interface{}

var savedDate = time.Date(2000, time.January, 2, 3, 4, 5, 6, time.UTC)
var values = map[int]S{
	INT64:   S{int64(-123456), int64(0), int64(123456)},
	INT:     S{int(-123456), int(0), int(123456)},
	INT32:   S{int32(-123456), int32(0), int32(123456)},
	INT16:   S{int16(-123), int16(0), int16(123)},
	INT8:    S{int8(-123), int8(0), int8(123)},
	UINT64:  S{uint64(654321), uint64(0), uint64(123456)},
	UINT:    S{uint(65432), uint(0), uint(123456)},
	UINT32:  S{uint32(654321), uint32(0), uint32(123456)},
	UINT16:  S{uint16(6543), uint16(0), uint16(123)},
	UINT8:   S{uint8(213), uint8(0), uint8(123)},
	BYTE:    S{byte(128), byte(0), byte(64)},
	FLOAT32: S{float32(-123.456), float32(0.0), float32(123.456)},
	FLOAT64: S{float64(-123.456), float64(0.0), float64(123.456)},
	BOOL:    S{true, false},
	STRING:  S{"", "123456", "mystring", "true", "-123.456", "2000-01-02T03:04:05.000000006Z", "2000-01-02T03:04:05Z", "Sun Jan  2 03:04:05 2000", "Sun Jan  2 03:04:05 UTC 2000", "Sun Jan 02 03:04:05 +0000 2000", "02 Jan 00 03:04 UTC", "02 Jan 00 03:04 +0000", "Sunday, 02-Jan-00 03:04:05 UTC", "Sun, 02 Jan 2000 03:04:05 UTC", "Sun, 02 Jan 2000 03:04:05 +0000", "2000-01-02 03:04:05", "Jan  2 03:04:05.000", "Jan  2 03:04:05", "3:04AM", "2000-01-02"},
	TIME:    S{time.Time{}, savedDate},
	BYTES:   S{[]byte(nil), []byte{}, []byte{0x01, 0x02, 0x03}},
	SLICE:   S{[]interface{}(nil), []interface{}{}, []interface{}{"a", 2, true}},
	MAP:     S{map[string]interface{}(nil), map[string]interface{}{}, map[string]interface{}{"a": 2, "b": false}},
}

var res = map[int]map[int][]*Result{}

func main() {
	for reader, _ := range values {
		res[reader] = map[int][]*Result{}
		for val_types, vals := range values {
			s := []*Result{}
			for _, value := range vals {
				switch reader {
				case INT64:
					r, ok := eival.Int64(value)
					s = append(s, &Result{V: r, Ok: ok})
				case INT:
					r, ok := eival.Int(value)
					s = append(s, &Result{V: r, Ok: ok})
				case INT32:
					r, ok := eival.Int32(value)
					s = append(s, &Result{V: r, Ok: ok})
				case INT16:
					r, ok := eival.Int16(value)
					s = append(s, &Result{V: r, Ok: ok})
				case INT8:
					r, ok := eival.Int8(value)
					s = append(s, &Result{V: r, Ok: ok})
				case UINT64:
					r, ok := eival.Uint64(value)
					s = append(s, &Result{V: r, Ok: ok})
				case UINT:
					r, ok := eival.Uint(value)
					s = append(s, &Result{V: r, Ok: ok})
				case UINT32:
					r, ok := eival.Uint32(value)
					s = append(s, &Result{V: r, Ok: ok})
				case UINT16:
					r, ok := eival.Uint16(value)
					s = append(s, &Result{V: r, Ok: ok})
				case UINT8:
					r, ok := eival.Uint8(value)
					s = append(s, &Result{V: r, Ok: ok})
				case BYTE:
					r, ok := eival.Byte(value)
					s = append(s, &Result{V: r, Ok: ok})
				case FLOAT32:
					r, ok := eival.Float32(value)
					s = append(s, &Result{V: r, Ok: ok})
				case FLOAT64:
					r, ok := eival.Float64(value)
					s = append(s, &Result{V: r, Ok: ok})
				case BOOL:
					r, ok := eival.Bool(value)
					s = append(s, &Result{V: r, Ok: ok})
				case STRING:
					r, ok := eival.String(value)
					s = append(s, &Result{V: r, Ok: ok})
				case TIME:
					r, ok := eival.Time(value)
					s = append(s, &Result{V: r, Ok: ok})
				case BYTES:
					r, ok := eival.Bytes(value)
					s = append(s, &Result{V: r, Ok: ok})
				case SLICE:
					r, ok := eival.Slice(value)
					s = append(s, &Result{V: r, Ok: ok})
				case MAP:
					r, ok := eival.Map(value)
					s = append(s, &Result{V: r, Ok: ok})
				}
			}
			res[reader][val_types] = s
		}
	}

	fmt.Println("var readValues = map[int]map[int]S{")
	for reader, vs := range res {
		fmt.Printf("\t%s: map[int]S{\n", names[reader])
		for value_type, vals := range vs {
			s := []string{}
			for _, value := range vals {
				if reader == TIME {
					if value.Ok {
						//t, _ := value.V.(time.Time)
						s = append(s, fmt.Sprintf("R(time.Date(%d, %d, %d, %d, %d, %d, %d, loc *Location),%t)", types[reader], value.V, value.Ok))
					} else {
						s = append(s, fmt.Sprintf("R(%#v,%t)", value.V, value.Ok))
					}
				} else {
					if value.Ok {
						s = append(s, fmt.Sprintf("R(%s(%#v),%t)", types[reader], value.V, value.Ok))
					} else {
						s = append(s, fmt.Sprintf("R(%#v,%t)", value.V, value.Ok))
					}
				}
			}
			fmt.Println(fmt.Sprintf("\t\t%s: S{ %s },", names[value_type], strings.Join(s, ", ")))
		}
		fmt.Println("\t},")
	}
	fmt.Println("}")

}
