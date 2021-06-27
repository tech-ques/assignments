package pkg

import "fmt"

type Notifications struct {
}

func (n Notifications) Send() (bool, error) {
	return true, nil
}

type Sms struct {
	Notifications
}

func (s Sms) Send(u User) (bool, error) {
	if u.ID%100 == 0 {
		Retry(u)
	} else {
		fmt.Println("SMS sent to user: ", u.FirstName)
	}
	return true, nil
}

type Call struct {
	Notifications
}

func (s Call) Send(u User) (bool, error) {
	if u.ID%100 == 0 {
		Retry(u)
	} else {
		fmt.Println("Call done to user: ", u.FirstName)
	}
	return true, nil
}

type Email struct {
	Notifications
}

func (e Email) Send(u User) (bool, error) {
	if u.ID%100 == 0 {
		Retry(u)
	} else {
		fmt.Println("Email sent to user: ", u.FirstName)
	}
	return true, nil
}

func Notify(u User) {
	switch u.CommPreference {
	case "sms":
		s := Sms{}
		s.Send(u)
	case "call":
		c := Call{}
		c.Send(u)
	case "email":
		e := Email{}
		e.Send(u)
	}
}
