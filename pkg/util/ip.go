package util

import (
	"errors"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

var IP = new(ipUtil)

type ipUtil struct{}

// 获取用户发送请求的 IP 地址
// 如果服务器不经过代理, 直接把自己 IP 暴露出去, c.Request.RemoteAddr 就可以直接获取 IP
// 目前流行的架构中, 请求经过服务器前基本会经过代理 (Nginx 最常见), 此时直接获取 IP 拿到的是代理服务器的 IP
func (*ipUtil) GetIpAddress(c *gin.Context) (ipAddress string) {
	// c.ClientIP() 获取的是代理服务器的 IP (Nginx)

	// X-Real-IP: Nginx 服务代理, 本项目明确使用 Nginx 作代理, 因此优先获取这个
	ipAddress = c.Request.Header.Get("X-Real-IP")

	// X-Forwarded-For 经过 HTTP 代理或 负载均衡服务器时会添加该项
	// X-Forwarded-For 格式: client1,proxy1,proxy2
	// 一般情况下，第一个 IP 为客户端真实 IP，后面的为经过的代理服务器 IP
	if ipAddress == "" || len(ipAddress) == 0 || strings.EqualFold("unknown", ipAddress) {
		ips := c.Request.Header.Get("X-Forwarded-For") // "ip1,ip2,ip3"
		splitIps := strings.Split(ips, ",")            // ["ip1", "ip2", "ip3"]
		if len(splitIps) > 0 {
			ipAddress = splitIps[0]
		}
	}

	// Pdoxy-Client-IP: Apache 服务代理
	if ipAddress == "" || len(ipAddress) == 0 || strings.EqualFold("unknown", ipAddress) {
		ipAddress = c.Request.Header.Get("Proxy-Client-IP")
	}

	// WL-Proxy-Client-IP: Weblogic 服务代理
	if ipAddress == "" || len(ipAddress) == 0 || strings.EqualFold("unknown", ipAddress) {
		ipAddress = c.Request.Header.Get("WL-Proxy-Client-IP")
	}

	// RemoteAddr: 发出请求的远程主机的 IP 地址 (经过代理会设置为代理机器的 IP)
	if ipAddress == "" || len(ipAddress) == 0 || strings.EqualFold("unknown", ipAddress) {
		ipAddress = c.Request.RemoteAddr
	}

	// 检测到是本机 IP, 读取其局域网 IP 地址
	if strings.HasPrefix(ipAddress, "127.0.0.1") || strings.HasPrefix(ipAddress, "[::1]") {
		ip, err := externalIP()
		if err != nil {
		}
		ipAddress = ip.String()
	}

	if ipAddress != "" && len(ipAddress) > 15 {
		if strings.Index(ipAddress, ",") > 0 {
			ipAddress = ipAddress[:strings.Index(ipAddress, ",")]
		}
	}
	return ipAddress
}

// 获取非 127.0.0.1 的局域网 IP
func externalIP() (net.IP, error) {
	// 获取服务器的网络接口列表
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		// 不在活动状态
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		// 环回
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		// 单播接口地址列表
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network")
}
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil
	}
	return ip
}
