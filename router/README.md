Lily Router
============

Given a path return a controller. It support parameters in path. It parse the 
yaml file letter by letter creating a route map (# for regex). to find the path
the parser goes letter by letter. When it don't find a path, attempts a regex if
any.

```go
   
    package main
    //  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefgh
    // ijklmnopqrstuvwxyz{|}~

    import (
	    "fmt"
    )

    func main() {
        result := []byte{}
        for i := byte(32); i < 127; i++ {
                result = append(result, i)
        }
	    fmt.Println(string(result))
    }
```