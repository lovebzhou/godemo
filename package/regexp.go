package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func init() {
	fmt.Println("==================== BEGIN regexp =================")
}

func demoPattern() {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	fmt.Println("searchIn:", searchIn)
	fmt.Println("pat:", pat)

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

func demoSubMatch() {
	var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	m := validPath.FindStringSubmatch("/edit/123")

	for i, k := range m {
		fmt.Println(i, k)
	}
}

func demoRegexp() {
	demoPattern()
	demoSubMatch()
}