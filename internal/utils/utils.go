package utils

import (
	"fmt"
	"net/url"
	"strings"
)

type FlagURL struct {
	URLs []*url.URL
}

func (f *FlagURL) String() string {
	return fmt.Sprint(f.URLs)
}
func (f *FlagURL) Set(s string) error {
	urls := strings.Split(s, "")
	fmt.Println("Set", urls)
	return nil
}
