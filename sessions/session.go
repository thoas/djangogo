package sessions

import (
	"github.com/thoas/djangogo/serializers"
	"github.com/thoas/djangogo/sessions/store"
	"strings"
)

type Session struct {
	Key        string
	data       map[string]interface{}
	Serializer serializers.Serializer
	Store      store.Store
}

func NewSession(key string, serializer serializers.Serializer, store store.Store) *Session {
	return &Session{
		Key:        key,
		Serializer: serializer,
		Store:      store,
	}
}

func (s *Session) Get(key string) (interface{}, error) {
	if s.data == nil {
		err := s.Load()

		if err != nil {
			return "", err
		}
	}

	return s.data[key], nil
}

func (s *Session) Load() error {
	data, err := s.Store.Get(s.Key)

	if err != nil {
		return err
	}

	session, err := s.Serializer.Loads(strings.NewReader(data))

	if err != nil {
		return err
	}

	s.data = session

	return nil
}
