package hmondbsqlite

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "hermawan-monitora/hmonglobal"
)

type ServerPorts struct {
  Server MPServer
  Ports []MPPort
}

type result struct {
    ServerGroup string
    IP string
    ServerName string
    Port string
    PortName string
}

func GetDb() (*gorm.DB, error) {
    return gorm.Open(sqlite.Open(hmonglobal.DbFilepath), &gorm.Config{})
}
