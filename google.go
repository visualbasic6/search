package main
 
import (
        "bufio"
        "fmt"
        "io/ioutil"
        "net/http"
        "net/url"
        "os"
        "strings"
)
 
func main() {
        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print("Enter search query: ")
                query, _ := reader.ReadString('\n')
                query = url.QueryEscape(strings.TrimSpace(query))
                resp, err := http.Get("https://www.google.com/alerts/preview?params=[null,[null,null,null,[null,%22" + query + "%22,%22com%22,[null,%22en%22,%22US%22],null,null,null,0,1],null,3,[[null,1,%22user@example.com%22,[],1,%22en-US%22,null,null,null,null,null,%220%22,null,null,%22%22]],null,null,[2,1]],0]")
                if err != nil {
                        fmt.Println(err)
                        return
                }
                defer resp.Body.Close()
                response, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                        fmt.Println(err)
                        return
                }
                responseStr := string(response)
                urls := make(map[string]bool)
                startIndex := strings.Index(responseStr, "url=")
                for startIndex != -1 {
                        endIndex := strings.Index(responseStr[startIndex:], "&amp;")
                        if endIndex == -1 {
                                break
                        }
                        endIndex += startIndex
                        url := responseStr[startIndex+4 : endIndex]
                        if !urls[url] {
                                fmt.Println(url)
                                urls[url] = true
                        }
                        responseStr = responseStr[endIndex:]
                        startIndex = strings.Index(responseStr, "url=")
                }
        }
}
