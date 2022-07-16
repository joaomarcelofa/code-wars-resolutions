package stringsmix

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type stringData struct {
	letter    rune
	frequency int
	sequency  string
}

type charsMap map[rune]int

func Mix(s1 string, s2 string) string {
	s1Map := transformToCharsMap(s1)
	s2Map := transformToCharsMap(s2)
	stringDataList := buildStringDataList(s1Map, s2Map)
	sortedDataList := sortStringData(stringDataList)
	return mountFullString(sortedDataList)
}

func transformToCharsMap(s string) charsMap {
	dataMap := charsMap{}
	lowerCaseRegex := regexp.MustCompile(`(?m)[a-z]`)
	for _, char := range s {
		if lowerCaseRegex.MatchString(string(char)) {
			dataMap[char] += 1
		}
	}

	filteredDataMap := charsMap{}
	for k, v := range dataMap {
		if v > 1 {
			filteredDataMap[k] = v
		}
	}

	return filteredDataMap
}

func buildStringDataList(s1CharMap charsMap, s2CharMap charsMap) []stringData {
	stringDataMap := map[rune]stringData{}

	// Initializing the list with s1
	for char, frequency := range s1CharMap {
		stringDataMap[char] = stringData{char, frequency, mountStringToPrint("1", string(char), frequency)}
	}

	// Comparing the s1 list with s2
	for char, frequency := range s2CharMap {
		s1Frequency, existsInS1 := s1CharMap[char]
		if !existsInS1 {
			stringDataMap[char] = stringData{char, frequency, mountStringToPrint("2", string(char), frequency)}
			continue
		}

		majorFrequency := frequency
		prefixChar := ""
		// Checking if is greater, or equal
		switch {
		case frequency > s1Frequency:
			prefixChar = "2"
		case frequency == s1Frequency:
			prefixChar = "="
		case frequency < s1Frequency:
			prefixChar = "1"
			majorFrequency = s1Frequency
		}
		stringDataMap[char] = stringData{char, majorFrequency, mountStringToPrint(prefixChar, string(char), majorFrequency)}
	}

	// Building the string data list
	stringDataList := []stringData{}
	for _, stringData := range stringDataMap {
		stringDataList = append(stringDataList, stringData)
	}

	return stringDataList
}

func sortStringData(stringDataList []stringData) []stringData {
	sort.Slice(stringDataList, func(i, j int) bool {
		a := stringDataList[i]
		b := stringDataList[j]

		isFrequencyGreater := a.frequency > b.frequency
		isFrequencyEqual := a.frequency == b.frequency
		isLettersLower := strings.Compare(a.sequency, b.sequency) == -1

		// Reasons to swap:
		// Frequency is bigger
		if isFrequencyGreater {
			return true
		}

		// Frequency is equal, so we have to use the second criteria: lexical order
		if isFrequencyEqual && isLettersLower {
			return true
		}

		return false
	})

	return stringDataList
}

func mountStringToPrint(s string, letter string, frequency int) string {
	return fmt.Sprintf("%s:%s", s, strings.Repeat(letter, int(frequency)))
}

func mountFullString(sortedDataList []stringData) string {
	strs := []string{}
	for _, s := range sortedDataList {
		strs = append(strs, s.sequency)
	}
	return strings.Join(strs, "/")
}
