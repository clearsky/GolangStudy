package nonrepeating

import "testing"

func TestSubStr(t *testing.T){
	tests := []struct{
		s string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinede support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
	}

	for _, tt := range tests{
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans{
			t.Errorf("lengthOfNonRepeatingSubStr(%s);get:%d;expexted:%d;",tt.s, actual, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B){
	s := "黑化肥挥发发灰会花飞灰花肥挥发发黑会飞花"
	ans := 6

	b.ResetTimer()
	for i:=0; i < b.N;i++{
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans{
			b.Errorf("lengthOfNonRepeatingSubStr(%s);get:%d;expexted:%d;",s, actual, ans)
		}
	}
}