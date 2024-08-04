package main

import (
	"embed"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"webscale/app"
	"webscale/external/pixabay"

	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}

	arg := os.Args[1]
	router.DB_FLAVOR = "sqlite"

	pixabay.Init()

	if arg == "import" {
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "pixabay" {
		q := os.Args[2]
		images := pixabay.ImageSearch(q)
		var offset = 1
		offsetParamInt, _ := strconv.Atoi(os.Args[3])
		offset = offsetParamInt
		for i, p := range images {
			there := fmt.Sprintf("other/%s", p.ID)
			_, err := os.Stat(there)
			if err == nil {
				fmt.Println("SKIPPING", p.ID, p.URL)
				continue
			}
			fmt.Println(p.User)

			asBytes := pixabay.DownloadImage(p.URL)
			if asBytes == nil {
				continue
			}
			filename := fmt.Sprintf("data/%04d.jpg", i+offset)
			fmt.Println(len(asBytes))
			file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
			file.Write(asBytes)
			file.Close()
			app.Resize(filename)
			pixabay.WriteCredit(q, p.User, filename)

			file, _ = os.OpenFile(there, os.O_CREATE|os.O_WRONLY, 0644)
			file.WriteString("")
			file.Close()
		}
	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.Welcome
		r.Paths["core"] = app.Core
		//r.Paths["api"] = app.HandleApi
		//r.Paths["login"] = app.Login
		//r.Paths["register"] = app.Register
		//r.Paths["admin"] = app.Admin
		r.Paths["markup"] = router.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
	}
}
