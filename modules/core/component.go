package core

type IUpdatable interface {
	Update(timestamp int64)
}

type IComponent interface {
	IUpdatable
	Id() string
	Name() string
}
