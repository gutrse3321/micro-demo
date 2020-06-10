package oauth

import (
	"fmt"
	"github.com/RangelReale/osin"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/10 11:11
 * @Title: store
 * --- --- ---
 * @Desc:
 */

//implement osin.Storage
type TestStorage struct {
	clients   map[string]osin.Client
	authorize map[string]*osin.AuthorizeData
	access    map[string]*osin.AccessData
	refresh   map[string]string
}

func NewTestStorage() (r *TestStorage) {
	r = &TestStorage{
		clients:   make(map[string]osin.Client),
		authorize: make(map[string]*osin.AuthorizeData),
		access:    make(map[string]*osin.AccessData),
		refresh:   make(map[string]string),
	}

	r.clients["1234"] = &osin.DefaultClient{
		Id:          "1234",
		Secret:      "aabbccdd",
		RedirectUri: "http://localhost:14000/app/auth/code",
	}

	return
}

func (s *TestStorage) Clone() osin.Storage {
	return s
}

func (s *TestStorage) Close() {
}

//客户端 setter
func (s *TestStorage) SetClient(id string, client osin.Client) error {
	fmt.Println("SetClient:", id)
	s.clients[id] = client
	return nil
}

//客户端 getter
func (s *TestStorage) GetClient(id string) (osin.Client, error) {
	fmt.Println("GetClient:", id)
	if c, ok := s.clients[id]; ok {
		return c, nil
	}
	return nil, osin.ErrNotFound
}

//授权信息save
func (s *TestStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	fmt.Println("SaveAuthorize:", data.Code)
	s.authorize[data.Code] = data
	return nil
}

//授权信息加载
func (s *TestStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	fmt.Println("LoadAuthorize:", code)
	if d, ok := s.authorize[code]; ok {
		return d, nil
	}
	return nil, osin.ErrNotFound
}

//删除授权信息
func (s *TestStorage) RemoveAuthorize(code string) error {
	fmt.Println("RemoveAuthorize:", code)
	delete(s.authorize, code)
	return nil
}

//保存access信息
func (s *TestStorage) SaveAccess(data *osin.AccessData) error {
	fmt.Println("SaveAccess:", data.AccessToken)
	s.access[data.AccessToken] = data
	if data.RefreshToken != "" {
		s.refresh[data.RefreshToken] = data.AccessToken
	}
	return nil
}

//加载access信息
func (s *TestStorage) LoadAccess(token string) (*osin.AccessData, error) {
	fmt.Println("LoadAccess:", token)
	if d, ok := s.access[token]; ok {
		return d, nil
	}
	return nil, osin.ErrNotFound
}

//移除access信息
func (s *TestStorage) RemoveAccess(token string) error {
	fmt.Println("RemoveAccess:", token)
	delete(s.access, token)
	return nil
}

//刷新access token
func (s *TestStorage) LoadRefresh(token string) (*osin.AccessData, error) {
	fmt.Println("LoadRefresh:", token)
	if d, ok := s.refresh[token]; ok {
		return s.LoadAccess(d)
	}
	return nil, osin.ErrNotFound
}

//移除刷新token
func (s *TestStorage) RemoveRefresh(token string) error {
	fmt.Println("RemoveRefresh:", token)
	delete(s.refresh, token)
	return nil
}
