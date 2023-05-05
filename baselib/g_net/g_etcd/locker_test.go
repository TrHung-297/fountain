package g_etcd

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var gInstance *GEtcd

func InitEtcdConnection(t *testing.T) {
	gInstance = NewEtcdDiscovery(WithEtcdAddrs("http://localhost:2379"))

	if gInstance == nil {
		t.Error("Can not create etcd instance")
	}
}

func TestSingleLocker(t *testing.T) {
	InitEtcdConnection(t)

	locker := gInstance.NewLocker("key1")
	err := locker.Lock()
	if err != nil {
		log.Printf("TestSingleLocker - Error: %v", err)
	}
	assert.Nil(t, err, err)

	time.Sleep(5 * time.Second)
	err = locker.Unlock()
	assert.Nil(t, err, err)
}

func TestMultiLocker(t *testing.T) {
	InitEtcdConnection(t)

	locker := gInstance.NewLocker("key1")

	idx := []string{"init locker"}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		idx = append(idx, "1 prepare lock")
		defer wg.Done()

		err := locker.Lock()
		if err != nil {
			t.Logf("TestSingleLocker - Error: %v", err)
		}
		assert.Nil(t, err, err)
		idx = append(idx, fmt.Sprintf("1 err: %+v", err))
		idx = append(idx, "1 lock done")
		time.Sleep(10 * time.Second)

		err = locker.Unlock()
		idx = append(idx, fmt.Sprintf("1 err: %+v", err))
		idx = append(idx, "1 unlock done")
		assert.Nil(t, err, err)
	}()

	wg.Add(1)
	go func() {
		time.Sleep(500 * time.Millisecond)
		idx = append(idx, "2 prepare lock")
		defer wg.Done()

		err := locker.Lock()
		if err != nil {
			t.Logf("TestSingleLocker - Error: %v", err)
		}
		idx = append(idx, fmt.Sprintf("2 err: %+v", err))
		assert.Nil(t, err, err)

		idx = append(idx, "2 lock done")

		err = locker.Unlock()
		assert.Nil(t, err, err)
		idx = append(idx, fmt.Sprintf("2 err: %+v", err))
		idx = append(idx, "2 unlock done")
	}()

	wg.Wait()

	t.Logf("Idx: %+v", idx)

}
