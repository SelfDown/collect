package collect

import (
	"fmt"
	"reflect"
	text_template "text/template"

	utils "collect.mod/src/collect/utils"
	uuid "github.com/satori/go.uuid"
)

/*
* 将配置文件的所有模板替换成模板变量
* 1. 将params 里面template 替换成成模板
* 2. 将check 里面template 替换
 */

type handlerTpl interface {
	SetServiceName(string)
	SetDataFrom(string)
	SetTarget(interface{})
	SetFromField(string)
	SetToField(string)
	SetTemplate(Template)
	SetField(string)
	SetThirdField(string)
	SetFields([]ThirdField)
	SetThirdFields([]ThirdField)
	SetThirdArrayField(string)
	SetThirdArrayFields([]ThirdField)
	handler() interface{}
}
type baseTplHandler struct {
	service          string      // 服务名称
	template         Template    // 模板
	dataFrom         string      // 数据来源内容
	fromField        string      // 来源字段
	toField          string      //目标中的字段
	target           interface{} // 目标对象
	field            string      // 字段
	thirdField       string      //三级字段
	fields           []ThirdField
	thirdFields      []ThirdField
	thirdArrayField  string // 三级数组字段
	thirdArrayFields []ThirdField
	handlerTpl
}

func (t *baseTplHandler) _handle_simple_field(data reflect.Value) {
	if utils.IsValueEmpty(t.fromField) {
		return
	}
	dataFrom := data.FieldByName(t.fromField).String()
	if utils.IsValueEmpty(dataFrom) {
		return
	}
	tpl, _ := _load_template(dataFrom)
	data.FieldByName(t.toField).Set(reflect.ValueOf(tpl))
}
func (t *baseTplHandler) _handler_array_field(data reflect.Value) {
	for _, f := range t.fields {
		dataFrom := data.FieldByName(f.From).String()
		if utils.IsValueEmpty(dataFrom) {
			continue
		}
		tpl, _ := _load_template(dataFrom)
		data.FieldByName(f.To).Set(reflect.ValueOf(tpl))
	}
}
func (t *baseTplHandler) _handler_second_array_field(data reflect.Value) {

	if utils.IsValueEmpty(t.thirdArrayField) {
		return
	}
	third := data.FieldByName(t.thirdArrayField)
	// 新生成三级字段
	thirdNew := reflect.New(third.Type()).Elem()

	utils.CopyRecursive(third, thirdNew)

	for i := 0; i < thirdNew.Len(); i++ {
		thirdItem := thirdNew.Index(i)
		thirdItemNew := reflect.New(thirdItem.Type()).Elem()
		utils.CopyRecursive(thirdItem, thirdItemNew)
		for _, f := range t.thirdArrayFields {
			fromContent := thirdItem.FieldByName(f.From).String()
			if utils.IsValueEmpty(fromContent) {
				continue
			}
			tpl, _ := _load_template(fromContent)

			thirdItemNew.FieldByName(f.To).Set(reflect.ValueOf(tpl))

		}
		thirdNew.Index(i).Set(thirdItemNew)

	}

	data.FieldByName(t.thirdArrayField).Set(thirdNew)
}
func (t *baseTplHandler) _handler_second_field(data reflect.Value) {

	if utils.IsValueEmpty(t.thirdField) {
		return
	}

	third := data.FieldByName(t.thirdField)
	// 新生成三级字段
	thirdNew := reflect.New(third.Type()).Elem()
	utils.CopyRecursive(third, thirdNew)
	for _, f := range t.thirdFields {
		fromContent := thirdNew.FieldByName(f.From).String()
		if utils.IsValueEmpty(fromContent) {
			continue
		}
		tpl, _ := _load_template(fromContent)
		thirdNew.FieldByName(f.To).Set(reflect.ValueOf(tpl))
	}
	data.FieldByName(t.thirdField).Set(thirdNew)
}

func (t *baseTplHandler) _handler_config_field(data reflect.Value) {
	//简单处理一级字段,data_file
	t._handle_simple_field(data)
	// 循环处理一级字段，数组类型
	t._handler_array_field(data)
	// 处理二级字段.check.template、check.err_msg
	t._handler_second_field(data)
	// 处理二级数组列表,处理handler_params
	t._handler_second_array_field(data)

}

func (t *baseTplHandler) SetServiceName(service string) {
	t.service = service

}

func (t *baseTplHandler) SetTemplate(template Template) {
	t.template = template

}

func (t *baseTplHandler) SetDataFrom(dataFrom string) {
	t.dataFrom = dataFrom

}

func (t *baseTplHandler) SetFromField(fromField string) {
	t.fromField = fromField
}

func (t *baseTplHandler) SetToField(toField string) {
	t.toField = toField

}

func (t *baseTplHandler) SetTarget(target interface{}) {
	t.target = target

}

func (t *baseTplHandler) SetField(field string) {
	t.field = field

}
func (t *baseTplHandler) SetThirdField(field string) {
	t.thirdField = field

}
func (t *baseTplHandler) SetFields(fields []ThirdField) {
	t.fields = fields

}
func (t *baseTplHandler) SetThirdFields(fields []ThirdField) {
	t.thirdFields = fields

}

func (t *baseTplHandler) SetThirdArrayField(field string) {
	t.thirdArrayField = field
}
func (t *baseTplHandler) SetThirdArrayFields(fields []ThirdField) {
	t.thirdArrayFields = fields
}

/*
* 处理字段的模板转换
* @tplContent 表示模板内容
* @toField 表示转换的目标对象的字段
* @target  表示目标对象
 */
func (t *baseTplHandler) _handler_value_tpl(tplContent string, toField string, target interface{}) error {
	if utils.IsValueEmpty(tplContent) {
		return nil
	}
	// 根据来源文件转，文件内容
	tpl, err := _load_template(tplContent)
	if err != nil {
		return err

	}
	utils.SetDataValue(toField, tpl, target)
	return nil
}

/*
* 处理模板
 */
func (t *baseTplHandler) handler_value_tpl(tplContent string, target interface{}) {
	err := t._handler_value_tpl(tplContent, t.toField, target)
	t.handlerTplErr(t.fromField, tplContent, err)
}
func (t *baseTplHandler) handlerTplErr(fieldFrom string, dataFrom string, err error) {
	if err == nil {
		return
	}
	msg := "\n服务:" + t.service
	if utils.IsValueEmpty(fieldFrom) {
		msg += "\n字段：" + fieldFrom
	}
	msg += "\n内容：" + dataFrom + "\n错误：" + err.Error()
	t.template.LogErr(msg)

}
func (t *baseTplHandler) get_target() interface{} {
	return t.target
}

/*
* 简单字段的处理
 */
type fieldTplHandler struct {
	baseTplHandler
}

func (t *fieldTplHandler) handler() interface{} {
	t.handler_value_tpl(t.dataFrom, t.get_target())
	return t.get_target()
}

/*
* 字典字段的处理
 */
type mapTplHandler struct {
	baseTplHandler
}

//func (t *map_tpl_handler) get_target() interface{} {
//	return reflect.ValueOf(t.target).Elem().Elem().Interface()
//}

/*
  - 这里需要紧急已经非常棘手的事情，
  - 就是改结构体里面的map,如果正向修改，需要取里面的内容
  - 结构体字段都是没有地址的，不能通过反射直接赋值，取出来的变量，也是不是结构体的值，而是复制来的变量
    fmt.Println(&username)
    field := service.Params["username"]
    field.Name = "hello"
    service.Params["username"] = field
*/
func (t *mapTplHandler) handler() interface{} {
	dv := reflect.ValueOf(t.get_target()).Elem()
	out := dv.FieldByName(t.field)
	// 拷贝新数据
	dataNew := reflect.New(out.Type()).Elem()
	utils.CopyRecursive(out, dataNew)
	iter := out.MapRange()
	for iter.Next() {
		iv := iter.Value()
		iv_new := reflect.New(iv.Type()).Elem()
		utils.CopyRecursive(iv, iv_new)
		t._handler_config_field(iv_new)
		dataNew.SetMapIndex(iter.Key(), iv_new)

	}
	utils.SetDataValue(t.field, dataNew.Interface(), t.target)
	return t.target

}

/*
* 字典字段的处理
 */
type arrTplHandler struct {
	baseTplHandler
}

func (t *arrTplHandler) handler() interface{} {
	dv := reflect.ValueOf(t.get_target()).Elem()
	out := dv.FieldByName(t.field)
	// 拷贝新数据
	dataNew := reflect.New(out.Type()).Elem()
	utils.CopyRecursive(out, dataNew)
	for i := 0; i < out.Len(); i++ {
		iv := dataNew.Index(i)
		t._handler_config_field(iv)
		dataNew.Index(i).Set(iv)

	}
	dv.FieldByName(t.field).Set(dataNew)
	return t.get_target()

}

func (t *PluginLoader) LoadRouterAllEnable(config Plugin, template Template, routerAll *RouterAll) {
	for _, field := range config.Fields { // 将里面的字段转换成模板。
		t.handler("", "router_all", routerAll, field, template)
	}
}
func (t *PluginLoader) handler(dataFrom string, serviceName string, target interface{}, field SubField, template Template) {
	var _handler handlerTpl
	rule := field.Rule
	if rule == "simple_field" { // 如果是一级字段直接转值
		// 处理service 第一层的字段，根据来源文件转，文件内容
		_handler = &fieldTplHandler{}
		_handler.SetTarget(target)

	} else if rule == "map_field" { //如果是map 里面字段
		/*
		* 处理第二层的字段
		* 处理params 里面的template 处理成模板
		 */
		_handler = &mapTplHandler{}
		_handler.SetTarget(target)

	} else if rule == "array_field" { // 处理字典二级数据
		_handler = &arrTplHandler{}
		_handler.SetTarget(target)
	}

	if _handler == nil {
		template.LogErr("没有找到模板加载规则" + rule)
		return
	}
	// 设置属性
	_handler.SetServiceName(serviceName)
	_handler.SetDataFrom(dataFrom)
	_handler.SetFromField(field.From)
	_handler.SetToField(field.To)
	_handler.SetTemplate(template)
	_handler.SetField(field.Field)
	_handler.SetFields(field.Fields)
	// 设置三级解析字段
	_handler.SetThirdField(field.ThirdField)
	_handler.SetThirdFields(field.ThirdFields)
	// 设置三级数组字段
	_handler.SetThirdArrayField(field.ThirdArrayField)
	_handler.SetThirdArrayFields(field.ThirdArrayFields)
	// 调用处理方法
	_handler.handler()

}

/*
* 加载data_file 文件内容,
* 在启动的时候，就转换好模板，避免在运行的时候，转换模板，减少cpu 使用
 */
func (t *PluginLoader) LoadDataFileTpl(config Plugin, template Template, routerAll *RouterAll) {
	serviceList := routerAll.GetRegisterServices()
	for _, service := range serviceList {

		for _, field := range config.Fields { // 将里面的字段转换成模板。
			// 根据from 获取来源字段名称,h
			dataFrom := utils.GetDataValueStr(field.From, *service)
			serviceName := service.Service
			//if utils.IsValueEmpty(dataFrom) { // 如果没有配置来源字段则，直接返回
			//	continue
			//}
			// field 字段表示来源哪个字段中，处理二级字段

			/*
			* 主要处理sql_file 和count_file,转换成模板
			 */
			t.handler(dataFrom, serviceName, service, field, template)

		}

	}

}

func Uuid() string {
	u4 := uuid.NewV4()
	return u4.String()
	// return ""
}
func IsEmpty(value interface{}) bool {
	return utils.IsValueEmpty(value)
}
func Must(value interface{}) bool {
	return !utils.IsValueEmpty(value)
}

/*
* 文件内容转成模板
 */
func _load_template(fileData string) (*text_template.Template, error) {
	collectTemplate := text_template.New(fileData).Funcs(text_template.FuncMap{"uuid": Uuid, "must": Must, "is_empty": IsEmpty})
	tpl, error_info := collectTemplate.Parse(fileData)
	if error_info != nil {
		fmt.Println(error_info)
	}
	return tpl, error_info

}
