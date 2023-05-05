/* !!
 * File: util.go
 * File Created: Thursday, 20th May 2021 10:33:24 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 20th May 2021 10:35:10 am
 
 */

package elastic_client

import (
	"fmt"
	"strings"
)

func GetPrefixIndex() string {
	if esConf == nil {
		err := fmt.Errorf("need config for elastic client first")
		panic(err)
	}

	return strings.ToLower(fmt.Sprintf("%s.%s-%s", esConf.AppKey, esConf.PlatformDefault, esConf.LogTypeDefault))
}
