package db

import (
	"fmt"
	"github.com/overlordyy/MiniIam/api/account"
)

func SaveAccount(acc account.Account) error {
	result, err := dbClient.NamedExec(`INSERT INTO accounts () values ()`, acc)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	fmt.Println(result)
	return err
}
