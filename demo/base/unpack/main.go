package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const Msg_Header = "12345678"

func Encode(bytesBuffer io.Writer, content string) error {

	//msg_header+content_len+content  8+4+content_len

	/**func Write(w io.Writer, order ByteOrder, data interface{}) error
		将data的binary编码格式写入w，data必须是定长值、定长值的切片、定长值的指针。order指定写入数据的字节序，写入结构体时，名字中有'_'的字段会置为0。
	**/
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err
	}

	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	MagicBug := make([]byte, len(Msg_Header))

	/**
	func ReadFull(r Reader, buf []byte) (n int, err error)
	ReadFull从r精确地读取len(buf)字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。只有没有读取到字节时才可能返回EOF；
	如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。 只有返回值err为nil时，返回值n才会等于len(buf)。
	*/
	if _, err := io.ReadFull(bytesBuffer, MagicBug); err != nil {
		return nil, err
	}
	if string(MagicBug) != Msg_Header {
		return nil, errors.New("msg_header error")
	}

	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf = make([]byte, length)
	if _, err := io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}

	return bodyBuf, err
}
