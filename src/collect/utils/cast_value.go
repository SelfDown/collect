package collect

import (
	"github.com/demdxx/gocast"
	"strings"
)

// CastValue 根据类型转换值
func CastValue(value interface{}, dataType string) interface{} {
	if IsValueEmpty(dataType) {
		return value
	}
	dataType = strings.ToLower(dataType)
	isPtr := false
	if strings.HasPrefix(dataType, "*") { //处理指针类型
		dataType = dataType[1:]
		isPtr = true
	}
	switch dataType {
	case "string":
		valueNew := gocast.ToString(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	case "int32":
		valueNew := gocast.ToInt32(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	case "int64":
		valueNew := gocast.ToInt64(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	case "bigint":
		fallthrough
	case "int":
		valueNew := gocast.ToInt(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	case "bool":
		valueNew := gocast.ToBool(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	case "float":
		valueNew := gocast.ToFloat(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	case "time.time":
		fallthrough
	case "time":
		valueNew := ToTime(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	case "sql.nulltime":
		valueNew := ToSqlNullTime(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	default:
		valueNew := gocast.ToString(value)
		if isPtr {
			return &valueNew
		}
		return valueNew
	}
	//return value
}
