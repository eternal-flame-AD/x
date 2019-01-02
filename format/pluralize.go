package format

import (
	"strconv"

	"github.com/jinzhu/inflection"
)

// AutoPluralizer implements fmt.Stringer interface and pluralizes words automatically
type AutoPluralizer struct {
	Quantity  int
	Qualifier string
}

func (c AutoPluralizer) String() string {
	return AutoPluralize(c.Quantity, c.Qualifier)
}

// AutoPluralize pluralizes words automatically
func AutoPluralize(quantity int, qualifier string) string {
	if quantity == 1 {
		return strconv.Itoa(quantity) + " " + inflection.Singular(qualifier)
	}
	return strconv.Itoa(quantity) + " " + inflection.Plural(qualifier)
}
