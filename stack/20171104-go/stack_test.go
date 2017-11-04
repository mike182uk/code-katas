package stack

import "testing"

func TestStackInitialisedEmpty(t *testing.T) {
	s := NewStack(1)

	expected := 0
	result := s.Size()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPushElementOnToStack(t *testing.T) {
	s := NewStack(1)

	_ = s.Push(1)

	expected := 1
	result := s.Size()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPopElementOffOfStack(t *testing.T) {
	s := NewStack(1)

	element := 1

	_ = s.Push(element)

	popResult, _ := s.Pop()

	if popResult != element {
		t.Errorf("Expected %d, got %d", popResult, element)
	}

	size := s.Size()
	expectedSize := 0

	if size != expectedSize {
		t.Errorf("Expected %d, got %d", size, expectedSize)
	}
}

func TestLIFO(t *testing.T) {
	s := NewStack(2)

	element1 := 1
	element2 := 2

	_ = s.Push(element1)
	_ = s.Push(element2)

	firstPopResult, _ := s.Pop()
	secondPopResult, _ := s.Pop()

	if firstPopResult != element2 {
		t.Errorf("Expected %d, got %d", element1, firstPopResult)
	}

	if secondPopResult != element1 {
		t.Errorf("Expected %d, got %d", element2, secondPopResult)
	}
}

func TestErrorReturnedIfPopWhenEmpty(t *testing.T) {
	s := NewStack(1)

	expectedResult := -1
	expectedErrMsg := "Stack empty"

	result, err := s.Pop()

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	errMsg := err.Error()

	if errMsg != expectedErrMsg {
		t.Errorf(`Expected error message to be "%s", got "%s"`, expectedErrMsg, errMsg)
	}
}

func TestErrorReturnedIfPushWhenFull(t *testing.T) {
	s := NewStack(2)

	element1 := 1
	element2 := 2
	element3 := 3

	_ = s.Push(element1)
	_ = s.Push(element2)
	err := s.Push(element3)

	expectedErrMsg := "Stack full"

	if err == nil {
		t.Error("Expected error, got nil")
	}

	errMsg := err.Error()

	if errMsg != expectedErrMsg {
		t.Errorf(`Expected error message "%s", got "%s"`, expectedErrMsg, errMsg)
	}
}

func TestPeek(t *testing.T) {
	stackSize := 2
	s := NewStack(stackSize)

	element1 := 1
	element2 := 2

	_ = s.Push(element1)
	_ = s.Push(element2)

	result, _ := s.Peek()

	if result != element2 {
		t.Errorf("Expected %d, got %d", element2, result)
	}

	size := s.Size()

	if size != stackSize {
		t.Errorf("Expected %d, got %d", stackSize, size)
	}
}

func TestErrorReturnedIfPeekWhenEmpty(t *testing.T) {
	s := NewStack(1)

	expectedResult := -1
	expectedErrMsg := "Stack empty"

	result, err := s.Peek()

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	errMsg := err.Error()

	if errMsg != expectedErrMsg {
		t.Errorf(`Expected error message to be "%s", got "%s"`, expectedErrMsg, errMsg)
	}
}
