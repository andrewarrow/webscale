package pixabay

import (
	"fmt"
	"os/exec"
)

func WriteCredit(q, user, filename string) {
	/*
		convert data/0004.jpg -gravity South -pointsize 36 -stroke '#000C' -strokewidth 2 -annotate +0+20 'pixabay.com/users/geof08-36360517' -stroke none -fill white -annotate +0+20 'pixabay.com/users/geof08-36360517' output.png

	*/
	slug := fmt.Sprintf("%s pixabay.com/users/%s", q, user)
	cmd := exec.Command(
		"magick",
		filename,
		"-fill", "black",
		"-draw", "rectangle %[fx:(w/2)-400],%[fx:h-60] %[fx:(w/2)+400],%[fx:h-20]",
		filename,
	)
	o, err := cmd.CombinedOutput()
	cmd = exec.Command(
		"convert",
		filename,
		"-gravity", "South",
		"-fill", "green",
		"-font", "/System/Library/Fonts/Supplemental/Andale Mono.ttf",
		"-pointsize", "18",
		"-stroke", "#00AA00",
		"-strokewidth", "1",
		"-annotate", "+0+20", slug,
		filename,
	)
	o, err = cmd.CombinedOutput()
	_ = o
	_ = err
	//fmt.Println(string(o), err)
}
