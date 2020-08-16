package bot

import (
	"github.com/spf13/pflag"
)

func InitialiseFlags() {
	pflag.StringP("token", "t", "", "bot token")
	pflag.BoolP("verbose", "v", false, "verbose mode, responds to every command")
}
