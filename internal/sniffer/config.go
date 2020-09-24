package sniffer

import (
	"github.com/spf13/pflag"
)

func InitialiseFlags() {
	pflag.StringP("device", "d", "eth0", "device to sniff")
	pflag.String("guild_id", "", "guild id for muting")
	pflag.String("report_endpoint", "http://localhost:8080/event", "reporting endpoint for bot auto muting")
	pflag.IntP("server_port", "p", 22023, "port of the server")
}
