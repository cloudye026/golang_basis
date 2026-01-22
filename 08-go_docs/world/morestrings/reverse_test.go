package morestrings

import "testing"

func TestIsstring(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{"hello", "是字符串"},
		{"世界", "是字符串"},
		{"", "是字符串"},
		{123, "不是字符串"},
		{45.67, "不是字符串"},
		{true, "不是字符串"},
		{[]int{1, 2, 3}, "不是字符串"},
		{nil, "不是字符串"},
	}

	for _, c := range cases {
		got := IsString(c.in)
		if got != c.want {
			t.Errorf("Isstring(%v) == %q, want %q", c.in, got, c.want)
		}
	}
}
