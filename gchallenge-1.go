package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func repeatStr(times int, s string, level int) string {
	fmt.Printf("REPEATING \"%s\" %d TIME/S @STATE-%d\n", s, times, level)
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

func decompressInner(i *int, s string, level *int) string {
	fmt.Printf("DECOPRESS INNER FUNC with i:%d @STATE-%d\n", *i, *level)
	result := ""
	for *i < len(s) {
		matched, _ := regexp.MatchString("[a-z]", string(s[*i]))
		if matched { //state::letter
			result += string(s[*i])
		}
		matched, _ = regexp.MatchString("[0-9]", string(s[*i]))
		if matched { //state::number
			num := getNum(&*i, s)
			if s[*i] == '[' {
				*level++
				result += repeatStr(num, decompressInner(&*i, s, &*level), *level)
			}
		}
		matched, _ = regexp.MatchString("[", string(s[*i]))
		if matched { //state::OpenBrackets
			*i++
			result += string(s[*i])
		}
		matched, _ = regexp.MatchString("]", string(s[*i]))
		if matched { //state::CloseBrackets
			if *level >= 1 { //state::CloseBrackets for top state/s
				*level++
				return result
			}
			*level-- //state::CloseBrackets for higher state
		}
		*i++
	}
	return result
}

func decompress(s string) string {
	fmt.Printf("DECOPRESS FUNC for string:%s @STATE-0\n", s)
	result := ""
	level := 0
	i := 0
	for i < len(s) {
		matched, _ := regexp.MatchString("[a-z]", string(s[i]))
		if matched { //state::letter
			result += string(s[i])
		}
		matched, _ = regexp.MatchString("[0-9]", string(s[i]))
		if matched { //state::number
			num := getNum(&i, s)
			if s[i] == '[' {
				result += repeatStr(num, decompressInner(&i, s, &level), level)
			}
		}
		matched, _ = regexp.MatchString("[", string(s[i]))
		if matched { //state::OpenBrackets
			i++
			result += string(s[i])
		}
		level-- //state::CloseBrackets
		i++
	}
	return result
}

func main() {
	input := "y2[a2[x3[h2[v4[d]]]y]b]c"
	//input := "xyv3[abc]"
	fmt.Printf("\ninput: %s\noutput: %s\n\n", input, decompress(input))
}
