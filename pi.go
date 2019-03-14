package pi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	invalid = -1
)

var (
	once sync.Once
	g    Getter
)

func init() {
	once.Do(func() {
		g = NewGetter()
	})
}

func Get(start, numberOfDigits int) ([]byte, int, error) {
	return g.Get(start, numberOfDigits)
}

// getter implements Getter interface
type getter func(start, numberOfDigits int) string

// response is used to unmarshal http response from the REST API call
type response struct {
	Content string `json:"content"`
}

// NewGetter gets a new instance of pi Getter
func NewGetter() Getter {
	g := getter(func(start, numberOfDigits int) string {
		return fmt.Sprintf("https://api.pi.delivery/v1/pi?start=%d&numberOfDigits=%d",
			start, numberOfDigits)
	})

	return g
}

// Get fetches digits of pi and indicates location of decimal value in second returned value.
// Absence of decimal value returns -1.
func (g getter) Get(start, numberOfDigits int) ([]byte, int, error) {
	if start < 0 || numberOfDigits < 0 {
		return nil, invalid, fmt.Errorf("error in input args, cannot be negative")
	}

	url := g(start, numberOfDigits)

	resp, err := http.Get(url)
	if err != nil {
		return nil, invalid, fmt.Errorf("error fetching pi digits:%v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, invalid, fmt.Errorf("error reading pi api get response body:%v", err)
	}

	response := new(response)
	if err := json.Unmarshal(b, response); err != nil {
		return nil, invalid, err
	}

	out := make([]byte, 0, numberOfDigits)
	decimal := -1
	for i := 0; i < len(response.Content); i++ {
		switch v := response.Content[i]; v {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			out = append(out, v-'0')
		case '.':
			decimal = i
		default:
			return nil, invalid, fmt.Errorf("error in received data, seems corrupted data")
		}
	}

	return out, decimal, nil
}
