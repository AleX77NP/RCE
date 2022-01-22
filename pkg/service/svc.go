package service

import (
	"log"
	"os"

	exec "rce.amopdev/m/v2/pkg/executor"
)

type RemoteService struct {
	lang string
	code string
	file string
}

func NewRemoteService(lang, code string) *RemoteService {
	return &RemoteService{
		lang: lang,
		code: code,
		file: "",
	}
}

func (svc *RemoteService) SetSettings(lang, code string) {
	svc.lang = lang
	svc.code = code
}

// File name to be dynamic based on lang
func (svc *RemoteService) RunCode(code string) (error, string) {
	f, err := svc.createFileAndAddCode("aco", code)
	if err != nil {
		panic(err)
	}
	fileName := "code/" + f
	codePath := getCodeDir() + "/code"
	executor := exec.NewExecutor(svc.lang, svc.code, fileName, codePath)
	return executor.ExecuteCode()
}

// Dynamic for each user should it be !
func (svc *RemoteService) createFileAndAddCode(user string, code string) (string, error) {
	fName := getCodeDir() + "/code/" + user + getExtension("python")
	file, err := os.Create(fName)
	defer file.Close()
	if err != nil {
		return "", err
	}
	_, err2 := file.WriteString(code)
	if err2 != nil {
		return "", err2
	}
	svc.file = fName
	return user + getExtension("python"), err
}

func (svc *RemoteService) RemoveFile() {
	err := os.Remove(svc.file)
	if err != nil {
		log.Fatal(err)
	}
}

// make these dynamic for other language
func getExtension(lang string) string {
	return ".py"
}

// change this 
func getCodeDir() string {
	return "/Users/aleksandar77np/Desktop/rce/backend"
}
