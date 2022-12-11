package collect

import (
	"database/sql"
	"github.com/demdxx/gocast"
	"reflect"
	"time"
)

func ToInt32(value interface{}) int32 {
	_, ok := value.(int32)
	if !ok { // 如果断言失败，这强制转换
		value = gocast.ToInt32(value)
	}
	return value.(int32)
}

func ToTime(value interface{}) time.Time {
	_, ok := value.(time.Time)
	if !ok {
		value, _ = time.ParseInLocation("2006-01-02 15:04:05", Strval(value), time.Local) //这里按照当前时区转
	}

	return value.(time.Time)
}
func ToSqlNullTime(value interface{}) sql.NullTime {
	valueNew := sql.NullTime{Time: ToTime(value), Valid: true}
	return valueNew
}

// ParseValueByField 根据model 里面的字段类型，转值
func ParseValueByField(field reflect.Value, value interface{}) interface{} {
	typeStr := field.Type().String()
	valueNew := CastValue(value, typeStr)
	return valueNew

}
