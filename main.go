package main

import (
    "fmt"
    "net"
    "net/http"
    "io/ioutil"
    "sync"
)

func handleConnection(c net.Conn, cond *sync.Cond) {
    defer c.Close()
    for {
        cond.L.Lock()
        cond.Wait()
        c.Write([]byte("<password>\n"))
        cond.L.Unlock()
    }
}
func homePage(w http.ResponseWriter, r *http.Request, cond *sync.Cond){
    buff, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return
    }
    if string(buff) != "<password>" {
        w.WriteHeader(400)
        return
    }
    cond.Broadcast()
    w.WriteHeader(200)
    w.Write([]byte("Boot signal sent\n"))
}

func main() {
    addr, err := net.ResolveTCPAddr("tcp", ":8080")
    if err != nil {
        fmt.Printf("Unable to resolve IP")
    }
    ln, err := net.ListenTCP("tcp", addr)
    if err != nil {
        fmt.Println("Baad")
    }

    m := &sync.Mutex{}
    cond := sync.NewCond(m)

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
        homePage(w, r, cond)
    })
    go http.ListenAndServeTLS(":6969", "/home/ubuntu/files/server.crt", "/home/ubuntu/files/server.key", nil)
    for {
            conn, err := ln.AcceptTCP()
        fmt.Println("Accepted")
            if err != nil {
                // handle error
            }
        _ = conn.SetKeepAlive(true)
            go handleConnection(conn, cond)
    }
}

