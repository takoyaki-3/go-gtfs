package gtfs

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(src, dest string) (error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	if err = os.MkdirAll(dest,777);err!=nil{
		return err
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func LoadFromUnzipGTFS(fileName string, filter map[string]bool) (*GTFS, error) {
	// 一時フォルダに展開
	dirName := "./tmpGTFS"
	err := Unzip(fileName, dirName)
	if err != nil {
		return &GTFS{}, err
	}
	g, err := Load(dirName, filter)
	if err != nil {
		return g, err
	}
	err = os.RemoveAll(dirName)

	return g, err
}
