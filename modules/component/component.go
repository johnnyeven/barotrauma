package component

import "github.com/johnnyeven/barotrauma/modules/common/id_generator"

// Component 基类
type Component struct {
	id   string
	name string
}

func NewComponent(name string) Component {
	return Component{
		id:   id_generator.GetGenerator().GenerateUniqueID(),
		name: name,
	}
}

func (c *Component) Id() string {
	return c.id
}

func (c *Component) Name() string {
	return c.name
}
