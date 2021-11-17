package main

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
)

var (
	ErrorWrongFileName = errors.New("wrong file name")
)

type SearchTarget struct {
	Name   string
	Path   string
	Size   int64
	wg     *sync.WaitGroup
	logger *zap.SugaredLogger
}

type Checker interface {
	Check(info os.FileInfo, filePath string) bool
	WgAdd()
	WgDone()
	WgWait()
}

func (s *SearchTarget) Check(fi os.FileInfo, filePath string) bool {
	s.logger.Debug("Check method call")
	if s.Path == filePath {
		s.logger.Infof("Same dirs %s == %s", s.Path, filePath)
		return false
	}

	s.logger.Infof("Different dirs %s == %s", s.Path, filePath)

	return s.isCopy(fi)
}

func (s *SearchTarget) isCopy(fi os.FileInfo) bool {
	s.logger.Debug("isCopy method call")
	return s.Name == fi.Name() && s.Size == fi.Size()
}

func (s *SearchTarget) WgAdd() {
	s.logger.Debug("WgAdd method call")
	s.wg.Add(1)
}

func (s *SearchTarget) WgDone() {
	s.logger.Debug("WgDone method call")
	s.wg.Done()
}

func (s *SearchTarget) WgWait() {
	s.logger.Debug("WgWait method call")
	s.wg.Wait()
}

func NewSearchTarget(fileName string, logger *zap.SugaredLogger) (*SearchTarget, error) {
	logger.Debug("NewSearchTarget method call")
	if fileName == "" {
		logger.Error("wrong file name: ", fileName)
		return nil, ErrorWrongFileName
	}

	fi, err := os.Stat(fileName)
	if err != nil {
		logger.Error("file does not exist: ", err)
		return nil, err
	}

	path, err := filepath.Abs(fileName)
	if err != nil {
		logger.Error("cannot get file path: ", err)
		return nil, err
	}

	return &SearchTarget{
		Name:   fi.Name(),
		Path:   path,
		Size:   fi.Size(),
		wg:     new(sync.WaitGroup),
		logger: logger,
	}, nil
}
