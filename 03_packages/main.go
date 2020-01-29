package main

import (
	"fmt"

	"math"

	"github.com/farhan9689/go_crash_course/03_packages/strutil"
)

func main() {

	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.4))
	fmt.Println(math.Sqrt(16))

	fmt.Println(strutil.Reverse("farhan"))

}
