package config

type Mpotato struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultMpotatoConfig returns default config values
func DefaultMpotatoConfig() Mpotato {
	return Mpotato{
	HelloWorldMessage:
		"hello from the default config",
	}
}
