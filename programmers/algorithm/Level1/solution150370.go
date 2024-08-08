import (
    "strconv"
    "strings"
    "time"
)

func solution(today string, terms []string, privacies []string) []int {
    result := []int{}
    
    todayDate, _ := time.Parse("2006.01.02", today)

    termsMap := make(map[string]int, len(terms))
    for _, v := range terms {
        term := strings.Split(v, " ")
        termDate, _ := strconv.Atoi(term[1])
        termsMap[term[0]] = termDate
    }
    
    for i, v := range privacies {
        privacy := strings.Split(v, " ")
        pDate, _ := time.Parse("2006.01.02", privacy[0])
        if todayDate.After(pDate.AddDate(0, termsMap[privacy[1]], -1)) {
            result = append(result, i+1)
        }
    }
    
    return result
}
