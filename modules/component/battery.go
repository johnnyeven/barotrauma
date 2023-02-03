package component

// Battery 电池
type Battery struct {
	// 最大蓄电量
	storage int64
	// 最大每秒充电量
	maxRecharge int64
	// 最大每秒输出量
	maxOutput int64
	// 充电率
	rechargeRate float64

	Wiring
}

func NewBattery(durability int64, storage int64) *Battery {
	return &Battery{
		storage: storage,
		Wiring:  NewWiring("电池", 0, durability),
	}
}

func (b *Battery) Storage() int64 {
	return b.storage
}

func (b *Battery) SetStorage(storage int64) {
	b.storage = storage
}

func (b *Battery) MaxRecharge() int64 {
	return b.maxRecharge
}

func (b *Battery) SetMaxRecharge(maxRecharge int64) {
	b.maxRecharge = maxRecharge
}

func (b *Battery) MaxOutput() int64 {
	return b.maxOutput
}

func (b *Battery) SetMaxOutput(maxOutput int64) {
	b.maxOutput = maxOutput
}

func (b *Battery) RechargeRate() float64 {
	return b.rechargeRate
}

func (b *Battery) SetRechargeRate(rechargeRate float64) {
	if rechargeRate > 1 {
		rechargeRate = 1
	}
	b.rechargeRate = rechargeRate
}

func (b *Battery) Update(timestamp int64) {
	b.Wiring.Update(timestamp)
}
