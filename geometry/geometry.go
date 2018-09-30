// geometry.go
package main

import (
    "fmt"
    "geometry/rectangle" //导入自定义包
    "log"
)
/*
 * 1. 包级别变量
*/
var rectLen,rectWidth float64 = 6,7

/*
 * 2. init函数检查长和宽是否大于0
*/

func init() {
    fmt.Println("main package initialized")
    if rectLen < 0 {
        log.Fatal("length is less than zero")
    }
    if rectWidth < 0 {
        log.Fatal("width is less than zero")
    }
}


func main() {
    fmt.Println("Geometrical shape properties")
    /*Area function of rectangle package used*/
    fmt.Printf("Area of rectangle %.2f\n",rectangle.Area(rectLen,rectWidth))
    /*Diagonal function of rectangle package used*/
    fmt.Printf("diagonal of the rectangle %.2f", rectangle.Diagonal(rectLen,rectWidth))
}
    
