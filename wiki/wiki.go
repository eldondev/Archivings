package main

import (
  "net"
  "flag"
  "os"
  "log"
  "bufio"
  "io"
  "fmt"
  "time"
  "strings"
)

func main() {
  flag.Parse()
  var servername string
  server := flag.Args()[0]
  nick := flag.String("nick", "goirseeds", "The nickname for the bot to use")
  user := flag.String("user", "goirseedsuser", "The username for the bot to use")
  name := flag.String("name", "goirseedsname", "The name for the bot to use")
  conn, err := net.Dial("tcp", server)
  if err != nil {
    log.Fatal(err)
  }

  pipereader, pipewriter := io.Pipe()
  sockin := io.MultiWriter(os.Stdout, pipewriter)
  myscan := bufio.NewScanner(pipereader)
  go io.Copy(sockin, conn)
  defer conn.Close()
  io.WriteString(conn, fmt.Sprintf("NICK %s\nUSER %s * * :%s\n", *nick, *user, *name))
  msgs := make(chan string, 2)
  pings := make(chan string)
  commands := make(chan string)
  sends := make(chan string, 1000)

  go func() {
   for myscan.Scan() {
     msgs <- myscan.Text()
   }
  }()
  go func() {
    channels := make([]string, 0)
    for command := range(commands) {
      switch {
        case strings.HasPrefix(command, " 321 "): {
          channels = make([]string, 0)
        }
        case strings.HasPrefix(command, " 322 "): {
          command = strings.TrimPrefix(command, fmt.Sprintf(" 322 %s ", *nick))
          channel := strings.Split(command, " ")[0]
          channels = append(channels, channel)
        }
        case strings.HasPrefix(command, " 323 "): {
          go func() {
            for _, channel := range(channels) {
              sends <- fmt.Sprintf("JOIN %s\n", channel)
              time.Sleep(3 * time.Second)
            }
          }()
        }
        case strings.HasPrefix(command, " 376 ") || strings.HasPrefix(command, " 422 "): {
          sends <- "LIST"
        }
       }
    }
  }()
  go func() {
        for resp := range(sends) {
          io.WriteString(conn, resp + "\n")
          time.Sleep(2 * time.Second);
        }
  }()
  go func() {
    for msg := range(msgs) {
      switch {
        case servername == "" && strings.HasPrefix(msg,":"):
          servername = strings.Split(msg, " ")[0]
          msgs <- msg
        case strings.HasPrefix(msg, "PING") :
          pings <- msg
        case strings.HasPrefix(msg, servername) :
         commands <- strings.TrimPrefix(msg, servername)
      }
    }
  }()
  for {
    select {
      case p := <- pings:
        sends <- "PONG " + p[6:]
      case <-time.After(time.Second * 11):
              fmt.Println("timeout 2")
      }
    }
  }
