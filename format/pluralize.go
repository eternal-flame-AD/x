package format

import (
	"strconv"

	"github.com/jinzhu/inflection"
)

func AutoPluralize(quantity int, qualifier string) string {
	if quantity == 1 {
		return strconv.Itoa(quantity) + " " + inflection.Singular(qualifier)
	}
	return strconv.Itoa(quantity) + " " + inflection.Plural(qualifier)
}
