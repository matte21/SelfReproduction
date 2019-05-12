package main

import "fmt"

func main() {
    s := "package main\n\nimport \"fmt\"\n\nfunc main() {\n    s := %q\n    fmt.Printf(s, s)\n}\n\n"
    fmt.Printf(s, s)
}

