package main
import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func comPrefix(arr []string) string {
    if len(arr) == 0 {
        return ""
    }
    prefix := arr[0]
    result := ""

    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr[i]) && j < len(prefix); j++ {
            if arr[i][j] != prefix[j] {
                result = prefix[:j]
                break
            }
        }
    }
    return result
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the words separated with spaces: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    arr := strings.Split(input, " ")
    char := comPrefix(arr)
    fmt.Println("Longest common prefix:", char)
}
