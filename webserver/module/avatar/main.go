package avatar

import (
    "os"
    "fmt"
    "path/filepath"
    "hermawan-monitora/hmonglobal"
)


func GetAvatarFromFile(username string) ([]byte, error) {
    var err error
    var data []byte
    var workDirLocation string
    workDirLocation, err = os.Getwd()
    if err != nil {
        return nil, err
    }
    sysFileLocation := filepath.Join(workDirLocation,
                                     hmonglobal.AvatarDirPath,
                                     fmt.Sprintf("%s.jpg", username))
    _, err = os.Stat(sysFileLocation)
    if err != nil {
        return nil, err
    }
    data, err = os.ReadFile(sysFileLocation)
    return data, err
}

func GetDefaultAvatarFromFile() ([]byte, error) {
    var err error
    var data []byte
    var workDirLocation string
    workDirLocation, err = os.Getwd()
    if err != nil {
        return nil, err
    }
    sysFileLocation := filepath.Join(workDirLocation, "/static/img/icon/icons8-male-user-96.png")
    _, err = os.Stat(sysFileLocation)
    if err != nil {
        return nil, err
    }
    data, err = os.ReadFile(sysFileLocation)
    return data, err
}
