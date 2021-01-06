package wxgameod

import (
	"fmt"
	"github.com/birjemin/wxgameod/utils"
	"log"
)

// RemoveUserStorage model
type RemoveUserStorage struct {
	AccessToken string
	SessionKey  string
	OpenID      string
	SigMethod   string
	Key         string
	HTTPRequest *utils.HTTPClient
}

// RespRemoveUserStorage response
type RespRemoveUserStorage struct {
	CommonError
}

// RemoveUserStorage remove user storage
func (r *RemoveUserStorage) RemoveUserStorage() (*RespRemoveUserStorage, error) {
	return r.doRemoveUserStorage(wechatDomain)
}

// doRemoveUserStorage handle this action
func (r *RemoveUserStorage) doRemoveUserStorage(domain string) (*RespRemoveUserStorage, error) {

	url := fmt.Sprintf("%s%s?access_token=%s&signature=%s&openid=%s&sig_method=%s",
		domain,
		r.getRemoveUserStorageURI(),
		r.AccessToken,
		r.generateSignature(),
		r.OpenID,
		r.getSigMethod())

	var resp = new(RespRemoveUserStorage)

	if err := r.HTTPRequest.HTTPPostJSON(url, r.Key); err != nil {
		log.Println("[remove_user_storage]do, post failed", err)
		return resp, err
	}

	if err := r.HTTPRequest.GetResponseJSON(resp); err != nil {
		log.Println("[remove_user_storage]do, response json failed", err)
		return resp, err
	}
	return resp, nil
}

// getSigMethod get sig method
func (r *RemoveUserStorage) getSigMethod() string {
	var sigMethod = "hmac_sha256"
	if r.SigMethod != "" {
		sigMethod = r.SigMethod
	}
	return sigMethod
}

// generateSignature cal signature
func (r *RemoveUserStorage) generateSignature() string {
	return utils.GenerateSha256(r.SessionKey, r.Key)
}

// getInteractiveDataUrl get set_interactive_data URI
func (r *RemoveUserStorage) getRemoveUserStorageURI() string {
	return removeUserStorageURI
}
