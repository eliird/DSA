package main

import "fmt"

func romanToInt(s string) int {
	number := 0
	skip_flag := false
	myDict := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	for i, char := range s {
		fmt.Println(number)
		if skip_flag {
			skip_flag = false
			if i == len(s)-2 {
				number += myDict[s[i+1]]
				break
			}
			continue
		}
		if len(s) == 1 {
			return myDict[s[0]]
		}
		// fmt.Println(skip_flag, i)
		switch char {
		case 'I':
			if s[i+1] == 'V' {
				number += 4
				skip_flag = true
			} else if s[i+1] == 'X' {
				number += 9
				skip_flag = true
			} else {
				number += 1
			}
		case 'V':
			number += 5
		case 'X':
			if s[i+1] == 'L' {
				number += 40
				skip_flag = true
			} else if s[i+1] == 'C' {
				number += 90
				skip_flag = true
			} else {
				number += 10
			}
		case 'L':
			number += 50
		case 'C':
			if s[i+1] == 'D' {
				number += 400
				skip_flag = true
			} else if s[i+1] == 'M' {
				number += 900
				skip_flag = true
			} else {
				number += 100
			}
		case 'D':
			number += 500
		case 'M':
			number += 1000
		}
		// fmt.Println(number, len(s), i)
		if i == len(s)-2 && !skip_flag {
			number += myDict[s[i+1]]
			break
		}

	}
	return number
}

func longestCommonPrefix(strs []string) string {
	min_len := 10000
	for _, str := range strs {
		min_len = min(min_len, len(str))
	}
	first_str := strs[0]

	glob_min := 10000
	for _, str := range strs {
		fmt.Println(str)
		local_min := 10000
		for j := 0; j < min(len(first_str), len(str)); j++ {
			if first_str[j] != str[j] {
				local_min = j - 1
			} else {
				local_min = j
			}
		}
		fmt.Println(local_min, glob_min)
		glob_min = min(local_min, glob_min)
		if glob_min == 10000 || glob_min == -1 {
			return ""
		}
	}
	return first_str[:glob_min-1]
}
func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
