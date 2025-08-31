package hmonenv

import (
    "os"
    "github.com/joho/godotenv"
)

func Get(key string) string {
    err := godotenv.Load(".env")

    if err != nil {
        panic("error read .env")
    }

    return os.Getenv(key)
}
