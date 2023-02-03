package core

type IComponent interface {
	Id() string
	Name() string
	Update(timestamp int64)
}
