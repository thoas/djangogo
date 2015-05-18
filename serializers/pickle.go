package serializers

import (
	"github.com/hydrogen18/stalecucumber"
	"strings"
)

type PickleSerializer struct{}

func (p *PickleSerializer) Loads(reader *strings.Reader) (map[string]interface{}, error) {
	return stalecucumber.DictString(stalecucumber.Unpickle(reader))
}
