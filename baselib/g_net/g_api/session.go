/* !!
 * File: session.go
 * File Created: Saturday, 20th November 2021 2:52:24 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Saturday, 20th November 2021 2:52:25 pm
 
 */

package g_api

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// StoreInSession stores a specified key/value pair in the session.
func StoreInSession(ctx *fiber.Ctx, key, value string) error {
	session, err := instanceGAPI.Store.Get(ctx)
	if err != nil {
		log.Printf("StoreInSession - get session for path: %s, error: %+v", ctx.BaseURL(), err)
		return err
	}

	log.Printf("StoreInSession - session.ID(): %s for path: %s", session.ID(), ctx.BaseURL())

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(value)); err != nil {
		return err
	}
	if err := gz.Flush(); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}

	session.Set(key, b.String())
	return session.Save()

}

// GetFromSession retrieves a previously-stored value from the session.
// If no value has previously been stored at the specified key, it will return an error.
func GetFromSession(ctx *fiber.Ctx, key string) (string, error) {
	session, err := instanceGAPI.Store.Get(ctx)
	if err != nil {
		log.Printf("GetFromSession - Store get error: %v", err)
		return "", err
	}

	log.Printf("GetFromSession - session.ID(): %s for path: %s", session.ID(), ctx.BaseURL())

	value := session.Get(key)
	if value == nil {
		return "", fmt.Errorf("could not find a matching session for this request")
	}

	rData := strings.NewReader(value.(string))
	r, err := gzip.NewReader(rData)
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(s), nil
}
