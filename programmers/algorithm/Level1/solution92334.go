import (
    "strings"
)

func solution(id_list []string, report []string, k int) []int {
    userReport := make(map[string][]string, len(id_list))
    userPenalty := make(map[string]int, len(id_list))
    penaltyList := []string{}
    res_list := make([]int, len(id_list))
    
    for _, r := range report {
        report := strings.Split(r, " ")
        chk := false
        for _, ur := range userReport[report[0]] {
            if ur == report[1] {
                chk = true
                break
            }
        }
        if !chk {
            userReport[report[0]] = append(userReport[report[0]], report[1])
            userPenalty[report[1]]++
        }
        
        if userPenalty[report[1]] >= k {
            penaltyList = append(penaltyList, report[1])
        }
    }

    for i, user := range id_list {
        ur := userReport[user]
        for _, r := range ur {
            for _, p := range penaltyList {
                if p == r {
                    res_list[i]++
                    break
                }
            }
        }
    }
    
    return res_list
}
