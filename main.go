package gotool

import (
	"bufio"
	"fmt"
	"os"
)

//GetStr :
func GetStr() string {
	var str string
	fmt.Scanf("%s", &str)
	return str
}

//CountAlphabet :
func CountAlphabet(str string) [26]int {
	var ret [26]int
	for i := 0; i < len(str); i++ {
		valInt := (int)(str[i] - 97)
		// fmt.Println(valInt)
		ret[valInt]++
	}
	return ret
}

//GetStringListBySpace :
func GetStringListBySpace() []string {
	var ret []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		target := scanner.Text()
		ret = append(ret, target)
	}
	return ret
}

//InputStr :
func InputStr() string {
	var input string
	fmt.Scanln(&input)
	return input
}

//InputAry :
func InputAry(size int) []string {
	var retAry []string
	for i := 0; i < size; i++ {
		val := InputStr()
		if len(val) > 0 {
			retAry = append(retAry, val)
		} else {
			i--
		}

	}

	return retAry
}
