/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2021/4/27 16:10
*/

package main

import (
	"fmt"
	"rosefintech-rosl2/common"
	"rosefintech-rosl2/contractInterface"
	"rosefintech-rosl2/server"
	"rosefintech-rosl2/tools"
)

func main() {
	tools.LogInit(common.Conf)
	var ch = make(chan int)
	for _, v := range common.Conf.Contract.FundManager {
		contractInterface.UploadOpenPosition(v)
		contractInterface.UploadOutPosition(v)
		go contractInterface.WatchOpenPosition(v)
		go contractInterface.WatchOutPosition(v)
	}

	httpService := service.NewService(":1234")
	for {
		select {
		case <-ch:
			fmt.Println("service success")
			err := httpService.Close()
			if err != nil {
				panic(err)
			}
		}
	}

}
