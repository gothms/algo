package main

import (
    "strings"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func mergeAlternately(word1 string, word2 string) string {
    var sb strings.Builder
    i,j := 0,0
    for ; i <len(word1)&&j<len(word2) ; i,j=i+1,j+1 {
        sb.WriteByte(word1[i])
        sb.WriteByte(word2[j])
    }
    if i<len(word1){
        sb.WriteString(word1[i:])
    }else{
        sb.WriteString(word2[j:])
    }
    return sb.String()
}
//leetcode submit region end(Prohibit modification and deletion)
