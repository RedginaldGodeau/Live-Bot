package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var video = []string{".mp4", ".webpm", ".mov", ".avi"}
var image = []string{".png", ".jpg", ".webp"}

type contentFile struct {
	FileName string
	FileType string
	Duration int
}

func getVideoDuration(filePath string) (int, error) {
	cmd := exec.Command("ffprobe", "-i", filePath, "-show_entries", "format=duration", "-v", "quiet", "-of", "csv=p=0")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, err
	}
	durationStr := strings.TrimSpace(out.String())
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, err
	}
	s := fmt.Sprintf("%.0f", duration)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func UploadDiscordFile(fileURL string) (*contentFile, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	extension := strings.Split(filepath.Ext(fileURL), "?")
	fileName := uuid.New().String() + extension[0]
	filePath := os.Getenv("UPLOAD_DIRECTORY") + "/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	var typeFile = "picture"
	var duration = 15
	if slices.Contains(video, extension[0]) {
		typeFile = "video"
		dur, err := getVideoDuration(filePath)
		if err != nil {
			return nil, err
		}
		duration = dur
	}

	return &contentFile{
		FileName: fileName,
		FileType: typeFile,
		Duration: duration,
	}, nil
}
