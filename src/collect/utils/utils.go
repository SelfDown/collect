package collect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/demdxx/gocast"
	engine "github.com/dengsgo/math-engine/engine"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	text_template "text/template"
	"time"
)

func ParentDirName(filePath string) string {
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	arr := strings.Split(filePath, "/")
	return strings.Join(arr[:len(arr)-1], "/") + "/"
}

/**
* 读取文件内容
 */
func ReadFileContent(filePath string) (string, bool) {
	content, success := ReadFileBytes(filePath)
	if !success {
		return Strval(content), false
	}
	return Strval(content), true
}

/**
* 读取文件内容
 */
func ReadFileBytes(filePath string) ([]byte, bool) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		msg := filePath + "文件不存在"
		log.Println(msg)
		return []byte(msg), false
	}
	return content, true

}

type Interface interface {
	DeepCopy() interface{}
}

func RenderTplExec(Tpl *text_template.Template, params map[string]interface{}, exec bool) string {
	s := RenderTpl(Tpl, params)
	if !exec {
		return s
	}
	r, err := engine.ParseAndExec(s)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return gocast.ToString(r)

}
func RenderTplBool(Tpl *text_template.Template, params map[string]interface{}) bool {
	value := RenderTpl(Tpl, params)
	return gocast.ToBool(value)
}

//根据模板渲染数据，优选取参数里面的字段
func RenderTplData(Tpl *text_template.Template, params map[string]interface{}) interface{} {

	name := Tpl.Name()
	// 取一级变量
	if v, ok := params[name]; ok {
		return v
	}
	// todo 先取一级，后算再考虑取二级
	//if strings.Contains(name, ".") {
	//	fields := strings.Split(name, ".")
	//	first := ""
	//	second := ""
	//	if len(fields) >= 2 {
	//		first = fields[0]
	//		second = fields[1]
	//	}
	//	if v, ok := params[first]; ok {
	//		if v, ok := params[first] {
	//
	//		}
	//	}
	//}
	//模板渲染
	return RenderTpl(Tpl, params)
}

// 根据模板渲染数据
func RenderTpl(Tpl *text_template.Template, params map[string]interface{}) string {
	var buf bytes.Buffer
	err := Tpl.Execute(&buf, params)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return buf.String()
}

func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}
	original := reflect.ValueOf(src)
	cpy := reflect.New(original.Type()).Elem()
	CopyRecursive(original, cpy)

	value := cpy.Interface()
	return value
}

func CopyRecursive(src, dst reflect.Value) {
	if src.CanInterface() {
		if copier, ok := src.Interface().(Interface); ok {
			dst.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	switch src.Kind() {
	// 这里将指针拷贝去掉，直接引用，如果需要将指针重新分配对象，这里可以改造一个方法，将此段开发
	// 目前主要处理模板指针对象，发现模板函数uuid 方法没有挂载上，经过排除发现template 有个common 对象
	// 没有赋值，还nil,所以改成直接引用指针对象
	//case reflect.Ptr:
	//	originalValue := src.Elem()
	//
	//	if !originalValue.IsValid() {
	//		return
	//	}
	//	dst.Set(reflect.New(originalValue.Type()))
	//	CopyRecursive(originalValue, dst.Elem())

	case reflect.Interface:
		if src.IsNil() {
			return
		}
		originalValue := src.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		CopyRecursive(originalValue, copyValue)
		dst.Set(copyValue)

	case reflect.Struct:
		t, ok := src.Interface().(time.Time)
		if ok {
			dst.Set(reflect.ValueOf(t))
			return
		}
		for i := 0; i < src.NumField(); i++ {
			if src.Type().Field(i).PkgPath != "" {
				continue
			}
			CopyRecursive(src.Field(i), dst.Field(i))
		}

	case reflect.Slice:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			CopyRecursive(src.Index(i), dst.Index(i))
		}

	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {
			originalValue := src.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			CopyRecursive(originalValue, copyValue)
			copyKey := Copy(key.Interface())
			dst.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		dst.Set(src)
	}
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
func GetJSONData(value interface{}) string {
	newValue, _ := json.Marshal(value)
	return string(newValue)
}

/**
* 根据字符串，获取结构体里面字段值
 */
func GetDataValueStr(name string, data interface{}) string {
	dv := reflect.ValueOf(data)
	value := dv.FieldByName(name).String()
	return value
}
func IsMap(name string, data interface{}) bool {
	dv := reflect.ValueOf(data)
	value := dv.FieldByName(name).Kind().String()
	return value == "map"

}
func GetDataValueMapIter(name string, data interface{}) *reflect.MapIter {
	dv := reflect.ValueOf(data)
	value := dv.FieldByName(name).MapRange()
	return value
}

/*
* @name 字段名称
* @value 值
* @data 目标对象，可以是指针地址，也可以是目标对象
**/
func SetDataValue(name string, value interface{}, data interface{}) interface{} {
	dv := reflect.ValueOf(data)
	if dv.Kind() != reflect.Ptr { // 如果不是指针，则取地址
		dv = reflect.ValueOf(&data)
	}
	elem := dv.Elem()
	if dv.Type().String() == "*interface {}" { // 如果外面还包了一层interface ,还原类型，否则结构体不能设置值
		// v := reflect.ValueOf(data).Elem()

		// Allocate a temporary variable with type of the struct.
		//    v.Elem() is the vale contained in the interface.
		tmp := reflect.New(elem.Elem().Type()).Elem()
		// Copy the struct value contained in interface to
		// the temporary variable.
		tmp.Set(elem.Elem())
		// Set the field.
		field := tmp.FieldByName(name)
		if !field.IsValid() {
			log.Println("name：【" + name + "】字段不存在，请检查配置")
			log.Println(data)
		}
		field.Set(reflect.ValueOf(value))
		// Set the interface to the modified struct value.
		elem.Set(tmp)
	} else {
		field := elem.FieldByName(name)

		if !field.IsValid() {
			log.Println("name：" + name + "字段不存在")
			log.Println(data)
		}
		field.Set(reflect.ValueOf(value))
	}
	return data

}

/**
* 判断对象里面值是否为空
**/
func GetSafeData(name string, data map[string]interface{}) interface{} {
	if data == nil {
		return ""
	}
	if v, ok := data[name]; ok {
		return v
	} else {
		return nil
	}
}
func GetMapKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}

func GetMapValues(data map[string]interface{}) []interface{} {
	values := make([]interface{}, 0, len(data))
	for k := range data {
		values = append(values, data[k])
	}
	return values
}

/**
* 判断对象里面值是否为空
**/
func IsEmpty(name string, data map[string]interface{}) bool {
	if data == nil {
		return true
	}
	value := data[name]
	return IsValueEmpty(value)
}

func IsMultipleField(field string) bool {
	if strings.Contains(".", field) {
		return true
	} else {
		return false
	}

}

/**
* 判断值是否为空
**/
func IsValueEmpty(value any) bool {

	switch s := value.(type) {
	case bool:
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case float32:
	case float64:
		return false
	case nil:
		return true
	default:
		if reflect.ValueOf(s).Len() == 0 {

			return true
		}
	}
	return false
}

/**
* 判断是否为数组
**/
func IsArray(value any) bool {
	if value == nil {
		return false
	}
	kind := reflect.TypeOf(value).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

/**
* 将字符串先乘以几次，然后以逗号拼接转成一个字符串
**/
func MultiplyJoinComma(content string, times int) string {
	return MultiplyJoin(content, times, ",")
}

/**
* 将字符串先乘以几次，然后以sep转成一个字符串
**/
func MultiplyJoin(content string, times int, sep string) string {
	if times <= 0 {
		return content
	}
	if IsValueEmpty(content) {
		return content
	}

	elem := []string{}
	for i := 0; i < times; i++ {
		elem = append(elem, content)
	}
	return strings.Join(elem, sep)
}
func CopyMap(src map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range src {
		m[k] = v
	}
	return m
}
