package hw03frequencyanalysis

import (
	"math"
	"sort"
	"strings"
)

type WordStat struct {
	word      string
	frequency int
}

func Top10(text string) []string {
	wordFrequencies := make(map[string]int)
	for _, word := range strings.Fields(text) {
		wordFrequencies[word]++
	}

	wordStats := make([]WordStat, 0, len(wordFrequencies))
	for word, frequency := range wordFrequencies {
		wordStats = append(wordStats, WordStat{word, frequency})
	}

	sort.Slice(wordStats, func(i, j int) bool {
		if wordStats[i].frequency == wordStats[j].frequency {
			return wordStats[i].word < wordStats[j].word
		}
		return wordStats[i].frequency > wordStats[j].frequency
	})

	top := make([]string, 0, 10)
	topBound := int(math.Min(10, float64(len(wordStats))))
	for _, stat := range wordStats[:topBound] {
		top = append(top, stat.word)
	}

	return top
}
