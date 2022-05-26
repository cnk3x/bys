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
	"亚优萃",
	"亚图芙",
	"丝亚兰",
	"香草亚",
	"酵亚霜",
	"亚洁芙",
	"亚花诗",
	"雅爱亚",
	"丽亚图",
	"香丽亚",
	"亚尔香",
	"霜诗亚",
	"亚爱兰",
	"亚清美",
	"颜伊亚",
	"兰颜亚",
	"亚爱洁",
	"丝亚雅",
	"渍医亚",
	"丽亚珀",
	"香乐亚",
	"雅丽亚",
	"颜薇亚",
	"亚蜜米",
	"亚朵娜",
	"优亚多",
	"亚草美",
	"亚光子",
	"香亚姿",
	"亚朵诗",
	"菲亚图",
	"优亚云",
	"香思亚",
	"美亚元",
	"诗图亚",
	"香乐亚",
	"亚兰美",
	"亚丽医",
	"亚家佳",
	"云亚斯",
	"莱亚森",
	"亚家森",
	"亚生方",
	"记莱亚",
	"亚百蜜",
	"亚欧兰",
	"亚乐蜜",
	"光亚云",
	"雅菲亚",
	"亚士康",
	"亚安森",
	"蜜乐亚",
	"亚康美",
	"舒亚多",
	"贝亚子",
	"兰莱亚",
	"迪森亚",
	"亚医安",
	"亚欧泉",
	"兰妮亚",
	"亚百灵",
	"亚乐云",
	"西雪亚",
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
