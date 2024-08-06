package main

import (
        "fmt"
        "strings"
        "unicode"
)

func wordFrequency(s string ) map[string]int{
        freq := make(map[string]int)
        s = strings.ToLower(s)
        var wordBuilder strings.Builder
        for _, char := range s{
                if unicode.IsLetter(char)|| unicode.IsDigit(char) || unicode.IsSpace(char){
                        wordBuilder.WriteRune(char)
                } 
        }
        words := strings.Fields(wordBuilder.String())
        for _, word := range words{
                freq[word]++
        }
        return freq
}

func main() {
        s := "Hello, World! Hello world. This is a test. Test, this is."
        freq := wordFrequency(s)
        for k, v := range freq {
                fmt.Printf("%s: %d\n", k, v)
        }
}
