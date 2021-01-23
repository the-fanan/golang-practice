package concurrency

import (
	"testing"
)

func TestGoRoutinesNotExecutedInOrder(t *testing.T){
	RoutinesAreUnordered()
}