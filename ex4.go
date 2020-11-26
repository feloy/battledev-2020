/*******
 * Read input from os.Stdin
 * Use: fmt.Println to ouput your result to STDOUT.
 * Use: os.Stderr.WriteString to ouput debugging information to STDERR.
 * ***/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contestResponse() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	if len(parts) < 2 {
		fmt.Fprintf(os.Stderr, "ERREUR 1")
	}
	N, _ := strconv.Atoi(parts[0])

	scanner.Scan()
	line := scanner.Text()
	stroctets := strings.Split(line, " ")
	octets := make([]int, len(stroctets))
	for i, o := range stroctets {
		octets[i], _ = strconv.Atoi(o)
	}

	result := make([][]int, N)
	for i := 0; i < N; i++ {
		result[i] = make([]int, N)
	}

	for x := 0; x < N; x++ {
		r := octets[x]
		result[x][x] = r
		for y := x + 1; y < N; y++ {
			r = r ^ octets[y]
			result[x][y] = r
		}
	}

	resultat := make([]int, 256)

	for scanner.Scan() {
		line = scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			fmt.Fprintf(os.Stderr, "ERREUR 2")
		}
		L, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "err\n")
		}
		R, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "err\n")
		}
		resultat[result[L][R]]++
	}

	i := 0
	for _, v := range resultat {
		if i == 0 {
			fmt.Printf("%d", v)
		} else {
			fmt.Printf(" %d", v)
		}
		i++
	}
}

