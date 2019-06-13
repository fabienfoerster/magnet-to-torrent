package magnettotorrent
import "fmt"
import "net/http"
import "strings"

func cachesURL(hash string) []string {
	hash = strings.ToUpper(hash)
	var urls []string 
	urls = append(urls,fmt.Sprintf("http://itorrents.org/torrent/%s.torrent",hash))
	//urls = append(urls,fmt.Sprintf("http://btcache.me/torrent/%s",hash))
	
	return urls
}

func isValidURL(url string) int {
	resp,err := http.Get(url)
	if err != nil {
		return 0
	}
	return resp.StatusCode
}

func GetTorrent(hash string) string {
	cachesURL := cachesURL(hash)
	for _,c := range cachesURL {
		if isValidURL(c) != 0 {
			return c
		}
	}
	return ""
}







