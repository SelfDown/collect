package collect

import (
	"database/sql"
	"fmt"
	"log"

	common "collect.mod/src/collect/common"
	utils "collect.mod/src/collect/utils"
	"github.com/demdxx/gocast"
	_ "github.com/go-sql-driver/mysql"

	// service_imp "test.mod/src/collect/service_imp"
	text_template "text/template"

	config "collect.mod/src/collect/config"
	handler_template "collect.mod/src/collect/service_imp/sql/handler_template"
)

type SqlService struct {
}

var db0 *sql.DB

func get_datasource() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dev_user:!QAZ2wsx#EDC4rfv2022@{xx[,*.]}.com2022@.com@tcp(172.26.0.13:3306)/devops?charset=utf8")
	if err != nil {
		log.Fatal("数据库打开出现了问题：", err)
		return nil, err
	}
	db.Ping()
	if err != nil {
		log.Fatal("数据库连接出现了问题：", err)
		return nil, err
	}
	return db, err

}

var SqlServerPlugin SqlService

func (s *SqlService) Result(template *config.Template) (*common.Result, error) {
	r := common.Result{}
	var err error
	db0, err = get_datasource()
	if err != nil {
		msg := err.Error()
		return r.NotOk(msg), err
	}
	//获取文件数据
	// fileData := template.GetFileData()
	params := template.GetParams()
	if template.Log {
		template.LogData("服务请求参数:")
		template.LogData(params)
	}
	// 生成执行SQL和参数
	sql, real_values := getSQLByTpl(template.FileDataTpl, params)
	// 执行SQL
	if template.Log {
		template.LogData("执行数据SQL:")
		template.LogData(sql)
		template.LogData("数据SQL参数:")
		template.LogData(real_values)
	}
	// 执行结果
	maps, _ := sqlToData(sql, real_values...)
	count := 0
	if template.CountFileDataTpl != nil {
		// count 设置不分页
		params[template.Pagination] = false
		// 生成执行SQL和参数
		count_sql, count_real_values := getSQLByTpl(template.CountFileDataTpl, params)
		// 执行SQL
		if template.Log {
			template.LogData("执行count SQL:")
			template.LogData(count_sql)
			template.LogData("count SQL参数:")
			template.LogData(count_real_values)
		}
		// 执行结果
		count_maps, _ := sqlToData(count_sql, count_real_values...)

		if len(count_maps) != 0 {
			count_data := count_maps[0]
			var count_value interface{}

			if !utils.IsEmpty("count", count_data) { // 获取小写的count
				count_value = utils.GetSafeData("count", count_data)
			} else if !utils.IsEmpty("COUNT", count_data) { //获取大写的count
				count_value = utils.GetSafeData("COUNT", count_data)
			} else { //获取第一个key 的值
				count_value = utils.GetMapValues(count_data)[0]

			}
			count = gocast.ToInt(count_value)

		}

	}

	msg := "执行成功"
	r = *r.OkWithCount(maps, &msg, count)
	return &r, nil
}

func getSQLByTpl(tpl *text_template.Template, params map[string]interface{}) (string, []interface{}) {
	// 渲染第一次，将二级变量处理成一级变量。第一遍，根据模块转换
	t := handler_template.NewSqlTemplateByTpl(tpl)
	sql, sql_params, _ := t.Content2Sql(params, true)
	// 渲染第二次,获取实际值，第二步根据模板转换的结果，重新渲染
	t = handler_template.NewSqlTemplate(sql)
	sql, _, real_values := t.Content2Sql(sql_params, false)
	return sql, real_values
}

func getSQL(sql_data string, params map[string]interface{}) (string, []interface{}) {

	// 渲染第一次，将二级变量处理成一级变量
	t := handler_template.NewSqlTemplate(sql_data)
	sql, sql_params, _ := t.Content2Sql(params, true)
	// 渲染第二次,获取实际值
	t = handler_template.NewSqlTemplate(sql)
	sql, _, real_values := t.Content2Sql(sql_params, false)
	return sql, real_values
}

func sqlToData(sql_template string, params ...any) ([]map[string]interface{}, error) {
	rows, err := db0.Query(sql_template, params...)
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
	switch colType {
	case "BIGINT":
		fallthrough
	case "INT":
		return gocast.ToInt(value)
	default:
		return gocast.ToString(value)
	}

}
