/**
    @author: cloudy
    @date: 2022-07-25
    @note:
**/
package hosts

import "net"

///
//  GetLocalIP
//  @Description: 获取本机IP Addresses
//  @return string
///
func GetLocalIP() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addr {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok &&
			!ipnet.IP.IsUnspecified() &&
			!ipnet.IP.IsLoopback() &&
			!ipnet.IP.IsMulticast() &&
			!ipnet.IP.IsLinkLocalMulticast() &&
			!ipnet.IP.IsInterfaceLocalMulticast() &&
			!ipnet.IP.IsLinkLocalUnicast() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
