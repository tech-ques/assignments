package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadUsersFromDisk(filelocation string) ([]User, error) {
	jsonfile, err := os.Open(filelocation)
	if err != nil {
		fmt.Println("Error in ReadJSONData: while opening file, err: ", err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		fmt.Println("Error in ReadJSONData: while reading file, err: ", err)
		return nil, err
	}

	users := UserList{}
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		fmt.Println("Error in ReadJSONData: while unmarshaling, err: ", err)
		return nil, err
	}

	return users.Users, nil
}
