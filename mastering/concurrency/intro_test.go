package concurrency

import (
	"testing"
)

func TestGoRoutinesNotExecutedInOrder(t *testing.T){
	RoutinesAreUnordered()
}

func TestWaitForGoRoutines(t *testing.T){
	WaitForGoRoutinesToFinishExecution()
}

func TestNoWaitForGoRoutines(t *testing.T){
	DoNotWaitForGoRoutinesToFinishExecution()
}