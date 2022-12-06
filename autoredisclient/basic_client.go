package autoredisclient

// Code commented out to avoid conflicts with the finallib
/*

type RedisClient struct {
	conn net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewRedisClient(ip string, port int) (*RedisClient, error) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &RedisClient{
		conn: conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}, nil
}

func (c *RedisClient) Set(key, value string) (string, error) {
	if err := c.sendCommand("SET", key, value); err != nil {
		return "", err
	}

	res, err := c.readSimpleString()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (c *RedisClient) sendCommand(command string, args ...string) error {

	c.writer.WriteByte('*')
	c.writer.WriteString(strconv.Itoa(len(args) + 1))
	c.writer.WriteByte('\r')
	c.writer.WriteByte('\n')

	c.writer.WriteByte('$')
	c.writer.WriteString(strconv.Itoa(len(command)))
	c.writer.WriteByte('\r')
	c.writer.WriteByte('\n')

	c.writer.WriteString(command)
	c.writer.WriteByte('\r')
	c.writer.WriteByte('\n')

	for _, arg := range args {
*/
