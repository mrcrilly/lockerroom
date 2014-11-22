
package lockerroom

import (
  "time"
  "encoding/base64"
)

type LockerRoom interface {
  Encrypt([]byte, []byte) ([]byte, []byte, []byte, error)
  Decrypt([]byte, []byte, []byte, []byte) ([]byte, error)
  Secret([]byte) ([]byte, error)

  Lock() (error) // Scrubs data
}

type Locker struct {
  Path []byte `json:"path"`

  Locked bool `json:"locked"`

  SecuredSecret []byte `json:"securedsecret"`
  Shelves []LockerShelf `json:"shelves"`

  Created time.Time `json:"created"`
  Accessed time.Time `json:"accessed"`
  Modified time.Time `json:"modified"`
}

type LockerShelf struct {
  Locked bool `json:"locked"`
  Manipulated bool `json:"-"`

  SecuredSecret []byte `json:"securedsecret"`
  Items []LockerItem `json:"items"`

  Created time.Time `json:"created"`
  Accessed time.Time `json:"accessed"`
  Modified time.Time `json:"modified"`
}

type LockerItem struct {
  Secret []byte `json:"secret"`
  Notes []byte `json:"notes"`

  Manipulated bool `json:"-"`
}

func (lkr *LockerShelf) Encrypt(password, plaintext []byte) (ciphertext, salt, nonce []byte, err error) {
  return nil, nil, nil, nil
}

func (lkr *LockerShelf) Decrypt(password, salt, none, ciphertext []byte) (plaintext []byte, err error) {
  return nil, nil 
}

func (lkr *LockerShelf) Secret(password []byte) (secret []byte, err error) {
  return nil, nil
}

func (lkr *LockerShelf) Lock() (err error){
  return nil
}
