package elasticsearch

import (
	"context"
	"github.com/olivere/elastic/v7"
	"ihome/service/house/conf"
	"ihome/service/utils"
	"log"
	"os"
	"reflect"
	"time"
)

type HouseElasticSearch struct {
	Client  *elastic.Client
	Index   string
	Mapping string
}

var HouseES *HouseElasticSearch
var ctx = context.Background()

//初始化es驱动
func init() {
	HouseES = &HouseElasticSearch{}
	errorLog := log.New(os.Stdout, "app", log.LstdFlags)
	var err error
	client, err := elastic.NewClient(
		elastic.SetInfoLog(errorLog),
		elastic.SetURL(conf.HouseESHost),
		elastic.SetSniff(false),
		//elastic.SetHealthcheck(false),
	)
	if err != nil {
		utils.NewLog().Error("elastic.NewClient error:", err)
		return
	}
	info, code, err := client.Ping(conf.HouseESHost).Do(context.Background())
	if err != nil {
		utils.NewLog().Error("elastic.NewClient error:", err)
		return
	}
	utils.NewLog().Infof("Es return with code %d and version %s \n", code, info.Version.Number)
	esVersionCode, err := client.ElasticsearchVersion(conf.HouseESHost)
	if err != nil {
		utils.NewLog().Error("elastic.NewClient error:", err)
	}
	utils.NewLog().Infof("es version %s\n", esVersionCode)
	//创建House index
	mappingTpl := `{
    "mappings": {
        "properties": {
            "user_id": {
                "type": "long"
            },
            "user_name": {
                "type": "text"
            },
            "address": {
                "type": "text"
            },
            "area_id": {
                "type": "long"
            },
            "area_name": {
                "type": "keyword"
            },
            "acreage": {
                "type": "integer"
            },
            "unit": {
                "type": "text"
            },
            "capacity": {
                "type": "integer"
            },
            "beds": {
                "type": "keyword"
            },
            "ctime": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
            },
            "deposit": {
                "type": "integer"
            },
            "min_days": {
                "type": "integer"
            },
            "max_days": {
                "type": "integer"
            },
            "house_id": {
                "type": "long"
            },
            "img_url": {
                "type": "text"
            },
            "order_count": {
                "type": "integer"
            },
            "price": {
                "type": "integer"
            },
            "room_count": {
                "type": "integer"
            },
            "title": {
                "type": "text"
            },
            "user_avatar": {
                "type": "text"
            }
        }
    }
}`
	index := "ihome_house"
	//创建
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		utils.NewLog().Errorf("userEs init exist failed err is %s\n", err)
		return
	}
	if !exists {
		_, err := client.CreateIndex(index).Body(mappingTpl).Do(ctx)
		if err != nil {
			utils.NewLog().Errorf("houseEs init exist failed err is %s\n", err)
			return
		}
	}
	HouseES.Client = client
	HouseES.Index = index
	HouseES.Mapping = mappingTpl

}

func (es *HouseElasticSearch) BatchAdd(ctx context.Context, housePos []*ESHousePo) error {
	var err error
	for i := 0; i < conf.ESTryTimeLimit; i++ {
		if err = es.batchAdd(ctx, housePos); err != nil {
			utils.NewLog().Info("batch add failed: ", err)
			time.Sleep((2 << i) * time.Second)
			continue
		}
		return err
	}
	return err
}
func (es *HouseElasticSearch) batchAdd(ctx context.Context, housePos []*ESHousePo) error {
	req := es.Client.Bulk().Index(es.Index)
	for _, house := range housePos {
		utils.NewLog().Debug("house:", house)
		doc := elastic.NewBulkIndexRequest().Type("_doc").Id(utils.IntToString(house.HouseId)).Doc(house)
		req.Add(doc)
	}
	if req.NumberOfActions() < 0 {
		utils.NewLog().Debug("req.NumberOfActions...")
		return nil
	}
	if _, err := req.Do(ctx); err != nil {
		return err
	}
	utils.NewLog().Debug("batchAdd ends...")
	return nil
}

func (es *HouseElasticSearch) BatchUpdate(ctx context.Context, housePos []map[string]interface{}) error {
	var err error
	for i := 0; i < conf.ESTryTimeLimit; i++ {
		if err = es.batchUpdate(ctx, housePos); err != nil {
			utils.NewLog().Info("batch add failed: ", err)
			time.Sleep((2 << i) * time.Second)
			continue
		}
		return err
	}
	return err
}

func (es *HouseElasticSearch) batchUpdate(ctx context.Context, housePos []map[string]interface{}) error {
	req := es.Client.Bulk().Index(es.Index)
	for _, house := range housePos {
		doc := elastic.NewBulkUpdateRequest().Id(utils.IntToString(house["house_id"])).Doc(house)
		req.Add(doc)
	}

	if req.NumberOfActions() < 0 {
		return nil
	}
	if _, err := req.Do(ctx); err != nil {
		return err
	}
	return nil
}

func (es *HouseElasticSearch) BatchDel(ctx context.Context, housePos []*ESHousePo) error {
	var err error
	for i := 0; i < conf.ESTryTimeLimit; i++ {
		if err = es.batchDel(ctx, housePos); err != nil {
			utils.NewLog().Info("batch add failed: ", err)
			time.Sleep((2 << i) * time.Second)
			continue
		}
		return err
	}
	return err
}

func (es *HouseElasticSearch) batchDel(ctx context.Context, housePos []*ESHousePo) error {
	req := es.Client.Bulk().Index(es.Index)
	for _, house := range housePos {
		doc := elastic.NewBulkDeleteRequest().Id(utils.IntToString(house.HouseId))
		req.Add(doc)
	}

	if req.NumberOfActions() < 0 {
		return nil
	}

	if _, err := req.Do(ctx); err != nil {
		return err
	}
	return nil
}

//ToFilter 查询
func (r *HouseSearchReq) ToFilter() *ESSearch {

	var search ESSearch
	if r.AreaId != 0 {
		search.MustQuery = append(search.MustQuery, elastic.NewMatchQuery("area_id", r.AreaId))
	}
	if r.Days != 0 {
		// min_days<=days<=max_days
		search.MustQuery = append(search.MustQuery, elastic.NewRangeQuery("min_days").Lte(r.Days))
		search.MustQuery = append(search.MustQuery, elastic.NewRangeQuery("max_days").Gte(r.Days))

	}
	if r.MinPrice != 0 {
		search.MustQuery = append(search.MustQuery, elastic.NewRangeQuery("price").Gte(r.MinPrice))
	}
	if r.MaxPrice != 0 {
		search.MustQuery = append(search.MustQuery, elastic.NewRangeQuery("price").Lte(r.MaxPrice))
	}

	search.From = (r.Page - 1) * r.PageSize
	search.Size = r.PageSize
	return &search
}

func (es *HouseElasticSearch) Search(ctx context.Context, filter *ESSearch) ([]*ESHousePo, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(filter.MustQuery...)
	boolQuery.MustNot(filter.MustNotQuery...)
	boolQuery.Should(filter.ShouldQuery...)
	boolQuery.Filter(filter.Filters...)

	// 当should不为空时，保证至少匹配should中的一项
	if len(filter.MustQuery) == 0 && len(filter.MustNotQuery) == 0 && len(filter.ShouldQuery) > 0 {
		boolQuery.MinimumShouldMatch("1")
	}
	service := es.Client.Search().Index(es.Index).Query(boolQuery).SortBy(filter.Sorters...).From(filter.From).Size(filter.Size)
	resp, err := service.Do(ctx)
	if err != nil {
		return nil, err
	}
	if resp.TotalHits() == 0 {
		return nil, nil
	}
	houseES := make([]*ESHousePo, 0)
	for _, e := range resp.Each(reflect.TypeOf(&ESHousePo{})) {
		house := e.(*ESHousePo)
		houseES = append(houseES, house)
	}
	return houseES, nil
}
