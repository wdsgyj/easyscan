// conv
package easyscan

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func any2String(i interface{}) (string, bool) {
	switch va := i.(type) {
	case []byte:
		return string(va), true
	case string:
		return va, true
	default:
		return fmt.Sprintf("%v", va), true
	}
}

func any2Bytes(i interface{}) ([]byte, bool) {
	switch va := i.(type) {
	case []byte:
		return va, true
	case string:
		return []byte(va), true
	default:
		return []byte(fmt.Sprintf("%v", va)), true
	}
}

func any2Int(i interface{}) (int64, bool) {
	switch va := i.(type) {
	case []byte:
		rs, err := strconv.ParseInt(string(va), 10, 64)
		if err != nil {
			return 0, false
		} else {
			return rs, true
		}
	case string:
		rs, err := strconv.ParseInt(va, 10, 64)
		if err != nil {
			return 0, false
		} else {
			return rs, true
		}
	case float64:
		return int64(va), true
	case bool:
		if va {
			return 1, true
		} else {
			return 0, true
		}
	case int64:
		return va, true
	case time.Time:
		return va.Unix()*1000 + va.UnixNano()/1000000, true
	default:
		panic(fmt.Sprintf("%T:%v", i, i))
	}
}

func any2Float(i interface{}) (float64, bool) {
	switch va := i.(type) {
	case []byte:
		rs, err := strconv.ParseFloat(string(va), 64)
		if err != nil {
			return 0, false
		} else {
			return rs, true
		}
	case string:
		rs, err := strconv.ParseFloat(va, 64)
		if err != nil {
			return 0, false
		} else {
			return rs, true
		}
	case int64:
		return float64(va), true
	case bool:
		if va {
			return 1.0, true
		} else {
			return 0, true
		}
	case float64:
		return va, true
	case time.Time:
		return float64(va.Unix()*1000 + va.UnixNano()/1000000), true
	default:
		panic(fmt.Sprintf("%T:%v", i, i))
	}
}

func any2Bool(i interface{}) (bool, bool) {
	switch va := i.(type) {
	case []byte:
		switch string(va) {
		case "1", "t", "T", "true", "TRUE", "True":
			return true, true
		case "0", "f", "F", "false", "FALSE", "False":
			return false, true
		default:
			panic("error string to bool")
		}
	case string:
		switch va {
		case "1", "t", "T", "true", "TRUE", "True":
			return true, true
		case "0", "f", "F", "false", "FALSE", "False":
			return false, true
		default:
			panic("error string to bool")
		}
	case bool:
		return va, true
	case int64:
		if va != 0 {
			return true, true
		} else {
			return false, true
		}
	case float64:
		if math.Abs(va-0) < 0.000001 {
			return false, true
		} else {
			return true, true
		}
	case time.Time:
		return !va.IsZero(), true
	default:
		panic(fmt.Sprintf("%T:%v", i, i))
	}
}
