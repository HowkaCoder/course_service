package usecase

import (
	"course_service/internal/app/entity"
	"course_service/internal/app/repository"
)

type ActUsecase interface {
	GetActsByAnswerID(id uint) ([]entity.Act, error)
	CreateAct(act *entity.Act) error
	UpdateAct(act *entity.Act, id uint) error
	DeleteAct(id uint) error
}

type actUsecase struct {
	actRepository repository.ActRepository
}

func NewActUsecase(actRepository repository.ActRepository) *actUsecase {
	return &actUsecase{actRepository: actRepository}
}

func (u *actUsecase) GetActsByAnswerID(id uint) ([]entity.Act, error) {
	return u.actRepository.GetActsByAnswerID(id)
}

func (u *actUsecase) CreateAct(act *entity.Act) error {
	return u.actRepository.CreateAct(act)
}

func (u *actUsecase) UpdateAct(act *entity.Act, id uint) error {
	return u.actRepository.UpdateAct(act, id)
}

func (u *actUsecase) DeleteAct(id uint) error {
	return u.actRepository.DeleteAct(id)
}
