package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Converts a hex color code to rgb or rgba string
func hexToRGB(hex string) string {
	hex = strings.TrimPrefix(hex, "#")
	length := len(hex)
	var r, g, b int

	//  Default alpha value for rgba
	var a float64 = -1

	switch length {
		// single rgb means rrggbb,
		// so that's why i have to use 17 to multiply to get the int value for that hex char.
	case 3: // #rgb
		r = hexCharToInt(hex[0]) * 17
		g = hexCharToInt(hex[1]) * 17
		b = hexCharToInt(hex[2]) * 17
	case 4: // #rgba
		r = hexCharToInt(hex[0]) * 17
		g = hexCharToInt(hex[1]) * 17
		b = hexCharToInt(hex[2]) * 17
		a = float64(hexCharToInt(hex[3]) * 17) / 255
	case 6: // #rrggbb
		r = mustHexToInt(hex[0:2])
		g = mustHexToInt(hex[2:4])
		b = mustHexToInt(hex[4:6])
	case 8: // #rrggbbaa
		r = mustHexToInt(hex[0:2])
		g = mustHexToInt(hex[2:4])
		b = mustHexToInt(hex[4:6])
		a = float64(mustHexToInt(hex[6:8])) / 255
	default:
		return hex // Invalid length
	}

	if a >= 0 {
		return fmt.Sprintf("rgba(%d %d %d / %.5f)", r, g, b, a)
	}
	return fmt.Sprintf("rgb(%d %d %d)", r, g, b)
}

// Converts a single hex char like 'a' or 'f' to integer
func hexCharToInt(c byte) int {
	val, _ := strconv.ParseInt(string(c), 16, 0)
	return int(val)
}

// Converts hex substring to int with panic on error
func mustHexToInt(hexStr string) int {
	val, err := strconv.ParseInt(hexStr, 16, 0)
	if err != nil {
		panic(fmt.Sprintf("invalid hex value: %s", hexStr))
	}
	return int(val)
}

func processLine(line string) string{
	re := regexp.MustCompile(`#([a-fA-F0-9]{3,8})`)
	return re.ReplaceAllStringFunc(line, func(match string) string {
	 	return hexToRGB(match)

	})}


func convertMode(){
	scanner:= bufio.NewScanner(os.Stdin)

	// Read lines from standard input
	// ex: cat simple.css | go run convert.go -c
	for scanner.Scan(){
		line:= scanner.Text()
		fmt.Println(processLine(line))
	}
}


func main(){


	if len(os.Args) > 1 && os.Args[1] == "-c"{
		convertMode()
	}else{
		fmt.Println("Usage: go run convert.go -c")
	}

}