package spc

import (
	"reflect"
	"testing"
	"time"
)

func TestFirstDuration(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		newFunc := FirstDuration(time.Hour * 3)
		result := newFunc(
			Block{newTime(9, 0, 0), newTime(17, 0, 0), OverlapNone},
			newTime(9, 0, 0),
			newTime(17, 0, 0),
		)
		expected := Block{newTime(9, 0, 0), newTime(12, 0, 0), OverlapOutside}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Got %v, want %v", result, expected)
		}
	})

	t.Run("With focus boundaries (inside)", func(t *testing.T) {
		newFunc := FirstDuration(time.Hour * 3)
		result := newFunc(
			Block{newTime(9, 0, 0), newTime(17, 0, 0), OverlapNone},
			newTime(10, 0, 0),
			newTime(11, 0, 0),
		)
		expected := Block{newTime(10, 0, 0), newTime(11, 0, 0), OverlapInside}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Got %v, want %v", result, expected)
		}
	})
}

func TestDurationAfter(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		newFunc := DurationAfter(time.Hour * 3)
		result := newFunc(
			Block{newTime(9, 0, 0), newTime(17, 0, 0), OverlapNone},
			newTime(9, 0, 0),
			newTime(17, 0, 0),
		)
		expected := Block{newTime(12, 0, 0), newTime(17, 0, 0), OverlapOutside}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Got %v, want %v", result, expected)
		}
	})
}

func TestDurationFromTo(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		newFunc := DurationFromTo(time.Hour*3, time.Hour*4)
		result := newFunc(
			Block{newTime(9, 0, 0), newTime(17, 0, 0), OverlapNone},
			newTime(9, 0, 0),
			newTime(17, 0, 0),
		)
		expected := Block{newTime(12, 0, 0), newTime(13, 0, 0), OverlapOutside}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Got %v, want %v", result, expected)
		}
	})

	t.Run("With focus", func(t *testing.T) {
		newFunc := DurationFromTo(time.Hour*3, time.Hour*6)
		result := newFunc(
			Block{newTime(9, 0, 0), newTime(17, 0, 0), OverlapNone},
			newTime(13, 0, 0),
			newTime(17, 0, 0),
		)
		expected := Block{newTime(13, 0, 0), newTime(15, 0, 0), OverlapStart}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Got %v, want %v", result, expected)
		}
	})
}
