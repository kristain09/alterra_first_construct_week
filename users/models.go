package users

import (
	"database/sql"
	"log"
)

type UsersModels struct {
	conn *sql.DB
}

func (um *UsersModels) SetConnUsersModels(db *sql.DB) {
	um.conn = db
}

func (um UsersModels) GetUserByID(id int, password string) (result *Users, err error) {
	result = &Users{}
	resultQuery := um.conn.QueryRow("SELECT id, username FROM users WHERE id = ? AND password = ? AND deleted_at IS NULL", id, password)
	if resultQuery.Err() != nil {
		log.Println("Look up data error", resultQuery.Err().Error())
		return nil, resultQuery.Err()
	}
	err = resultQuery.Scan(&result.ID, &result.UserName)
	if err != nil {
		log.Println("Scan result error", err.Error())
		return nil, err
	}
	return result, nil
}

func (um *UsersModels) InsertDataToUsers(newUser Users) error {
	result, err := um.conn.Exec("Insert into users (username, password) values (?,?)", newUser.UserName, newUser.password)

	if err != nil {
		log.Println("Error melakukan insert data", err.Error())
		return err
	}
	resultAff, err := result.RowsAffected()

	if err != nil {
		log.Println("error after inserting data", err.Error())
		return err
	}

	if resultAff <= 0 {
		log.Println("Tidak ada perubahan yang dilakukan", err.Error())
		return err
	}
	return nil
}

func (um *UsersModels) DeleteDataFromUsers(userID int) error {
	result, err := um.conn.Exec("UPDATE users SET deleted_at = current_timestamp() WHERE id = ?", userID)
	if err != nil {
		log.Print("error before delete")
		return err
	}

	resultAff, err := result.RowsAffected()
	if err != nil {
		log.Print("Error after delete")
		return err
	}

	if resultAff <= 0 {
		log.Println("terjadi kesalahan dalam melakukan delete pengguna!\nHubungi pihak IT")
	}

	return nil
}
