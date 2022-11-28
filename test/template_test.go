package test

import (
	"bytes"
	"fmt"
	"testing"

	// "log"

	"time"

	uuid "github.com/satori/go.uuid"

	text_template "text/template"
)

func Uuid() string {
	u4 := uuid.NewV4()
	return u4.String()
	// return ""
}
func Test1(t *testing.T) {
	// 这里是测试代码，如果是模板渲染，1w个，golang只需要96毫秒，如果是python 渲染需要19s
	// 这里96ms
	// 经过测试模板如果能提前parse,100w 个只要765ms,100w个生产uuuid 2.1秒
	tpl := text_template.New("sql-template").Funcs(text_template.FuncMap{"uuid": Uuid})
	tpl2, _ := tpl.Parse("{{.nick}}1,{{uuid}}]")
	startTime := time.Now()
	for i := 0; i < 1; i++ {
		params := make(map[string]interface{})
		params["nick"] = "张治"
		params["a"] = "张治"
		var buf bytes.Buffer
		tpl2.Execute(&buf, params)
		fmt.Println(buf.String())

	}
	fmt.Println(time.Since(startTime))
	/*
		这里运行只要19s
			import  time
			start = time.time()
			for i in range(100000):
				params={
					"nick":"张治"
				}
				sql_templ="{{nick}}"
				t = env.from_string(sql_templ)
				t.render(**params)
			end = time.time()
			print end-start
	*/

	/*
		    这样写100w只要3.4秒，前面的写法都在parse,只是耗时很高，之前估计是写的自定义函数，uuid，每次都不能生成
			不知道golang 是否支持
			import  time
			start = time.time()
			sql_templ = "{{nick}}"
			t = env.from_string(sql_templ)
			for i in range(1000000):
				params={
					"nick":"张治"
				}
				t.render(**params)
			end = time.time()
			print end-start
	*/

	/*
	* 进过测试，python 的filter ，只有第一次注册的时候会触发uuid,这个就是致命的问题，因为我要靠模板生成主键，主键每次必须不一样
	* 而python 的目标每次生成一样的，会造成主键冲突。
	* 如果每次新生成模板，耗时巨大
	* {{''|uuid}} 它好像是根据传的值，取的缓存，但是我没有找到这个缓存开关
	 */

	/**
	 * 进过研究发现有个globals 参数，可以直接调用
	 def uuid():
		import uuid
		return str(uuid.uuid4())
	 sql_templ = "{{nick}}_{{uuid()}}"
	 t = env.from_string(sql_templ,globals={"uuid":uuid})

	 这样生成100w需要12.9 秒,而golang 生成只要1.3 秒，相差9倍

	 总结，golang 和python 都能够支持，先编译模板，然后再拿模板取渲染值。python 里面直接调用函数，是使用globals 传参，之前传个空字符串，写法有点
	 问题。也只是针对uuid 有影响，针对其他时间格式化，yyyy-mm-dd 还是一样能用。
	 如果都改成模板先编译，运行是渲染，代码性能优化个10倍以上

	*/

}
