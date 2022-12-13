package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"runtime"
)

var db0 *sql.DB
var gormDb *gorm.DB

type BaseHandler struct {
}

func (s *BaseHandler) RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func (s *BaseHandler) GetGormDb() *gorm.DB {
	//如果连接过则直接，返回对象
	if gormDb != nil {
		return gormDb
	}
	db, _ := s.GetDatasource()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	return gormDB
}

// GetDatasource 获取数据库连接
func (s *BaseHandler) GetDatasource() (*sql.DB, error) {
	if db0 != nil {
		return db0, nil
	}
	// 获取连接信息
	dataSourceName := utils.GetAppKey("dataSourceName")
	// 获取驱动
	driverName := utils.GetAppKey("driverName")

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("数据库打开出现了问题：", err)
		return nil, err
	}
	db.Ping()
	if err != nil {
		log.Fatal("数据库连接出现了问题：", err)
		return nil, err
	}
	db0 = db
	return db, err
}

func (s *BaseHandler) Result(template *config.Template) *common.Result {
	return common.Ok(nil, "")
}

func (s *BaseHandler) HandlerData(template *config.Template, handlerParam *config.HandlerParam) *common.Result {
	return common.Ok(nil, "")
}
func (s *BaseHandler) UpdateFields(params map[string]interface{}, modelData interface{}, ignoreFields []string, updateFields []string) (interface{}, []string) {
	_, names := utils.SetDataValueByParams(params, modelData, ignoreFields, updateFields)
	return modelData, names
}
