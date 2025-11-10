package services

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type DivinationOutcome struct {
	PrimaryName   string
	SecondaryName string
	UpperTrigram  string
	LowerTrigram  string
	Commentary    string
	Lines         []int
	ChangingLines []int
}

type trigram struct {
	Key         string
	ChineseName string
	Element     string
	Attribute   string
	Meaning     string
}

type hexagramInsight struct {
	Name    string
	Summary string
}

var trigramByBinary = map[string]trigram{
	"111": {Key: "Qian", ChineseName: "乾", Element: "天", Attribute: "阳刚", Meaning: "创始与领导"},
	"011": {Key: "Dui", ChineseName: "兑", Element: "泽", Attribute: "喜悦", Meaning: "交流与欢悦"},
	"101": {Key: "Li", ChineseName: "离", Element: "火", Attribute: "光明", Meaning: "洞察与热情"},
	"001": {Key: "Zhen", ChineseName: "震", Element: "雷", Attribute: "震动", Meaning: "行动与觉醒"},
	"110": {Key: "Xun", ChineseName: "巽", Element: "风", Attribute: "渗透", Meaning: "顺势与传播"},
	"010": {Key: "Kan", ChineseName: "坎", Element: "水", Attribute: "险阻", Meaning: "智慧与深沉"},
	"100": {Key: "Gen", ChineseName: "艮", Element: "山", Attribute: "止息", Meaning: "稳重与界限"},
	"000": {Key: "Kun", ChineseName: "坤", Element: "地", Attribute: "柔顺", Meaning: "承载与滋养"},
}

var hexagramInsights = map[string]hexagramInsight{
	"Qian-Qian": {Name: "乾为天", Summary: "上下皆乾，纯阳之象，强调自强不息与领导力。"},
	"Kun-Kun":   {Name: "坤为地", Summary: "上下皆坤，纯阴之象，提醒以柔顺承载，广纳百川。"},
	"Qian-Kun":  {Name: "天地否", Summary: "阳在上阴在下，气不相交，宜守正待时。"},
	"Kun-Qian":  {Name: "地天泰", Summary: "阴居上阳居下，天地交泰，万物通达。"},
	"Kan-Li":    {Name: "水火既济", Summary: "水在上火在下，动静协调，提醒慎终如始。"},
	"Li-Kan":    {Name: "火水未济", Summary: "火在上水在下，尚未调和，需持之以恒。"},
	"Zhen-Xun":  {Name: "雷风恒", Summary: "雷动风行，恒久之道在于上下相应。"},
	"Gen-Dui":   {Name: "山泽损", Summary: "山上有泽，损益相生，教人损有余补不足。"},
	"Xun-Zhen":  {Name: "风雷益", Summary: "风行雷动，相辅相成，象征增益。"},
	"Kan-Kan":   {Name: "坎为水", Summary: "重坎之象，险阻频仍，需以智慧谨慎前行。"},
	"Li-Li":     {Name: "离为火", Summary: "重离之象，光明炽烈，提醒守正不骄。"},
}

// GenerateDivination produces a stochastic Mei Hua Yi Shu reading.
func GenerateDivination(subject, method string) DivinationOutcome {
	rand.Seed(time.Now().UnixNano())

	lines := make([]int, 6)
	var changing []int
	for i := range lines {
		value := randomLineValue()
		lines[i] = value
		if value == 6 || value == 9 {
			changing = append(changing, i+1) // Mei Hua lines count from bottom as 1
		}
	}

	lower := resolveTrigram(lines[:3])
	upper := resolveTrigram(lines[3:])
	primary := describeHexagram(upper, lower)

	secondaryLines := applyChanging(lines)
	secondaryLower := resolveTrigram(secondaryLines[:3])
	secondaryUpper := resolveTrigram(secondaryLines[3:])
	secondary := describeHexagram(secondaryUpper, secondaryLower)

	commentary := buildCommentary(subject, method, primary, secondary, upper, lower, changing)

	return DivinationOutcome{
		PrimaryName:   primary.Name,
		SecondaryName: secondary.Name,
		UpperTrigram:  fmt.Sprintf("%s · %s", upper.ChineseName, upper.Meaning),
		LowerTrigram:  fmt.Sprintf("%s · %s", lower.ChineseName, lower.Meaning),
		Commentary:    commentary,
		Lines:         lines,
		ChangingLines: changing,
	}
}

func randomLineValue() int {
	seed := rand.Intn(16)
	switch {
	case seed < 4:
		return 6
	case seed < 8:
		return 7
	case seed < 12:
		return 8
	default:
		return 9
	}
}

func applyChanging(lines []int) []int {
	next := make([]int, len(lines))
	copy(next, lines)
	for i, v := range next {
		switch v {
		case 6:
			next[i] = 7
		case 9:
			next[i] = 8
		}
	}
	return next
}

func resolveTrigram(lines []int) trigram {
	binary := make([]rune, len(lines))
	for i, v := range lines {
		if v == 7 || v == 9 {
			binary[len(lines)-1-i] = '1'
		} else {
			binary[len(lines)-1-i] = '0'
		}
	}
	key := string(binary)
	if tri, ok := trigramByBinary[key]; ok {
		return tri
	}
	return trigram{Key: "Unknown", ChineseName: "未知", Element: "", Attribute: "平衡", Meaning: "潜藏之势"}
}

func describeHexagram(upper, lower trigram) hexagramInsight {
	key := fmt.Sprintf("%s-%s", upper.Key, lower.Key)
	if insight, ok := hexagramInsights[key]; ok {
		return insight
	}
	summary := fmt.Sprintf("上卦为%s，象征%s；下卦为%s，象征%s。两者互动提醒关注%s与%s之间的平衡。",
		upper.ChineseName, defaultMeaning(upper), lower.ChineseName, defaultMeaning(lower),
		defaultAttribute(upper), defaultAttribute(lower))
	return hexagramInsight{
		Name:    fmt.Sprintf("%s上%s下", upper.ChineseName, lower.ChineseName),
		Summary: summary,
	}
}

func defaultMeaning(t trigram) string {
	if t.Meaning != "" {
		return t.Meaning
	}
	return "内在潜能"
}

func defaultAttribute(t trigram) string {
	if t.Attribute != "" {
		return t.Attribute
	}
	return "动静"
}

func buildCommentary(subject, method string, primary, secondary hexagramInsight, upper, lower trigram, changing []int) string {
	var b strings.Builder
	if subject != "" {
		b.WriteString(fmt.Sprintf("【主题】%s\n", subject))
	}
	if method != "" {
		b.WriteString(fmt.Sprintf("【起卦方式】%s\n", method))
	}
	b.WriteString(fmt.Sprintf("【本卦】%s：%s\n", primary.Name, primary.Summary))
	b.WriteString(fmt.Sprintf("【互卦信息】上卦%s（%s），下卦%s（%s）。\n", upper.ChineseName, defaultMeaning(upper), lower.ChineseName, defaultMeaning(lower)))
	b.WriteString(fmt.Sprintf("【变卦】%s：%s\n", secondary.Name, secondary.Summary))
	if len(changing) > 0 {
		b.WriteString(fmt.Sprintf("【动爻】第%s爻提示顺势调整。\n", formatLines(changing)))
	} else {
		b.WriteString("【动爻】本卦未出现动爻，宜以本卦为主。\n")
	}
	b.WriteString("【建议】结合当前形势，保持内外呼应，循序渐进。")
	return b.String()
}

func formatLines(lines []int) string {
	parts := make([]string, len(lines))
	for i, v := range lines {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(parts, "、")
}
