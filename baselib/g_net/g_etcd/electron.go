package g_etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/TrHung-297/fountain/baselib/g_log"
	"go.etcd.io/etcd/client/v3/concurrency"
)

const (
	KDefaultElectronTimeout  = 30 * time.Second
	KElectionNamespacePrefix = "election"
)

// GLocker type
type GElection struct {
	*GEtcd
	*concurrency.Election
	ElectionName      string
	LockerActionError error
}

func (g *GEtcd) NewElectron(name ...string) *GElection {
	var (
		session *concurrency.Session
		err     error
	)

	gElection := &GElection{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	lease, err := g.Client.Lease.Grant(ctx, 30)
	if err != nil {
		g_log.V(1).Errorf("GEtcd::NewLocker - Grant Lease Error: %+v", err)
		return nil
	}

	if lease != nil {
		session, err = concurrency.NewSession(g.Client, concurrency.WithLease(lease.ID))
		if err != nil {
			g_log.V(1).Errorf("GEtcd::NewLocker - NewSession Error: %+v", err)
			return nil
		}
	}

	if session == nil {
		g_log.V(1).Errorf("GEtcd::NewLocker - Can not create session")
		return nil
	}

	keyElection := fmt.Sprintf("%s/%s", KNameProjectDir, KElectionNamespacePrefix)
	if len(name) != 0 && name[0] != "" {
		keyElection = fmt.Sprintf("%s/%s", keyElection, name[0])
	}

	el := concurrency.NewElection(session, keyElection)

	gElection.GEtcd = g
	gElection.Election = el
	gElection.ElectionName = keyElection

	return gElection
}

// Raise func;
// Tự tin ứng cử làm leader;
func (e *GElection) Raise(missionName string) error {
	if missionName == "" {
		missionName = KElectionNamespacePrefix
	}

	ctx, cancel := context.WithTimeout(context.Background(), KDefaultLockTimeout)
	defer cancel()

	return e.Election.Campaign(ctx, missionName)
}

func (e *GElection) Resign() error {
	ctx, cancel := context.WithTimeout(context.Background(), KDefaultLockTimeout)
	defer cancel()

	err := e.Election.Resign(ctx)
	return err
}
