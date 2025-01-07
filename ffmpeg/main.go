package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
)

type History struct {
	files []string
	folder string
}

func main() {
	dir := os.Args[1]
	dirPath := fmt.Sprintf("./%s", dir)

	videos, err := os.ReadDir(dirPath)

	if err != nil {
		log.Fatal(err)
	}
	videos = filter(videos, "mp4")

	filesList := []string{}

	for _, dirE := range videos {
		filesList = append(filesList, fmt.Sprintf("%s/%s", dirPath, dirE.Name()));
	}

	history := NewHistory(dirPath)
	filesNotShowed := getNotShowedFiles(filesList, history)	
	
	if len(filesNotShowed) > 0 {
		fileToShow := filesNotShowed[0];
		showFile(fileToShow)
		history.putInHistory(fileToShow)
	} else {
		history.reset()
		filesNotShowed := getNotShowedFiles(filesList, history);
		fileToShow := filesNotShowed[0];
		showFile(fileToShow)
		history.putInHistory(fileToShow)
	}
}

func getNotShowedFiles(filesList []string, history *History) []string {

	filesNotShowed := []string{}
	
	for _, file := range filesList {
		if !inSlice(file, history.files) {
			filesNotShowed = append(filesNotShowed, file)
		}
	}

	return filesNotShowed
}

func showFile(fileToShow string) {
	fmt.Println(fileToShow)
}


func filter(files []os.DirEntry, ext string) ([]os.DirEntry) {
	r := []os.DirEntry{}

	for _, dirE := range files {
		if !dirE.IsDir() && isNeedExt(dirE.Name(), ext) {
			r = append(r, dirE)
		}
	}

	return r
}

func isNeedExt(filename string, ext string) bool {
	r, _ := regexp.Compile(fmt.Sprintf(`\.%s$`, ext))

	return r.MatchString(filename);
}

func inSlice(filename string, list []string) bool {
	for _, b := range list {
        if b == filename {
            return true
        }
    }
    return false
}

func NewHistory(folder string) *History {

	filePath := fmt.Sprintf("%s/history.json", folder);

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), 0644)
	}

	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}
	files := []string{}

	json.Unmarshal(data, &files)

	return &History{files: files, folder: folder}
}

func (h *History) reset() {
	h.files = []string{}
	os.WriteFile(h.getFilePath(), []byte("[]"), 0644)
}

func (h *History) getFilePath() (string) {
	return fmt.Sprintf("%s/history.json", h.folder);
}

func (h *History) putInHistory(filePath string) {
	h.files = append(h.files, filePath)

	json, err := json.Marshal(h.files)

	if err != nil {
		log.Fatal(err)
	}
	
	os.WriteFile(h.getFilePath(), json, 0644)
}