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
	if strings.HasPrefix(dataType, "*") { //处理指针类型
		switch dataType {
		// 处理指针
		case "*string":
			ptr := gocast.ToString(value)
			return &ptr
		// 处理指针
		case "*int32":
			ptr := gocast.ToInt32(value)
			return &ptr
		case "*int64":
			ptr := gocast.ToInt64(value)
			return &ptr
		case "*bigint":
			fallthrough
		case "*int":
			ptr := gocast.ToInt(value)
			return &ptr
		case "*bool":
			ptr := gocast.ToBool(value)
			return &ptr
		case "*float":
			ptr := gocast.ToFloat(value)
			return &ptr
		case "*time.time":
			fallthrough
		case "*time":
			ptr := ToTime(value)
			return &ptr
		case "*sql.nulltime":
			ptr := ToSqlNullTime(value)
			return &ptr
		default:
			value = gocast.ToString(value)
		}
	}
	switch dataType {
	case "int32":
		value = gocast.ToInt32(value)
		break
	case "int64":
		value = gocast.ToInt64(value)
		break
	case "bigint":
		fallthrough
	case "int":
		value = gocast.ToInt(value)
		break
	case "bool":
		value = gocast.ToBool(value)
		break
	case "float":
		value = gocast.ToFloat(value)
		break
	case "time.time":
		fallthrough
	case "time":
		value = ToTime(value)
		break
	case "sql.nulltime":
		value = ToSqlNullTime(value)
		break
	default:
		value = gocast.ToString(value)
	}
	return value
}
