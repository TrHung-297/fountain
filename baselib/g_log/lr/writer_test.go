/* !!
 * File: writer_test.go
 * File Created: Monday, 12th July 2021 6:04:09 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 12th July 2021 7:27:04 pm
 
 */

package lr

import (
	"log"
	"net/http"
)

func ExampleLogger_Writer_httpServer() {
	logger := New()
	w := logger.Writer()
	defer w.Close()

	srv := http.Server{
		// create a stdlib log.Logger that writes to
		// Logger.
		ErrorLog: log.New(w, "", 0),
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func ExampleLogger_Writer_stdlib() {
	logger := New()
	logger.Formatter = &JSONFormatter{}

	// Use lr for standard log output
	// Note that `log` here references stdlib's log
	// Not lr imported under the name `log`.
	log.SetOutput(logger.Writer())
}
