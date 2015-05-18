package serializers

import (
	"strings"
)

type Serializer interface {
	Loads(reader *strings.Reader) (map[string]interface{}, error)
}
