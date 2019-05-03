package storage

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type userInfo struct {
	Email            string
	RegistrationDate int64
}

func (s *Storage) NewUser(email string) string {
	hash := md5.Sum([]byte(email))
	token := fmt.Sprintf("%x", hash)

	if s.UserExists(token) {
		return token
	}

	s.Lock()
	s.users[token] = userInfo{
		Email:            email,
		RegistrationDate: time.Now().Unix(),
	}
	s.Unlock()

	return token
}

func (s *Storage) GetUserDevices(token string) (collection []deviceInfo) {
	s.RLock()
	for _, d := range s.userDevices[token] {
		collection = append(collection, s.devices[d])
	}
	s.RUnlock()
	return
}

func (s *Storage) UserExists(token string) bool {
	s.RLock()
	defer s.RUnlock()
	return s.users[token].Email != ""
}

func (s *Storage) UserExistsByEmail(email string) bool {
	hash := md5.Sum([]byte(email))
	token := fmt.Sprintf("%x", hash)

	return s.UserExists(token)
}

func (s *Storage) SaveUsers(fileName string) {
	u, err := json.Marshal(s.users)
	if err != nil {
		fmt.Printf("'%s' is not saved by error: %v\n", fileName, err)
		return
	}

	err = ioutil.WriteFile(fileName, u, 0644)
	if err != nil {
		fmt.Printf("'%s' is not saved by error: %v\n", fileName, err)
		return
	}
}

func (s *Storage) SaveUserDevices(fileName string) {
	u, err := json.Marshal(s.userDevices)
	if err != nil {
		fmt.Printf("'%s' is not saved by error: %v\n", fileName, err)
		return
	}

	err = ioutil.WriteFile(fileName, u, 0644)
	if err != nil {
		fmt.Printf("'%s' is not saved by error: %v\n", fileName, err)
		return
	}
}

func (s *Storage) addUserDevice(email string, deviceToken string) {
	hash := md5.Sum([]byte(email))
	token := fmt.Sprintf("%x", hash)

	s.Lock()
	s.userDevices[token] = append(s.userDevices[token], deviceToken)
	s.Unlock()
}
