package stringutils

func StringOccurrence(s []string) (occurrenceMap map[string]int) {
  occurrenceMap = make(map[string]int)
  
  for _, s := range s {
    count, exist := occurrenceMap[s]
    if exist {
      count++
    } else {
      count = 1
    }
    occurrenceMap[s] = count
  }
  
  return occurrenceMap
}