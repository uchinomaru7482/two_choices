package persistence

import (
	"two_choices/pkg/rdb"
	"two_choices/src/domain"
	"two_choices/src/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

// Save - 保存
func (up userPersistence) Save(scope *rdb.SessionScope, user *domain.User) error {
	if err := scope.Conn.Model(user).Save(user).Error; err != nil {
		return err
	}
	return nil
}

// GetByUserID - 取得/ユーザID
func (up userPersistence) GetByUserID(scope *rdb.SessionScope, userID uint64) (*domain.User, error) {
	// domainのメソッドとして定義すると、domain層に技術的関心を書く事になるのでNG？
	domUser := &domain.User{}
	if err := scope.Conn.Where("id=?", userID).First(domUser).Error; err != nil {
		return nil, err
	}
	return domUser, nil
}

// GetByEmail - 取得/UID
func (up userPersistence) GetByEmail(scope *rdb.SessionScope, email string) (*domain.User, error) {
	domUser := &domain.User{}
	if err := scope.Conn.Where("email=?", email).First(domUser).Error; err != nil {
		return nil, err
	}
	return domUser, nil
}
