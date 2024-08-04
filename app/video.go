package app

import (
	"os/exec"
)

/*

ffmpeg -framerate 1 -i %04d.jpg -t 194 -pix_fmt yuv420p black.mp4

ffmpeg -framerate 1/2 -pattern_type glob -i '*.jpg' -vf "eq=brightness=0:contrast=1:saturation=1"  -t 87 -pix_fmt yuv420p ../bright.mp4


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
