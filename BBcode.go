package qalam

import (
	"regexp"
	"strings"
)

type splittedTxt struct {
	t    string
	code string
}

type formattedTxt struct {
	t              string
	inheritedCodes []uint
}

func splitTags(s string) []splittedTxt {
	firstTag := regexp.MustCompile(openingTagRegex)
	splits := []splittedTxt{}
OUTER:
	for {
		openTagLoc := firstTag.FindAllIndex([]byte(s), -1)
		for i, v := range openTagLoc {
			code := s[v[0]+1 : v[1]-1]
			closeTag := "[/" + code + "]"
			closeTagLoc := strings.Index(s, closeTag)
			if closeTagLoc == -1 {
				nextI := 0
				if i == len(openTagLoc)-1 {
					nextI = len(s)
				} else {
					nextI = openTagLoc[i+1][0]
				}
				splits = append(splits, splittedTxt{t: s[:nextI]})
				s = s[nextI:]
				continue OUTER
			}
			closeTagLoc += len(closeTag)
			beforeTag := splittedTxt{t: s[:v[0]]}
			afterTag := splittedTxt{t: s[closeTagLoc:]}
			if beforeTag.t != "" {
				splits = append(splits, beforeTag)
			}
			splits = append(splits, splittedTxt{t: s[v[1] : closeTagLoc-len(closeTag)], code: code})
			if afterTag.t != "" && shouldAppendAfterTag(openTagLoc, closeTagLoc, i) {
				splits = append(splits, afterTag)
			}
			s = s[closeTagLoc:]
			continue OUTER
		}
		if len(splits) == 0 {
			splits = append(splits, splittedTxt{t: s})
		}
		break
	}
	return splits
}
func shouldAppendAfterTag(i [][]int, closingLoc int, currentIndex int) bool {
	if currentIndex == len(i)-1 {
		return true
	}
	currentLoc := i[currentIndex][0]
	for i, v := range i {
		if i == currentIndex || currentLoc > v[0] {
			continue
		}
		if closingLoc < v[1] {
			return false
		}
	}
	return true
}
func format(s *formattedTxt) string {
	splits := splitTags(s.t)
	haveNested := regexp.MustCompile(openingTagRegex)
	n := ""
	for _, v := range splits {
		codes := append(s.inheritedCodes, getStyles(v.code)...)
		if v.code == "" {
			n += applyStyles(v.t, codes)
			continue
		}
		if ok := haveNested.Match([]byte(v.t)); ok {
			n += format(&formattedTxt{
				t:              v.t,
				inheritedCodes: codes,
			})
			continue
		}
		n += applyStyles(v.t, codes)
	}
	return n
}
