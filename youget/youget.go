package youget

import (
	"encoding/json"
	"os/exec"
)

var BinPath string

type YouGet struct {
	bin string
}

type (
	Info struct {
		Url     string            `json:"url"`
		Title   string            `json:"title"`
		Site    string            `json:"site"`
		Streams map[string]Stream `json:"streams"`
		//Extra   interface{}       `json:"extra"`
	}
	Stream struct {
		Container string `json:"container"`
		Quality   string `json:"quality"`
		Size      int    `json:"size"`
		//src []string
	}
)

func NewYouGet() *YouGet {
	y := new(YouGet)
	if BinPath != "" {
		y.bin = BinPath
	} else {
		y.bin = "you-get"
	}
	return y
}

func (f *YouGet) Info(url string) (*Info, error) {
	i := new(Info)
	err := f.cmd2(i, "-i", url, "--json")
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (f *YouGet) Download(url string, format string, outputDir string) (string, error) {
	data, err := f.cmd("--format="+format, "--output-dir="+outputDir, url)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (f *YouGet) cmd(arg ...string) ([]byte, error) {
	cmd := exec.Command(f.bin, arg...)
	return cmd.Output()
}

func (f *YouGet) cmd2(v interface{}, arg ...string) error {
	cmd := exec.Command(f.bin, arg...)
	data, err := cmd.Output()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
