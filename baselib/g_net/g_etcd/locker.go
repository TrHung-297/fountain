/* !!
 * File: locker.go
 * File Created: Monday, 14th June 2021 3:06:47 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 14th June 2021 3:06:47 pm
 
 */

package g_etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/TrHung-297/fountain/baselib/g_log"
	"go.etcd.io/etcd/client/v3/concurrency"
)

const (
	KDefaultLockTimeout    = 10 * time.Second
	KLockerNamespacePrefix = "lock"
)

// GLocker type
type GLocker struct {
	*GEtcd
	*concurrency.Mutex
	keyLocker         string
	LockerActionError error
	CancelFunc        context.CancelFunc
}

func (g *GEtcd) NewLocker(key string) *GLocker {
	ctx, cancel := context.WithTimeout(context.Background(), KDefaultLockTimeout)
	defer cancel()

	session, err := concurrency.NewSession(g.Client, concurrency.WithTTL(30), concurrency.WithContext(ctx))
	if err != nil {
		g_log.V(1).Errorf("GEtcd::NewLocker - Error: %+v", err)
		return nil
	}

	keyLocker := fmt.Sprintf("%s/%s/%s", KNameProjectDir, KLockerNamespacePrefix, key)
	lck := concurrency.NewMutex(session, keyLocker)

	return &GLocker{
		GEtcd:     g,
		Mutex:     lck,
		keyLocker: keyLocker,
	}
}

func (l *GLocker) Lock() *GLocker {
	ctx, cancel := context.WithTimeout(context.Background(), KDefaultLockTimeout)
	l.CancelFunc = cancel

	l.LockerActionError = l.Mutex.Lock(ctx)
	return l
}

func (l *GLocker) Unlock() *GLocker {
	ctx, cancel := context.WithTimeout(context.Background(), KDefaultLockTimeout)
	l.CancelFunc = cancel

	l.LockerActionError = l.Mutex.Unlock(ctx)
	return l
}
