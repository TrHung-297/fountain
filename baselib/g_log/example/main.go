/* !!
 * File: main.go
 * File Created: Wednesday, 7th July 2021 3:48:18 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 7th July 2021 3:48:18 pm
 
 */

package main

import (
	"fmt"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
)

func main() {
	g_log.LogFileLevel(5)
	g_log.LogPrintLevel(3)
	g_log.LogDir("./")
	g_log.ServiceName("test")

	g_log.V(1).Infof("Time: %s", time.Now().Format("02-01-2006 15:04:05.000000"))

	// for i := 0; i < 10000; i++ {
	g_log.V(1).WithError(fmt.Errorf("error ne")).Infof("log loi ne")
	g_log.Infof("haha")
	g_log.Infof("haha without level")
	g_log.V(1).WithError(fmt.Errorf("error ne")).Infof("log info voi loi ne voi level = 1")
	g_log.V(2).WithError(fmt.Errorf("error ne")).Infof("log info voi loi ne voi level = 2")
	g_log.V(3).WithError(fmt.Errorf("error ne")).Infof("log info voi loi ne voi level = 3")
	g_log.V(4).WithError(fmt.Errorf("error ne")).Infof("log info voi loi ne voi level = 4")
	g_log.V(5).WithError(fmt.Errorf("error ne")).Infof("log info voi loi ne voi level = 5")
	g_log.Infof("<p>haha without level again</p>")

	g_log.V(1).WithError(fmt.Errorf("error ne")).Errorf("log loi ne voi level = %d", 1)
	g_log.V(2).WithError(fmt.Errorf("error ne")).Errorf("log loi ne voi level = %d", 2)
	g_log.V(3).WithError(fmt.Errorf("error ne")).Errorf("log loi ne voi level = %d", 3)
	g_log.V(4).WithError(fmt.Errorf("error ne")).Errorf("log loi ne voi level = %d", 4)
	g_log.V(5).WithError(fmt.Errorf("error ne")).Errorf("log loi ne voi level = %d", 5)

	g_log.V(1).Errorf("log k loi ne voi level = %d", 1)
	g_log.V(2).Errorf("log k loi ne voi level = %d", 2)
	g_log.V(3).Errorf("log k loi ne voi level = %d", 3)
	g_log.V(4).Errorf("log k loi ne voi level = %d", 4)
	g_log.V(5).Errorf("log k loi ne voi level = %d", 5)
	// }

}
