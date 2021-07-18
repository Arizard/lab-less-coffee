package spc

import "time"

type PaySpec string

// Translate de-serialises a pay spec and creates a pay rule func
func Translate(spec PaySpec) PayRuleFunc {
	return DurationAfter(3 * time.Hour)
}
