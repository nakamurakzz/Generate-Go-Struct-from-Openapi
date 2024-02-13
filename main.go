package main

import (
	"fmt"

	"github.com/nakamurakzz/generate-go-sruct/presentation/response/group"
	"github.com/nakamurakzz/generate-go-sruct/presentation/response/person"
)

func main() {
	person := person.Person{
		Id:   1,
		Name: "person name",
	}

	group := group.Group{
		GroupName: "group name",
	}

	fmt.Printf("Id: %d, Name: %s\n", person.Id, person.Name)
	fmt.Printf("GroupName: %s\n", group.GroupName)
}
