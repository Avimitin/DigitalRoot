package main

import (
	"reflect"
	"testing"
)

func TestCountDigit(t *testing.T) {
	got := countNum(114514)
	want := 6
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDigitRoot(t *testing.T) {
	t.Run("test 16", func(t *testing.T) {
		got := DigitalRoot(16)
		want := 7

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("942", func(t *testing.T) {
		got := DigitalRoot(942)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("132189", func(t *testing.T) {
		got := DigitalRoot(132189)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestInsertDigit(t *testing.T) {
	t.Run("get 16", func(t *testing.T) {
		got := insertDigit(16, 2)
		want := []int{1, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("get 114514", func(t *testing.T) {
		got := insertDigit(114514, 6)
		want := []int{1, 1, 4, 5, 1, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestGetBit(t *testing.T) {
	t.Run("get 16", func(t *testing.T) {
		got := getBit(16, 2)
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("get 114514", func(t *testing.T) {
		got := getBit(114514, 6)
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("get bit", func(t *testing.T) {
		got := getBit(6, 1)
		want := 6
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestPow(t *testing.T) {
	got := pow(6)
	want := 1000000
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetPart(t *testing.T) {
	t.Run("get part 114514", func(t *testing.T) {
		got := getPart(114514, 6)
		want := 14514
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("get part 16", func(t *testing.T) {
		got := getPart(16, 2)
		want := 6
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
