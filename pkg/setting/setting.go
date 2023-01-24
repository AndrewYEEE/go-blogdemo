package setting

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ServerENV struct {
	HTTPPort     int
	PageSize     int
	JwtSecret    string
	RunMode      string
	DBType       string
	DBName       string
	User         string
	Password     string
	Host         string
	TablePrefix  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ENV = &ServerENV{}

func Setup() {
	err := godotenv.Load(filepath.Join("", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file :", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadDB()
}

func LoadBase() {
	ENV.RunMode = getStrEnv("RUN_MODE")
}

func LoadServer() {
	val := getStrEnv("HTTP_PORT")
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic("HTTP_PORT error")
	}
	ENV.HTTPPort = ret

	val = getStrEnv("READ_TIMEOUT")
	read, err := strconv.Atoi(val)
	if err != nil {
		panic("READ_TIMEOUT error")
	}

	val = getStrEnv("WRITE_TIMEOUT")
	write, err := strconv.Atoi(val)
	if err != nil {
		panic("WRITE_TIMEOUT error")
	}

	ENV.ReadTimeout = time.Duration(read) * time.Second
	ENV.WriteTimeout = time.Duration(write) * time.Second
}

func LoadApp() {
	ENV.JwtSecret = getStrEnv("JWT_SECRET")
	val := getStrEnv("PAGE_SIZE")
	page, err := strconv.Atoi(val)
	if err != nil {
		panic("PAGE_SIZE error")
	}
	ENV.PageSize = page
}

func LoadDB() {
	ENV.DBType = getStrEnv("TYPE")
	ENV.DBName = getStrEnv("NAME")
	ENV.User = getStrEnv("USER")
	ENV.Password = getStrEnv("PASSWORD")
	ENV.Host = getStrEnv("HOST")
	ENV.TablePrefix = getStrEnv("TABLE_PREFIX")
}

func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("some error msg")
	}
	return val
}
