// encrypt pw, store encypted value and not the actual the pw
// compare pw recieved enrypted with encrypted pw stored to see if it matches
// server-side encryption
// bcrypt or scrypt is common

package main

import (

  "log"
  "golang.org/x/crypto/bcrypt"
)

func main(){
  // obviously don't use this or any hardcode value- example purposes only
  pass := "012345678910"
  hashedPass, err := hashPw(pass)
  err = comparePassword(pass, hashedPass)
  if err != nil{
    log.Fatalln("doesn't match")
  }
  log.Println("matches, log in")
}

func hashPw(password string)([]byte, error) {
  return  bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// compares the hash
// safe function that doesn't make the comparison itself vulnerable
func comparePassword(password string,hashedPass []byte) error {
  err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
  if err != nil {
    return err
  }
  return nil
}
