package component

import "github.com/sirupsen/logrus"

type IOHandler func() error

// IO 输入输出管理
type IO struct {
	Input    []*Wire
	Output   []*Wire
	handlers []IOHandler
}

func NewIO() IO {
	return IO{
		Input:  make([]*Wire, 0),
		Output: make([]*Wire, 0),
	}
}

func (io *IO) AddHandler(h func() error) *IO {
	io.handlers = append(io.handlers, h)
	return io
}

func (io *IO) AttachInput(w *Wire) *IO {
	io.Input = append(io.Input, w)
	return io
}

func (io *IO) AttachOutput(w *Wire) *IO {
	io.Output = append(io.Output, w)
	return io
}

func (io *IO) Update(timestamp int64) {
	for _, handler := range io.handlers {
		if err := handler(); err != nil {
			logrus.Errorf("%v", err)
		}
	}
}
