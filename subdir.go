package subdir

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Dir string

func (dir Dir) String() string {
	return string(dir)
}

func New(path string) (Dir, error) {
	fullpath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(fullpath, os.ModeDir)
	if err != nil {
		return "", err
	}
	return Dir(fullpath), nil
}

func (dir Dir) getPath(path string) string {
	return filepath.Join(dir.String(), path)
}

func (dir Dir) SubDir(path string) (Dir, error) {
	fullpath := dir.getPath(path)
	err := os.MkdirAll(fullpath, os.ModeDir)
	if err != nil {
		return "", err
	}
	return Dir(fullpath), nil
}

func (dir Dir) Open(name string) (*os.File, error) {
	return os.Open(dir.getPath(name))
}

func (dir Dir) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(dir.getPath(name), flag, perm)
}

func (dir Dir) Create(name string) (*os.File, error) {
	return os.Create(dir.getPath(name))
}

func (dir Dir) Remove(name string) error {
	return os.Remove(dir.getPath(name))
}

func (dir Dir) RemoveAll(name string) error {
	return os.RemoveAll(dir.getPath(name))
}

func (dir Dir) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(dir.getPath(path), perm)
}

func (dir Dir) List() ([]os.DirEntry, error) {
	return os.ReadDir(dir.String())
}

func (dir Dir) Walk(fn fs.WalkDirFunc) error {
	return filepath.WalkDir(dir.String(), fn)
}

func (dir Dir) DeleteSelf() error {
	return os.RemoveAll(dir.String())
}

func (dir Dir) Stat(name string) (os.FileInfo, error) {
	return os.Stat(dir.getPath(name))
}
