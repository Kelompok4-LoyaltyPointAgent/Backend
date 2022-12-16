package config

import "os"

type XenditConfig struct {
	Secret string
}

func LoadXenditConfig() XenditConfig {
	return XenditConfig{
		Secret: os.Getenv("XENDIT_SECRET"),
	}
}
