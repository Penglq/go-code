package stringCombineTest

import "testing"

func TestStringsJoin(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringsJoin(tt.args.strs); got != tt.want {
				t.Errorf("StringsJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringsString(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringsString(tt.args.strs); got != tt.want {
				t.Errorf("StringsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringBuffer(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringBuffer(tt.args.strs); got != tt.want {
				t.Errorf("StringBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringByte(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringByte(tt.args.strs); got != tt.want {
				t.Errorf("StringByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
