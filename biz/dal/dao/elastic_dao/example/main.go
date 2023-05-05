
package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/elastic_client"
)

func main() {
	esConf := elastic_client.SetESConfig("gtv_log_event", "g_plus2", "event_log")
	timeStart := 0 // 1522221200
	timeEnd := 0   // 1620061200

	dateStart := time.Unix(int64(timeStart), 0)
	dateStartYear, dateStartMonth, dateStartDay := dateStart.Date()
	log.Printf("dateStartYear: %d, dateStartMonth: %d, dateStartDay: %d", dateStartYear, dateStartMonth, dateStartDay)

	dateEnd := time.Unix(int64(timeEnd), 0)
	dateEndYear, dateEndMonth, dateEndDay := dateEnd.Date()
	log.Printf("dateEndYear: %d, dateEndMonth: %d, dateEndDay: %d", dateEndYear, dateEndMonth, dateEndDay)

	listQueryString := make([]string, 0)
	if dateStartYear == dateEndYear {
		if dateStartMonth == dateEndMonth {
			for i := dateStartDay; i <= dateEndDay; i++ {
				listQueryString = append(listQueryString, strings.ToLower(fmt.Sprintf("%s.%s-%s-%d.%s.%s", esConf.AppKey, esConf.PlatformDefault, esConf.LogTypeDefault, dateEndYear, base.StandardizedNumber(int(dateEndMonth), 2), base.StandardizedNumber(i, 2))))
			}
		} else {
			for i := dateStartMonth; i <= dateEndMonth; i++ {
				listQueryString = append(listQueryString, strings.ToLower(fmt.Sprintf("%s.%s-%s-%d.%s.*", esConf.AppKey, esConf.PlatformDefault, esConf.LogTypeDefault, dateEndYear, base.StandardizedNumber(int(i), 2))))
			}
		}
	} else {
		for i := dateStartYear; i <= dateEndYear; i++ {
			listQueryString = append(listQueryString, strings.ToLower(fmt.Sprintf("%s.%s-%s-%s.*", esConf.AppKey, esConf.PlatformDefault, esConf.LogTypeDefault, base.StandardizedNumber(i, 4))))
		}
	}

	log.Printf("listDay: %s", base.JSONDebugDataString(listQueryString))
}
