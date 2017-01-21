package benchmarks

import (
	"sync"
)

type sharedBuffer struct {
	sync.Mutex
	rcond  *sync.Cond
	wcond  *sync.Cond
	closed bool
	read   int
	write  int
	count  int
	size   int
	buffer []byte
}

func newSharedBuffer(size int) *sharedBuffer {
	s := &sharedBuffer{
		buffer: make([]byte, size),
		size:   size,
	}
	s.rcond = sync.NewCond(s)
	s.wcond = sync.NewCond(s)
	return s
}

func (s *sharedBuffer) close() {
	s.Lock()
	s.closed = true
	s.rcond.Broadcast()
	s.Unlock()
}

func (s *sharedBuffer) put(val byte) {
	s.Lock()

	// If the buffer is full we need to wait for space to appear
	for s.count == s.size {
		s.wcond.Wait()
	}

	// s.write tells us the next space that's free to write to. If we reach the
	// end of the buffer we wrap around to the start
	s.buffer[s.write] = val
	s.write++
	if s.write == s.size {
		s.write = 0
	}

	// If the buffer was empty, then signal to anyone that's waiting as there's
	// now data to read
	if s.count == 0 {
		s.rcond.Signal()
	}
	s.count++
	s.Unlock()
}

func (s *sharedBuffer) get() (byte, bool) {
	s.Lock()

	// If the buffer is empty then we need to wait for some data
	for s.count == 0 {
		if s.closed {
			s.Unlock()
			return 0, true
		}
		s.rcond.Wait()
	}

	// s.read tells us where the next byte to read is. If we reach the end of
	// the buffer we wrap around to the beginning
	val := s.buffer[s.read]
	s.read++
	if s.read == s.size {
		s.read = 0
	}

	// If the buffer was full, then signal to anyone waiting to write as there is
	// now space
	if s.count == s.size {
		s.wcond.Signal()
	}
	s.count--

	s.Unlock()
	return val, false
}
