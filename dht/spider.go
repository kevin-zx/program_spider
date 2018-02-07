package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/shiyanhui/dht"
	"net/http"
	_ "net/http/pprof"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"flag"
)

type file struct {
	Path   []interface{} `json:"path"`
	Length int           `json:"length"`
}

type bitTorrent struct {
	InfoHash string `json:"infohash"`
	Name     string `json:"name"`
	Files    []file `json:"files,omitempty"`
	Length   int    `json:"length,omitempty"`
}



func main() {
	// 原始爬虫中是1024
	requestQueueSize := flag.Int("r",4,"requestQueueSize")
	// 原始爬虫中是256
	workQueueSize := flag.Int("w",4,"workQueueSize")
	flag.Parse()
	var mu = mysqlutil.MysqlUtil{}
	err := mu.InitMySqlUtil("111.231.24.24",3306, "remote","Iknowthat@@!221","crawler_data")
	if err != nil{
		panic(err)
	}

	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	w := dht.NewWire(65536, *requestQueueSize, *workQueueSize)
	go func() {
		for resp := range w.Response() {
			metadata, err := dht.Decode(resp.MetadataInfo)
			if err != nil {
				continue
			}
			info := metadata.(map[string]interface{})

			if _, ok := info["name"]; !ok {
				continue
			}

			bt := bitTorrent{
				InfoHash: hex.EncodeToString(resp.InfoHash),
				Name:     info["name"].(string),
			}

			if v, ok := info["files"]; ok {
				files := v.([]interface{})
				bt.Files = make([]file, len(files))

				for i, item := range files {
					f := item.(map[string]interface{})
					bt.Files[i] = file{
						Path:   f["path"].([]interface{}),
						Length: f["length"].(int),
					}
				}
			} else if _, ok := info["length"]; ok {
				bt.Length = info["length"].(int)
			}

			data, err := json.Marshal(bt)
			data2, err := json.Marshal(bt.Files)
			if err == nil {
			//magnet:?xt=urn:btih:8d8b60a4b86c8e2cdd6b07f53f219249f23cf1eb
				err := mu.Insert("INSERT IGNORE INTO thunder_data (`link`,`name`,`length`,`file_count`,`file_data`) VALUE (?,?,?,?,?) ",
					fmt.Sprintf("magnet:?xt=urn:btih:%s", bt.InfoHash),bt.Name,bt.Length,len(bt.Files),data2)
				if err != nil {
					println(err.Error())
				}
				fmt.Printf("%s\n\n", data)
			}
		}
	}()
	go w.Run()

	config := dht.NewCrawlConfig()
	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		w.Request([]byte(infoHash), ip, port)
	}
	d := dht.New(config)

	d.Run()
}
