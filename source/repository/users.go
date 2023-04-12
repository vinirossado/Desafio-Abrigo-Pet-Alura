package repository

import "abrigos/source/domain/entities"

func FindUsers() {

}

func FindUserByUsername(username string, tx ...*TransactionalOperation) (*entities.User, error) {
	user := &entities.User{}
	return user, WithTransaction(tx).Where("username = ?", username).First(user).Error
}

func FindUserById() {

}

func CreateUser(user *entities.User, tx ...*TransactionalOperation) error {
	return WithTransaction(tx).Create(user).Error
}

func UpdateUser() {

}

func DeleteUser() {

}
