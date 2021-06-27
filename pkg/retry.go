package pkg

import "fmt"

var (
	retryList []AttemptList
)

const (
	ATTEMPT int = 5
)

type AttemptList struct {
	AttemptUser  User
	RetryAttempt int
}

func Retry(user User) bool {
	al := AttemptList{}
	al.AttemptUser = user
	al.RetryAttempt = 1
	retryList = append(retryList, al)
	return true
}

func RetryAll() bool {
	userlist := []User{}
	for i := 0; i < len(retryList); i++ {
		if retryList[i].RetryAttempt == ATTEMPT {
			fmt.Println("removing:", retryList[i].AttemptUser.FirstName)
			userlist = append(userlist[:i], userlist[i+1])
			continue
		}
		retryList[i].RetryAttempt += 1
		userlist = append(userlist, retryList[i].AttemptUser)
	}
	fmt.Println("****** Retrying failed attempts ******")
	StartProcess(userlist)
	return true
}
