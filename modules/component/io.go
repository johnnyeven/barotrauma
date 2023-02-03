package component

// IO 输入输出管理
type IO struct {
	Input  []*Wire
	Output []*Wire
}

func NewIO() IO {
	return IO{
		Input:  make([]*Wire, 0),
		Output: make([]*Wire, 0),
	}
}

func (io *IO) AttachInput(w *Wire) {
	io.Input = append(io.Input, w)
}

func (io *IO) AttachOutput(w *Wire) {
	io.Output = append(io.Output, w)
}
