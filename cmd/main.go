package main

import (
	"fmt"
	"github.com/OpenIMSDK/getcdv3"
)

func main() {
	opID := ""
	etcdConn, err := getcdv3.NewEtcdConn("", "", "", "", "", 10000, 10)
	err = etcdConn.RegisterEtcd(opID, "group")
	if err != nil {
		panic(err.Error())
	}

	etcdConn.SetDefaultEtcdConfig("user", []int{9000, 9001})
	userRpc := etcdConn.GetConn(opID, "user")
	fmt.Println(userRpc)

	gates := etcdConn.GetDefaultGatewayConn4Unique(opID, "gateway")
	fmt.Println(gates)

	etcdConn, err = getcdv3.NewEtcdConn("", "", "", "", "", 10000, 10)
	err = etcdConn.RegisterEtcd4Unique("", "gateway")

}
