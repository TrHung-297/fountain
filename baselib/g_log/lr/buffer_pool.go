/* !!
 * File: buffer_pool.go
 * File Created: Monday, 12th July 2021 6:04:09 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 12th July 2021 7:24:47 pm
 
 */

package lr

import (
	"bytes"
	"sync"
)

var (
	bufferPool BufferPool
)

type BufferPool interface {
	Put(*bytes.Buffer)
	Get() *bytes.Buffer
}

type defaultPool struct {
	pool *sync.Pool
}

func (p *defaultPool) Put(buf *bytes.Buffer) {
	p.pool.Put(buf)
}

func (p *defaultPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

func getBuffer() *bytes.Buffer {
	return bufferPool.Get()
}

func putBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

// SetBufferPool allows to replace the default lr buffer pool
// to better meets the specific needs of an application.
func SetBufferPool(bp BufferPool) {
	bufferPool = bp
}

func init() {
	SetBufferPool(&defaultPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	})
}
