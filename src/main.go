type CacheClient struct {
  conn net.Conn
}

func NewCacheClient(addr string, port int) (*CacheClient, error) {
  conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
  if err != nil {
    return nil, err
  }

  return &CacheClient{
    conn: conn,
  }, nil
}

func (c *CacheClient) Set(key, value string) error {
  cmd := fmt.Sprintf("SET %s %s", key, value)
  _, err := c.conn.Write([]byte(cmd))
  if err != nil {
    return err
  }

  resp, err := c.conn.Read([]byte(1024))
  if err != nil {
    return err
  }

  if string(resp) != "OK" {
    return fmt.Errorf("failed to set cache entry: %s", string(resp))
  }

  return nil
}

func (c *CacheClient) Get(key string) (string, error) {
  cmd := fmt.Sprintf("GET %s", key)
  _, err := c.conn.Write([]byte(cmd))
  if err != nil {
    return "", err
  }

  resp, err := c.conn.Read([]byte(1024))
  if err != nil {
    return "", err
  }

  return string(resp), nil
}

func (c *CacheClient) Delete(key string) error {
  cmd := fmt.Sprintf("DEL %s", key)
  _, err := c.conn.Write([]byte(cmd))
  if err != nil {
    return err
  }

  resp, err := c.conn.Read([]byte(1024))
  if err != nil {
    return err
  }

  if string(resp) != "1" {
    return fmt.Errorf("failed to delete cache entry: %s", string(resp))
  }

  return nil
}

func (c *CacheClient) Close() error {
  return c.conn.Close()
}
