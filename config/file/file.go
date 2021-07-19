package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dayu-go/gkit/config"
)

type file struct {
	path string
}

// NewSource new a file source.
func NewSource(path string) config.Source {
	return &file{path: path}
}

func (f *file) Load() (kvs []*config.KeyValue, err error) {
	fi, err := os.Stat(f.path)
	if err != nil {
		return
	}
	if fi.IsDir() {
		return loadDir(f.path)
	}
	kv, err := loadFile(f.path)
	if err != nil {
		return
	}
	return []*config.KeyValue{kv}, nil
}

func loadFile(path string) (kv *config.KeyValue, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	info, err := file.Stat()
	if err != nil {
		return
	}
	kv = &config.KeyValue{
		Key:    info.Name(),
		Value:  data,
		Format: format(info.Name()),
	}
	return
}

func loadDir(path string) (kvs []*config.KeyValue, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() || strings.HasPrefix(file.Name(), ".") {
			continue
		}
		kv, err := loadFile(filepath.Join(path, file.Name()))
		if err != nil {
			return nil, err
		}
		kvs = append(kvs, kv)
	}
	return
}

func format(name string) string {
	if p := strings.Split(name, "."); len(p) > 1 {
		return p[len(p)-1]
	}
	return ""
}
