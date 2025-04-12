package assets

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadFontEx(contents *embed.FS, fileName string,
	fontSize int32, codePoints []rune) (font rl.Font, err error) {
	font = rl.Font{}

	fontFile, err := contents.Open(fileName)
	if err != nil {
		return
	}
	defer fontFile.Close()

	fileData, err := loadFileData(&fontFile)
	if err != nil {
		return
	}

	ext, err := getFileExtension(fileName)
	if err != nil {
		return
	}

	if len(fileData) != 0 {
		font = rl.LoadFontFromMemory(ext, fileData, fontSize, codePoints)
	}

	return
}

func LoadImage(contents *embed.FS, fileName string) (img *rl.Image, err error) {
	var fType string

	imageFile, err := contents.Open(fileName)
	if err != nil {
		err = fmt.Errorf("Failed to open embed file '%s'", fileName)
		rl.TraceLog(rl.LogError, err.Error())
		return
	}
	defer imageFile.Close()

	rl.SetTraceLogLevel(rl.LogWarning)
	fileData, err := loadFileData(&imageFile)
	if err != nil {
		err = fmt.Errorf("IMAGE: Failed to read data: %v", err)
		rl.TraceLog(rl.LogWarning, err.Error())
		return
	}

	if len(fileData) != 0 {
		fType, err = getFileExtension(fileName)
		if err != nil {
			return
		}
		rl.TraceLog(rl.LogWarning, "File type: "+fType)
		img = rl.LoadImageFromMemory(fType, fileData, int32(len(fileData)))
	} else {
		err = fmt.Errorf("IMAGE: Invalid file data '%s'", fileName)
		rl.TraceLog(rl.LogWarning, err.Error())
	}

	return
}

func loadFileData(iFile *fs.File) (fileData []byte, err error) {
	fileData = make([]byte, 0)

	// I think it is not good idea but whatever
	fileData, err = io.ReadAll(*iFile)

	return
}

func getFileExtension(fileName string) (ext string, err error) {
	dotInd := strings.LastIndex(fileName, ".")
	if dotInd == -1 {
		err = fmt.Errorf("MissingFileExtension")
		rl.SetTraceLogLevel(rl.LogWarning)
		rl.TraceLog(rl.LogWarning,
			fmt.Sprintf("Missing file extension '%s'", fileName),
		)
		return
	}
	ext = fileName[dotInd:]
	return
}
