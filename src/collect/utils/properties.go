package collect

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Properties struct {
	FileName string                 // 文件路径
	Data     map[string]interface{} // 数据
}

func init() {
	// 由于启动有要连接数据库，直接先进行导入配置了
	LoadAppProperties("./conf/application.properties")
}

var p Properties

func LoadAppProperties(path string) {
	properties := Properties{FileName: path}
	properties.Load()
	p = properties

}
func GetAppKeyWithDefault(key string, defaultValue string) string {
	if !p.HasKey(key) {
		return defaultValue
	}
	return p.GetKey(key)
}
func GetAppKey(key string) string {
	if !p.HasKey(key) {
		log.Println("没有找到配置" + key)
	}
	return p.GetKey(key)

}

/**
** 获取配置
 */
func (p *Properties) Load() {

	srcFile, err := os.OpenFile(p.FileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Println(p.FileName + "文件不存在")
		return
	}
	srcReader := bufio.NewReader(srcFile)
	p.Data = make(map[string]interface{})
	for {
		str, _, err := srcReader.ReadLine()
		content := Strval(str)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if IsValueEmpty(content) {
			continue
		}
		content = strings.Trim(content, " ")
		if strings.Contains(content, "=") && !strings.HasPrefix(content, "#") {
			content = strings.ReplaceAll(content, "\\n", "\n")
			content = strings.ReplaceAll(content, "\\t", "\t")
			contentArr := strings.Split(content, "=")
			key := contentArr[0]
			value := strings.Join(contentArr[1:], "=")
			p.Data[key] = value
		}

	}
	defer srcFile.Close()

}

/**
** 获取配置
 */
func (p *Properties) GetKey(key string) string {
	return Strval(GetSafeData(key, p.Data))
}

/*
** 判断是否有key
 */
func (p *Properties) HasKey(key string) bool {
	if _, ok := p.Data[key]; !ok {
		return false
	} else {
		return true
	}

}
