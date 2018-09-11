package main


import (
"fmt"
"bufio"
"os"
"time"
)

var words map[string]int

func findFunnels(word string) []string {
    funnels := make([]string, 1)
    r := []rune(word)
    wordAsArray := r[:]
    for i:=0; i < len(wordAsArray); i++ {
        tempArray := make([]rune, len(wordAsArray))
        copy(tempArray[:], wordAsArray[:])
        copy(tempArray[i:], tempArray[i+1:])
        tempArray = tempArray[:len(tempArray)-1]
        funnels = append(funnels, string(tempArray))
    }
    return funnels
}

func wordExists(word string) bool {
    existingWords := make(map[string]int)
    existingWords["gnash"] = 1
    existingWords["what"] = 1
    
    _, exists := words[word]
    return exists
}

func readWords() {
    file, err := os.Open("/tmp/lol.txt")
    if err != nil {
       fmt.Println(err) 
    }
    defer file.Close()

    words = make(map[string]int)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //fmt.Println(scanner.Text())
        word := scanner.Text()
        words[word] = 1 
    }
}

//func findDepth(word string, depth int) int {
//    
//}

func findFunnelLenght(word string) int {
    funnels := findFunnels(word)

    longest := 1
    for i:=0; i < len(funnels); i++ {
        if wordExists(funnels[i]) {
            temp := 1+findFunnelLenght(funnels[i])
            if temp > longest {
                longest = temp
            }
        }
    }
    return longest
}

func main() {
    fmt.Println("word funnel")
    readWords()

    //words := []string{"gnash", "princesses", "turntables", "implosive",
      //                  "programmer"}
    start := time.Now()
    for k := range words {
        lent := findFunnelLenght(k)
        if lent==10 {
            fmt.Println(k + " funnels => ",lent)
            break
        }
    }
    end := time.Now()
    fmt.Println(start.Unix())
    fmt.Println(end.Unix())
    fmt.Println(end.Unix()-start.Unix())
 

}



