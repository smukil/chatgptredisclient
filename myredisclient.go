package main

import (
	"fmt"
	"os"
	"chatgpt/libs/myredisclient/autoredisclient"
)

func main() {
    sharded_replicas := [][]string{
	    []string{"127.0.0.1:6380", "127.0.0.1:6381"},
	    []string{"127.0.0.1:6382", "127.0.0.1:6383"},
    }
    autoclient, err := autoredisclient.NewRedisClient(sharded_replicas)
    if err != nil {
      fmt.Println("Error: ", err)
      os.Exit(3)
    }
    retval,_ := autoclient.Set("testkey", "testxyz")
    fmt.Println("SET Return value is: ", retval)

    retval,err = autoclient.Get("testkey")
    fmt.Println("GET Return value is: ", retval)
    fmt.Println("GET Error is: ", retval)

    retval,_ = autoclient.Set("b_secondkey", "testhaha")
    fmt.Println("SET Return value is: ", retval)

    retval,err = autoclient.Get("b_secondkey")
    fmt.Println("GET Return value is: ", retval)
    fmt.Println("GET Error is: ", retval)
    
}
