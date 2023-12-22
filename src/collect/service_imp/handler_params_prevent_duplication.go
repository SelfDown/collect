package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	cacheHandler "github.com/SelfDown/collect/src/collect/service_imp/cache_handler"
	"github.com/demdxx/gocast"
)

//处理缓存
type PreventDuplication struct {
	BaseHandler
}

func (hc *PreventDuplication) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	params := template.GetParams()
	handler := cacheHandler.CacheHandler{}
	fieldList := hc.GetFieldNames(handlerParam)
	dataKey := handler.GetCacheKey(handlerParam.Room, fieldList, params)
	data, ok := handler.Get(dataKey)
	if ok { //如果获取到了直接返回

		result := data.(common.Result)
		ttl, has := handler.GetTTl(dataKey)
		left := ttl.Milliseconds()
		if has {
			result.Count = gocast.ToInt64(left)
		}

		return &result
	} else {
		tmp := common.Ok("1", "缓存成功")
		ok := handler.SetMini(dataKey, *tmp, handlerParam.Second)
		if !ok {
			template.LogData("缓存设置失败" + dataKey)
		} else {
			handler.Wait()
		}
	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
