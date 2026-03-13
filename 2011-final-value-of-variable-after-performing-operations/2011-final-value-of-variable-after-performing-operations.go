func finalValueAfterOperations(operations []string) int {
    var result int

    for _, operation := range operations {
        switch operation {
            case "--X", "X--":
            result-- 

            case "++X", "X++":
            result++
        }
    }

    return result
}

//Time complexity O(n)
//Space complexity O(1)