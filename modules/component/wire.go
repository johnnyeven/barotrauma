package component

// Wire 电线
type Wire struct {
	signal int64
	data   string
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
