package fizzbuzz

import "testing"

func Test1ShouldTranslateTo1(t *testing.T) {
	result := Translate(1)
	expected := "1"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test2ShouldTranslateTo2(t *testing.T) {
	result := Translate(2)
	expected := "2"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test3ShouldTranslateToFizz(t *testing.T) {
	result := Translate(3)
	expected := "Fizz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test4ShouldTranslateTo4(t *testing.T) {
	result := Translate(4)
	expected := "4"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test5ShouldTranslateToBuzz(t *testing.T) {
	result := Translate(5)
	expected := "Buzz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test6ShouldTranslateToFizz(t *testing.T) {
	result := Translate(6)
	expected := "Fizz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test9ShouldTranslateToFizz(t *testing.T) {
	result := Translate(9)
	expected := "Fizz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test10ShouldTranslateToBuzz(t *testing.T) {
	result := Translate(10)
	expected := "Buzz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test20ShouldTranslateToBuzz(t *testing.T) {
	result := Translate(20)
	expected := "Buzz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test15ShouldTranslateToFizzBuzz(t *testing.T) {
	result := Translate(15)
	expected := "FizzBuzz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test30ShouldTranslateToFizzBuzz(t *testing.T) {
	result := Translate(30)
	expected := "FizzBuzz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGeneratingFizzBuzzUptoLimit(t *testing.T) {
	result := Generate(30)
	expected := "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
