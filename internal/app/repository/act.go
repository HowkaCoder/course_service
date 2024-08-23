package repository

import (
	"course_service/internal/app/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ActRepository interface {
	GetActsByAnswerID(id uint) ([]entity.Act, error)
	CreateAct(act *entity.Act) error
	UpdateAct(act *entity.Act, id uint) error
	DeleteAct(id uint) error
}

type actRepostiory struct {
	rdb *redis.Client
	db  *gorm.DB
}

func NewActRepository(db *gorm.DB, rdb *redis.Client) *actRepostiory {
	return &actRepostiory{db: db, rdb: rdb}
}

func (r *actRepostiory) GetActsByAnswerID(id uint) ([]entity.Act, error) {
	var acts []entity.Act
	err := r.db.Order("position asc").Where("answer_id = ?", id).Find(&acts).Error
	return acts, err
}

func (r *actRepostiory) CreateAct(act *entity.Act) error {
	return r.db.Create(act).Error
}

func (r *actRepostiory) UpdateAct(act *entity.Act, id uint) error {
	var eAct entity.Act
	if err := r.db.First(&eAct, id).Error; err != nil {
		return err
	}
	if act.ActText != "" {
		eAct.ActText = act.ActText
	}
	if act.ImageUrl != "" {
		eAct.ImageUrl = act.ImageUrl
	}
	if act.Position != 0 {
		eAct.Position = act.Position
	}
	if act.AnswerID != 0 {
		eAct.AnswerID = act.AnswerID
	}
	return r.db.Save(act).Error
}

func (r *actRepostiory) DeleteAct(id uint) error {
	return r.db.Delete(&entity.Act{}, id).Error
}
