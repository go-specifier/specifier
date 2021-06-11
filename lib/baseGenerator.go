package lib

type Specification struct {
	Entities []Entity
}

func (d *Specification) Add(entity Entity) {
	d.Entities = append(d.Entities, entity)
}
