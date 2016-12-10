package goscan

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Text() string {

	scan := bufio.NewScanner(os.Stdin)
	fmt.Scan(scan.Scan())
	return scan.Text()
}

func Number() int {

	scan := bufio.NewScanner(os.Stdin)
	fmt.Scan(scan.Scan())
	num, _ := strconv.Atoi(scan.Text())
	return num
}
