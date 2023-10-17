package thyme

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func track(waitTime int) {
	// Here's a description of some of the ffmpeg parameters being used:
	// -vframes -> Number of frames to output.
	// -q:v -> (alias: qscale) Quality of the output. 100 means as lossy as
	// possible (resulting in smaller filesize).
	// -vf scale=1280:-1 -> Rescale to 1280 width (and automatically keep aspect
	// ratio). Uses ffmpeg filtering.

	baseArgs := []string{
		"-f", "avfoundation",
		"-framerate", "1",
		"-i", "1:0",
		"-vframes", "1",
		"-q:v", "100",
		"-vf", "scale=1280:-1",
	}

	for {
		time.Sleep(time.Duration(waitTime) * time.Second)
		filename := "output" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpeg"

		args := append(baseArgs, filename)

		cmd := exec.Command("ffmpeg", args...)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func watch() {
	// TODO
}

func main() {
	cmd := flag.String("cmd", "track", "Command to execute. Can be either track or watch.")
	waitTime := flag.Int("waittime", 60, "How many seconds to wait between each capture.")
	flag.Parse()

	switch *cmd {
	case "track":
		track(*waitTime)
	case "watch":
		watch()
	default:
		fmt.Println("Unknown command:", *cmd)
	}
}
