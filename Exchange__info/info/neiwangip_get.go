package info

import (
	"Exchange__info/logger"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"strings"
)

func Get_IP(URL string, proxyAddr string) error {
	var conn net.Conn
	var err error

	if proxyAddr != "" {
		logger.Log.InfoMsaf("[*] 使用代理连接：%s", proxyAddr)
		conn, err = net.Dial("tcp", proxyAddr)
		if err != nil {
			return logger.Log.ErrorMsaf("连接代理失败: %v", err)
		}

		connectRequest := fmt.Sprintf("CONNECT %s:443 HTTP/1.1\r\nHost: %s\r\n\r\n", URL, URL)
		_, err = conn.Write([]byte(connectRequest))
		if err != nil {
			return logger.Log.ErrorMsaf("发送 CONNECT 请求失败: %v", err)
		}

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			return logger.Log.ErrorMsaf("读取 CONNECT 响应失败: %v", err)
		}
		resp := string(buf[:n])

		if !strings.Contains(resp, "200 Connection established") {
			return logger.Log.ErrorMsaf("代理未成功建立连接")
		}
	} else {
		logger.Log.InfoMsaf("[*] 不使用代理，直连目标地址：%s", URL)
		conn, err = net.Dial("tcp", URL+":443")
		if err != nil {
			return logger.Log.ErrorMsaf("直连目标失败: %v", err)
		}
	}

	defer conn.Close()

	// TLS 握手
	tlsConn := tls.Client(conn, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         URL + ":443",
		MinVersion:         tls.VersionTLS10,
		MaxVersion:         tls.VersionTLS12,
	})
	err = tlsConn.Handshake()
	if err != nil {
		return logger.Log.ErrorMsaf("TLS 握手失败: %v", err)
	}

	rawRequest := "GET /ecp HTTP/1.0\r\n" +
		"User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:137.0) Gecko/20100101 Firefox/137.0\r\n" +
		"Accept: */*\r\n" +
		"Accept-Encoding: gzip, deflate\r\n" +
		"Connection: close\r\n\r\n"

	_, err = tlsConn.Write([]byte(rawRequest))
	if err != nil {
		return logger.Log.ErrorMsaf("发送 HTTP 请求失败: %v", err)
	}

	var fullResp strings.Builder

	buf := make([]byte, 4096)
	for {
		n, err := tlsConn.Read(buf)
		if n > 0 {
			chunk := string(buf[:n])
			fullResp.WriteString(chunk)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return logger.Log.ErrorMsaf("读取响应失败: %v", err)
		}
	}

	rawHeaders := fullResp.String()
	scanner := strings.Split(rawHeaders, "\r\n")
	for _, line := range scanner {
		if strings.HasPrefix(strings.ToLower(line), "location:") {
			location := strings.TrimSpace(line[len("location:"):])
			split := strings.Split(location, "/")
			logger.Log.InfoMsaf("探测的内网IP地址为：%s", split[2])
			break
		}
	}

	return nil
}
