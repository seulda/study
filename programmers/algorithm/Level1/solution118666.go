func solution(survey []string, choices []int) string {

    typeCnt := map[string]int{
        "R" : 0,
        "T" : 0,
        "C" : 0,
        "F" : 0,
        "J" : 0,
        "M" : 0,
        "A" : 0,
        "N" : 0,
    }
    
    for i, v := range survey {
        choice := choices[i]
        if choice == 4 {
            continue
        } else if choice > 4 {
            typeCnt[string(v[1])] += (choice - 4)
        } else if choice < 4 {
            typeCnt[string(v[0])] += (4 - choice)
        }
    }
    
    result := typeCheck(typeCnt)

    return result
}

func typeCheck(typeCnt map[string]int) string {

    rt, cf, jm, an := "", "", "", ""

    if typeCnt["R"] >= typeCnt["T"] {
        rt = "R"
    } else {
        rt = "T"
    }
    if typeCnt["C"] >= typeCnt["F"] {
        cf = "C"
    } else {
        cf = "F"
    }
    if typeCnt["J"] >= typeCnt["M"] {
        jm = "J"
    } else {
        jm = "M"
    }
    if typeCnt["A"] >= typeCnt["N"] {
        an = "A"
    } else {
        an = "N"
    }

    return rt + cf + jm + an
}
