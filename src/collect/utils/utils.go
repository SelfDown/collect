package collect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/demdxx/gocast"
	engine "github.com/dengsgo/math-engine/engine"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	text_template "text/template"
	"time"
)

func CreateDirs(path string) error {
	if !IsPathExist(path) {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}
func ParentDirName(filePath string) string {
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	arr := strings.Split(filePath, "/")
	return strings.Join(arr[:len(arr)-1], "/") + "/"
}
func FileName(filePath string) string {
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	arr := strings.Split(filePath, "/")
	return arr[len(arr)-1]
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

func CurrentDateTime() string {
	timeStamp := time.Now().Unix()
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(timeStamp, 0).Format(timeLayout)
	return timeStr
}
func CurrentDateFormat(timeLayout string) string {
	timeStamp := time.Now().Unix()
	if IsValueEmpty(timeLayout) {
		timeLayout = "2006-01-02 15:04:05"
	}

	timeStr := time.Unix(timeStamp, 0).Format(timeLayout)
	return timeStr
}

/*
* @param exec 是计算，是否执行运算结果
* 比如计算分页
      start:
        template: " ({{.page}}-1) * {{.size}}"
        exec: true
        type: int
*/
func RenderTplExec(Tpl *text_template.Template, params map[string]interface{}, exec bool) interface{} {
	s := RenderTplData(Tpl, params)
	if !exec {
		return s
	}
	r, err := engine.ParseAndExec(Strval(s))
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

// RenderTplDataWithType 执行结果转类型
func RenderTplDataWithType(Tpl *text_template.Template, params map[string]interface{}, dataType string) interface{} {

	value := RenderTplData(Tpl, params)
	if value == nil {
		return value
	}
	t := reflect.TypeOf(value)
	//如果是非字符串类型，直接返回
	if t.Kind().String() != "string" {
		return value
	}
	return CastValue(value, dataType)

}
func IsRenderVar(name string) bool {
	if strings.HasPrefix(name, "[") && strings.HasSuffix(name, "]") {
		return true
	}
	return false
}

//取参数变量、或者取值
func RenderVarOrValue(name interface{}, params map[string]interface{}) interface{} {
	pKey, ok := name.(string)
	if ok && IsRenderVar(pKey) {
		return RenderVar(pKey, params)
	}
	return pKey
}
func RenderVarToArrMap(name string, params map[string]interface{}) ([]map[string]interface{}, string) {
	// 直接渲染变量
	data := RenderVar(name, params)
	dataList, ok := data.([]map[string]interface{})
	if ok {
		return dataList, ""
	}

	// 在深层次取，挨个转换
	tmp, ok := data.([]interface{})
	if ok {
		dataList = make([]map[string]interface{}, 0)
		for _, item := range tmp {
			itemData, itemOk := item.(map[string]interface{})
			if itemOk { //如果能直接转，就直接转，转不了包一层
				dataList = append(dataList, itemData)
			} else {
				itemTmp := make(map[string]interface{})
				itemTmp["item"] = item
				dataList = append(dataList, itemTmp)
			}

		}
		return dataList, ""
	}
	if IsArray(data) {
		original := reflect.ValueOf(data)
		for i := 0; i < original.Len(); i++ {
			dataItem := make(map[string]interface{})
			iter := original.Index(i).MapRange()
			for iter.Next() {
				key := iter.Key().String()
				value := iter.Value().Interface()
				dataItem[key] = value
			}
			dataList = append(dataList, dataItem)
		}
	} else {
		return dataList, "非数组对象"
	}

	//test := original.
	//	test.Next()
	//t := test.Value().Interface()
	//fmt.Printf("%#v", t)
	//fmt.Println(test)
	//if tmp == nil {
	//	return nil, name + "对象非数组"
	//}

	return dataList, ""
}

func GetRenderVarName(name string) string {
	varName := strings.Replace(name, "[", "", -1)
	// 替换右边括号
	varName = strings.Replace(varName, "]", "", -1)
	return varName
}
func GetFieldValueList(fields []string, params map[string]interface{}) []string {
	valueList := make([]string, 0)
	for _, item := range fields {
		value := gocast.ToString(RenderVarOrValue(item, params))
		valueList = append(valueList, value)
	}
	return valueList
}
func RenderVar(name string, params map[string]interface{}) interface{} {

	varName := GetRenderVarName(name)
	// 取一级变量
	tmpArr := strings.Split(varName, ".")
	first := tmpArr[0]
	v, _ := params[first]
	if len(tmpArr) == 1 {
		return v
	}
	second := tmpArr[1]
	if len(tmpArr) == 2 { // 处理2级
		// 如果是map[string]字符串类型转换成功
		param2, ok := v.(map[string]string)
		if ok {
			v2, _ := param2[second]
			return v2
		}
		// 如果map[interface] 类型
		param3, ok := v.(map[string]interface{})
		if ok {
			v3, _ := param3[second]
			return v3
		}
	}

	return nil
}

// RenderTplData 根据模板渲染数据，优选取参数里面的字段
func RenderTplData(Tpl *text_template.Template, params map[string]interface{}) interface{} {

	name := Tpl.Name()
	// 如果匹配占位变量
	if IsRenderVar(name) {
		return RenderVar(name, params)
	}
	//模板渲染
	return RenderTpl(Tpl, params)
}
func RenderTplDataBool(Tpl *text_template.Template, params map[string]interface{}) bool {
	data := RenderTplData(Tpl, params)
	return gocast.ToBool(data)
}

// RenderTpl 根据模板渲染数据
func RenderTpl(Tpl *text_template.Template, params map[string]interface{}) string {
	var buf bytes.Buffer
	err := Tpl.Execute(&buf, params)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return strings.TrimSpace(buf.String())
}

/*
* 指针拷贝引用
 */
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

/*
* 指针也同时生成一份
 */
func CopyWithPtr(src interface{}) interface{} {
	if src == nil {
		return nil
	}
	original := reflect.ValueOf(src)
	cpy := reflect.New(original.Type()).Elem()
	CopyRecursivePtr(original, cpy)

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
	//这里将指针拷贝去掉，直接引用，如果需要将指针重新分配对象，这里可以改造一个方法，将此段开发
	//目前主要处理模板指针对象，发现模板函数uuid 方法没有挂载上，经过排除发现template 有个common 对象
	//没有赋值，还nil,所以改成直接引用指针对象

	//case reflect.Ptr:
	//if withPtr {
	//	originalValue := src.Elem()
	//
	//	if !originalValue.IsValid() {
	//		return
	//	}
	//	dst.Set(reflect.New(originalValue.Type()))
	//	CopyRecursive(originalValue, dst.Elem(), withPtr)
	//}

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

func CopyRecursivePtr(src, dst reflect.Value) {
	if src.CanInterface() {
		if copier, ok := src.Interface().(Interface); ok {
			dst.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	switch src.Kind() {
	//这里将指针拷贝去掉，直接引用，如果需要将指针重新分配对象，这里可以改造一个方法，将此段开发
	//目前主要处理模板指针对象，发现模板函数uuid 方法没有挂载上，经过排除发现template 有个common 对象
	//没有赋值，还nil,所以改成直接引用指针对象

	case reflect.Ptr:
		originalValue := src.Elem()
		if !originalValue.IsValid() {
			return
		}
		dst.Set(reflect.New(originalValue.Type()))
		CopyRecursivePtr(originalValue, dst.Elem())

	case reflect.Interface:
		if src.IsNil() {
			return
		}
		originalValue := src.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		CopyRecursivePtr(originalValue, copyValue)
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
			CopyRecursivePtr(src.Field(i), dst.Field(i))
		}

	case reflect.Slice:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			CopyRecursivePtr(src.Index(i), dst.Index(i))
		}

	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {
			originalValue := src.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			CopyRecursivePtr(originalValue, copyValue)
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
	if v, ok := value.(string); ok {
		return v
	}
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

func toCamelCase(input string) string {
	//titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	//camel := strings.Replace(titleSpace, " ", "", -1)
	c := cases.Title(language.Und, cases.NoLower)
	titleSpace := c.String(strings.ReplaceAll(input, "_", " "))
	camel := strings.ReplaceAll(titleSpace, " ", "")
	return camel
}
func StringArrayContain(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// 处理id 转大写，将gen源码直接拷贝过来
var commonInitialisms = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}

func ToSchemaName(name string) string {
	result := strings.ReplaceAll(strings.Title(strings.ReplaceAll(name, "_", " ")), " ", "")
	for _, initialism := range commonInitialisms {
		result = regexp.MustCompile(strings.Title(strings.ToLower(initialism))+"([A-Z]|$|_)").ReplaceAllString(result, initialism+"$1")
	}
	return result
}

func SetDataValueByParams(params map[string]interface{}, data interface{}, ignoreFields []string, updateFields []string, optionFields []string) (map[string]interface{}, []string) {
	dv := reflect.ValueOf(data)
	if dv.Kind() != reflect.Ptr { // 如果不是指针，则取地址
		dv = reflect.ValueOf(&data)
	}
	elem := dv.Elem()
	// 先只实现指针对象
	result := make(map[string]interface{})
	fieldNames := make([]string, 0)
	if dv.Type().String() == "*interface {}" { // 如果外面还包了一层interface ,还原类型，否则结构体不能设置值
		tmp := reflect.New(elem.Elem().Type()).Elem()
		tmp.Set(elem.Elem())
		for name, value := range params {
			fieldName := ToSchemaName(name)
			field := tmp.FieldByName(fieldName)
			if !field.IsValid() {
				continue
			}
			//如果在ignoreFields 则跳过，如果ignoreFields不为空
			if ignoreFields != nil && StringArrayContain(ignoreFields, name) {
				continue
			}
			//如果不在updateFields 则跳过，如果updateFields不为空
			if updateFields != nil && !StringArrayContain(updateFields, name) {
				continue
			}
			// 如果不在optionFields 则跳过，如果optionFields 不能为空
			if optionFields != nil && !StringArrayContain(optionFields, name) {
				continue
			}
			fieldNames = append(fieldNames, name)
			value = ParseValueByField(field, value)
			field.Set(reflect.ValueOf(value))
			result[fieldName] = value
		}
		// 设置map的剩余字段
		dt := reflect.TypeOf(tmp.Interface())
		for i := 0; i < tmp.NumField(); i++ {
			fieldName := dt.Field(i).Name
			if _, ok := result[fieldName]; !ok { //将剩下的元素复制
				result[fieldName] = tmp.Field(i).Interface()
			}
		}
		elem.Set(tmp)
	} else {

	}
	return result, fieldNames
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
		tmp := reflect.New(elem.Elem().Type()).Elem()
		tmp.Set(elem.Elem())
		field := tmp.FieldByName(name)
		if !field.IsValid() {
			log.Println("name：【" + name + "】字段不存在，请检查配置")
			log.Println(data)
		}
		field.Set(reflect.ValueOf(value))
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
	case string:
		if len(s) == 0 {
			return true
		}
		return false
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
