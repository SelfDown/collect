package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"encoding/json"
	"github.com/go-ldap/ldap/v3"
	"log"
	"strings"
)

type LdapService struct {
	BaseHandler
}
type Ldap struct {
	Connection     LdapConnection
	Method         string
	SearchParams   ldap.SearchRequest
	AddParams      ldap.AddRequest
	ModifyParams   ldap.ModifyRequest
	DeleteParams   ldap.DelRequest
	ModifyDnParams ldap.ModifyDNRequest
	Split          string
}

//type LdapRequest struct {
//	ldap.SearchRequest
//}
type LdapConnection struct {
	Server   string
	User     string
	Password string
}

func handlerAdd(l *ldap.Conn, template *config.Template, ldapConfig *Ldap) *common.Result {

	tmp := ldap.AddRequest(ldapConfig.AddParams)
	err := l.Add(&tmp)
	if err != nil {
		template.LogData(err)
		return common.NotOk(err.Error())
	}
	return common.Ok(nil, "成功")
}

func handlerModify(l *ldap.Conn, template *config.Template, ldapConfig *Ldap) *common.Result {
	tmp := ldap.ModifyRequest(ldapConfig.ModifyParams)
	err := l.Modify(&tmp)
	if err != nil {
		template.LogData(err)
		return common.NotOk(err.Error())
	}
	return common.Ok(nil, "成功")
}
func handlerDelete(l *ldap.Conn, template *config.Template, ldapConfig *Ldap) *common.Result {
	tmp := ldap.DelRequest(ldapConfig.DeleteParams)
	err := l.Del(&tmp)
	if err != nil {
		template.LogData(err)
		return common.NotOk(err.Error())
	}
	return common.Ok(nil, "成功")
}
func handlerModifyDn(l *ldap.Conn, template *config.Template, ldapConfig *Ldap) *common.Result {
	tmp := ldap.ModifyDNRequest(ldapConfig.ModifyDnParams)
	err := l.ModifyDN(&tmp)
	if err != nil {
		template.LogData(err)
		return common.NotOk(err.Error())
	}
	return common.Ok(nil, "成功")
}
func handlerSearch(l *ldap.Conn, template *config.Template, ldapConfig *Ldap) *common.Result {
	tmp := ldap.SearchRequest(ldapConfig.SearchParams)
	sr, err := l.Search(&tmp)
	if err != nil {
		template.LogData(err)
		return common.NotOk(err.Error())

	}
	result := make([]map[string]interface{}, 0)
	split := ldapConfig.Split
	if utils.IsValueEmpty(split) {
		split = ","
	}
	for _, item := range sr.Entries {
		obj := make(map[string]interface{})
		for _, attr := range item.Attributes {
			obj[attr.Name] = strings.Join(attr.Values, split)
		}
		result = append(result, obj)
	}
	return common.Ok(result, "成功")
}

/**
* 连接信息直接json 转的
 */
func (s *LdapService) Result(template *config.Template, ts *TemplateService) *common.Result {

	params := template.GetParams()
	dataContent := utils.RenderTpl(template.FileDataTpl, params)
	if template.Log {
		template.LogData(dataContent)
	}
	ldapConfig := Ldap{}
	json.Unmarshal([]byte(dataContent), &ldapConfig)
	connect := ldapConfig.Connection
	if utils.IsValueEmpty(connect.Server) {
		return common.NotOk("ldap登陆失败！ ldap 服务器地址不能为空")
	}
	l, err := ldap.DialURL(connect.Server)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: connect.User,
		Password: connect.Password,
	})
	if err != nil {
		if template.Log {
			template.LogData("登陆失败")
			template.LogData(connect)
		}
		return common.NotOk("ldap登陆失败！" + err.Error())
	}
	// 如果没有执行方法，就直接返回
	if utils.IsValueEmpty(ldapConfig.Method) {
		return common.Ok("ldap 登陆成功", "")
	}
	// 处理查询
	if ldapConfig.Method == "search" {
		return handlerSearch(l, template, &ldapConfig)
	} else if ldapConfig.Method == "add" {
		return handlerAdd(l, template, &ldapConfig)
	} else if ldapConfig.Method == "modify" {
		return handlerModify(l, template, &ldapConfig)
	} else if ldapConfig.Method == "delete" {
		return handlerDelete(l, template, &ldapConfig)
	} else if ldapConfig.Method == "modifyDn" {
		return handlerModifyDn(l, template, &ldapConfig)
	}
	////search := ldapConfig.Params
	////这里转一下，是为了ldapConfig.Params 以后可以定义更多的参数，比如controller[] 不好处理
	//tmp := ldap.SearchRequest(ldapConfig.SearchParams)
	//sr, err := l.Search(&tmp)
	//if err != nil {
	//	template.LogData(err)
	//}
	//result := make([]map[string]string, 0)
	//for _, item := range sr.Entries {
	//	obj := make(map[string]string)
	//	for _, attr := range item.Attributes {
	//		obj[attr.Name] = attr.Values[0]
	//	}
	//	result = append(result, obj)
	//}
	return common.Ok(nil, "成功")
}
