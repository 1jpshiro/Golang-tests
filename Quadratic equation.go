package main
import (
    "fmt"
    "math"
    "os"
    "bufio"
)

func findRoots(a, b, c float64) (float64, float64) {
    de := func(a, b, c float64) float64 {
        return math.Sqrt(b*b-4*a*c)
    }
    i := de(a, b, c)
    return (-b-i) / (2*a), (-b+i) / (2*a)
}
    
func main() {
    fmt.Print("Enter the coefficients separated with spaces: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input := strings.TrimSpace(input)
    nums := strings.Split(input, " ")
    a, b, c := nums[0], nums[1], nums[2]
    x1, x2 := findRoots(a, b, c)
    fmt.Printf("Roots: %v, %v", int(x1), int(x2))
}
