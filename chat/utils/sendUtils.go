package utils

import (
	"errors"
	"encoding/binary"
	"net"
	// "log"
	"io"
)

type Send struct {
	Conn net.Conn
}

// write json-encode data by connection 
// in: json-encode data which need to be send
// out: write error of connection
func (this *Send) WriteBytes(data []byte) (err error) {
	buffer := make([]byte, 4)
	// get length of data 
	n := uint32(len(data))
	// uint32 to byte[]
	binary.BigEndian.PutUint32(buffer, n)
	// send length of data to check the connection is stable or not
	_, err = this.Conn.Write(buffer)
	if err != nil {
		err = errors.New("[ERR_SEND]:WRITE_NUM (" + err.Error() + ")")
		return
	}

	// send data
	_, err = this.Conn.Write(data)
	if err != nil {
		err = errors.New("[ERR_SEND]:WRITE_DATA (" + err.Error() + ")")
		return
	}

	// log.Println("[_TO_]:", this.Conn.RemoteAddr(), n)
	return
}

// read json-encode data by connection
// out: json-encode data read from connection. read error of connection
func (this *Send) ReadBytes() (data []byte, err error) {
	buffer := make([]byte, 4)
	// get length of data to check the connection is stable or not
	n, err := this.Conn.Read(buffer)
	if err != nil {
		if err == io.EOF{
			return
		}
		err = errors.New("[ERR_SEND]:READ_NUM (" + err.Error() + ")")
		return
	}
	
	// []byte to uint32
	binary.BigEndian.Uint32(buffer)
	lens := int(buffer[n-1])
	
	buffer = make([]byte, 1024)
	// read data
	n, err = this.Conn.Read(buffer)
	// stability test
	if n != lens {
		err = errors.New("[ERR_SEND]:DATA_LOSS")
		return
	}
	if err != nil {
		err = errors.New("[ERR_SEND]:READ_DATA (" + err.Error() + ")")
		return
	}
	data = buffer[:n]
	// log.Println("[FROM]:", this.Conn.RemoteAddr(), lens)
	return
}