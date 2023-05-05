/* !!
 * File: code.go
 * File Created: Wednesday, 26th May 2021 11:43:19 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 26th May 2021 11:43:20 am
 
 */

package gerr

type Error struct {
	Error error
	Code  uint32
	Line  string
}

func New(code uint32, err error, line string) *Error {
	return &Error{
		Error: err,
		Code:  code,
		Line:  line,
	}
}

/********************************************************************/
/* Client-side Error Code											*/
/********************************************************************/
const (
	ErrorBindData         uint32 = 400000
	ErrorValidData        uint32 = 400001
	ErrTokenInvalid       uint32 = 400016
	ErrorBracketCreated   uint32 = 430001
	ErrorPlayerRegistered uint32 = 430002
	ErrorWrongEmailFormat uint32 = 400012
	ErrorGetDataFromDB    uint32 = 400013
)

/********************************************************************/
/* Server-side Error Code											*/
/********************************************************************/
const (
	ErrorConnect      uint32 = 50000
	ErrorSaveData     uint32 = 50001
	ErrorRetrieveData uint32 = 50002
	ErrorLogin        uint32 = 50003
	ErrorNotFound     uint32 = 50004
	ErrorGetUserInfo  uint32 = 50005
)
