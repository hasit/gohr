package gohr

import (
	"bytes"
	"testing"
)

func Test_drawPatterns(t *testing.T) {
	type args struct {
		w        int
		patterns []string
	}
	tests := []struct {
		name   string
		args   args
		wantWr string
	}{
		{
			"Default pattern, even width",
			args{
				w:        4,
				patterns: []string{},
			},
			"####\n",
		},
		{
			"Default pattern, odd width",
			args{
				w:        3,
				patterns: []string{},
			},
			"###\n",
		},
		{
			"Custom pattern (single character), even width",
			args{
				w:        4,
				patterns: []string{"-"},
			},
			"----\n",
		},
		{
			"Custom pattern (single character), odd width",
			args{
				w:        5,
				patterns: []string{"-"},
			},
			"-----\n",
		},
		{
			"Custom pattern (multiple characters), even width",
			args{
				w:        4,
				patterns: []string{"-+"},
			},
			"-+-+\n",
		},
		{
			"Custom pattern (multiple characters), odd width",
			args{
				w:        5,
				patterns: []string{"-+"},
			},
			"-+-+-\n",
		},
		{
			"Custom pattern 2 (multiple characters), odd width",
			args{
				w:        7,
				patterns: []string{"-+="},
			},
			"-+=-+=-\n",
		},
		{
			"Custom patterns, even width",
			args{
				w:        6,
				patterns: []string{"-+", "0"},
			},
			"-+-+-+\n000000\n",
		},
		{
			"Custom patterns, odd width",
			args{
				w:        9,
				patterns: []string{"//", "\\"},
			},
			"/////////\n\\\\\\\\\\\\\\\\\\\n",
		},
	}
	for _, tt := range tests {
		wr := &bytes.Buffer{}
		drawPatterns(tt.args.w, tt.args.patterns, wr)
		if gotWr := wr.String(); gotWr != tt.wantWr {
			t.Errorf("%q. drawPatterns() = %v, want %v", tt.name, gotWr, tt.wantWr)
		}
	}
}
