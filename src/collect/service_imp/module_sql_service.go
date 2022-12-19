package collect

import (
	common "collect.mod/src/collect/common"
	handler_template "collect.mod/src/collect/service_imp/handler_template"
	utils "collect.mod/src/collect/utils"
	"database/sql"
	"fmt"
	"github.com/demdxx/gocast"
	"log"
	text_template "text/template"

	config "collect.mod/src/collect/config"
)

type SqlService struct {
	BaseHandler
}

//var db0 *sql.DB

func (s *SqlService) Result(template *config.Template, ts *TemplateService) *common.Result {
	r := common.Result{}
	var err error
	db0, err = s.GetDatasource()
	if err != nil {
		msg := err.Error()
		return r.NotOk(msg)
	}
	//获取文件数据
	// fileData := template.GetFileData()
	params := template.GetParams()
	if template.Log {
		template.LogData("服务请求参数:")
		template.LogData(params)
	}
	// 生成执行SQL和参数
	sql, realValues := getSQLByTpl(template.FileDataTpl, params)
	// 执行SQL
	if template.Log {
		template.LogData("执行数据SQL:")
		template.LogData(sql)
		template.LogData("数据SQL参数:")
		template.LogData(realValues)
	}
	// 执行结果
	maps, _ := sqlToData(sql, realValues...)
	var count int64
	if template.CountFileDataTpl != nil {
		// count 设置不分页
		params[template.Pagination] = false
		// 生成执行SQL和参数
		countSql, countRealValues := getSQLByTpl(template.CountFileDataTpl, params)
		// 执行SQL
		if template.Log {
			template.LogData("执行count SQL:")
			template.LogData(countSql)
			template.LogData("count SQL参数:")
			template.LogData(countRealValues)
		}
		// 执行结果
		countMaps, _ := sqlToData(countSql, countRealValues...)
		if len(countMaps) != 0 {
			countData := countMaps[0]
			var countValue interface{}

			if !utils.IsEmpty("count", countData) { // 获取小写的count
				countValue = utils.GetSafeData("count", countData)
			} else if !utils.IsEmpty("COUNT", countData) { //获取大写的count
				countValue = utils.GetSafeData("COUNT", countData)
			} else { //获取第一个key 的值
				countValue = utils.GetMapValues(countData)[0]

			}
			count = gocast.ToInt64(countValue)

		}

	}
	t := r.OkWithCount(maps, "执行成功", count)
	return t
}

func getSQLByTpl(tpl *text_template.Template, params map[string]interface{}) (string, []interface{}) {
	// 渲染第一次，将二级变量处理成一级变量。第一遍，根据模块转换
	t := handler_template.NewSqlTemplateByTpl(tpl)
	sql, sqlParams, _ := t.Content2Sql(params, true)
	// 渲染第二次,获取实际值，第二步根据模板转换的结果，重新渲染
	t = handler_template.NewSqlTemplate(sql)
	sql, _, realValues := t.Content2Sql(sqlParams, false)
	return sql, realValues
}

func getSQL(sqlData string, params map[string]interface{}) (string, []interface{}) {

	// 渲染第一次，将二级变量处理成一级变量
	t := handler_template.NewSqlTemplate(sqlData)
	sql, sqlParams, _ := t.Content2Sql(params, true)
	// 渲染第二次,获取实际值
	t = handler_template.NewSqlTemplate(sql)
	sql, _, realValues := t.Content2Sql(sqlParams, false)
	return sql, realValues
}

func sqlToData(sqlTemplate string, params ...any) ([]map[string]interface{}, error) {
	rows, err := db0.Query(sqlTemplate, params...)
	if err != nil {
		fmt.Println("出错了", err)
		return nil, err
	}
	//转换成map
	maps := convertMaps(rows)
	return maps, nil
}

func convertMaps(rows *sql.Rows) []map[string]interface{} {

	colNames, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()
	var cols = make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		cols[i] = new(interface{})
	}
	var maps = make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(cols...)
		if err != nil {
			log.Fatal(err.Error())
		}
		var rowMap = make(map[string]interface{})
		for i := 0; i < len(colNames); i++ {
			rowMap[colNames[i]] = convertRowByCol(colTypes[i].DatabaseTypeName(), *(cols[i].(*interface{})))
		}
		maps = append(maps, rowMap)
	}
	return maps

}
func convertRowByCol(colType string, value any) any {
	return utils.CastValue(value, colType)

}
