package file

import (
	"errors"
	"os"
)

type File struct {
}

func (f *File) CreateFile(nameFile, path, content string) error {
	if nameFile == "" {
		return errors.New("nameFile is empty")
	}
	err := os.MkdirAll(path, 0777)
	if path == "" {
		err = errors.New("path is empty")
	}
	if err != nil {
		return err
	}
	_, err = os.Create(path + "/" + nameFile)
	if content == "" {
		err = errors.New("content is empty")
	}
	if err != nil {
		return err
	}
	f.UpdateFile(nameFile, path, content)
	return nil
}

func (f *File) ReadFile(nameFile, path string) (string, error) {
	file, err := os.Open(path + "/" + nameFile)
	if path == "" {	
		err = errors.New("path is empty")
	}
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func (f *File) UpdateFile(nameFile, path, content string) error {
	file, err := os.OpenFile(path+"/"+nameFile, os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Truncate(0)
	_, err = file.WriteString(content)
	if content == "" {
		err = errors.New("content is empty")
	}
	if err != nil {
		return err
	}
	return nil
}

func (f *File) DeleteFile(nameFile, path string) error {
	err := os.Remove(path + "/" + nameFile)
	if path == "" {
		err = errors.New("path is empty")
	}
	if err != nil {
		return err
	}
	return nil
}

/*
go test ./file/ -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

go test ./file/ -run=TestReadFile -v
*/
