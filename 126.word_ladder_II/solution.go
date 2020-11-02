package leetcode

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := make(map[string]struct{})
	current, trace := make(map[string]struct{}), make(map[string][]string)
	for _, s := range wordList {
		dict[s] = struct{}{}
		trace[s] = make([]string, 0)
	}
	dict[beginWord] = struct{}{}
	trace[beginWord] = make([]string, 0)
	current[beginWord] = struct{}{}
	_, ok := current[endWord]
	for len(current) != 0 && !ok {
		for word := range current {
			delete(dict, word)
		}
		next := make(map[string]struct{})
		for word := range current {
			for i := range word {
				for _, c := range "abcdefghijklmnopqrstuvwxyz" {
					candidate := word[:i] + string(c) + word[i+1:]
					if _, ok := dict[candidate]; ok {
						trace[candidate] = append(trace[candidate], word)
						next[candidate] = struct{}{}
					}
				}
			}
		}
		current = next
		_, ok = current[endWord]
	}
	var results [][]string
	if len(current) != 0 {
		backtrace(&results, trace, []string{}, endWord)
	}
	return results
}

func backtrace(results *[][]string, trace map[string][]string, path []string, word string) {
	if len(trace[word]) == 0 {
		*results = append(*results, append([]string{word}, path...))
	} else {
		for _, prev := range trace[word] {
			backtrace(results, trace, append([]string{word}, path...), prev)
		}
	}
}
