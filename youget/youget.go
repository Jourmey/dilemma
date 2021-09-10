package youget

import (
	"context"
	"dilemma/tool"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/trace"
	"os/exec"
)

var BinPath string

type YouGet struct {
	logx.Logger
	ctx context.Context
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

func NewYouGet(ctx context.Context) *YouGet {
	y := new(YouGet)
	ctx, _ = trace.StartClientSpan(ctx, "dilemma", "youget")
	y.ctx = ctx
	y.Logger = logx.WithContext(y.ctx)

	if BinPath != "" {
		y.bin = BinPath
	} else {
		y.bin = "you-get"
	}
	return y
}

func (f *YouGet) Info(url string) (*Info, error) {
	i := new(Info)
	err := f.cmd2(i, url, "--json")
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (f *YouGet) Download(url string, format string, outputDir string) (string, error) {
	data, err := f.cmd("--format="+format, "--output-dir="+outputDir, url)
	if err != nil {
		return string(data), err
	}
	return string(data), nil
}

func (f *YouGet) cmd(arg ...string) ([]byte, error) {
	f.Logger.Info(tool.Start, f.bin, arg)
	cmd := exec.Command(f.bin, arg...)
	b, err := cmd.Output()
	if err != nil {
		f.Logger.Error(tool.Failed)
	} else {
		f.Logger.Info(tool.Success)
	}
	return b, err
}

func (f *YouGet) cmd2(v interface{}, arg ...string) error {
	data, err := f.cmd(arg...)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
