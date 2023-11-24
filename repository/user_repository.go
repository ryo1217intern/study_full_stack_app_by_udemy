package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

//usecase層にて使用するメソッドをinterfaceとして定義する
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

//userRepositoryの構造体でdbという名前のgorm.DBのポインター型を定義する
type userRepository struct {
	db *gorm.DB
}

//コンストラクター. 引数はuserRepositoryで定義した型.　戻り値はIUserRepositoryで定義した
func NewUserRespository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

//GetUserByEmailの実際の処理を記述.emailの情報からUserのロウを取得する.
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	//emailは一意なので
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}