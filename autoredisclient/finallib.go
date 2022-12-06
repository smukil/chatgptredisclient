package autoredisclient

import (
	"bufio"
  "io"
	"errors"
	"fmt"
	"net"
	"strconv"
)

import jump "github.com/lithammer/go-jump-consistent-hash"
import "github.com/cespare/xxhash"

type RedisClient struct {
	conns   [][]net.Conn
	readers [][]*bufio.Reader
	writers [][]*bufio.Writer
}

func NewRedisClient(replicas [][]string) (*RedisClient, error) {
	var conns [][]net.Conn
	var readers [][]*bufio.Reader
	var writers [][]*bufio.Writer
	for _, replica := range replicas {
		var replicaConns []net.Conn
		var replicaReaders []*bufio.Reader
		var replicaWriters []*bufio.Writer
		for _, server := range replica {
			conn, err := net.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			replicaConns = append(replicaConns, conn)
			replicaReaders = append(replicaReaders, bufio.NewReader(conn))
			replicaWriters = append(replicaWriters, bufio.NewWriter(conn))
		}
		conns = append(conns, replicaConns)
		readers = append(readers, replicaReaders)
		writers = append(writers, replicaWriters)
	}

	return &RedisClient{conns, readers, writers}, nil
}

func FindShard(key string, numServers int32) int {
	h := xxhash.New()
	if _, err := h.Write([]byte(key)); err != nil {
		return 0
	}
	hash := h.Sum64()
	return int(jump.Hash(hash, numServers))
}


func (c *RedisClient) Set(key, value string) (string, error) {
	shard := FindShard(key, int32(len(c.conns[0])))
	var result string
	var err error

  fmt.Println("Set shard for key", key, shard)
	for i, replicaConns := range c.conns {
		conn := replicaConns[shard]
		writer := c.writers[i][shard]
		if err = sendCommand(conn, writer, "SET", key, value); err != nil {
			continue
		}
		result, err = readSimpleString(conn, c.readers[i][shard])
	}
	return result, err
}

func (c *RedisClient) Delete(key string) (string, error) {
	shard := FindShard(key, int32(len(c.conns[0])))
	var result string
	var err error
	for i, replicaConns := range c.conns {
		conn := replicaConns[shard]
		writer := c.writers[i][shard]
		if err = sendCommand(conn, writer, "DELETE", key); err != nil {
			continue
		}
		result, err = readSimpleString(conn, c.readers[i][shard])
	}
	return result, err
}

func (c *RedisClient) Get(key string) (string, error) {
	shard := FindShard(key, int32(len(c.conns[0])))
	var result string
	var err error
	for i, replicaConns := range c.conns {
		conn := replicaConns[shard]
		writer := c.writers[i][shard]
		if err = sendCommand(conn, writer, "GET", key); err != nil {
			continue
		}
		result, err = readBulkString(conn, c.readers[i][shard])
		if err == nil {
			break
		}
	}
	return result, err
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
	for _, replicaConns := range c.conns {
		for _, conn := range replicaConns {
			if err1 := conn.Close(); err1 != nil {
				err = err1
			}
		}
	}
	return err
}

