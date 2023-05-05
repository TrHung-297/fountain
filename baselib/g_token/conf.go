/* !!
 * File: conf.go
 * File Created: Friday, 25th June 2021 10:06:52 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Friday, 25th June 2021 10:06:52 am
 
 */

package g_token

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

type JWTConfig struct {
	SecretKey       string `json:"SecretKey,omitempty"`
	AccessTokenTTL  int    `json:"AccessTokenTTL,omitempty"`
	RefreshTokenTTL int    `json:"RefreshTokenTTL,omitempty"`
}

var conf *JWTConfig

func getJWTConfigFromEnv() {
	conf = &JWTConfig{}

	if err := viper.UnmarshalKey("OpenIDJwt", conf); err != nil {
		err := fmt.Errorf("not found config name with env %q for OpenIDJwt with error: %+v", "OpenIDJwt", err)
		g_log.V(1).WithError(err).Errorf("getJWTConfigFromEnv - Error: %v", err)

		return
	}
}
