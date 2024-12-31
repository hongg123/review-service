package snowflake

import (
	"errors"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

// 利用雪花算法生成id

/*
在公司中如果有生成唯一ID的服务可以直接对接，
本质上只是为了获取一个全局唯一ID。
*/

var (
	InvalidInitParamErr  = errors.New("snowflake初始化失败，无效的startTime或machineID")
	InvalidTimeFormatErr = errors.New("snowflake初始化失败，无效的startTime格式")
)

var node *sf.Node

// Init 雪花算法初始化配置
/* machineID：由于实现的是分布式服务，所以需要传入节点机器ID */
func Init(startTime string, machineID int64) (err error) {
	if len(startTime) == 0 || machineID <= 0 {
		return InvalidInitParamErr
	}
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return InvalidTimeFormatErr
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// 调用此方法 GenID 生一个ID
func GenID() int64 {
	return node.Generate().Int64()
}
