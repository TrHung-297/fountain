

package main

import "gitlab.gplay.vn/gtv-backend/fountain/baselib/tracing"

func main() {
	tracing.InitTracing("Test")
	tracing.GetTracer()
}
