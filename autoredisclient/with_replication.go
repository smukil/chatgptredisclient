package autoredisclient

// Code commented out to avoid conflicts with the finallib
/*
import (
	"bufio"
  "io"
	"errors"
	//"fmt"
	"net"
	"strconv"
)

type RedisClient struct {
	conns   []net.Conn
	readers []*bufio.Reader
	writers []*bufio.Writer
}

func NewRedisClient(servers []string) (*RedisClient, error) {
	var conns []net.Conn
	for _, server := range servers {
		conn, err := net.Dial("tcp", server)
		if err != nil {
			return nil, err
		}
		conns = append(conns, conn)
	}

	c := &RedisClient{conns: conns}
	for _, conn := range conns {
		c.readers = append(c.readers, bufio.NewReader(conn))
		c.writers = append(c.writers, bufio.NewWriter(conn))
	}

	return c, nil
}

func (c *RedisClient) Get(key string) (string, error) {
	for i, conn := range c.conns {
		writer := c.writers[i]
		reader := c.readers[i]
		if err := sendCommand(conn, writer, "GET", key); err != nil {
			return "", err
		}

		res, err := readBulkString(conn, reader)
		if err != nil {
			continue
		}

		return res, nil
	}

	return "", errors.New("failed to get value from all servers")
}

func (c *RedisClient) Set(key, value string) (string, error) {
	var res string
	var err error
	for i, conn := range c.conns {
		writer := c.writers[i]
		reader := c.readers[i]
		if err := sendCommand(conn, writer, "SET", key, value); err != nil {
			return "", err
		}

		r, e := readSimpleString(conn, reader)
		if e != nil {
			err = e
		} else {
			res = r
		}
	}

	if res != "" {
		return res, nil
	}
	return "", err
}

func (c *RedisClient) Delete(key string) (string, error) {
	var res string
	var err error
	for i, conn := range c.conns {
		writer := c.writers[i]
		reader := c.readers[i]
		if err := sendCommand(conn, writer, "DELETE", key); err != nil {
			return "", err
		}

		r, e := readSimpleString(conn, reader)
		if e != nil {
			err = e
		} else {
			res = r
		}
	}

	if res != "" {
		return res, nil
	}
	return "", err
}

func sendCommand(conn net.Conn, writer *bufio.Writer, command string, args ...string) error {
	writer.WriteByte('*')
	writer.WriteString(strconv.Itoa(len(args) + 1))
	writer.WriteByte('\r')
	writer.WriteByte('\n')

	writer.WriteByte('$')
	writer.WriteString(strconv.Itoa(len(command)))
	writer.WriteByte('\r')
	writer.WriteByte('\n')

	writer.WriteString(command)
	writer.WriteByte('\r')
	writer.WriteByte('\n')

	for _, arg := range args {
		writer.WriteByte('$')
		writer.WriteString(strconv.Itoa(len(arg)))
		writer.WriteByte('\r')
		writer.WriteByte('\n')
		writer.WriteString(arg)
		writer.WriteByte('\r')
		writer.WriteByte('\n')
	}

	return writer.Flush()
}

func readSimpleString(conn net.Conn, reader *bufio.Reader) (string, error) {
	prefix, err := reader.ReadByte()
	if err != nil {
		return "", err
	}

	if prefix == '-' {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		return "", errors.New(line[:len(line)-2])
	}

	if prefix != '+' {
		return "", errors.New("unexpected prefix")
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return line[:len(line)-2], nil
}

func readBulkString(conn net.Conn, reader *bufio.Reader) (string, error) {
	prefix, err := reader.ReadByte()
	if err != nil {
		return "", err
	}

	if prefix == '-' {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		return "", errors.New(line[:len(line)-2])
	}

	if prefix != '$' {
		return "", errors.New("unexpected prefix")
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	n, err := strconv.Atoi(line[:len(line)-2])
	if err != nil {
		return "", err
	}

	if n == -1 {
		return "", nil
	}

	bytes := make([]byte, n)
	if _, err := io.ReadFull(reader, bytes); err != nil {
		return "", err
	}

	if _, err := reader.ReadByte(); err != nil {
		return "", err
	}

	if _, err := reader.ReadByte(); err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c *RedisClient) Close() error {
	var err error
	for _, conn := range c.conns {
		if e := conn.Close(); e != nil {
			err = e
		}
	}
	return err
}
*/
