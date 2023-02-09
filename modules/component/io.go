package component

import "github.com/sirupsen/logrus"

type IOHandler func() error

// IO 输入输出管理
type IO struct {
	Input    map[int]*Wire
	Output   map[int][]*Wire
	handlers []IOHandler
}

func NewIO() IO {
	return IO{
		Input:  make(map[int]*Wire),
		Output: make(map[int][]*Wire),
	}
}

func (io *IO) AddHandler(h func() error) *IO {
	io.handlers = append(io.handlers, h)
	return io
}

func (io *IO) AttachInput(slot int, w *Wire) *IO {
	io.Input[slot] = w
	return io
}

func (io *IO) AttachOutput(slot int, w *Wire) *IO {
	if _, ok := io.Output[slot]; !ok {
		io.Output[slot] = make([]*Wire, 0)
	}
	io.Output[slot] = append(io.Output[slot], w)
	return io
}

func (io *IO) Update(timestamp int64) {
	for _, handler := range io.handlers {
		if err := handler(); err != nil {
			logrus.Errorf("%v", err)
		}
	}
}
