package app

import (
	"os/exec"
)

/*
ffmpeg -f lavfi -i color=c=black:s=1280x720 -framerate 1 -i %04d.jpg  -t 00:03:01 black.mp4

*/

func Resize(filename string) {
	cmd := exec.Command(
		"convert",
		filename,
		"-resize", "1280x720^",
		"-gravity", "center",
		"-extent", "1280x720",
		"-background", "black",
		filename,
	)
	o, err := cmd.CombinedOutput()
	_ = o
	_ = err
	//fmt.Println(string(o), err)
}
