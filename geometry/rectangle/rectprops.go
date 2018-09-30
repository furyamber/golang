// rectprops.go
package rectangle

import (
    "math"
    "fmt"
)

func init() {
    fmt.Println("rectangle package initial")
}

func Area(length, width float64) float64{
    area := length * width
    return area
}

func Diagonal(length,width float64) float64 {
    diagonal := math.Sqrt((length * length)+(width * width))
    return diagonal
}
