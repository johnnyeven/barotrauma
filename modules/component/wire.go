package component

import "github.com/johnnyeven/barotrauma/modules/core"

// Wire 电线
type Wire struct {
	input  core.IComponent
	output core.IComponent
	// 电力负载
	load int64
	// 信号量
	signal int64
	// 信号值
	data string
}

func (w *Wire) Input() core.IComponent {
	return w.input
}

func (w *Wire) SetInput(input core.IComponent) {
	w.input = input
}

func (w *Wire) Output() core.IComponent {
	return w.output
}

func (w *Wire) SetOutput(output core.IComponent) {
	w.output = output
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
