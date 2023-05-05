

package elastic_dao

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/elastic_client"
)

// ESResponse type;
// Include count, hits Total and aggregations
type ESResponse struct {
	Count int `json:"count,omitempty"`
	Hits  struct {
		Total struct {
			Value int `json:"value,omitempty"`
		} `json:"Total,omitempty"`
	} `json:"hits,omitempty"`
	Aggregations struct {
		CountBy struct {
			Buckets []struct {
				KeyAsString string `json:"key_as_string,omitempty"`
				Key         int    `json:"key,omitempty"`
				DocsCount   int    `json:"doc_count,omitempty"`
			} `json:"buckets,omitempty"`
		} `json:"count_by,omitempty"`
	} `json:"aggregations,omitempty"`
}

// ElasticLogEventDAO type
type ElasticLogEventDAO struct {
	client *elastic.Client
}

// NewElasticLogEventDAO func
func NewElasticLogEventDAO(client *elastic.Client) *ElasticLogEventDAO {
	return &ElasticLogEventDAO{client}
}

func getElasticLogEventIndex(timeStart, timeEnd int32) []string {
	esConf := elastic_client.GetESConfig()
	patten := elastic_client.GetPrefixIndex()

	appKey := esConf.AppKey
	if appKey == "" {
		appKey = "*"
	}

	platform := esConf.PlatformDefault
	if platform == "" {
		platform = "*"
	}

	logType := esConf.LogTypeDefault
	if logType == "" {
		logType = "*"
	}

	if timeStart == timeEnd && timeEnd == 0 {
		return []string{fmt.Sprintf("%s-*", patten)}
	}

	if timeStart != 0 && timeEnd == 0 {
		timeEnd = int32(time.Now().Unix())
	}

	dateStart := time.Unix(int64(timeStart), 0)
	dateStartYear, dateStartMonth, dateStartDay := dateStart.Date()
	g_log.V(5).Infof("dateStartYear: %d, dateStartMonth: %d, dateStartDay: %d", dateStartYear, dateStartMonth, dateStartDay)

	dateEnd := time.Unix(int64(timeEnd), 0)
	dateEndYear, dateEndMonth, dateEndDay := dateEnd.Date()
	g_log.V(5).Infof("dateEndYear: %d, dateEndMonth: %d, dateEndDay: %d", dateEndYear, dateEndMonth, dateEndDay)

	listQueryString := make([]string, 0)
	if dateStartYear == dateEndYear {
		if dateStartMonth == dateEndMonth {
			// for i := dateStartDay; i <= dateEndDay; i++ {
			// listQueryString = append(listQueryString, fmt.Sprintf("%s-%d.%s.%s", patten, dateEndYear, base.StandardizedNumber(int(dateEndMonth), 2), base.StandardizedNumber(i, 2)))
			// }
			listQueryString = append(listQueryString, fmt.Sprintf("%s-%d.%s.*", patten, dateEndYear, base.StandardizedNumber(int(dateEndMonth), 2)))
		} else {
			for i := dateStartMonth; i <= dateEndMonth; i++ {
				listQueryString = append(listQueryString, fmt.Sprintf("%s-%d.%s.*", patten, dateEndYear, base.StandardizedNumber(int(i), 2)))
			}
		}
	} else {
		for i := dateStartYear; i <= dateEndYear; i++ {
			listQueryString = append(listQueryString, fmt.Sprintf("%s-%s.*", patten, base.StandardizedNumber(i, 4)))
		}
	}

	return listQueryString
}

// CountLogEvent func;
// Params timeStart, timeEnd is second range, queries is query like `EventType AND InstallUTM`; `EventType AND InstallUTM` like "EventType": "InstallUTM" in json
func (dao *ElasticLogEventDAO) CountLogEvent(timeStart, timeEnd int32, bodyQuery io.Reader) int {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	listQueries := make([]func(*esapi.CountRequest), 0)

	queriesString := getElasticLogEventIndex(timeStart, timeEnd)
	g_log.V(5).Infof("ElasticLogEventDAO::CountLogEvent - QueriesString: %s", base.JSONDebugDataString(queriesString))

	for _, query := range queriesString {
		listQueries = append(listQueries, dao.client.Count.WithIndex(query))
	}

	listQueries = append(listQueries, dao.client.Count.WithBody(bodyQuery))
	listQueries = append(listQueries, dao.client.Count.WithContext(myCtx))

	res, err := dao.client.Count(listQueries...)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("ElasticLogEventDAO::CountLogEvent - Count request error: %+v", err)

		return 0
	}

	dataLog, err := ioutil.ReadAll(res.Body)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("ElasticLogEventDAO::CountLogEvent - Can not parse the body of response error: %+v", err)

		return 0
	}

	if res.IsError() {
		g_log.V(1).WithError(err).Errorf("ElasticLogEventDAO::CountLogEvent - Count request is error: %s", dataLog)

		return 0
	}

	result := &ESResponse{}
	if err := json.Unmarshal(dataLog, result); err != nil {
		g_log.V(1).WithError(err).Errorf("ElasticLogEventDAO::CountLogEvent - Can not parse the response error: %+v", err)

		return 0
	}

	return result.Count

}

// CountQueryLogEvent func;
// Params timeStart, timeEnd is second range, queries is query like `EventType AND InstallUTM`; `EventType AND InstallUTM` like "EventType": "InstallUTM" in json
func (dao *ElasticLogEventDAO) CountQueryLogEvent(timeStart, timeEnd int32, bodyQuery io.Reader) map[string]int {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	// myBuf := make([]byte, 0)
	// w := bytes.NewBuffer(myBuf)
	// cl, e := io.Copy(w, bodyQuery)
	// g_log.V(1).Infof("ElasticLogEventDAO::CountQueryLogEvent - bodyQuery: %+v", bodyQuery)
	// g_log.V(1).Infof("ElasticLogEventDAO::CountQueryLogEvent - Writer: %s", myBuf)
	// g_log.V(1).Infof("ElasticLogEventDAO::CountQueryLogEvent - Clone: %d", cl)
	// g_log.V(1).Infof("ElasticLogEventDAO::CountQueryLogEvent - Error: %+v", e)

	listQueries := make([]func(*esapi.SearchRequest), 0)

	queriesString := getElasticLogEventIndex(timeStart, timeEnd)
	g_log.V(5).Infof("ElasticLogEventDAO::CountQueryLogEvent - QueriesString: %s", base.JSONDebugDataString(queriesString))
	for _, query := range queriesString {
		listQueries = append(listQueries, dao.client.Search.WithIndex(query))
	}

	listQueries = append(listQueries, dao.client.Search.WithBody(bodyQuery))
	listQueries = append(listQueries, dao.client.Search.WithContext(myCtx))

	result := make(map[string]int)

	res, err := dao.client.Search(listQueries...)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("ElasticLogEventDAO::CountQueryLogEvent - Count request error: %+v", err)

		return result
	}

	dataLog, err := ioutil.ReadAll(res.Body)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("ElasticLogEventDAO::CountQueryLogEvent - Can not parse the body of response error: %+v", err)

		return result
	}

	if res.IsError() {
		// q, e := ioutil.ReadAll(w)
		// g_log.V(1).Infof("ElasticLogEventDAO::CountQueryLogEvent Query - Error: %v", e)
		// g_log.V(1).Infof("ElasticLogEventDAO::CountQueryLogEvent Query - Query: %s", q)

		g_log.V(1).WithError(err).Errorf("ElasticLogEventDAO::CountQueryLogEvent - Query count request is error: %s", dataLog)

		return result
	}

	resp := new(ESResponse)
	if err := json.Unmarshal(dataLog, resp); err != nil {
		g_log.V(1).WithError(err).Errorf("ElasticLogEventDAO::CountQueryLogEvent - Can not parse the response error: %+v", err)

		return result
	}

	result["Total"] = resp.Hits.Total.Value
	for _, agg := range resp.Aggregations.CountBy.Buckets {
		result[fmt.Sprintf("%d", agg.Key)] = agg.DocsCount
	}

	return result
}
