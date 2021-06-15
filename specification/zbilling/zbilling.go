package zbilling

import (
	"github.com/go-specifier/specifier/lib"
	"github.com/go-specifier/specifier/specification/zbilling/elastic"
	"github.com/go-specifier/specifier/specification/zbilling/mariaDb"
)

type Billing struct {
	lib.SpecParam

	Elastic *elastic.Spec
	MariaDb *mariaDb.Spec
}

func (b *Billing) Setup() {

}
