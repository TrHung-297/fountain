package kstats

import (
	"context"

	newRelic "github.com/newrelic/go-agent"
)

type ctxKeyClientSegment struct{}

func setSegment(ctx context.Context, seg *newRelic.Segment) context.Context {
	return context.WithValue(ctx, ctxKeyClientSegment{}, seg)
}

func getSegment(ctx context.Context) (st *newRelic.Segment, ok bool) {
	if v := ctx.Value(ctxKeyClientSegment{}); v != nil {
		st, ok = v.(*newRelic.Segment)
	}
	return
}
