package component

// Wire 电线
type Wire struct {
	// 电力负载
	load int64
	// 信号量
	signal int64
	// 信号值
	data string
}

func (w *Wire) SetLoadOutput(l int64) {
	w.load = l
}

func (w *Wire) GetLoadInput() int64 {
	return w.load
}

func (w *Wire) SetSignalInput(s int64) {
	w.signal = s
}

func (w *Wire) GetSignalOutput() int64 {
	return w.signal
}

func (w *Wire) SetDataInput(d string) {
	w.data = d
}

func (w *Wire) GetDataOutput() string {
	return w.data
}
