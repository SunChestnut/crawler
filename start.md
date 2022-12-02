1. 启动数据持久化的 Server 端：

    ```go
    go run distributed/persist/server/itemsaver.go
    ```

2. 启动 worker 的服务端，服务端个数可自定义

    ```go
    go run distributed/worker/server/worker.go -port 9000
    go run distributed/worker/server/worker.go -port 9001
    ```

3. 运行 main 主函数