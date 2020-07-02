package config

type (
	// Service is the structure for the service information configuration.
	Service struct {
		Group   string
		Name    string
		Version string
	}

	// DB is the structure for the main database configuration.
	DB struct {
		Type       string
		Host       string
		Port       int
		User       string
		Password   string
		Database   string
		Log        bool
		Migrations struct {
			Enabled            bool
			Drop               bool
			SingularTableNames bool
			Seed               bool
		}
	}

	// Log is the structure for the logger configuration.
	// If not present, the Machinery will use a default logger provided
	// by the "gm-log" package.
	Log struct {
		Path       string
		Filename   string
		Console    bool
		Level      string
		JSON       bool
		MaxSize    int
		MaxBackups int
		MaxAge     int
		Compress   bool
		Caller     bool
	}

	// API is the structure for the Http API server and app configuration.
	API struct {
		Endpoint struct {
			Port            int
			BaseRoutingPath string
		}
		// Cors defines the cors allowed resources struct.
		Cors struct {
			Origin  []string
			Methods []string
			Headers []string
		}
		Security struct {
			Enabled bool
			Jwt     struct {
				Secret     string
				Expiration struct {
					Enabled bool
					Minutes int32
				}
			}
		}
	}

	// Configuration describe the type for the configuration file
	Configuration struct {
		Service Service
		API     API
		DB      DB
		Log     Log
	}
)

// Get .
func (c Configuration) Get() {
	// do nothing
}
