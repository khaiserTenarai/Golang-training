package utility

import (
 "errors"
 "time"
 "github.com/golang-jwt/jwt/v5"
)
type Claims struct { UserID int64 `json:"user_id"`; Username string `json:"username"`; Role string `json:"role"`; jwt.RegisteredClaims }
func GenerateToken(secret string,expiryMinutes int,userID int64,username,role string)(string,error){claims:=Claims{UserID:userID,Username:username,Role:role,RegisteredClaims:jwt.RegisteredClaims{Subject:username,IssuedAt:jwt.NewNumericDate(time.Now()),ExpiresAt:jwt.NewNumericDate(time.Now().Add(time.Duration(expiryMinutes)*time.Minute)),Issuer:"secure-gin-api"}};return jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString([]byte(secret))}
func ParseToken(secret,tokenString string)(*Claims,error){t,err:=jwt.ParseWithClaims(tokenString,&Claims{},func(t *jwt.Token)(any,error){if t.Method.Alg()!=jwt.SigningMethodHS256.Alg(){return nil,errors.New("unexpected signing method")};return []byte(secret),nil});if err!=nil{return nil,err};c,ok:=t.Claims.(*Claims);if !ok||!t.Valid{return nil,errors.New("invalid token")};return c,nil}
