package pkg

// import "fmt"
type WordInfo struct {
	freq  int
	posit int
}

// BEGIN (write your solution here)
func MostPopularWord(words []string) string {
	uniqWords := make(map[string]WordInfo, len(words))
	for i, word := range words {
		value, key_exist := uniqWords[word]
		if key_exist {
			uniqWords[word] = WordInfo{value.freq + 1, value.posit}
		} else {
			uniqWords[word] = WordInfo{1, i}
		}
	}
	keys := make([]string, len(uniqWords))
	i := 0
	for k := range uniqWords {
		keys[i] = k
		i++
	}

	var res string = keys[0]
	for j := 1; j < len(keys); j++ {
		if uniqWords[res].freq < uniqWords[keys[j]].freq {
			res = keys[j]
		}
		if uniqWords[res].freq == uniqWords[keys[j]].freq {
			if uniqWords[res].posit > uniqWords[keys[j]].posit {
				res = keys[j]
			}
		}
	}

	return res
}

func MostPopularWord2(words []string) string {
    mpw := ""
    max := 0 
    count := make(map[string]int, len(words))
    for i := len(words) - 1; i >= 0; i-- {
        count[words[i]] += 1
        if count[words[i]] >= max {
            max = count[words[i]]
            mpw = words[i]
        }
    }
    return mpw
}
// END
