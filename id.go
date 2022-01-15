package hid

import (
	"github.com/hlib-go/snowflake"
	"math/big"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var g *snowflake.Node

func init() {
	g, _ = snowflake.NewNode(getLocalIP() % 31)
}

//ID SnowFlake 生成长度20位的数字编号
func ID() string {
	return g.Generate().String() + strconv.Itoa(rand.Intn(10))
}

// 获取所有ip转int64后的和
func getLocalIP() (num int64) {
	var (
		err     error
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet
		isIpNet bool
	)
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil || num == 0 {
			num = time.Now().UnixNano()
		}
	}()
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	for _, addr = range addrs {
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				var ipv4 = ipNet.IP.String() // 192.168.1.1
				num = num + inetAtoN(ipv4)
			}
		}
	}
	return
}

func inetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}
