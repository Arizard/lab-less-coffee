package spc

import (
	"time"
)

type OverlapType int

const (
	OverlapDefault OverlapType = iota
	OverlapNone
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
		} else if b1.End.After(b2.Start) {
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

func MultiSplit(bBig Block, bSmls []Block) []Block {
	results := []Block{}
	sml := bSmls[0]

	for _, splitBlock := range Split(bBig, sml) {
		if len(bSmls) > 1 {
			for _, result := range MultiSplit(splitBlock, bSmls[1:]) {
				results = append(results, result)
			}
		} else {
			results = append(results, splitBlock)
		}
	}

	return results

}

func Split(bBig, bSml Block) []Block {
	if bSml.Start.After(bBig.End) || bSml.End.Before(bBig.Start) {
		return []Block{bBig}
	}
	return []Block{
		Block{Start: bBig.Start, End: bSml.Start},
		Block{Start: bSml.End, End: bBig.End},
	}
}

type BlockTimeline struct {
	blocks          []Block
	blockOrder      []int
	blockOrderIndex int
}

func (bs *BlockTimeline) Restart() {
	bs.blockOrderIndex = 0
}

func (bs *BlockTimeline) GetCurrentBlock() Block {
	return Block{}
}

func (bs *BlockTimeline) Next() {
	if bs.blockOrderIndex += 1; bs.blockOrderIndex > len(bs.blockOrder)-1 {
		bs.blockOrderIndex = len(bs.blockOrder) - 1
	}
}
func (bs *BlockTimeline) Previous() {
	if bs.blockOrderIndex -= 1; bs.blockOrderIndex < 0 {
		bs.blockOrderIndex = 0
	}
}

func NewBlockSet(blocks []Block) *BlockTimeline {
	return &BlockTimeline{
		blocks:          blocks,
		blockOrder:      make([]int, len(blocks)),
		blockOrderIndex: 0,
	}
}
