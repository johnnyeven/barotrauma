package component

// Battery 电池
type Battery struct {
	// 最大蓄电量
	capacity int64
	// 当前电量
	quantity int64
	// 最大每秒充电量
	maxRecharge int64
	// 当前每秒充电量
	currentRecharge int64
	// 最大每秒输出量
	maxOutput int64
	// 当前每秒输出量
	currentOutput int64
	// 充电率
	rechargeRate float64

	Wiring
}

func NewBattery(durability int64, storage int64) *Battery {
	instance := &Battery{
		capacity: storage,
		Wiring:   NewWiring("电池", 0, durability),
	}
	instance.
		AddHandler(instance.handleInputRechargeRate).
		AddHandler(instance.handleOutputCapacity).
		AddHandler(instance.handleOutputQuantity).
		AddHandler(instance.handleOutputChargeRate).
		AddHandler(instance.handleOutputCurrentOutput).
		AddHandler(instance.handleOutputCurrentRecharge)
	return instance
}

func (b *Battery) Capacity() int64 {
	return b.capacity
}

func (b *Battery) SetCapacity(capacity int64) {
	b.capacity = capacity
}

func (b *Battery) Quantity() int64 {
	return b.quantity
}

func (b *Battery) SetQuantity(quantity int64) {
	b.quantity = quantity
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
	if !b.IsWorking {
		return
	}

	// TODO 查询终端用电器

	// TODO 根据用电器每秒耗电量计算当前时间帧耗电量

	// TODO 依据剩余电量设置用电器工作状态（启用、停用）

	b.Wiring.Update(timestamp)
}

// 输入充电率
func (b *Battery) handleInputRechargeRate() error {
	return nil
}

// 输出当前电量
func (b *Battery) handleOutputQuantity() error {
	return nil
}

// 输出最大电量
func (b *Battery) handleOutputCapacity() error {
	return nil
}

// 输出每秒充电量
func (b *Battery) handleOutputCurrentRecharge() error {
	return nil
}

// 输出每秒输出量
func (b *Battery) handleOutputCurrentOutput() error {
	return nil
}

// 输出当前充电率
func (b *Battery) handleOutputChargeRate() error {
	return nil
}
