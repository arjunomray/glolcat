package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func rgbify(i int) (int, int, int) {
    f := 0.1

    return int(math.Sin(f*float64(i)+0)*127 +128),
        int(math.Sin(f*float64(i)+2)*127 +128),
        int(math.Sin(f*float64(i)+4)*127 +128)
}

func print(output []rune){
    for j:= range output{
        
        r, g, b := rgbify(j)
        fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
    }
    fmt.Println()
}

func main(){
    info, _ := os.Stdin.Stat()

    var output []rune
    if info.Mode()&os.ModeCharDevice != 0 {
        fmt.Println("The command is intended to work with pipes")
        fmt.Println("Usage: fortune | glolcat")
    }

    reader := bufio.NewReader(os.Stdin)
    for {
        input, _, err := reader.ReadRune()
        if err != nil && err == io.EOF{
            break
        }
        output = append(output, input)
    }

    print(output)
}
