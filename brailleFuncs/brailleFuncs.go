package brailleFuncs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const alph = " a c bif e d hjg k m lsp o n rtq              w  u x v   z y    "

// Take a number in binary format, return a string representation of the braille.
func numToBraille(num int64) string {
	s := strconv.FormatInt(num, 2)
	s = strings.Repeat("0", 6-len(s)) + s
	a, b, c := s[4:], s[2:4], s[:2]

	a = strings.Replace(a, "0", ".", -1)
	a = strings.Replace(a, "1", "O", -1)

	b = strings.Replace(b, "0", ".", -1)
	b = strings.Replace(b, "1", "O", -1)

	c = strings.Replace(c, "0", ".", -1)
	c = strings.Replace(c, "1", "O", -1)

	return fmt.Sprintf("%v%v%v", c, b, a)
}

// Take in a string representation of braille and return the corresponding binary.
func brailleToNum(name string) ([]string, []int64) {
	var m []string
	var n []int64
	j := 0

	file, _ := os.Open(name)
	defer file.Close()
	bfile := bufio.NewScanner(file)

	for bfile.Scan() {
		line := bfile.Text()
		tokens := strings.Split(line, ` `)

		if j == 0 {
			m = make([]string, len(tokens))
			n = make([]int64, len(tokens))
			j += 1
		}
		for i, v := range tokens {
			switch v {
			case "..":
				m[i] = "00" + m[i]
			case "O.":
				m[i] = "01" + m[i]
			case ".O":
				m[i] = "10" + m[i]
			case "OO":
				m[i] = "11" + m[i]
			}
		}
	}
	for i, v := range m {
		n[i], _ = strconv.ParseInt(v, 2, 0)
	}
	return m, n
}

// Take the string representation of braille and print to screen.
func printBraille(b []string) {
	// Convert this to only print the entire string to terminal width
	// as currently the output breaks if the message is long
	var d, e, f string
	var d2, e2, f2 string // Temp variables if message is long
	width := terminalWidth()

	for i, _ := range b {
		d += string(b[i][5]) + string(b[i][4]) + " "
		e += string(b[i][3]) + string(b[i][2]) + " "
		f += string(b[i][1]) + string(b[i][0]) + " "
	}
	for {
		if len(d) < width {
			width = len(d)
			if len(d) <= 0 {
				break
			}
		}
		d, d2 = d[width:], d[:width]
		e, e2 = e[width:], e[:width]
		f, f2 = f[width:], f[:width]
		fmt.Printf("%v\n%v\n%v\n\n", d2, e2, f2)
	}
}

// Take braille input and convert to english.
func BrailleToMessage(s string) {
	var msg string

	binary, decimal := brailleToNum(s)

	for i := 0; i < len(binary); i++ {
		msg += string(alph[decimal[i]])
	}
	fmt.Printf("%v\n", msg)
}

// Take english input message and convert to braille.
func MessageToBraille(s string) {
	braille := make([]string, len(s))
	for i, _ := range braille {
		curLet := strings.Index(alph, string((s)[i]))
		braille[i] = numToBraille(int64(curLet))
	}
	printBraille(braille)
}

func terminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	widthStr := strings.Split(string(out), " ")[1]
	width, _ := strconv.Atoi(strings.TrimRight(widthStr, "\n"))
	return width
}
