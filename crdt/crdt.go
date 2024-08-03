package crdt

import "fmt"

type CRDT interface {
	Insert(position int, value string) (string error)
	Delete(position int) string
}

func IsCrdt(c CRDT) {
	//temporary code to check if crdt works.
	fmt.Println(c.Insert(1, "a"))
}
