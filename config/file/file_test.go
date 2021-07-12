package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dayu-go/gkit/config"
)

const (
	_testJSON = `
{
    "test":{
        "settings":{
            "int_key":1000,
            "float_key":1000.1,
            "duration_key":10000,
            "string_key":"string_value"
        },
        "server":{
            "addr":"127.0.0.1",
            "port":8000
        }
    },
    "foo":[
        {
            "name":"nihao",
            "age":18
        },
        {
            "name":"nihao",
            "age":18
        }
    ]
}`
)

// go test -v *.go -test.run=^TestLoadFile$
func TestLoadFile(t *testing.T) {
	var (
		path     = filepath.Join(os.TempDir(), "test_config")
		filename = filepath.Join(path, "test.json")
		data     = []byte(_testJSON)
	)
	defer os.Remove(path)
	if err := os.MkdirAll(path, 0700); err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile(filename, data, 0666); err != nil {
		t.Error(err)
	}

	kv, err := loadFile(filename)
	t.Logf("kv:%+v, err:%v", kv, err)
}

// go test -v *.go -test.run=^TestLoadDir$
func TestLoadDir(t *testing.T) {
	var (
		path     = filepath.Join(os.TempDir(), "test_config")
		filename = filepath.Join(path, "test.json")
		data     = []byte(_testJSON)
	)
	defer os.Remove(path)
	if err := os.MkdirAll(path, 0700); err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile(filename, data, 0666); err != nil {
		t.Error(err)
	}

	kvs, err := loadDir(path)
	if err != nil {
		t.Fatalf("err:%s", err.Error())
	}
	for _, v := range kvs {
		t.Logf("key:%s, value:%s, format:%s\n", v.Key, string(v.Value), v.Format)
	}
}

// go test -v *.go -test.run=^TestLoad$
func TestLoad(t *testing.T) {
	var (
		path     = filepath.Join(os.TempDir(), "test_config")
		filename = filepath.Join(path, "test.json")
		data     = []byte(_testJSON)
	)
	defer os.Remove(path)
	if err := os.MkdirAll(path, 0700); err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile(filename, data, 0666); err != nil {
		t.Error(err)
	}

	f := file{path}
	kvs, err := f.Load()
	if err != nil {
		t.Fatalf("err:%s", err.Error())
	}
	for _, v := range kvs {
		t.Logf("key:%s, value:%s, format:%s\n", v.Key, string(v.Value), v.Format)
	}
}

// go test -v *.go -test.run=^TestConfig$
func TestConfig(t *testing.T) {
	var (
		path     = filepath.Join(os.TempDir(), "test_config")
		filename = filepath.Join(path, "test.json")
		data     = []byte(_testJSON)
	)
	defer os.Remove(path)
	if err := os.MkdirAll(path, 0700); err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile(filename, data, 0666); err != nil {
		t.Error(err)
	}

	c := config.New(config.WithSource(
		NewSource(filename),
	))
	testScan(t, c)
}

func testScan(t *testing.T, c config.Config) {
	type TestJSON struct {
		Test struct {
			Settings struct {
				IntKey      int     `json:"int_key"`
				FloatKey    float64 `json:"float_key"`
				DurationKey int     `json:"duration_key"`
				StringKey   string  `json:"string_key"`
			} `json:"settings"`
			Server struct {
				Addr string `json:"addr"`
				Port int    `json:"port"`
			} `json:"server"`
		} `json:"test"`
		Foo []struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		} `json:"foo"`
	}

	var conf TestJSON
	if err := c.Load(); err != nil {
		t.Error(err)
	}
	if err := c.Scan(&conf); err != nil {
		t.Error(err)
	}
}
