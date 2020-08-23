package pack

import (
	"encoding/binary"
	"errors"
	"io"
)

const Msg_Header = "zxr"

func Encode(writer io.Writer, content string) (err error) {
	if err = binary.Write(writer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err
	}
	if err = binary.Write(writer, binary.BigEndian, int32(len(content))); err != nil {
		return
	}
	if err = binary.Write(writer, binary.BigEndian, []byte(content)); err != nil {
		return
	}
	return nil
}

func Decode(conn io.Reader) (body []byte, err error) {
	headerBuf := make([]byte, len(Msg_Header))
	if _, err = io.ReadFull(conn, headerBuf); err != nil {
		return nil, err
	}
	if string(headerBuf) != Msg_Header {
		return nil, errors.New("msg header error")
	}
	lenBuf := make([]byte, 4)
	if _, err = io.ReadFull(conn, lenBuf); err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lenBuf)
	bodyBuf := make([]byte, length)
	if _, err = io.ReadFull(conn, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, nil
}
