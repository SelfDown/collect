package collect

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func init() {
	// 如果集群重复，这里machineID 可以改成从配置文件取
	if err := Init("2021-12-03", 1); err != nil {
		fmt.Println("Init() failed, err = ", err)
		return
	}

}
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// SnowID 生成 64 位的 雪花 ID
func SnowID() int64 {

	return node.Generate().Int64()
}
