package leetcode535

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	dataBase map[int]string
	id       int
}

func Constructor() Codec {
	return Codec{dataBase: make(map[int]string), id: 0}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.id++
	this.dataBase[this.id] = longUrl
	return fmt.Sprintf("http://tinyurl.com/%d", this.id)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	i := strings.LastIndexByte(shortUrl, '/')
	id, _ := strconv.Atoi(shortUrl[i+1:])
	return this.dataBase[id]
}
