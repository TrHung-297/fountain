/* !!
 * File: doc.go
 * File Created: Monday, 12th July 2021 6:04:08 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 12th July 2021 7:24:49 pm
 
 */

/*
package lr is a structured logger for Go, completely API compatible with the standard library logger.


The simplest way to use Lr is simply the package-level exported logger:

  package main

  import (
    log "github.com/TrHung-297/fountain/baselib/g_log/lr"
  )

  func main() {
    log.WithFields(log.Fields{
      "animal": "walrus",
      "number": 1,
      "size":   10,
    }).Info("A walrus appears")
  }

Output:
  time="2015-09-07T08:48:33Z" level=info msg="A walrus appears" animal=walrus number=1 size=10

For a full guide visit https://github.com/TrHung-297/fountain/baselib/g_log/lr
*/
package lr
