package auth
import "golang.org/x/crypto/bcrypt"


func HashPasswod(password string)(string,error){
  hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
  if err!=nil{
	return "",err
  }
  return  string(hash),nil
}

func ComparePassword(hasedPassword string,password string)(bool){
  err:=bcrypt.CompareHashAndPassword([]byte(hasedPassword),[]byte(password))
  if err!=nil{
    return  false
  }
  return true
}