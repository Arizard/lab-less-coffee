package spc

import "time"

type PayRuleFunc func(shift Block, focusStart time.Time, focusEnd time.Time) (result Block)

func Stack(rules ...PayRuleFunc) PayRuleFunc {
	return func(shift Block, focusStart time.Time, focusEnd time.Time) (result Block) {
		result = Block{shift.Start, shift.End, OverlapNone}
		for _, rule := range rules {
			result = rule(result, shift.Start, shift.End)
		}
		return Overlap(result, Block{Start: focusStart, End: focusEnd})
	}
}

func FirstDuration(x time.Duration) PayRuleFunc {
	return func(shift Block, focusStart time.Time, focusEnd time.Time) (result Block) {
		window := Block{shift.Start, shift.Start.Add(x), OverlapNone}
		return Overlap(
			Overlap(shift, window),
			Block{Start: focusStart, End: focusEnd},
		)
	}
}

func DurationAfter(x time.Duration) PayRuleFunc {
	return func(shift Block, focusStart time.Time, focusEnd time.Time) (result Block) {
		window := Block{shift.Start.Add(x), shift.End, OverlapNone}
		return Overlap(
			Overlap(shift, window),
			Block{Start: focusStart, End: focusEnd},
		)
	}
}

func DurationFromTo(from time.Duration, to time.Duration) PayRuleFunc {
	return Stack(
		DurationAfter(from),
		FirstDuration(to-from),
	)
}
