package main

import (
	"strconv"
	"strings"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type Codec struct {
	data map[int]string
	url  map[string]int
}

func Constructor() Codec {
	return Codec{make(map[int]string), make(map[string]int)}
}

const k1, k2 = 1117, 1e9 + 7

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if key, ok := this.url[longUrl]; ok {
		return "http://tinyurl.com/" + strconv.Itoa(key)
	}
	key, base := 0, 1
	for _, c := range longUrl {
		key = (key + int(c)*base) % k2
		base = base * k1 % k2
	}
	for this.data[key] != "" {
		key = (key + 1) % k2
	}
	this.data[key] = longUrl
	this.url[longUrl] = key
	return "http://tinyurl.com/" + strconv.Itoa(key)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	i := strings.LastIndex(shortUrl, "/")
	key, _ := strconv.Atoi(shortUrl[i+1:])
	return this.data[key]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
//leetcode submit region end(Prohibit modification and deletion)
