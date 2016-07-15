package controllers
 
 import (
 	"time"
 	"fmt"
 
 	"github.com/astaxie/beego"
 	"github.com/astaxie/beego/logs"
 	jwt "github.com/dgrijalva/jwt-go"
 
 	//"confignodoapi/configs"
 )
 
 // oprations for Jwt
 type JwtController struct {
 	beego.Controller
 }
 
 func (this *JwtController) URLMapping() {
 	this.Mapping("IssueToken", this.IssueToken)
 }
 
 // @Title IssueToken
 // @Description Issue a Json Web Token
 // @Success 200 string
 // @Failure 403 no privilege to access
 // @Failure 500 server inner error
 // @router /issue-token [get]
 func (this *JwtController) IssueToken() {
 	log := logs.NewLogger(10000)
 	log.SetLogger("console", "")
 
 	mySigningKey := []byte("AllYourBase")
 
 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	    "ID": "This is my super fake ID",
	    "exp": time.Now().Unix() + 36000,
	})

 	// The claims object allows you to store information in the actual token.
 	tokenString, err := token.SignedString(mySigningKey)

 	// tokenString Contains the actual token you should share with your client.
 	this.Data["json"] = map[string]string{"token": tokenString}
 	
 	fmt.Println(tokenString, err)




tokenVal, errVal := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte("AllYourBase"), nil
})

if claims, ok := tokenVal.Claims.(jwt.MapClaims); ok && tokenVal.Valid {
    fmt.Println("You look nice today")
    fmt.Println(claims["ID"])
} else if ve, ok := errVal.(*jwt.ValidationError); ok {
    if ve.Errors&jwt.ValidationErrorMalformed != 0 {
        fmt.Println("That's not even a token")
    } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
        // Token is either expired or not active yet
        fmt.Println("Timing is everything")
    } else {
        fmt.Println("Couldn't handle this token:", errVal)
    }
} else {
    fmt.Println("Couldn't handle this token:", errVal)
}





 	//	log.Debug(err.Error())
 	this.ServeJSON()


 	
 	
 }