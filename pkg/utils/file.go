package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/pkg/uuid"
	cp "github.com/otiai10/copy"
)

// OpenFileWriter - create or read file for writer
func OpenFileWriter(file string) (*os.File, error) {
	if file == "" {
		return nil, nil
	}

	err := os.MkdirAll(filepath.Dir(file), 0755)
	if err != nil && err != os.ErrExist {
		return nil, err
	}

	return os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}

// RenameDir rename directory to new name
// args
//
//	dir: directory
//	location: location after dir
//	old: name of the old directory
//	new: new name
//
// returns
//
//	error that happens while renaming
func RenameDir(dir, location, old, new string) error {
	keyPrefix := old
	return os.Rename(fmt.Sprint(dir, location, keyPrefix), fmt.Sprint(dir, location, new))
}

// Exists checks if directory is exists
// args
//
//	dir: directory to be checked
//
// returns
//
//	ok: true if directory is exist
//	error: error while checking
func Exists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return true, err
	}

	return true, nil
}

// Copy copies directory from src to dest
// args
//
//	src: source directory
//	dest: destination directory
func Copy(src, dest string) error {
	return cp.Copy(src, dest)
}

// ReadFile read file specified by dir
// args
//
//	dir: location of file that want to be read
//
// returns
//
//	string result of reading file
//	error that happens while reading file
func ReadFile(dir string) (string, error) {
	b, err := ioutil.ReadFile(dir)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// WriteFile writes txt to file in dir
// args
//
//	dir: location of file that want to be written
//	txt: text that want to be written to file
//
// returns
//
//	error that happens while reading file
func WriteFile(dir, txt string) error {
	return os.WriteFile(dir, []byte(txt), 0644)
}

func GetFileNameWithoutExt(filePath string) string {
	fileName := filepath.Base(filePath)
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func GenerateUniqueFilePath(baseFilePath string, uuidmod uuid.UUID) string {
	/* Generate UUID to store file, but only with the last part to make it short.
	To prevent for the folder path being to long.*/
	if uuidmod == nil {
		uuidmod = uuid.New()
	}

	var time = time.Now()
	newUuid := uuidmod.GenUUID()
	uniqueFolderId := strings.Split(newUuid, "-")[4]
	uniqueFileId := strings.Split(newUuid, "-")[3]

	// Also add current date as folder to avoid collision with the UUID
	return filepath.Clean(filepath.Join(baseFilePath, time.Format("20060102"), uniqueFolderId, uniqueFileId))
}

func GetFileContentType(filePath string) (typ string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return typ, err
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return typ, err
	}

	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	return http.DetectContentType(buffer), nil
}

func GetFileSize(filePath string) (int64, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		return constant.DefaultInt64, err
	}

	return file.Size(), nil
}

// GetAllFilesInFolder get all file paths from a folder directory
// args
//
//	folderPath: location of folder
//
// returns
//
// filePaths: list of file paths in the folder
// error: that happens while reading file
func GetAllFilesInFolder(folderPath string) (filePaths []string, err error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return filePaths, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePaths = append(filePaths, path.Join(folderPath, file.Name()))
	}
	return filePaths, err
}

func inTrustedRoot(path string, trustedRoot string) error {
	for path != "/" {
		if path == trustedRoot {
			return nil
		}
		path = filepath.Dir(path)
	}
	return errors.New("path is outside of trusted root")
}

func SanitizePath(path string) (sanitized string, err error) {
	// TODO: Make trusted root configurable.
	trustedRoot, err := os.Getwd()
	if err != nil {
		return sanitized, err
	}

	// Get absolute path.
	absPath, err := filepath.Abs(path)
	if err != nil {
		return sanitized, err
	}

	// Check if path resolves to symlinks.
	sanitized, err = filepath.EvalSymlinks(absPath)
	if err != nil {
		return sanitized, err
	}

	// Check if path is inside trusted root.
	err = inTrustedRoot(sanitized, trustedRoot)
	if err != nil {
		return sanitized, err
	}

	return sanitized, nil
}
