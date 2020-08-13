package bot

import (
	"github.com/spf13/pflag"
)

func InitialiseFlags() {
	pflag.StringP("token", "t", "", "bot token")
}
