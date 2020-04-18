package simple_factory

import (
	"github.com/assertgo/assert"
	"testing"
)

func TestNewPerson(t *testing.T) {
	person := NewPerson(1)
	assert.New(t).That(person).IsNotNil()
	person.Say()
}
