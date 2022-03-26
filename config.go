package todoapp

import "fmt"

type APIConfig struct {
	Database DatabaseConfig `mapstructure:",squash"`
	Server   ServerConfig   `mapstructure:",squash"`
}

type CLIConfig struct {
	Server ServerConfig `mapstructure:",squash"`
}

type ServerConfig struct {
	Port int `mapstructure:"APPLICATION_PORT"`
}

type DatabaseConfig struct {
	Database        string `mapstructure:"DATABASE_DATABASE"`
	Host            string `mapstructure:"DATABASE_HOST"`
	User            string `mapstructure:"DATABASE_USER"`
	Password        string `mapstructure:"DATABASE_PASS"`
	DBName          string `mapstructure:"DATABASE_NAME"`
	Schema          string `mapstructure:"DATABASE_SCHEMA"`
	ApplicationName string `mapstructure:"DATABASE_APP_NAME"`
	MigrationPath   string `mapstructure:"DATABASE_MIGRATION_PATH"`
	Driver          string `mapstructure:"DATABASE_DRIVER"`
	Port            int    `mapstructure:"DATABASE_PORT"`
}

func (d *DatabaseConfig) DSN() string {
	const (
		dsnPattern = "host=%v port=%v user=%v password=%v dbname=%v sslmode=%v " +
			"application_name=%v search_path=%v"
		disableSSLMode = "disable"
	)

	return fmt.Sprintf(dsnPattern, d.Host, d.Port, d.User, d.Password, d.DBName, disableSSLMode, d.ApplicationName, d.Schema)
}
