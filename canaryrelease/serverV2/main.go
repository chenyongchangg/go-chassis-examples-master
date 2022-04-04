package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/canaryrelease/serverV2/resource"
	"github.com/go-mesh/openlogging"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rpc/server/

func main() {
	chassis.RegisterSchema("rest", &resource.RestFulRouterB{})
	if err := chassis.Init(); err != nil {
		openlogging.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
