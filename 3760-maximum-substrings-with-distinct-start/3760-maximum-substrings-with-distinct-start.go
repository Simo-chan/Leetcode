//In this case it will be enough to just count the amount of unique characters in the given string. 
//So with each character we check if it exists in the map, if not we append it. And in the end we return
// the lenght of the map.

func maxDistinct(s string) int {
    //Using struct for "set" behavior. 
    distincts:= make(map[rune]struct{})

    for _, r := range s {
        if _, exists:= distincts[r]; !exists {
            distincts[r] = struct{}{}
        }
    }

    return len(distincts)
}

//Time complexity is O(n), as we iterate over the given srting only once
//Space complexity is O(1), because even though we create a map, the maximum amount of unique characters
//will be 26 as the tasks states that the given string is consisting of lowercase English letters.