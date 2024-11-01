package collect

import "github.com/demdxx/gocast"

func Cast(value any, dataType string) any {

	if dataType == "int" {
		return gocast.ToInt(value)
	} else if dataType == "int64" {
		return gocast.ToInt64(value)
	} else if dataType == "float" {
		return gocast.ToFloat32(value)
	} else if dataType == "float64" {
		return gocast.ToFloat64(value)
	}
	return value
}
