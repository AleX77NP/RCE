package service

import (
	exec "rce.amopdev/m/v2/pkg/executor"
)

type RemoteService struct {
	lang string
	code string
}

func NewRemoteService(lang, code string) *RemoteService {
	return &RemoteService{
		lang,
		code,
	}
}

func (svc *RemoteService) SetSettings(lang, code string) {
	svc.lang = lang
	svc.code = code
}

// File name to be dynamic based on lang
func (svc *RemoteService) RunCode() (error, string) {
	executor := exec.NewExecutor(svc.lang, svc.code, "code/code.py", "/Users/aleksandar77np/Desktop/rce/backend/code")
	return executor.ExecuteCode();
}