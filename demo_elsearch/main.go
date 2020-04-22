package main

import (
	"demo_elsearch/moudle"
	"demo_elsearch/routers"
	"fmt"
	"reflect"
)

func main()  {

	//ping 连接测试
	//routers.PingNode()

	//校验 index 是否存在
	//result := routers.IndexExist("car_source", "test")
	//fmt.Println("all index exists: ", result)

	//var mapping = `{
	//				"settings":{
	//					"number_of_shards": 3,
	//					"number_of_replicas": 1
	//				},
	//				"mappings":{
	//					"doc":{
	//						"properties":{
	//							"user":{
	//								"type":"keyword"
	//							},
	//							"age":{
	//								"type": "integer"
	//							},
	//							"message":{
	//								"type":"text",
	//								"store": true,
	//								"fielddata": true
	//							},
	//							"image":{
	//								"type":"keyword"
	//							},
	//							"created":{
	//								"type":"date"
	//							},
	//							"tags":{
	//								"type":"keyword"
	//							},
	//							"location":{
	//								"type":"geo_point"
	//							},
	//							"suggest_field":{
	//								"type":"completion"
	//							}
	//						}
	//					}
	//				}
	//			}`

	//创建 index
	//routers.CreateIndex("twitter", mapping)

	//删除 index
	//result := routers.DelIndex("twitter")
	//fmt.Println("all index deleted: ", result)

	//批量插入
	//tweet1 := moudle.Tweet{User: "Jame1",Age: 23, Message: "Take One", Retweets: 1, Created: time.Now()}
	//tweet2 := moudle.Tweet{User: "Jame2",Age: 32, Message: "Take Two", Retweets: 0, Created: time.Now()}
	//tweet3 := moudle.Tweet{User: "Jame3",Age: 32, Message: "Take Three", Retweets: 0, Created: time.Now()}
	//routers.Batch("twitter", "doc", tweet1, tweet2, tweet3)

	//修改
	//routers.Update("twitter", "doc","1",map[string]interface{}{"age" : 88})

	//获取指定 Id 的文档
	//var tweet moudle.Tweet
	//data := routers.GetDoc("twitter", "2")
	//if err := json.Unmarshal(data, &tweet); err == nil {
	//	fmt.Println("data: %v\n", tweet)
	//}

	//简单分页
	//routers.List(1,2)

	//routers.Lists()

	//term 精确匹配查询
	//var tweet moudle.Tweet
	//result := routers.TermQuery("twitter", "doc", "user", "Take Two")
	//for _, item := range result.Each(reflect.TypeOf(tweet)) {
	//	if t, ok := item.(moudle.Tweet); ok {
	//		fmt.Println("tweet : %v\n", t)
	//	}
	//}

	//Range Query （范围查询）,Exist Query （是否存在查询）,Prefix Query （前缀查询）,Wildcard Query （通配符查询）
	//Search 是我们最常用的 API,模糊匹配 Match，字段判空 Exists，精准匹配 Term 和 Terms，范围匹配 Range
	result := routers.Search("twitter", "doc")
	fmt.Println(result)
	var tweet moudle.Tweet
	for _, item := range result.Each(reflect.TypeOf(tweet)) {
		if t, ok := item.(moudle.Tweet); ok {
			fmt.Println("tweet : %v\n", t)
		}
	}

	//routers.AggsSearch("twitter", "doc")
}



