package lib

type Specification struct {
	Params []SpecParam
}

func (d *Specification) Add(params ...SpecParam) {
	d.Params = append(d.Params, params...)
}
