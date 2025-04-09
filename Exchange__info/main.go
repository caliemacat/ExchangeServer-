package main

import (
	"Exchange__info/info"
	"Exchange__info/logger"
	"flag"
	"fmt"
	"log/slog"
)

var (
	url     string
	help    bool
	version bool
	fqdn    bool
	ip      bool
	proxy   string
)

func main() {
	logger.Init(
		logger.WithLevel(slog.LevelDebug),
		logger.WithTimeFormat("15:04:05"),
		logger.WithUseColor(true),
		logger.WithOutputJson(true))

	// 定义标志
	flag.StringVar(&url, "url", "", "请输入要处理的URL或IP地址。")
	flag.BoolVar(&version, "v", false, "显示应用程序的版本。")
	flag.BoolVar(&fqdn, "f", false, "获取给定URL或IP的完全限定域名（FQDN）。")
	flag.BoolVar(&ip, "ip", false, "获取给定URL或IP的内网IP地址。")
	flag.BoolVar(&help, "h", false, "显示此帮助信息。")
	flag.StringVar(&proxy, "p", "", "传入代理地址")

	// 自定义flag的Usage信息
	flag.Usage = func() {
		fmt.Println(`
用法: main.exe -url <IP>

选项:
  -url string   指定要处理的IP地址。
  -v            显示当前应用程序的版本。
  -f            获取给定URL或IP的完全限定域名（FQDN）。
  -ip           获取给定URL或IP的内网IP地址。
  -p            指定是否使用代理
  -h            显示此帮助信息。
`)
	}

	// 解析命令行参数
	flag.Parse()

	// 如果请求显示帮助或者URL未提供，则显示帮助信息
	if help || (url == "" && !version && !fqdn && !ip && proxy == "") {
		flag.Usage()
		return
	}

	// 根据不同的标志进行相应操作
	if version {
		_ = info.Get_Version(url, proxy)
	} else if fqdn {
		_ = info.Get_fqdn(url, proxy)
	} else if ip {
		_ = info.Get_IP(url, proxy)
	}
}
