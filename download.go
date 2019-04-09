package main

import "os"
import "fmt"
import "strconv"
import "net/http"
import "io/ioutil"

var pages = map[int]int{
    1:5,
    2:5,
    3:5,
    4:7,
    5:11,
    6:6,
}

var directories = map[int]string {
    1: "ch1-basic",
    2: "ch2-function",
    3: "ch3-derivative",
    4: "ch4-differential",
    5: "ch5-indefinite-integral",
    6: "ch6-matrix",
}

var Base = "http://www2.edu-edu.com.cn/lesson_crs78/self/j_0022/soft/"
var Suffix = ".html"

func downloadFile(url string, filename string) {
    resp, err := http.Get(url)
    if err!=nil {
        fmt.Println("Download failed at: ", url)
        return
    }
    defer resp.Body.Close()
    
    file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err!=nil {
        fmt.Println("Open file failed at: ", filename)
        return
    }
    defer file.Close()

    fmt.Println("Storing " + url + " to " + filename)
    body, err := ioutil.ReadAll(resp.Body)
    if err!=nil {
        fmt.Println("Can't store content into ", filename)
        return
    }
    file.WriteString(string(body))
    fmt.Println("Succeed in storing content...")
}

func main() {
    fmt.Println("Start to config...")

    fmt.Println("Creating directories...")
    success := true
    for _,v := range directories {
        fmt.Println("Creating directory " + v + "...")
        if err := os.Mkdir(v, os.ModePerm); err != nil {
            success = false
        }
    }
    if success == false {
        fmt.Println("Creating directories with errors, exiting...")
        os.Exit(-1)
    }
    fmt.Println("Finished creating directories")

    fmt.Println("Start to download urls...")
    for k,v := range pages {
        chapter := strconv.Itoa(k)
        for i:=1; i<=v; i++ {
            hierarchy := "ch0" + chapter + "0" + strconv.Itoa(i)
            url := Base + hierarchy + Suffix
            filename := directories[k] + "/" + hierarchy + Suffix

            downloadFile(url, filename)
        }
    }
    fmt.Println("Finished downloading urls...")
}
