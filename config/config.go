package envConf

import (
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once   sync.Once
	result *Env
)

type Env struct {
	Port     string `env:"APP_PORT"`
	Database string `env:"DB_DATABASE"`
	Host     string `env:"DB_HOST"`
	Db_port  string `env:"DB_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
}

func loadFromEnv() {
	godotenv.Load()

	env := &Env{}
	val := reflect.ValueOf(env).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		envKey := field.Tag.Get("env")
		if envKey == "" {
			envKey = field.Name
		}

		envValue := os.Getenv(envKey)
		val.Field(i).SetString(envValue)
	}

	result = env
}

func GetConfig() *Env {

	once.Do(loadFromEnv)

	return result
}
