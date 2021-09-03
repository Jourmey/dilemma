package youget

import "os/exec"

var BinPath string

type YouGet struct {
	bin string
}

func NewYouGet() *YouGet {
	y := new(YouGet)
	if BinPath != "" {
		y.bin = BinPath
	} else {
		y.bin = "you-get"
	}
	return y
}

func (f *YouGet) Info(url string) (interface{}, error) {
	data, err := f.cmd("-i", url, "--json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *YouGet) cmd(arg ...string) ([]byte, error) {
	cmd := exec.Command(f.bin, arg...)
	return cmd.Output()
}
