package app

import (
	"os/exec"
)

/*

ffmpeg -framerate 1 -i %04d.jpg -t 194 -pix_fmt yuv420p black.mp4

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
