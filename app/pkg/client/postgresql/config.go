package postgresql

import "fmt"

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (c Config) MakeDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s", c.Username, c.Password, c.Host, c.Port, c.Database,
	)
}
