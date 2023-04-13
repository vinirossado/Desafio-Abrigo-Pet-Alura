package repository

import "abrigos/source/domain/entities"

func FindUsers(tx ...*TransactionalOperation) ([]entities.User, error) {
	users := []entities.User{}
	return users, WithTransaction(tx).Find(&users).Error
}

func FindUserByUsername(username string, tx ...*TransactionalOperation) (*entities.User, error) {
	user := &entities.User{}
	return user, WithTransaction(tx).Where("username = ?", username).First(user).Error
}

func FindUserById(id int, tx ...*TransactionalOperation) (*entities.User, error) {
	user := &entities.User{}
	return user, WithTransaction(tx).Where("id = ?", id).First(user).Error
}

func CreateUser(user *entities.User, tx ...*TransactionalOperation) error {
	return WithTransaction(tx).Create(user).Error
}

func UpdateUser() {

}

func DeleteUser() {

}
