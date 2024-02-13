package response

import (
	"github.com/nakamurakzz/generate-go-sruct/presentation/response/group"
	"github.com/nakamurakzz/generate-go-sruct/presentation/response/person"
)

// NewGroup creates a new group
func NewGroup(name string) group.Group {
	return group.Group{
		GroupName: name,
	}
}

// NewPerson creates a new person
func NewPerson(id int, name string) person.Person {
	return person.Person{
		Id:   id,
		Name: name,
	}
}
