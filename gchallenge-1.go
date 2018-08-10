package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func repeatStr(times int, s string, level int) string {
	fmt.Printf("(\"%s\") X %d @LEVEL-%d\n", s, times, level)
	result := ""
	for i := 0; i < times; i++ {
		result += s
	}
	return result
}

func getNum(i *int, s string) int {
	num := string(s[*i])
	*i++
	for *i < len(s) {
		if s[*i] == '[' {
			result, _ := strconv.Atoi(num)
			return result
		}
		num += string(s[*i])
		*i++
	}
	return 0
}

func decompressInner(i *int, s string, level *int, closeBracket *int) string {
	result := ""
	//innerLevel := iLevel
	for *i < len(s) {
		matched, _ := regexp.MatchString("]", string(s[*i]))
		if matched { //state::CloseBrackets
			fmt.Printf("] - INNER CLOSE BRACKET RESULT:%s *Level:%d closeBracket:%d\n", result, *level, *closeBracket)

			if *level != *closeBracket {
				fmt.Printf("++++INNER RETURN i:%d, RESULT:%s, *Level:%d, closeBracket:%d\n", *i, result, *level, *closeBracket)
				*level--
				*closeBracket++
				//*i++
				return result
			}
			if *level == *closeBracket {
				return result
			}
			//*level--
			//innerLevel--
		}
		matched, _ = regexp.MatchString("[a-z]", string(s[*i]))
		if matched { //state::letter
			result += string(s[*i])
		}
		matched, _ = regexp.MatchString("[0-9]", string(s[*i]))
		if matched { //state::number
			num := getNum(&*i, s)
			if s[*i] == '[' {
				*level++
				fmt.Printf("INNER CALLING INNER: %d X \"s[%d]\" *Level:%d closeBracket:%d result:%s\n", num, *i, *level, *closeBracket, result)
				result += repeatStr(num, decompressInner(&*i, s, &*level, &*closeBracket), *level)
				// if *level >= innerLevel {
				// 	*i++
				// 	result += decompressInner(&*i, s, &*level, iLevel+1)
				// }
				//result += tmpResult
				fmt.Printf("INNER REPEATER END: Current s[%d] *Level:%d closeBracket:%d result:%s\n", *i, *level, *closeBracket, result)
			}
		}

		*i++
	}
	return result
}

func decompress(s string) string {
	fmt.Printf("DECOMPRESS (\"%s\") @LEVEL-0\n", s)
	result := ""
	level := 0
	closeBracket := 0
	i := 0
	for i < len(s) {
		matched, _ := regexp.MatchString("[a-z]", string(s[i]))
		if matched { //state::letter
			result += string(s[i])
		}
		matched, _ = regexp.MatchString("[0-9]", string(s[i]))
		if matched { //state::number
			num := getNum(&i, s)
			level = 1
			if s[i] == '[' {
				fmt.Printf("CALLING INNER: %d X \"s[%d]\" *Level:%d result:%s\n", num, i, level, result)
				closeBracket = 0
				result += repeatStr(num, decompressInner(&i, s, &level, &closeBracket), level)
				fmt.Printf("DECOMPRESS REPEATER END: Current s[%d]=\"%s\" *Level:%d result:%s\n", i, string(s[i]), level, result)
			}
		}
		fmt.Printf("DECOMPRESS MAIN LOOP: Current s[%d]=\"%s\" *Level:%d result:%s\n", i, string(s[i]), level, result)
		i++
	}
	return result
}

func main() {
	//input := "3[ab2[x3[i]]c]4[ab]c"
	//input := "3[ab2[x]c]4[ab]c"
	//input := "3[abc]4[ab]c"
	input := "tr3[u8[c]b9[q]]4[ab]c"
	//input := "3[ab2[w2[y]v]]z"
	//input := "xyv3[a8[x2[w]y]b2[o]c]2[asg]z"
	fmt.Printf("\ninput: %s\noutput: %s\n\n", input, decompress(input))
}
