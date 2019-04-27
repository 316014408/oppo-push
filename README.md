# oppo-push
oppo 消息推送 push server sdk for golang

```Go
package main

import (
	"fmt"

	oppopush "github.com/316014408/oppo-push"
)

var client = oppopush.NewClient("appKey", "masterSecret")

func main() {

	//保存通知栏消息内容体
	msg0 := oppopush.NewSaveMessageContent("hi baby1", "hi1").
		SetSubTitle("hahaha1")
	result, err := client.SaveMessageContent(msg0)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("MessageID:", result.Data.MessageID)
		//广播推送-通知栏消息
		broadcast := oppopush.NewBroadcast(result.Data.MessageID).
			SetTargetType(2).
			SetTargetValue("CN_2DFC77B0D34EFDACA377F92554CBE4AB")
		result, err := client.Broadcast(broadcast)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("MessageID:", result.Data.MessageID)
			fmt.Println("TaskId:", result.Data.TaskId)
		}
	}

	//单推
	unicast := oppopush.NewMessage("hi baby2", "hi2").
		SetSubTitle("hahaha2").
		SetTargetType(2).
		SetTargetValue("CN_2DFC77B0D34EFDACA377F92554CBE4AB")
	result2, err := client.Unicast(unicast)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("MessageID:", result2.Data.MessageID)
	}

	//批量单推
	msg1 := oppopush.NewMessage("hi baby1", "hi1").
		SetSubTitle("hahaha1").
		SetTargetType(2).
		SetTargetValue("CN_2DFC77B0D34EFDACA377F92554CBE4AB")
	msg2 := oppopush.NewMessage("hi baby2", "hi2").
		SetSubTitle("hahaha2").
		SetTargetType(2).
		SetTargetValue("CN_2DFC77B0D34EFDACA377F92554CBE4AB")
	msg := []oppopush.Message{}
	msg = append(msg, *msg1)
	msg = append(msg, *msg2)
	result3, err := client.UnicastBatch(msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result3.Data)
	}

	//获取失效的registration_id列表
	result4, err := client.FetchInvalidRegidList()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("RegistrationIds:", result4.Data.RegistrationIds)
		fmt.Println("TotalCount:", result4.Data.TotalCount)
	}
}
```