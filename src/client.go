package main

import (
    "fmt"
    _ "net/rpc"
    "time"
    "net/rpc/jsonrpc"
)

var quit chan int

func server(i int) {


    fmt.Println("这是第",i,"个进程")
    flag := false
    //for{
        time.Sleep(1e9)
        if(flag){
            return
        }
        var reply int64
        var beginTime int64
        beginTime =time.Now().UnixNano()
        fmt.Println("客户端启动时间 单位/ms： ",beginTime/1e6)
        fmt.Println("客户端启动时间 单位/s： ",time.Unix(beginTime/1e9,0))
        client,err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
        defer client.Close()

        if err != nil {
            fmt.Println("连接rpc服务器失败",err)
            flag = true
        }

        //验证安全性
        args :="sunwei"

        err = client.Call("Arith.Receive", &args, &reply) //批量执行
        if err != nil {
            fmt.Println("调用远程RPC服务失败",err)
            flag =true
        }
        endTime := time.Now().UnixNano()
        fmt.Println("结束时间 单位/ms： ",endTime/1e6)
        fmt.Println("结束时间 单位/s： ",time.Unix(endTime/1e9,0))

        intervalTime := endTime-beginTime
        fmt.Println("延迟： ",intervalTime)
        serverTime := (reply + intervalTime/2)/1e6

        fmt.Println("这是第",i,"客户端")
        fmt.Println("远程服务器时间 单位/ms： ",serverTime)
        fmt.Println("远程服务器时间标准化：  ",time.Unix(serverTime/1e3,0))

    //}
    quit <- 0
}

func  main()  {
    length :=1000
    quit = make(chan int)//无缓冲
    for i:=0;i<length;i++{
        go server(i)
    }

    for i:=0;i<length;i++{
        <- quit
    }
    //close(quit)
}