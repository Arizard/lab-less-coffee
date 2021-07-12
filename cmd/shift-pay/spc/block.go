package spc

import "time"

type OverlapType int

const (
	OverlapNone OverlapType = iota
	OverlapStart
	OverlapEnd
	OverlapInside
	OverlapOutside
	OverlapFull
)

func Overlap(b1 Block, b2 Block) (result Block) {

	if b1.Start.Before(b2.Start) {
		if b1.End.Before(b2.Start) {
			result.OverlapType = OverlapNone
		} else if b1.End.Equal(b2.Start) {
			result.OverlapType = OverlapNone
		}
		if b1.End.After(b2.Start) {
			if b1.End.Before(b2.End) {
				result.Start = b2.Start
				result.End = b1.End
				result.OverlapType = OverlapStart
			} else if b1.End.Equal(b2.End) {
				result.Start = b2.Start
				result.End = b1.End
				result.OverlapType = OverlapInside
			} else if b1.End.After(b2.End) {
				result.Start = b2.Start
				result.End = b2.End
				result.OverlapType = OverlapInside
			}
		}
	} else if b1.Start.Equal(b2.Start) {
		if b1.End.Before(b2.End) {
			result.Start = b2.Start
			result.End = b1.End
			result.OverlapType = OverlapOutside
		} else if b1.End.Equal(b2.End) {
			result.Start = b2.Start
			result.End = b1.End
			result.OverlapType = OverlapFull
		} else if b1.End.After(b2.End) {
			result.Start = b2.Start
			result.End = b2.End
			result.OverlapType = OverlapInside
		}
	} else if b1.Start.Before(b2.End) {
		if b1.End.Before(b2.End) {
			result.Start = b1.Start
			result.End = b1.End
			result.OverlapType = OverlapOutside
		} else if b1.End.Equal(b2.End) {
			result.Start = b1.Start
			result.End = b1.End
			result.OverlapType = OverlapOutside
		} else if b1.End.After(b2.End) {
			result.Start = b1.Start
			result.End = b2.End
			result.OverlapType = OverlapEnd
		}
	} else if b1.Start.Equal(b2.End) {
		result.OverlapType = OverlapNone
	} else if b1.Start.After(b2.End) {
		result.OverlapType = OverlapNone
	}

	return result
}

type Block struct {
	Start       time.Time
	End         time.Time
	OverlapType OverlapType
}
