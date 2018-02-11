package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	pb "github.com/while-loop/levit/users/proto"
)

const (
	table = "users"
)

type mySqlRepo struct {
	db *gorm.DB
}

func NewMySql(db *gorm.DB) *mySqlRepo {
	m := &mySqlRepo{db: db.Table(table)}
	m.db.AutoMigrate(&pb.User{})
	return m
}

func (r *mySqlRepo) Create(user *pb.User) (*pb.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, errors.Wrap(err, "unable to create user")
	}

	return user, nil
}

func (r *mySqlRepo) Get(id uint64) (*pb.User, error) {
	var user pb.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, errors.Wrapf(err, "unable to get user %d", id)
	}

	return &user, nil
}

func (r *mySqlRepo) GetAll(ids ...uint64) ([]*pb.User, error) {
	var users []*pb.User
	if err := r.db.Where("id IN (?)", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *mySqlRepo) Update(user *pb.User) (*pb.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, errors.Wrapf(err, "Unable to update user %d", user.Id)
	}

	return user, nil
}

func (r *mySqlRepo) Delete(id uint64) error {
	if err := r.db.Delete(&pb.User{Id: id}).Error; err != nil {
		return err
	}

	return nil
}
