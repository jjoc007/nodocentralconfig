package configs
 
 import "io/ioutil"
 
 var (
 	PrivateKey []byte
 )
 
 func Init() {
 	PrivateKey, _ = ioutil.ReadFile("conf/beeblog.rsa")
 }