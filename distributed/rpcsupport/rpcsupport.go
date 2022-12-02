package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// ServeRpc 创建 RPC Server，监听 host 端口
func ServeRpc(host string, service any) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}

	listen, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	log.Printf("🥰Listening on %s\n", host)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error : %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
	return nil
}

// NewClient ==> 创建一个新的 RPC 客户端去连接地址为 host 的服务端
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Printf("Error dial rpc server: %v", err)
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
