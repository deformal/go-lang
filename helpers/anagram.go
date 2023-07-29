package helpers

const (
	word1 string = "car"
	word2 string = "rac"
)

func wordCounter(word string) map[string]uint16 {
	wordCount := map[string]uint16{}
	for _, w := range word {
		_, ok := wordCount[string(w)]
		if !ok {
			wordCount[string(w)] = 1
		} else {
			wordCount[string(w)] += 1
		}
	}
	return wordCount
}

func Anagram() bool {
	if len(word1) != len(word2) {
		return false
	}
	wc1 := wordCounter(word1)
	wc2 := wordCounter(word2)

	for key, value := range wc1 {
		if val, ok := wc2[key]; !ok || val != value {
			return false
		}
	}
	return true
}
