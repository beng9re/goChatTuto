package main

//https://thebook.io/006806/ch09/03/02/ 가이드
import (
	"encoding/json"
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
)

const (
	currentUserKey  = "oauth2_current_user" // 세션에 저장되는 CurrentUser의 키
	sessionDuration = time.Hour             // 로그인 세션 유지 시간
)

/*User is chat User*/
type User struct {
	UID       string    `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"user"`
	AvatarURL string    `json:"avatar_url"`
	Expired   time.Time `json:"expired"`
}

/*Valid is 만료기간 체크 하는 함수*/
func (u *User) Valid() bool {

	// 만료기간 체크
	return u.Expired.Sub(time.Now()) > 0
}

/*Refresh 만료기간 갱신*/
func (u *User) Refresh() {

	//세션 길이 추가
	u.Expired = time.Now().Add(sessionDuration)

}

//GetCorrentUser 세션에서 CurrentUser 정보를 가져옴
func GetCorrentUser(r *http.Request) *User {
	s := sessions.GetSession(r)

	if s.Get(currentUserKey) == nil {
		return nil
	}

	data := s.Get(currentUserKey).([]byte)
	var u User
	json.Unmarshal(data, &u)
	return &u

}

//SetCurrentUser 현재의 유저를 셋팅 해주는 함수
func SetCurrentUser(r *http.Request, u *User) {
	if u != nil {
		// CurrentUser 만료 시간 갱신
		u.Refresh()
	}

	// 세션에 CurrentUser 정보를 json으로 저장
	s := sessions.GetSession(r)
	val, _ := json.Marshal(u)
	s.Set(currentUserKey, val)
}
