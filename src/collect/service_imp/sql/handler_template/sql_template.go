package collect

import (
	"bytes"
	"strings"
	text_template "text/template"
	"text/template/parse"

	utils "collect.mod/src/collect/utils"
)

/**
* sql 模板
**/
type SqlTemplate struct {
	SqlContent string
	Nodes      []parse.Node
	Tpl        text_template.Template
}

/**
* 初始化模板
**/
func NewSqlTemplate(sqlContent string) *SqlTemplate {
	tpl := text_template.New("sql-template")
	tpl, _ = tpl.Parse(sqlContent)

	nodes := tpl.Root.Nodes
	t := SqlTemplate{SqlContent: sqlContent, Nodes: nodes, Tpl: *tpl}
	return &t
}

/**
* 初始化模板
**/
func NewSqlTemplateByTpl(tpl *text_template.Template) *SqlTemplate {
	nodes := tpl.Root.Nodes
	t := SqlTemplate{Nodes: nodes, Tpl: *tpl}
	return &t
}

/**
** 通过递归获取，节点里面变量列表
**/
func varNameList(nodes []parse.Node, params map[string]interface{}) []BaseParam {
	param_name_list := []BaseParam{}
	// 或取子key
	get_child_keys := func(rNodes []parse.Node, item_name string) []string {
		childFieldList := make([]string, 0)

		for i := 0; i < len(rNodes); i++ {
			if rNodes[i].Type() != parse.NodeAction {
				continue
			}
			nodeAction := rNodes[i].(*parse.ActionNode)
			// actionNodeList = append(actionNodeList, *nodeAction)
			for k := 0; k < len(nodeAction.Pipe.Cmds); k++ {
				varName := nodeAction.Pipe.Cmds[k].String()
				// 将$item. 替换掉，取后面的名字
				varName = strings.ReplaceAll(varName, item_name+".", "")
				childFieldList = append(childFieldList, varName)
			}
		}
		return childFieldList

	}

	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		switch node.Type() {
		case parse.NodeRange:
			//for循环里面的变量

			nodeRange := node.(*parse.RangeNode)
			rNodes := nodeRange.List.Nodes
			// var_list := varNameList(rNodes)
			item_name := nodeRange.Pipe.Decl[1].Ident[0]
			cmds := nodeRange.Pipe.Cmds
			foreach := cmds[0].String()

			foreach = strings.Split(foreach, ".")[1]
			// 获取子变量
			childFieldList := get_child_keys(rNodes, item_name)
			// 循环获取里面的变量
			param_name_list = append(param_name_list, GetArrayObjParamObj(foreach, childFieldList, params))
		case parse.NodeField:
			nString := node.String()
			param_name_list = append(param_name_list, GetSimpleParamObj(nString, params))

		// 处理变量

		case parse.NodeAction:
			nodeAction := node.(*parse.ActionNode)
			nString := nodeAction.String()
			// 如果包含空格，则取里面的变量
			if !strings.Contains(nString, " ") {
				param_name_list = append(param_name_list, GetSimpleParamObj(nString, params))
			}

			// param_name_list = append(param_name_list, nString)
		// 处理if else  里面的变量
		case parse.NodeIf:
			nodeif := node.(*parse.IfNode)
			if_names := varNameList(nodeif.List.Nodes, params)
			param_name_list = append(param_name_list, if_names...)
			if nodeif.ElseList != nil {
				else_names := varNameList(nodeif.ElseList.Nodes, params)
				param_name_list = append(param_name_list, else_names...)
			}

		}

	}
	return param_name_list

}

/**
** 获取变量列表
**/
func (t *SqlTemplate) VarNameList(params map[string]interface{}) []BaseParam {
	params_result := varNameList(t.Nodes, params)
	return params_result
}

/**
*
 根据sql 模板渲染sql, 并且获取变量
 sql 模板内容转执行sql,sql 需要渲染2次。
 第一次处理if else ,处理显示和隐藏的变量，
 第二次将变量替换为sql 里面的站位符号，处理成预编译语句
 sql 只处理params第一层变量，不处理2、3层变量以及更深层次的变量，如果有2、3层变量将直接渲染，请注意否有sql 注入的风险
 也不能像写html 层级，各种标签
 @param sql_content SQL 内容
 @param params      请求参数
 @to_param_key      是否转参数字段，第一次需要处理if else ,第二次渲染预编译的sql
**/
func (t *SqlTemplate) Content2Sql(params map[string]interface{}, to_param_key bool) (string, map[string]interface{}, []interface{}) {
	// template_params := make(map[string]interface{})
	param_keys := t.VarNameList(params)
	template_params := make(map[string]interface{})
	real_values := make([]interface{}, 0)
	// 处理模板里面的变量，如果模板里面有变量，就添加到新变量里面去，直接替换对应模板变量或者占位符号
	for i := 0; i < len(param_keys); i++ {
		sql_param := param_keys[i]
		sql_param_key := sql_param.GetSqlParamParamKey()
		if utils.IsEmpty(sql_param_key, params) {
			continue
		}
		param_key_value := sql_param.GetSqlParamKeyValue(to_param_key)
		// 设置变量值
		template_params[sql_param_key] = param_key_value
		// 设置新的变量和值
		apList := sql_param.GetSqlParamParamKeyList()
		for j := 0; j < len(apList); j++ {
			ap := apList[j]
			template_params[ap.AttrName] = ap.AttrValue
		}

	}
	// 将请求参数里面剩余的变量添加进去，未计算if else 判断语句需要的值
	for key, v := range params {

		_, ok := template_params[key]
		if !ok {
			template_params[key] = v
		}
	}

	for _, sql_param := range param_keys {
		param_key := sql_param.GetSqlParamParamKey()
		if utils.IsEmpty(param_key, params) {
			continue
		}
		real_values = append(real_values, sql_param.GetValue()...)

	}

	var buf bytes.Buffer
	t.Tpl.Execute(&buf, template_params)

	return buf.String(), template_params, real_values

}
