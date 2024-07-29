package main
import "fmt"

func nonRepeat(s string) string {
    if len(s) == 0 {
        return ""
    }

    res := ""
    m := make(map[rune]bool, 0)

    for i := 0; i < len(s); i++ {
        if i == 0 {
            if s[i] == s[i+1] {
                m[rune(s[i])] = true
            }
        } else if i == len(s)-1 {
            if s[i] == s[i-1] {
                m[rune(s[i])] = true
            }
        } else {
            if s[i] == s[i-1] || s[i] == s[i+1] {
                m[rune(s[i])] = true
            }
        }
    }
    for j := 0; j < len(s); j++ {
        if !m[rune(s[j])] {
            res = string(s[j])
            break
        }
    }

    return res
}

func main() {
    var s string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&s)
    char := nonRepeat(s)
    fmt.Println("First non-repeating character:", char)
}
