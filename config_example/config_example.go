// change to package config
package configexample

import "os"

// SetVars func
func SetVars() {
	os.Setenv("DB_PARAMS", "user=admin password=password dbname=mf_soundboard-v1 host=db sslmode=disable")
	// os.Setenv("DB_PARAMS", "user=admin password=password dbname=<database-name-here> host=db<refers to 'db' service in docker-compose file> sslmode=disable")
}
