package config
import "os"
type Config struct { DBHost,DBPort,DBUser,DBPassword,DBName,DBSSLMode string }
func Load() Config { return Config{get("DB_HOST","localhost"),get("DB_PORT","5432"),get("DB_USER","postgres"),get("DB_PASSWORD","Root"),get("DB_NAME","architecture_db"),get("DB_SSLMODE","disable")} }
func (c Config) DatabaseURL() string { return "postgres://"+c.DBUser+":"+c.DBPassword+"@"+c.DBHost+":"+c.DBPort+"/"+c.DBName+"?sslmode="+c.DBSSLMode }
func get(k,d string)string{if v:=os.Getenv(k);v!=""{return v};return d}
