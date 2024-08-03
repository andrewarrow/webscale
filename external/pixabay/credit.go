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
		"convert",
		filename,
		"-gravity", "South",
		"-pointsize", "36",
		"-stroke", "#000C",
		"-strokewidth", "2",
		"-annotate", "+0+20", slug,
		"-stroke", "none",
		"-fill", "white",
		"-annotate", "+0+20", slug,
		filename,
	)
	o, err := cmd.CombinedOutput()
	fmt.Println(string(o), err)
}
