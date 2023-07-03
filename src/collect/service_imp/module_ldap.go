package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"encoding/json"
	"github.com/go-ldap/ldap/v3"
	"log"
)

type LdapService struct {
	BaseHandler
}
type Ldap struct {
	Connection LdapConnection
	Method     string
	Params     ldap.SearchRequest
}
type LdapRequest struct {
	ldap.SearchRequest
}
type LdapConnection struct {
	Server   string
	User     string
	Password string
}

/**
* 连接信息直接json 转的
 */
func (s *LdapService) Result(template *config.Template, ts *TemplateService) *common.Result {

	params := template.GetParams()

	dataContent := utils.RenderTpl(template.FileDataTpl, params)
	template.LogData(dataContent)
	ldapConfig := Ldap{}
	json.Unmarshal([]byte(dataContent), &ldapConfig)
	connect := ldapConfig.Connection
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

	//search := ldapConfig.Params
	//这里转一下，是为了ldapConfig.Params 以后可以定义更多的参数，比如controller[] 不好处理
	tmp := ldap.SearchRequest(ldapConfig.Params)
	sr, err := l.Search(&tmp)
	if err != nil {
		template.LogData(err)
	}
	result := make([]map[string]string, 0)
	for _, item := range sr.Entries {
		obj := make(map[string]string)
		for _, attr := range item.Attributes {
			obj[attr.Name] = attr.Values[0]
		}
		result = append(result, obj)
	}
	return common.Ok(result, "成功")
}
