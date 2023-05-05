

package main

import "github.com/TrHung-297/fountain/baselib/tracing"

func main() {
	tracing.InitTracing("Test")
	tracing.GetTracer()
}
