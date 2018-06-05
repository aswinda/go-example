
package variable

import "fmt"

func main() {
    // declaration variable int
    var a int
    a = 5
    fmt.Printf("%d\n", a)
    
    // or you can
    b := 15
    fmt.Printf("%d\n", b)

    // Multiple var declarations may also be grouped
    var (
        c int
        d int32
    )
    c = 1; d = 2
    fmt.Printf("%d %d\n", c, d)
}
