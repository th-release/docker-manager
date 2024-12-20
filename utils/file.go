package utils

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetFilePath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(dir, "public", file)
}

func CreateDirectoryIfNotExist(filePath string) error {
	dir := filepath.Dir(GetFilePath("") + filePath) // 파일 경로로부터 디렉토리 경로 추출

	// 디렉토리가 이미 존재하는지 확인
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 디렉토리가 존재하지 않으면 생성
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			// 디렉토리 생성 중 오류 발생
			return err
		}
	}

	return nil
}

func CreateJsonIfNotExist(filePath string, initialContent interface{}) error {
	_, err := os.Open(filePath)
	if err != nil {
		err := os.MkdirAll(GetFilePath(""), os.ModePerm)
		if err != nil {
			return fmt.Errorf("setup error: %v", err)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("setup error: %v", err)
		}
		defer file.Close()

		if initialContent != nil {
			data, err := json.Marshal(initialContent)
			if err != nil {
				return fmt.Errorf("error marshaling initial content: %v", err)
			}

			_, err = file.Write(data)
			if err != nil {
				return fmt.Errorf("error writing initial content to file: %v", err)
			}
		}
	}

	return nil
}

func CreateBuildContext(dirPath string) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// 디렉토리 탐색
	err := filepath.Walk(dirPath, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 파일 정보 가져오기
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// 상대 경로 설정
		relPath, err := filepath.Rel(dirPath, file)
		if err != nil {
			return err
		}
		header.Name = relPath

		// 파일 추가
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			defer data.Close()
			_, err = io.Copy(tw, data)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}

	return buf, nil
}
