package spc

import (
	"reflect"
	"testing"
)

func TestOverlap(t *testing.T) {
	type args struct {
		b1 Block
		b2 Block
	}
	tests := []struct {
		name       string
		args       args
		wantResult Block
	}{
		{
			"none (before)",
			args{
				Block{Start: newTime(9, 0, 0), End: newTime(17, 0, 0)},
				Block{Start: newTime(6, 0, 0), End: newTime(7, 0, 0)},
			},
			Block{OverlapType: OverlapNone},
		},
		{
			"none (boundary start)",
			args{
				Block{Start: newTime(9, 0, 0), End: newTime(17, 0, 0)},
				Block{Start: newTime(6, 0, 0), End: newTime(9, 0, 0)},
			},
			Block{OverlapType: OverlapNone},
		},
		{
			"none (boundary end)",
			args{
				Block{Start: newTime(9, 0, 0), End: newTime(17, 0, 0)},
				Block{Start: newTime(17, 0, 0), End: newTime(19, 0, 0)},
			},
			Block{OverlapType: OverlapNone},
		},
		{
			"none (after)",
			args{
				Block{Start: newTime(6, 0, 0), End: newTime(7, 0, 0)},
				Block{Start: newTime(9, 0, 0), End: newTime(17, 0, 0)},
			},
			Block{OverlapType: OverlapNone},
		},
		{
			"start",
			args{
				Block{Start: newTime(9, 0, 0), End: newTime(17, 0, 0)},
				Block{Start: newTime(12, 0, 0), End: newTime(19, 0, 0)},
			},
			Block{newTime(12, 0, 0), newTime(17, 0, 0), OverlapStart},
		},
		{
			"inside",
			args{
				Block{Start: newTime(9, 0, 0), End: newTime(17, 0, 0)},
				Block{Start: newTime(12, 0, 0), End: newTime(15, 0, 0)},
			},
			Block{newTime(12, 0, 0), newTime(15, 0, 0), OverlapInside},
		},
		{
			"outside",
			args{
				Block{Start: newTime(12, 0, 0), End: newTime(15, 0, 0)},
				Block{Start: newTime(9, 0, 0), End: newTime(17, 0, 0)},
			},
			Block{newTime(12, 0, 0), newTime(15, 0, 0), OverlapOutside},
		},
		{
			"full",
			args{
				Block{Start: newTime(12, 0, 0), End: newTime(15, 0, 0)},
				Block{Start: newTime(12, 0, 0), End: newTime(15, 0, 0)},
			},
			Block{newTime(12, 0, 0), newTime(15, 0, 0), OverlapFull},
		},
		{
			"end",
			args{
				Block{Start: newTime(12, 0, 0), End: newTime(19, 0, 0)},
				Block{Start: newTime(9, 0, 0), End: newTime(17, 0, 0)},
			},
			Block{newTime(12, 0, 0), newTime(17, 0, 0), OverlapEnd},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Overlap(tt.args.b1, tt.args.b2); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Overlap() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	bBig := Block{
		Start: newTime(9, 0, 0),
		End:   newTime(17, 0, 0),
	}
	bSml := Block{
		Start: newTime(12, 0, 0),
		End:   newTime(13, 0, 0),
	}
	want := [2]Block{
		{
			Start: bBig.Start,
			End:   bSml.Start,
		},
		{
			Start: bSml.End,
			End:   bBig.End,
		},
	}
	t.Run("Split", func(t *testing.T) {
		if got := Split(bBig, bSml); !reflect.DeepEqual(got, want) {
			t.Errorf("Split() = %v, want %v", got, want)
		}
	})
}

func TestMultiSplit(t *testing.T) {
	bBig := Block{
		Start: newTime(9, 0, 0),
		End:   newTime(17, 0, 0),
	}
	bSmls := []Block{
		{
			Start: newTime(10, 0, 0),
			End:   newTime(10, 30, 0),
		},
		{
			Start: newTime(12, 0, 0),
			End:   newTime(12, 30, 0),
		},
		{
			Start: newTime(15, 0, 0),
			End:   newTime(15, 30, 0),
		},
	}
	want := []Block{
		{
			Start: newTime(9, 0, 0),
			End:   newTime(10, 0, 0),
		},
		{
			Start: newTime(10, 30, 0),
			End:   newTime(12, 0, 0),
		},
		{
			Start: newTime(12, 30, 0),
			End:   newTime(15, 0, 0),
		},
		{
			Start: newTime(15, 30, 0),
			End:   newTime(17, 0, 0),
		},
	}
	t.Run("Split", func(t *testing.T) {
		if got := MultiSplit(bBig, bSmls); !reflect.DeepEqual(got, want) {
			t.Errorf("MultiSplit() = %v, want %v", got, want)
		}
	})
}
