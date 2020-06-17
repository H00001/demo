package tcp

import (
	"fmt"
	"net"
	test "test/internal"
	"test/internal/transports/tcp/controllers"
)

type TcpServer struct {
	fd  net.Listener
	app *test.App
}

func NewTcpServer(app *test.App) *TcpServer {
	t := TcpServer{app: app}
	return &t
}

func (gs *TcpServer) Start() {
	var err error
	gs.fd, err = net.Listen("tcp", "localhost:8999")
	if err != nil {
		fmt.Println("listener creat error", err)
	}
	fmt.Println("waiting for client")
	go func() {
		for {
			conn, err := gs.fd.Accept()
			if err != nil {
				fmt.Println("accept error", err)
			}
			go gs.handleConnection(conn)
		}
	}()

}

func (gs *TcpServer) handleConnection(conn net.Conn) {
	c := controllers.NewControllers(gs.app)
	for {
		fmt.Println("connection success")
		fmt.Println("client address: ", conn.RemoteAddr())
		buffer := make([]byte, 1024)
		recvLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error", err)
		}
		strBuffer := string(buffer[:recvLen])
		fmt.Println("Message: ", strBuffer)
		fmt.Println("Message len :", recvLen)
		c.ServerNode.Regisite(buffer[:recvLen])
		if err != nil {
			fmt.Println("send message error", err)
		}
		fmt.Println("send message success")
	}
}

func (gs *TcpServer) GracefulShutDown() {
	_ = gs.fd.Close()
}
