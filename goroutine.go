package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
)

func fetchSiteContent(channel chan string, uri string) {
    resp, err := http.Get(uri)

    if err != nil {
        log.Fatal(err)
        channel <- "Get site fail"
    } else {
        // 相当于 Ruby 里的 ensure
        defer resp.Body.Close()

        // 读取出内容
        _, err := ioutil.ReadAll(resp.Body)

        if err != nil {
            log.Fatal(err)
            channel <- "Ready reponse's body fail"
        }

        // 不返回内容，只需要看看是否完成了即可
        channel <- ("fetch done: " + uri)
    }
}

func main() {
    site_channel := make(chan string, 2)

    go fetchSiteContent(site_channel, "http://www.justqyx.me")
    go fetchSiteContent(site_channel, "http://www.baidu.com")

    fmt.Println(<-site_channel)
    fmt.Println(<-site_channel)
}
