package main

import (
	"flag"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	var (
		first  = flag.String("1", "", "first line")
		second = flag.String("2", "", "second line")

		shiri = [8]string{
			"／￣￣￣＼",
			"|　ー　ー|　  ／",
			"|　 ●　● | 　|　",
			"\\　　 □  /　<  ",
			" ＼　　 イ 　 ＼",
			"  ／　　　＼",
			"/　|　　　 \\\\　支離滅裂な",
			"|　|　　　 ||　思考・発言",
		}
	)

	flag.Parse()

	fNum := utf8.RuneCountInString(*first)
	sNum := utf8.RuneCountInString(*second)

	if fNum == 0 {
		fmt.Println("flag is not exist")
		os.Exit(0)
	}

	mostLong := fNum
	if sNum > fNum {
		mostLong = sNum
	}

	fmt.Println(mostLong)

	topLine := "￣"
	bottomLine := "＿"
	empty := "　"
	for i := 0; i < mostLong-1; i++ {
		topLine += "￣"
		bottomLine += "＿"
		empty += "　"
	}

	shiri[1] += topLine
	shiri[1] += "＼"

	shiri[2] += *first + "　|"
	if sNum != 0 {
		shiri[3] += *second + "　|"
	} else {
		shiri[3] += empty + "　|"
	}

	shiri[4] += bottomLine + "／"

	for _, str := range shiri {
		fmt.Println(str)
	}
}
