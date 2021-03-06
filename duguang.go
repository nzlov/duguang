package duguang

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"errors"
)

var (
	Err_NoQuantity = errors.New("数量用尽")
)

type Duguang struct {
	appcode string
	timeout time.Duration
}

func NewAppCodeDuguang(appcode string) *Duguang {
	return NewDuguang(appcode, time.Second*10)
}
func NewDuguang(appcode string, timeout time.Duration) *Duguang {
	return &Duguang{
		appcode: appcode,
		timeout: timeout,
	}
}

func (d *Duguang) SetAppcode(appcode string) {
	d.appcode = appcode
}

func (d *Duguang) req(host string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", host, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "APPCODE "+d.appcode)

	c := &http.Client{
		Timeout: d.timeout,
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	switch resp.StatusCode {
	case 403:
		return nil, Err_NoQuantity
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 通用文字识别－高精版接口文档
func (d *Duguang) Advanced(a Advanced) (*Result, error) {
	if a.Img == "" && a.URL == "" {
		return nil, ErrNoImg
	}
	if a.Img != "" && a.URL != "" {
		return nil, ErrImgRepeat
	}
	if len(a.Img) > IMGSIZE {
		return nil, ErrSize
	}

	data, err := json.Marshal(&a)
	if err != nil {
		return nil, err
	}

	data, err = d.req("https://ocrapi-advanced.taobao.com/ocrservice/advanced", data)
	if err != nil {
		return nil, err
	}

	r := &Result{}
	err = json.Unmarshal(data, r)
	return r, err
}

// 文档结构化还原识别接口文档
func (d *Duguang) Document(a Document) (*Result, error) {
	if a.Img == "" && a.URL == "" {
		return nil, ErrNoImg
	}
	if a.Img != "" && a.URL != "" {
		return nil, ErrImgRepeat
	}
	if len(a.Img) > IMGSIZE {
		return nil, ErrSize
	}

	data, err := json.Marshal(&a)
	if err != nil {
		return nil, err
	}

	data, err = d.req("https://ocrapi-document-structure.taobao.com/ocrservice/documentStructure", data)
	if err != nil {
		return nil, err
	}

	r := &Result{}
	err = json.Unmarshal(data, r)
	return r, err
}
