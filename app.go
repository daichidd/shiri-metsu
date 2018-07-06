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
			"|　 ●　● | 　|  ",
			"\\　　 □  /　<   ",
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
	f := true
	if sNum > fNum {
		mostLong = sNum
		f = false
	}

	lineEmpties := ""
	emptyLong := 0

	if f == false {
		emptyLong = mostLong - fNum
	} else {
		emptyLong = mostLong - sNum
	}

	for i := 0; i < emptyLong; i++ {
		lineEmpties += "　"
	}

	topLine := ""
	bottomLine := ""
	empty := ""
	for i := 0; i < mostLong; i++ {
		topLine += "￣"
		bottomLine += "＿"
		empty += "　"
	}

	shiri[1] += topLine
	shiri[1] += "＼"

	closeLine := " |"

	fLine := *first
	sLine := *second
	if f == false {
		fLine += lineEmpties
	} else {
		sLine += lineEmpties
	}

	fLine += closeLine
	sLine += closeLine
	shiri[2] += fLine

	shiri[3] += sLine
	shiri[4] += bottomLine + "／"

	for _, str := range shiri {
		fmt.Println(str)
	}
}
