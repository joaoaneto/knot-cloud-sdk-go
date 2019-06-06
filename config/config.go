package config

// KnotWSConfig represents Knot WebSocket configuration properties
type KnotWSConfig struct {
	Hostname string
	Port     int
	Protocol string
}

// NewKnotWSConfig creates a new KnotWSConfig instance
func New(hostname string, port int, protocol string) KnotWSConfig {
	return KnotWSConfig{hostname, port, protocol}
}
