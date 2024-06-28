package testfiles

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/assignment-amori/pkg/utils"
	"github.com/assignment-amori/pkg/uuid"
)

type module struct {
	baseProjectPath string
	currFolderPath  string
	testFolderPath  string
	tempFolderPath  string
	uuid            uuid.UUID
}

func New() *module {
	wd, err := os.Getwd()
	if err != nil {
		return nil
	}

	uuid := uuid.New()
	baseProjectPath := wd[0 : strings.LastIndex(wd, projectName)+len(projectName)]
	currFolderPath := wd[strings.LastIndex(wd, projectName)+len(projectName):]
	tempFolderPath := filepath.Dir(utils.GenerateUniqueFilePath(sampleBasePath+"temp/", uuid))

	return &module{
		baseProjectPath: baseProjectPath,
		currFolderPath:  currFolderPath,
		testFolderPath:  path.Join(baseProjectPath, sampleBasePath, currFolderPath),
		tempFolderPath:  path.Join(baseProjectPath, tempFolderPath),
		uuid:            uuid,
	}
}

func (m *module) GetContent(filename string) string {
	b, err := ioutil.ReadFile(m.GetTestFilePath(filename))
	if err != nil {
		slog.Error(err.Error())
	}

	dst := &bytes.Buffer{}
	err = json.Compact(dst, b)
	if err != nil {
		slog.Error(err.Error())
	}

	return dst.String()
}

func (m *module) GetBinaryContent(filename string) (data []byte) {
	f, err := os.Open(m.GetTestFilePath(filename))
	if err != nil {
		slog.Error(err.Error())
		return
	}

	reader := bufio.NewReader(f)
	data, err = ioutil.ReadAll(reader)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	return
}

func (m *module) GetFilesPath(dir string) (filenames []string) {
	files, err := ioutil.ReadDir(m.GetTestFilePath(dir))
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filenames = append(filenames, path.Join(dir, file.Name()))
	}

	return filenames
}

func (m *module) GetTestFilePath(filename string) string {
	return path.Join(m.testFolderPath, filename)
}

func (m *module) GetTempFilePath(filename string) string {
	return path.Join(m.tempFolderPath, filename)
}

func (m *module) GetTempFolderPath() string {
	return m.tempFolderPath
}

func (m *module) CopySampleToTemp() {
	slog.Info(m.tempFolderPath)

	os.MkdirAll(m.tempFolderPath, os.ModePerm)
	err := utils.Copy(m.testFolderPath, m.tempFolderPath)
	if err != nil {
		slog.Error(err.Error())
	}
}

func (m *module) Clean() {
	os.RemoveAll(m.tempFolderPath)
}
