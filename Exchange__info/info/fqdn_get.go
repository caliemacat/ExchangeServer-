package info

import (
	client2 "Exchange__info/client"
	"Exchange__info/logger"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strings"
	"unicode/utf16"
)

type NTLMChallenge struct {
	Signature        [8]byte
	MessageType      uint32
	DomainLen        uint16
	DomainMaxLen     uint16
	DomainOffset     uint32
	NegotiateFlags   uint32
	Challenge        [8]byte
	Reserved         [8]byte
	TargetInfoLen    uint16
	TargetInfoMaxLen uint16
	TargetInfoOffset uint32
	Version          []byte
}

func decodeUTF16LE(b []byte) string {
	if len(b)%2 != 0 {
		return ""
	}
	u16s := make([]uint16, len(b)/2)
	for i := 0; i < len(u16s); i++ {
		u16s[i] = binary.LittleEndian.Uint16(b[i*2:])
	}
	return string(utf16.Decode(u16s))
}

func parseTargetInfoFields(data []byte) map[uint16]string {
	result := make(map[uint16]string)
	offset := 0

	for offset+4 <= len(data) {
		avID := binary.LittleEndian.Uint16(data[offset:])
		avLen := binary.LittleEndian.Uint16(data[offset+2:])
		offset += 4

		if avID == 0 {
			break
		}

		if offset+int(avLen) > len(data) {
			break
		}

		value := data[offset : offset+int(avLen)]
		result[avID] = decodeUTF16LE(value)
		offset += int(avLen)
	}
	return result
}

func parseNTLMChallenge(data []byte) (*NTLMChallenge, error) {
	reader := bytes.NewReader(data)

	challenge := &NTLMChallenge{}

	// 1. Fixed length part
	if err := binary.Read(reader, binary.LittleEndian, &challenge.Signature); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.MessageType); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.DomainLen); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.DomainMaxLen); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.DomainOffset); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.NegotiateFlags); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.Challenge); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.Reserved); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.TargetInfoLen); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.TargetInfoMaxLen); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &challenge.TargetInfoOffset); err != nil {
		return nil, err
	}

	domainStart := int(challenge.DomainOffset)
	domainEnd := domainStart + int(challenge.DomainLen)
	if domainEnd <= len(data) {
		domainName := data[domainStart:domainEnd]
		logger.Log.InfoMsaf("目标域为: %s", decodeUTF16LE(domainName))
	}

	targetStart := int(challenge.TargetInfoOffset)
	targetEnd := targetStart + int(challenge.TargetInfoLen)

	if targetEnd <= len(data) {
		targetInfo := data[targetStart:targetEnd]
		avMap := parseTargetInfoFields(targetInfo)

		for avID, val := range avMap {
			switch avID {
			case 0x01:
				_ = fmt.Sprintf("NetBIOS 计算机名:%s", val)
			case 0x02:
				_ = fmt.Sprintf("NetBIOS 域名为:%s", val)
			case 0x03:
				logger.Log.InfoMsaf("机器名为：%s", val)
			case 0x04:
				logger.Log.InfoMsaf("域名为：%s", val)
			case 0x05:
				_ = fmt.Sprintf("邮箱域名 (DNS 树):%s", val)
			default:
				_ = fmt.Sprintf("AV_ID 0x%X: %s", avID, val)
			}
		}
	}

	return challenge, nil

}

func Get_fqdn(URL string, proxyAddr string) error {
	if proxyAddr == "" {
		Client := client2.Global()
		request := Client.R().SetHeader("Authorization", "NTLM TlRMTVNTUAABAAAAB4IIogAAAAAAAAAAAAAAAAAAAAAKAF1YAAAADw==")
		URL = "https://" + URL + "/rpc"
		resp, err := request.Get(URL)
		if err != nil {
			return logger.Log.ErrorMsaf("网络错误为%s", err)
		}
		//获取字段 然后解码即可
		authenticateHeader := resp.Header.Get("WWW-Authenticate")
		splitAuth := strings.Split(authenticateHeader, " ")
		//fmt.Println("Base64-encoded NTLM token:", splitAuth[1])
		decodeBytes, err2 := base64.StdEncoding.DecodeString(splitAuth[1])
		if err2 != nil {
			return logger.Log.ErrorMsaf("解码失败%s", err2)
		}
		logger.Log.InfoMsaf("获取机器名和域名信息")
		_, err3 := parseNTLMChallenge(decodeBytes)
		if err3 != nil {
			return logger.Log.ErrorMsaf("NTLM解析失败%s", err3)
		}

		return nil
	} else {
		Client := client2.Global().SetProxyURL("http://" + proxyAddr)
		request := Client.R().SetHeader("Authorization", "NTLM TlRMTVNTUAABAAAAB4IIogAAAAAAAAAAAAAAAAAAAAAKAF1YAAAADw==")
		URL = "https://" + URL + "/rpc"
		resp, err := request.Get(URL)
		if err != nil {
			return logger.Log.ErrorMsaf("网络错误为%s", err)
		}
		//获取字段 然后解码即可
		authenticateHeader := resp.Header.Get("WWW-Authenticate")
		splitAuth := strings.Split(authenticateHeader, " ")
		//fmt.Println("Base64-encoded NTLM token:", splitAuth[1])
		decodeBytes, err2 := base64.StdEncoding.DecodeString(splitAuth[1])
		if err2 != nil {
			return logger.Log.ErrorMsaf("解码失败%s", err2)
		}
		logger.Log.InfoMsaf("获取机器名和域名信息")
		_, err3 := parseNTLMChallenge(decodeBytes)
		if err3 != nil {
			return logger.Log.ErrorMsaf("NTLM解析失败%s", err3)
		}

		return nil
	}

}
