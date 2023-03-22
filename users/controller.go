package users

import (
	"fmt"
)

type UsersController struct {
	UsersModels UsersModels
}

func (uc *UsersController) SetConnectModels(um UsersModels) {
	uc.UsersModels = um
}

func (uc UsersController) Login() (user *Users, err error) {
	var (
		id       int
		password string
	)

	fmt.Println("Input your id!")
	_, err = fmt.Scanln(&id)
	if err != nil {
		return &Users{}, err
	}

	fmt.Println("Input your password!")
	_, err = fmt.Scanln(&password)
	if err != nil {
		return &Users{}, err
	}

	user, err = uc.UsersModels.GetUserByID(id, password)
	if err != nil {
		return &Users{}, err
	}

	return
}

func (uc *UsersController) Register() error {

	user := Users{}
	fmt.Println("Masukkan username cashier!")
	fmt.Scanln(&user.UserName)
	fmt.Println("Masukkan password cashier!")
	fmt.Scanln(&user.password)

	err := uc.UsersModels.InsertDataToUsers(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UsersController) DeleteUser() error {
	user := Users{}
	fmt.Println("Masukakn id user yang akan di delete!")
	fmt.Scanln(&user.ID)

	err := uc.UsersModels.DeleteDataFromUsers(user.ID)
	if err != nil {
		return err
	}
	return nil
}
