package tool

import (
	"fmt"
	"strconv"
	"strings"
)
func StringsExec(s string) {
	//第一次出现的索引
	fmt.Println(strings.Index(s, "l"))
	//最后一次出现的索引
	fmt.Println(strings.LastIndex(s, "l"))
	//是否以指定内容开头
	fmt.Println(strings.HasPrefix(s, "small"))
	//是否以指定内容结尾
	fmt.Println(strings.HasSuffix(s, "ming"))
	//是否包含指定字符串
	fmt.Println(strings.Contains(s, "mi"))
	//全变小写
	fmt.Println(strings.ToLower(s))
	//全大写
	fmt.Println(strings.ToUpper(s))
	//把字符串中前n个old子字符串替换成new字符串,如果n小于0表示全部替换.
	//如果n大于old个数也表示全部替换
	fmt.Println(strings.Replace(s, "m", "k", -1))
	//把字符串重复count遍
	fmt.Println(strings.Repeat(s, 2))
	//去掉字符串前后指定字符
	fmt.Println(strings.Trim(s, " ")) //去空格可以使用strings.TrimSpace(s)
	//根据指定字符把字符串拆分成切片
	fmt.Println(strings.Split(s, "m"))
	//使用指定分隔符把切片内容合并成字符串
	arr := []string{"small", "ming"}
	fmt.Println(strings.Join(arr, ""))
}

func String2Int(str5 string) int {
	int5, err := strconv.Atoi(str5)
	if err != nil {
		fmt.Println(err)
		return -1 //error
	} else {

		return int5
	}
}
