func isAnagram(s string, t string) bool {
    if len(s)!= len(t) {
        return false 
    }
    
    seen:= make(map[byte]int)
    for i:=0; i<len(s); i++{
      seen[s[i]]++
     seen[t[i]]--
    }

    for _, str:= range seen{
        if str!= 0 {
            return false 
        }
    }

    return true 
}