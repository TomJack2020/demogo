package draft

import (
	"fmt"
	"regexp"
	"strings"
)

func StringCheck() {

	// Declare a string variable

	a := "Hello World"
	b := "hello World"

	if strings.EqualFold(a, b) {
		fmt.Println("Both strings are equal after converting to lowercase.")
	} else {
		fmt.Println("Both strings are not equal after converting to lowercase.")
	}

	// 使用Compile来编译正则表达式
	text := "word This is a TEST string. We can find the WORDS here. word"
	word := "word"
	// 创建一个不区分大小写的正则表达式
	//re := regexp.MustCompile(`(?i)` + word)

	// 如果要精确匹配单词，可以在单词前后加上单词边界元字符 \b：
	re := regexp.MustCompile(`\b(?i)` + word + `\b`)
	// 查找所有匹配项
	match := re.FindAllString(text, -1)
	fmt.Println(match) // 输出: [word WORD]

}
