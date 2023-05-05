/* !!
 * File: md5_file.go
 * File Created: Thursday, 27th May 2021 10:19:57 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:30:31 am
 
 */

package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
)

func CalcMd5File(filename string) (string, error) {
	// fileName := core.NBFS_DATA_PATH + m.data.FilePath
	f, err := os.Open(filename)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("CalcMd5File - Error: %+v", err)
		return "", err
	}

	defer f.Close()

	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, f); err != nil {
		// fmt.Println("Copy", err)
		g_log.V(1).WithError(err).Errorf("CalcMd5File - Copy error: %+v", err)
		return "", err
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}
