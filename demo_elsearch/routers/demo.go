package routers

import (
	"context"
	"demo_elsearch/moudle"
	"fmt"
	"github.com/olivere/elastic"
	"reflect"
	"strconv"
	"time"
)

var client *elastic.Client

//初始化
func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200/"))
	if err != nil {
		fmt.Println("create client failed, err: %v", err)
	}
}

//ping 连接测试
func PingNode() {
	start := time.Now()
	info, code, err := client.Ping("http://127.0.0.1:9200/").Do(context.Background())
	if err != nil {
		fmt.Println("ping es failed, err: %v", err)
	}
	duration := time.Since(start)
	fmt.Println("cost time: %v\n", duration)
	fmt.Println("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

//校验 index 是否存在
func IndexExist(index ...string) bool {
	exists, err := client.IndexExists(index...).Do(context.Background())
	if err != nil {
		fmt.Println("%v\n", err)
	}
	return exists
}

//创建 index
func CreateIndex(index, mapping string) {
	result, err := client.Index().
		Index(index).
		Type("doc").
		Id("1").
		BodyJson(mapping).
		Do(context.Background())
	if err != nil {
		fmt.Println("create index failed, err: %v\n", err)
	}else{
		fmt.Println(result.Index,result.Type,result.Id)
	}
}

//删除 index
func DelIndex(index... string) bool {
	response, err := client.DeleteIndex(index...).Do(context.Background())
	if err != nil {
		fmt.Println("delete index failed, err: %v\n", err)
	}
	return response.Acknowledged
}

//批量插入
func Batch(index string, type_ string, datas... interface{})  {
	bulkRequest := client.Bulk()
	for i, data := range datas {
		//fmt.Println(i)
		//fmt.Println(data)
		doc := elastic.NewBulkIndexRequest().Index(index).Type(type_).Id(strconv.Itoa(i)).Doc(data)
		bulkRequest = bulkRequest.Add(doc)
	}
	response, err := bulkRequest.Do(context.TODO())
	if err != nil {
		panic(err)
	}
	failed := response.Failed()
	iter := len(failed)
	fmt.Println("error: %v, %v\n", response.Errors,  iter)
}

//修改
func Update(index,type_,id string,docs map[string]interface{})  {
	res, err := client.Update().
		Index(index).
		Type(type_).
		Id(id).
		Doc(docs).
		Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("update age %s\n", res)
}

//获取指定 Id 的文档
func GetDoc(index, id string) []byte {
	temp := client.Get().Index(index).Id(id)
	get, err := temp.Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get.Found {
		fmt.Println("Got document %s in version %d from index %s, type %s\n", get.Id, get.Version, get.Index, get.Type)
	}
	source, err := get.Source.MarshalJSON()
	if err != nil {
		fmt.Println("byte convert string failed, err: %v", err)
	}
	return source
}

//简单分页,获取所有不设置size，from参数，默认获取10条
func List(size,page int) {
	if size < 0 || page < 1 {
		fmt.Println("param error")
		return
	}
	res,err := client.Search("twitter").
		Type("doc").
		Size(size).
		From((page-1)*size).
		Do(context.Background())
	Querytwitter(res, err)

}

//查询到index/twitter
func Querytwitter(res *elastic.SearchResult, err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var typ moudle.Tweet
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(moudle.Tweet)
		fmt.Println("%#v\n", t)
	}
}

//简单分页,获取所有不设置size，from参数，默认获取10条
func Lists (){
	res,err := client.Search("twitter").
		Type("doc").
		Do(context.Background())
	Querytwitters(res, err)

}

//查询到index/twitter
func Querytwitters(res *elastic.SearchResult, err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var typ moudle.Tweet
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(moudle.Tweet)
		fmt.Println("%#v\n", t)
	}
}


//term 查询,精确匹配要采用TermQuery
func TermQuery(index, type_, fieldName, fieldValue string) *elastic.SearchResult {
		//termQuery：不会对搜索词进行分词处理，而是作为一个整体与目标字段进行匹配，若完全匹配，则可查询到。
		query := elastic.NewTermQuery(fieldName, fieldValue)
		//_ = elastic.NewQueryStringQuery(fieldValue) //关键字查询
		searchResult, err := client.Search().
		Index(index).Type(type_).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())

		if err != nil {
		panic(err)
	}
	    fmt.Println("query cost %d millisecond.\n", searchResult)
		fmt.Println("query cost %d millisecond.\n", searchResult.TookInMillis)
		fmt.Println("query cost %d millisecond.\n", searchResult.Clusters)

		return searchResult
	}
//条件查询,年龄大于30岁的
//Range Query （范围查询）,Exist Query （是否存在查询）,Prefix Query （前缀查询）,Wildcard Query （通配符查询）
//Search 是我们最常用的 API,模糊匹配 Match，字段判空 Exists，精准匹配 Term 和 Terms，范围匹配 Range
func Search(index, type_ string) *elastic.SearchResult {
	boolQuery := elastic.NewBoolQuery()
	//matchQuery：会将搜索词分词，再与目标查询字段进行匹配，若分词中的任意一个词与目标字段匹配上，则可查询到。
	boolQuery.Must(elastic.NewMatchQuery("user", "Jame1"))
	//（范围查询）
	boolQuery.Filter(elastic.NewRangeQuery("age").Gt("88"))
	searchResult, err := client.Search(index).
		Type(type_).Query(boolQuery).Pretty(true).Do(context.Background())
	if err != nil {
		panic(err)
	}

	return searchResult
}
func AggsSearch(index, type_ string) {

	minAgg := elastic.NewMinAggregation().Field("age")
	rangeAgg := elastic.NewRangeAggregation().Field("age").AddRange(0, 30).AddRange(30, 60).Gt(60)

	build := client.Search(index).Type(type_).Pretty(true)

	minResult, err := build.Aggregation("minAgg", minAgg).Do(context.Background())
	rangeResult, err := build.Aggregation("rangeAgg", rangeAgg).Do(context.Background())
	if err != nil {
		panic(err)
	}

	minAggRes, _ := minResult.Aggregations.Min("minAgg")
	fmt.Println("min: %v\n", *minAggRes.Value)

	rangeAggRes, _ := rangeResult.Aggregations.Range("rangeAgg")
	for _, item := range rangeAggRes.Buckets {
		fmt.Println("key: %s, value: %v\n", item.Key, item.DocCount)
	}
}


