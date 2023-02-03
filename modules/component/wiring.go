package component

// Wiring 电器基类
type Wiring struct {
	// 每秒耗电量
	PowerConsumption int64
	// 耐久度
	Durability int64
	IO
	Component
}

func NewWiring(name string, pc int64, d int64) Wiring {
	return Wiring{
		PowerConsumption: pc,
		Durability:       d,
		IO:               NewIO(),
		Component:        NewComponent(name),
	}
}

func (w *Wiring) Update(timestamp int64) {
}
