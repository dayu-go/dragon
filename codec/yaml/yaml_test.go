package yaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const (
	_testYAML = `app:
    http_addr: "0.0.0.0:9999"
    log_level: debug
    log_path: stdout
    env: test
service:
    domain:
        api: "http://47.102.215.71:9999"
        share: "http://ximei-h5-test.hubonews.com"
        article_share: "https://ximei-h5-test.hubonews.com"
    umeng:
        app_key: "25474633"
        app_secret: "dc40c4092074d8766ffae68e3932f1ec"
        app_code: "f877d03e329649c3b6ef39e1a4d9da89"
        android_app_key: "5ba8e0f2f1f556541600042a"
        ios_app_key: "5bc18affb465f52b5e000377"
    recommend_white_list: ["5122039","9113752"]
    token_expire_time: 31536000
    user_center_id: "http://10.0.3.4:7777/user-center/id"
storage:
    db:
        ximei:
            dsn: "sentiment:P6OWkX2uQUtdfjMe@tcp(rm-uf6h1bu41gmx48aq8rw.mysql.rds.aliyuncs.com:3306)/translate_summary_business"
        user_center:
            dsn: "tiger:Usercenter@Tiger@tcp(rm-uf6fqqm6x75xv7nv1.mysql.rds.aliyuncs.com:3306)/user-center2?parseTime=true&loc=Local"
    redis:
        ximei:
            host: "r-uf62e134a0c68014.redis.rds.aliyuncs.com"
            port: 6379
            password: "G3EsmjZ6ULW9p4Zw"
            db: 0`
)

func TestCodec_Marshal(t *testing.T) {
	value := map[string]string{"v": "hi"}
	got, err := (code{}).Marshal(value)
	if err != nil {
		t.Fatalf("should not return err")
	}
	if string(got) != "v: hi\n" {
		t.Fatalf("want \"v: hi\n\" return \"%s\"", string(got))
	}
}

// go test -v *.go -test.run=TestYamlUnmarshal
func TestYamlUnmarshal(t *testing.T) {
	var (
		path     = filepath.Join(os.TempDir(), "test_config")
		filename = filepath.Join(path, "test.yaml")
		data     = []byte(_testYAML)
	)
	t.Logf("create yaml file:%s", path)
	if err := os.MkdirAll(path, 0700); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filename, data, 0666); err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	data, err = ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	res := make(map[string]interface{})
	err = (code{}).Unmarshal(data, &res)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range res {
		fmt.Printf("k:%s v:%+v\n", k, v)
	}
}
