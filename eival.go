package eival

import (
	"math"
	"strconv"
	"time"
)

func Int64(v interface{}) (int64, bool) {
	switch n := v.(type) {
	case int64:
		return n, true
	case int:
		return int64(n), true
	case int32:
		return int64(n), true
	case uint:
		return int64(n), true
	case uint64:
		return int64(n), true
	case uint32:
		return int64(n), true
	case int16:
		return int64(n), true
	case int8:
		return int64(n), true
	case uint16:
		return int64(n), true
	case uint8:
		return int64(n), true
	case float32:
		return int64(n), true
	case float64:
		return int64(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseInt(n, 0, 64); err == nil {
			return s, true
		}
	case time.Time:
		return int64(n.Unix()), true
	}
	return 0, false
}

func Int64Z(v interface{}) int64 {
	r, _ := Int64(v)
	return r
}

func Int(v interface{}) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case int64:
		return int(n), true
	case int32:
		return int(n), true
	case uint:
		return int(n), true
	case uint64:
		return int(n), true
	case uint32:
		return int(n), true
	case int16:
		return int(n), true
	case int8:
		return int(n), true
	case uint16:
		return int(n), true
	case uint8:
		return int(n), true
	case float32:
		return int(n), true
	case float64:
		return int(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseInt(n, 0, 0); err == nil {
			return int(s), true
		}
	case time.Time:
		return int(n.Unix()), true
	}
	return 0, false
}

func IntZ(v interface{}) int {
	r, _ := Int(v)
	return r
}

func Int32(v interface{}) (int32, bool) {
	switch n := v.(type) {
	case int32:
		return n, true
	case int:
		return int32(n), true
	case int64:
		return int32(n), true
	case uint:
		return int32(n), true
	case uint32:
		return int32(n), true
	case uint64:
		return int32(n), true
	case int16:
		return int32(n), true
	case int8:
		return int32(n), true
	case uint16:
		return int32(n), true
	case uint8:
		return int32(n), true
	case float32:
		return int32(n), true
	case float64:
		return int32(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseInt(n, 0, 32); err == nil {
			return int32(s), true
		}
	case time.Time:
		return int32(n.Unix()), true
	}
	return 0, false
}

func Int32Z(v interface{}) int32 {
	r, _ := Int32(v)
	return r
}

func Int16(v interface{}) (int16, bool) {
	switch n := v.(type) {
	case int16:
		return n, true
	case int:
		return int16(n), true
	case uint16:
		return int16(n), true
	case int32:
		return int16(n), true
	case int8:
		return int16(n), true
	case uint:
		return int16(n), true
	case uint32:
		return int16(n), true
	case uint8:
		return int16(n), true
	case int64:
		return int16(n), true
	case uint64:
		return int16(n), true
	case float32:
		return int16(n), true
	case float64:
		return int16(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseInt(n, 0, 16); err == nil {
			return int16(s), true
		}
	case time.Time:
		return int16(n.Unix()), true
	}
	return 0, false
}

func Int16Z(v interface{}) int16 {
	r, _ := Int16(v)
	return r
}

func Int8(v interface{}) (int8, bool) {
	switch n := v.(type) {
	case int8:
		return n, true
	case uint8:
		return int8(n), true
	case int16:
		return int8(n), true
	case uint16:
		return int8(n), true
	case int:
		return int8(n), true
	case uint:
		return int8(n), true
	case int32:
		return int8(n), true
	case uint32:
		return int8(n), true
	case int64:
		return int8(n), true
	case uint64:
		return int8(n), true
	case float32:
		return int8(n), true
	case float64:
		return int8(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseInt(n, 0, 8); err == nil {
			return int8(s), true
		}
	case time.Time:
		return int8(n.Unix()), true
	}
	return 0, false
}

func Int8Z(v interface{}) int8 {
	r, _ := Int8(v)
	return r
}

func Uint64(v interface{}) (uint64, bool) {
	switch n := v.(type) {
	case uint64:
		return n, true
	case uint:
		return uint64(n), true
	case uint32:
		return uint64(n), true
	case int:
		return uint64(n), true
	case int64:
		return uint64(n), true
	case int32:
		return uint64(n), true
	case uint16:
		return uint64(n), true
	case uint8:
		return uint64(n), true
	case int16:
		return uint64(n), true
	case int8:
		return uint64(n), true
	case float32:
		return uint64(n), true
	case float64:
		return uint64(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseUint(n, 0, 64); err == nil {
			return s, true
		}
	case time.Time:
		return uint64(n.Unix()), true
	}
	return 0, false
}

func Uint64Z(v interface{}) uint64 {
	r, _ := Uint64(v)
	return r
}

func Uint(v interface{}) (uint, bool) {
	switch n := v.(type) {
	case uint:
		return n, true
	case uint64:
		return uint(n), true
	case uint32:
		return uint(n), true
	case int:
		return uint(n), true
	case int64:
		return uint(n), true
	case int32:
		return uint(n), true
	case uint16:
		return uint(n), true
	case uint8:
		return uint(n), true
	case int16:
		return uint(n), true
	case int8:
		return uint(n), true
	case float32:
		return uint(n), true
	case float64:
		return uint(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseUint(n, 0, 0); err == nil {
			return uint(s), true
		}
	case time.Time:
		return uint(n.Unix()), true
	}
	return 0, false
}

func UintZ(v interface{}) uint {
	r, _ := Uint(v)
	return r
}

func Uint32(v interface{}) (uint32, bool) {
	switch n := v.(type) {
	case uint32:
		return n, true
	case uint:
		return uint32(n), true
	case uint64:
		return uint32(n), true
	case int:
		return uint32(n), true
	case int32:
		return uint32(n), true
	case int64:
		return uint32(n), true
	case uint16:
		return uint32(n), true
	case uint8:
		return uint32(n), true
	case int16:
		return uint32(n), true
	case int8:
		return uint32(n), true
	case float32:
		return uint32(n), true
	case float64:
		return uint32(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseUint(n, 0, 32); err == nil {
			return uint32(s), true
		}
	case time.Time:
		return uint32(n.Unix()), true
	}
	return 0, false
}

func Uint32Z(v interface{}) uint32 {
	r, _ := Uint32(v)
	return r
}

func Uint16(v interface{}) (uint16, bool) {
	switch n := v.(type) {
	case uint16:
		return n, true
	case uint:
		return uint16(n), true
	case int16:
		return uint16(n), true
	case uint32:
		return uint16(n), true
	case uint8:
		return uint16(n), true
	case int:
		return uint16(n), true
	case int32:
		return uint16(n), true
	case int8:
		return uint16(n), true
	case uint64:
		return uint16(n), true
	case int64:
		return uint16(n), true
	case float32:
		return uint16(n), true
	case float64:
		return uint16(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseUint(n, 0, 16); err == nil {
			return uint16(s), true
		}
	case time.Time:
		return uint16(n.Unix()), true
	}
	return 0, false
}

func Uint16Z(v interface{}) uint16 {
	r, _ := Uint16(v)
	return r
}

func Uint8(v interface{}) (uint8, bool) {
	switch n := v.(type) {
	case uint8:
		return uint8(n), true
	case uint:
		return uint8(n), true
	case uint32:
		return uint8(n), true
	case uint64:
		return uint8(n), true
	case uint16:
		return uint8(n), true
	case int:
		return uint8(n), true
	case int32:
		return uint8(n), true
	case int64:
		return uint8(n), true
	case int8:
		return uint8(n), true
	case int16:
		return uint8(n), true
	case float32:
		return uint8(n), true
	case float64:
		return uint8(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseUint(n, 0, 8); err == nil {
			return uint8(s), true
		}
	case time.Time:
		return uint8(n.Unix()), true
	}
	return 0, false
}

func Uint8Z(v interface{}) uint8 {
	r, _ := Uint8(v)
	return r
}

func Byte(v interface{}) (byte, bool) {
	r, ok := Uint8(v)
	return byte(r), ok
}

func ByteZ(v interface{}) byte {
	r, _ := Byte(v)
	return r
}

func Float64(v interface{}) (float64, bool) {
	switch n := v.(type) {
	case float64:
		return n, true
	case float32:
		return float64(n), true
	case int:
		return float64(n), true
	case int32:
		return float64(n), true
	case int64:
		return float64(n), true
	case int8:
		return float64(n), true
	case int16:
		return float64(n), true
	case uint:
		return float64(n), true
	case uint32:
		return float64(n), true
	case uint64:
		return float64(n), true
	case uint8:
		return float64(n), true
		return float64(n), true
	case uint16:
		return float64(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseFloat(n, 64); err == nil {
			return s, true
		}
	case time.Time:
		return float64(n.Unix()), true
	}
	return 0, false
}

func Float64Z(v interface{}) float64 {
	r, _ := Float64(v)
	return r
}

func Float32(v interface{}) (float32, bool) {
	switch n := v.(type) {
	case float32:
		return n, true
	case float64:
		return float32(n), true
	case int:
		return float32(n), true
	case int32:
		return float32(n), true
	case int64:
		return float32(n), true
	case int8:
		return float32(n), true
	case int16:
		return float32(n), true
	case uint:
		return float32(n), true
	case uint32:
		return float32(n), true
	case uint64:
		return float32(n), true
	case uint8:
		return float32(n), true
	case uint16:
		return float32(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case string:
		if s, err := strconv.ParseFloat(n, 32); err == nil {
			return float32(s), true
		}
	case time.Time:
		return float32(n.Unix()), true
	}
	return 0, false
}

func Float32Z(v interface{}) float32 {
	r, _ := Float32(v)
	return r
}

func String(v interface{}) (string, bool) {
	switch n := v.(type) {
	case string:
		return n, true
	case []byte:
		return string(n), true
	case time.Time:
		return n.Format(time.RFC3339Nano), true
	case float64:
		return strconv.FormatFloat(float64(n), 'f', -1, 64), true
	case float32:
		return strconv.FormatFloat(float64(n), 'f', -1, 64), true
	case int:
		return strconv.FormatInt(int64(n), 10), true
	case int32:
		return strconv.FormatInt(int64(n), 10), true
	case int64:
		return strconv.FormatInt(int64(n), 10), true
	case int8:
		return strconv.FormatInt(int64(n), 10), true
	case int16:
		return strconv.FormatInt(int64(n), 10), true
	case uint:
		return strconv.FormatUint(uint64(n), 10), true
	case uint32:
		return strconv.FormatUint(uint64(n), 10), true
	case uint64:
		return strconv.FormatUint(uint64(n), 10), true
	case uint8:
		return strconv.FormatUint(uint64(n), 10), true
	case uint16:
		return strconv.FormatUint(uint64(n), 10), true
	case bool:
		return strconv.FormatBool(n), true
	}
	return "", false
}

func StringZ(v interface{}) string {
	r, _ := String(v)
	return r
}

func Time(v interface{}) (time.Time, bool) {
	switch n := v.(type) {
	case time.Time:
		return n, true
	case string:
		formats := []string{
			time.RFC3339Nano,
			time.RFC3339,
			time.ANSIC,
			time.UnixDate,
			time.RubyDate,
			time.RFC822,
			time.RFC822Z,
			time.RFC850,
			time.RFC1123,
			time.RFC1123Z,
			"2006-01-02 15:04:05", // UTC
			time.StampMilli,
			time.Stamp,
			time.Kitchen,
			"2006-01-02", // UTC
		}
		for _, f := range formats {
			t, err := time.Parse(f, n)
			if err == nil {
				return t, true
			}
		}
	case float64:
		sec, dec := math.Modf(n)
		return time.Unix(int64(sec), int64(dec*(1e9))), true
	case float32:
		sec, dec := math.Modf(float64(n))
		return time.Unix(int64(sec), int64(dec*(1e9))), true
	case int:
		return time.Unix(int64(n), 0), true
	case int32:
		return time.Unix(int64(n), 0), true
	case int64:
		return time.Unix(int64(n), 0), true
	case int8:
		return time.Unix(int64(n), 0), true
	case int16:
		return time.Unix(int64(n), 0), true
	case uint:
		return time.Unix(int64(n), 0), true
	case uint32:
		return time.Unix(int64(n), 0), true
	case uint64:
		return time.Unix(int64(n), 0), true
	case uint8:
		return time.Unix(int64(n), 0), true
	case uint16:
		return time.Unix(int64(n), 0), true
	}
	return time.Time{}, false
}

func TimeZ(v interface{}) time.Time {
	r, _ := Time(v)
	return r
}

func Bytes(v interface{}) ([]byte, bool) {
	switch n := v.(type) {
	case []byte:
		return n, true
	case string:
		return []byte(n), true
	}
	return nil, false
}

func BytesZ(v interface{}) []byte {
	r, _ := Bytes(v)
	return r
}

func BytesZI(v interface{}) []byte {
	r, _ := Bytes(v)
	if r == nil {
		return []byte{}
	}
	return r
}

func BoolZ(v interface{}) bool {
	r, _ := Bool(v)
	return r
}

func Bool(v interface{}) (bool, bool) {
	switch n := v.(type) {
	case bool:
		return n, true
	case uint:
		return n != 0, true
	case uint32:
		return n != 0, true
	case uint64:
		return n != 0, true
	case uint8:
		return n != 0, true
	case uint16:
		return n != 0, true
	case int:
		return n != 0, true
	case int32:
		return n != 0, true
	case int64:
		return n != 0, true
	case int8:
		return n != 0, true
	case int16:
		return n != 0, true
	case float32:
		return n != 0, true
	case float64:
		return n != 0, true
	case string:
		if s, err := strconv.ParseBool(n); err == nil {
			return s, true
		}
	}
	return false, false
}

func Slice(v interface{}) ([]interface{}, bool) {
	switch n := v.(type) {
	case []interface{}:
		return n, true
	}
	return nil, false
}

func SliceZ(v interface{}) []interface{} {
	r, _ := Slice(v)
	return r
}

func SliceZI(v interface{}) []interface{} {
	r, _ := Slice(v)
	if r == nil {
		return []interface{}{}
	}
	return r
}

func Map(v interface{}) (map[string]interface{}, bool) {
	switch n := v.(type) {
	case map[string]interface{}:
		return n, true
	}
	return nil, false
}

func MapZ(v interface{}) map[string]interface{} {
	r, _ := Map(v)
	return r
}

func MapZI(v interface{}) map[string]interface{} {
	r, _ := Map(v)
	if r == nil {
		return map[string]interface{}{}
	}
	return r
}

func Len(v interface{}) (int, bool) {
	switch n := v.(type) {
	case map[string]interface{}:
		return len(n), true
	case []interface{}:
		return len(n), true
	case string:
		return len(n), true
	case []byte:
		return len(n), true
	}
	return 0, false
}

func LenZ(v interface{}) int {
	l, ok := Len(v)
	if !ok {
		return -1
	}
	return l
}

func Key(v interface{}, key string) (interface{}, bool) {
	m, ok := Map(v)
	if ok {
		n, ok := m[key]
		return n, ok
	}
	return nil, false
}

func KeyZ(v interface{}, key string) interface{} {
	n, _ := Key(v, key)
	return n
}

func Idx(v interface{}, idx int) (interface{}, bool) {
	s, ok := Slice(v)
	if ok && idx < len(s) && idx >= 0 {
		return s[idx], true
	}
	return nil, false
}

func IdxZ(v interface{}, idx int) interface{} {
	n, _ := Idx(v, idx)
	return n
}

func KeyIdx(v interface{}, keyIdx interface{}) (interface{}, bool) {
	if m, ok := Map(v); ok {
		if key, ok := String(keyIdx); ok {
			n, ok := m[key]
			return n, ok
		}
	} else if s, ok := Slice(v); ok {
		if idx, ok := Int(keyIdx); ok && idx < len(s) && idx >= 0 {
			return s[idx], true
		}
	}
	return nil, false
}

func KeyIdxZ(v interface{}, keyIdx interface{}) interface{} {
	n, _ := KeyIdx(v, keyIdx)
	return n
}
