package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

var (
	ErrorFileDoesNotExist = errors.New("file does not exist")
	ErrorFileIsDir        = errors.New("the passed file is a dir")
	ErrorPathDoesNotExist = errors.New("path does not exist")
	ErrorPathIsNotDir     = errors.New("the passed path is not a dir")
)

func ValidateDirPath(path string, logger *zap.SugaredLogger) error {
	logger.Debug("ValidateDirPath method call")
	logger.Infof("Validate dir: %s", path)

	p, err := os.Stat(path)
	if err != nil {
		logger.Errorw("path does not exist: ", err)
		return ErrorPathDoesNotExist
	}

	if !p.IsDir() {
		logger.Error("path is not a dir: ", path)
		return ErrorPathIsNotDir
	}

	return nil
}

func ValidateFilePath(fileName string, logger *zap.SugaredLogger) error {
	logger.Debug("ValidateFilePath method call")
	logger.Infof("Validate file: %s", fileName)
	f, err := os.Stat(fileName)
	if err != nil {
		logger.Error("file does not exist: ", err)

		return ErrorFileDoesNotExist
	}

	if f.IsDir() {
		logger.Error("file is a dir: ", fileName)
		return ErrorFileIsDir
	}

	return nil
}

func FindDuplicate(path string, fileName string, logger *zap.SugaredLogger) ([]string, error) {
	logger.Debug("FindDuplicate method call")
	if err := ValidateDirPath(path, logger); err != nil {
		return nil, err
	}
	if err := ValidateFilePath(fileName, logger); err != nil {
		return nil, err
	}

	var (
		copyFilePathCh = make(chan string)
		copyList       = make([]string, 0)
		checker        Checker
	)
	checker, err := NewSearchTarget(fileName, logger)
	if err != nil {
		return nil, fmt.Errorf("error of init search target struct: %w", err)
	}
	go func() {
		logger.Debug("Goroutine call")
		for {
			select {
			case copyPath, ok := <-copyFilePathCh:
				if ok {
					logger.Info("Copy path: ", copyPath)
					copyList = append(copyList, copyPath)
				} else {
					break
				}
			case <-time.After(10 * time.Second):
				fmt.Println("Timeout. Channel reader is dead")
				return
			}
		}
	}()

	WalkInDir(path, checker, copyFilePathCh, logger)

	checker.WgWait()
	close(copyFilePathCh)
	return copyList, nil
}

func WalkInDir(path string, c Checker, fileCh chan<- string, logger *zap.SugaredLogger) {
	logger.Debug("WalkInDir method call")
	c.WgAdd()
	defer c.WgDone()

	logger.Info("try read dir ", path)
	dirList, err := os.ReadDir(path)
	if err != nil {
		logger.Error("cannot read dir ", err)

		return
	}

	logger.Infof("Dir list %v", dirList)
	for _, dirFile := range dirList {
		filePath, err := filepath.Abs(filepath.Join(path, dirFile.Name()))
		if err != nil {
			logger.Error("cannot get file path ", err)
			continue
		}

		fi, err := dirFile.Info()
		if err != nil {
			logger.Error("cannot get file info ", err)
			continue
		}

		if c.Check(fi, filePath) {
			fileCh <- filePath
		}

		if dirFile.IsDir() {
			panic(fmt.Sprintf("Sub dir found %s", filePath))
			//go WalkInDir(filePath, c, fileCh, logger)
			//continue
		}
	}
}
