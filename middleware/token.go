package middleware

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

//Claims struct
type Claims struct {
	MemberID uuid.UUID
	jwt.StandardClaims
}

//Token struct
type Token struct {
	Value    string    `valid:"stringlength(10|100000)~TOKEN_VALUE_INVALID" json:"value" groups:"login"`
	MemberID uuid.UUID `json:"memberID" groups:""`
	Expire   time.Time `json:"expire" groups:"login"`
}

//MemberToken struct
type MemberToken struct {
	Member model.Member
	Token  Token
}

//GenerateToken function
func GenerateToken(member model.Member) (MemberToken, error) {
	secretKey := viper.GetString("SecretKey")
	jwtKey := []byte(secretKey)
	var output Token
	var err error
	expMinutes := viper.GetInt(viper.GetString("Env") + ".TokenExpMinutes")
	expirationTime := time.Now().Add(time.Duration(expMinutes) * time.Minute)
	claims := &Claims{
		MemberID: member.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	output.Value = tokenString
	output.MemberID = member.ID
	output.Expire = expirationTime

	memberToken := MemberToken{
		Member: member,
		Token:  output,
	}

	return memberToken, err
}

//Validate function to validate token
func (token Token) Validate(app *App) (MemberToken, error) {
	secretKey := viper.GetString("SecretKey")
	jwtKey := []byte(secretKey)
	var err error
	var memberToken MemberToken
	memberToken.Token = token
	tkn, err := jwt.Parse(memberToken.Token.Value, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		var tm time.Time
		memberIDStr, ok := claims["MemberID"].(string)
		if !ok {
			err = errors.New("UNAUTHORIZED")
		}
		switch iat := claims["exp"].(type) {
		case float64:
			tm = time.Unix(int64(iat), 0)
		case json.Number:
			v, _ := iat.Int64()
			tm = time.Unix(v, 0)
		}
		memberToken.Token.Expire = tm
		if err == nil {
			memberToken.Member.ID, err = uuid.FromString(memberIDStr)
			if err != nil {
				err = errors.New("UNAUTHORIZED")
			}
			memberToken.Token.MemberID = memberToken.Member.ID
		}
	} else {
		err = errors.New("UNAUTHORIZED")
	}
	if err == nil {
		db := app.DB1
		memberToken.Member, err = db.GetMemberByID(memberToken.Member.ID)
		if err != nil {
			return memberToken, errors.New("UNAUTHORIZED")
		}
	}
	return memberToken, err
}

//Refresh function to refresh token - not done
func (token Token) Refresh(app *App) (MemberToken, error) {
	var err error
	var memberToken MemberToken
	if err == nil {
		if memberToken.Token.Expire.Sub(time.Now()) > 30*time.Second {
			return memberToken, errors.New("TOKEN_STILL_VALID")
		}
	}
	//refresh function here
	memberToken, err = GenerateToken(memberToken.Member)
	return memberToken, err
}
