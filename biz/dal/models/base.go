/* !!
 * File: base.go
 * File Created: Tuesday, 25th May 2021 4:10:04 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Tuesday, 25th May 2021 4:10:04 pm
 
 */

package models

type Model struct {
	CreatedTime int32  `db:"created_time" json:"created_time,omitempty"` // need set by application for localtion time in seconds
	UpdatedTime int32  `db:"updated_time" json:"updated_time,omitempty"` // need set by application for localtion time in seconds
	CreatedAt   string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   string `db:"updated_at" json:"updated_at,omitempty"`
}
