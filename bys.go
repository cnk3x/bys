package bys

import (
	"bytes"
	"embed"
	"encoding/csv"
	"strconv"
	"strings"
	"sync"
)

var (
	//go:embed dict/81.csv
	e81 []byte
	//go:embed dict/kx.csv
	eKX []byte
	//go:embed dict/fj.csv
	eFJ []byte
	_   embed.FS
	o81 sync.Once
	okx sync.Once
	m81 []*R81
	mkx map[rune]*RKx
)

func load81() {
	o81.Do(func() {
		r := csv.NewReader(bytes.NewReader(e81))
		// 数, 数理评分, 描述, 吉凶, 附加, 性格影响
		for l, err := r.Read(); err == nil; l, err = r.Read() {
			var (
				shu, _ = strconv.Atoi(strings.TrimSpace(l[0]))
				fen, _ = strconv.Atoi(strings.TrimSpace(l[1]))
			)
			m81 = append(m81, &R81{
				Shu: shu,
				Fen: fen,
				Jx:  l[2],
				Jxm: l[3],
				Zy:  l[4],
				Xg:  l[5],
				Xgm: l[6],
			})
		}
	})
}

func loadkx() {
	okx.Do(func() {
		mkx = map[rune]*RKx{}
		r := csv.NewReader(bytes.NewReader(eKX))
		// 二,2,火
		for l, err := r.Read(); err == nil; l, err = r.Read() {
			var (
				zi    = ([]rune(l[0]))[0]
				bi, _ = strconv.Atoi(l[1])
			)
			mkx[zi] = &RKx{
				Zi: string(zi),
				Bi: bi,
				Wu: l[2],
			}
		}

		r = csv.NewReader(bytes.NewReader(eFJ))
		// 繁,简
		for l, err := r.Read(); err == nil; l, err = r.Read() {
			var (
				f = ([]rune(l[0]))[0]
				j = ([]rune(l[1]))[0]
			)
			if cj, find := mkx[j]; find {
				mkx[f] = cj
			}
		}
	})
}

// RKx 康熙笔画数和五行属性
type RKx struct {
	Zi string `json:"zi,omitempty"` // 字符
	Bi int    `json:"bi,omitempty"` // 笔画
	Wu string `json:"wu,omitempty"` // 五行属性
}

// R81 八一数理结果，数理评分, 吉凶, 吉凶描述, 官财艺, 性格影响
type R81 struct {
	Shu int    `json:"shu,omitempty"` // 数
	Fen int    `json:"fen,omitempty"` // 数理评分
	Jx  string `json:"jx,omitempty"`  // 吉凶
	Jxm string `json:"jxm,omitempty"` // 吉凶描述
	Zy  string `json:"zy,omitempty"`  // 职业增益，官财艺
	Xg  string `json:"xg,omitempty"`  // 性格类型
	Xgm string `json:"xgm,omitempty"` // 性格描述
}

// Find81 查看数理说明
func Find81(shu int) *R81 {
	load81()
	if shu > 81 {
		shu = shu % 80
	}
	if shu == 0 {
		shu = 80
	}
	return m81[shu-1]
}

// FindKx 查找字符对应的康熙笔画数和五行属性
func FindKx(c rune) *RKx {
	loadkx()
	return mkx[c]
}

// AssessName 评测名称
func AssessName(name string) *NameReport {
	var n int32
	var wu string
	var chars []RKx
	for _, r := range name {
		var find bool
		switch {
		case r >= '0' && r <= '9':
			n += r - '0'
		case r >= 'a' && r <= 'z':
			n += r - 'a' + 1
		case r >= 'A' && r <= 'Z':
			n += r - 'A' + 1
		case r >= 0x4e00: // r >= 0x4e00 && r <= 0x9fa5
			if rkx := FindKx(r); rkx != nil {
				n += int32(rkx.Bi)
				if rkx.Wu != "" {
					wu = rkx.Wu
				}
				chars = append(chars, *rkx)
				find = true
			}
		}
		if !find {
			chars = append(chars, RKx{Zi: string(r)})
		}
	}

	if n == 0 {
		return nil
	}

	r81 := Find81(int(n))
	if r81 == nil {
		return nil
	}
	return &NameReport{
		Name:  name,
		Chars: chars,
		Wu:    wu,
		R81:   *r81,
	}
}

// NameReport 名称评测结果
type NameReport struct {
	Name  string `json:"name,omitempty"`
	Chars []RKx  `json:"chars,omitempty"`
	Wu    string `json:"wu,omitempty"` // 五行属性
	R81   `json:",inline"`
}

// AssessPhone 评测手机号码
func AssessPhone(phoneNum string) *PhoneReport {
	sn := phoneNum
	if len(sn) > 4 {
		sn = phoneNum[len(phoneNum)-4:]
	}
	n, _ := strconv.Atoi(sn)
	r81 := Find81(n)
	if r81 == nil {
		return nil
	}
	return &PhoneReport{
		Phone: phoneNum,
		R81:   *r81,
	}
}

// PhoneReport 手机号评测结果
type PhoneReport struct {
	Phone string `json:"phone,omitempty"`
	R81   `json:",inline"`
}
