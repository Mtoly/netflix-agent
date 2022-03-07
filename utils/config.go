package utils

import (
	"github.com/spf13/viper"
	"strings"
)

func GetConfig() *viper.Viper {
	c := viper.New()
	c.SetConfigType("yaml")
	c.SetConfigName("config")
	c.AddConfigPath(".")
	c.AutomaticEnv()

	c.SetDefault("net.ipv6.start", "2001:df3:2e80:4000::")
	c.SetDefault("net.ipv6.end", "2001:0df3:2e80:bfff:ffff:ffff:ffff:ffff")
	c.SetDefault("net.interface.name", "ens160")

	c.SetDefault("mode", "iptables")

	c.SetDefault("netflix.title", "70143836")

	replacer := strings.NewReplacer(".", "_")
	c.SetEnvKeyReplacer(replacer)
	c.ReadInConfig()
	return c
}
