package store

// PostgresConfig contains config data to connect to MySQL database
type PostgresConfig struct {
	Host                      string `json:"host"`
	Port                      int    `json:"port"`
	Schema                    string `json:"schema"`
	User                      string `json:"user"`
	Password                  string `json:"password"`
	Option                    string `json:"option"`
	ConnectionLifetimeSeconds int    `json:"connection_lifetime_seconds"`
	MaxIdleConnections        int    `json:"max_idle_connections"`
	MaxOpenConnections        int    `json:"max_open_connections"`
}
