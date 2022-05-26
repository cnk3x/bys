package bys

import (
	"strconv"
	"testing"
)

var names = []string{
	"五加一",
	"燕加元",
	"舒护侠",
	"半亩岛",
	"健优明",
	"谷味康",
	"究镜所",
	"VSZP",
	"vsA!123",
	"克罗诗丁",
	"汉方净",
}

func BenchmarkAssessName(b *testing.B) {
	b.ReportAllocs()
	l := len(names)
	for i := 0; i < b.N; i++ {
		AssessName(names[i%l])
	}
}

func BenchmarkAssessPhone(b *testing.B) {
	b.ReportAllocs()
	var nums []string
	for i := 0; i < 1000; i++ {
		nums = append(nums, strconv.Itoa(13012340000+i%10000))
	}
	l := len(nums)
	for i := 0; i < b.N; i++ {
		AssessPhone(nums[i%l])
	}
}

func TestByName(t *testing.T) {
	for _, n := range names {
		r := AssessName(n)
		t.Logf("%-*s %s", countText(n, 15), n, r)
	}
}

func TestByPhone(t *testing.T) {
	n := "13012345678"
	r := AssessPhone(n)
	t.Logf("%-*s %s", 13, n, r)
}
