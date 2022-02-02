package grpcutils

import (
	"bytes"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
)

const (
	BlockSize = 1024*1024*2 - 1024*2
)

type GrpcStrSender interface {
	Send(value *wrapperspb.StringValue) error
}

type GrpcStrRecv interface {
	Recv() (*wrapperspb.StringValue, error)
}

// SendBytesToStrServer 发送数据到grpc服务
func SendBytesToStrServer(server GrpcStrSender, data []byte) error {
	reader := bytes.NewReader(data)
	buf := make([]byte, BlockSize, BlockSize)
	for {
		readSize, err := reader.Read(buf)
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if err = server.Send(wrapperspb.String(string(buf[:readSize]))); err != nil {
			return err
		}
	}
}

// RecvStreamStrToStr 接收字符串流拼装为字符串
func RecvStreamStrToStr(server GrpcStrRecv) (string, error) {
	buf := &bytes.Buffer{}
	for {
		recv, err := server.Recv()
		if err == io.EOF {
			return buf.String(), nil
		}
		if err != nil {
			return "", err
		}
		buf.WriteString(recv.Value)
	}
}
