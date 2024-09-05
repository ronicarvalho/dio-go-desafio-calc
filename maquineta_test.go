package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	spected := float32(6)
	acc := maquineta(Operation{action: ADD, acc: 1}, 2, 3)
	if acc != spected {
		t.Error("Spected value:", spected, "Returned value", acc)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	spected := float32(5)
	acc := maquineta(Operation{action: ADD, acc: 1}, 2, 3)
	if acc == spected {
		t.Error("Spected value:", spected, "Returned value", acc)
	}
}

func TestShouldSubCorrect(t *testing.T) {
	spected := float32(6)
	acc := maquineta(Operation{action: SUB, acc: 12}, 3, 3)
	if acc != spected {
		t.Error("Spected value:", spected, "Returned value", acc)
	}
}

func TestShouldSubIncorrect(t *testing.T) {
	spected := float32(8)
	acc := maquineta(Operation{action: SUB, acc: 12}, 3, 3)
	if acc == spected {
		t.Error("Spected value:", spected, "Returned value", acc)
	}
}

func TestShouldDivCorrect(t *testing.T) {
	spected := float32(4)
	acc := maquineta(Operation{action: DIV, acc: 12}, 3)
	if acc != spected {
		t.Error("Spected value:", spected, "Returned value", acc)
	}
}

func TestShouldDivIncorrect(t *testing.T) {
	spected := float32(8)
	acc := maquineta(Operation{action: DIV, acc: 12}, 3)
	if acc == spected {
		t.Error("Spected value:", spected, "Returned value", acc)
	}
}

func TestShouldMultCorrect(t *testing.T) {
	spected := float32(10)
	acc := maquineta(Operation{action: MUL, acc: 2}, 5)
	if acc != spected {
		t.Error("Spected value:", spected, "Returned value", acc)
	}
}

func TestShouldMultIncorrect(t *testing.T) {
	spected := float32(6)
	acc := maquineta(Operation{action: MUL, acc: 2}, 5)
	if acc == spected {
		t.Error("Spected value:", spected, "Returned value", acc)
	}
}
