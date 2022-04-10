package jike

import (
	"context"
	"net/http"
)

type UsersService struct {
	jike   *Jike
	client *http.Client
	debug  bool
}

func NewUsersService(jike *Jike, c *http.Client, debug bool) *UsersService {
	return &UsersService{jike, c, debug}
}

func (u *UsersService) PasswordLogin(ctx context.Context, areaCode, phone, password string) (*LoginOutput, error) {
	input := map[string]interface{}{
		"areaCode":          areaCode,
		"mobilePhoneNumber": phone,
		"password":          password,
	}
	output := &LoginOutput{}
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `/1.0/users/loginWithPhoneAndPassword`,
		Input:      input,
		Output:     &output,
		Context:    ctx,
	}
	return output, req.send(u.jike)
}

func (u *UsersService) Profile(ctx context.Context, username string) (*ProfileOutput, error) {
	params := map[string]string{
		"username": username,
	}
	output := &ProfileOutput{}
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `/1.0/users/profile`,
		Params:     params,
		Output:     &output,
		Context:    ctx,
	}
	return output, req.send(u.jike)
}

func (u *UsersService) FollowingTimeline(ctx context.Context, limit int, loadMoreKey TimelineLoadMoreKey) (*TimelineOutput, error) {
	input := map[string]interface{}{
		"limit":       limit,
		"loadMoreKey": loadMoreKey,
	}
	output := &TimelineOutput{}
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `/1.0/personalUpdate/followingUpdates`,
		Input:      input,
		Output:     &output,
		Context:    ctx,
	}
	return output, req.send(u.jike)
}
