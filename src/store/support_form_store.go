package store

import (
	"context"

	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/pairswap-be/src/model"
	"gorm.io/gorm"
)

type SupportFormStore struct {
	db *gorm.DB
}

func NewSupportFormStore(db *gorm.DB) *SupportFormStore {
	return &SupportFormStore{db: db}
}

func (s *SupportFormStore) CreateSupportForm(ctx context.Context, supportForm *model.SupportForm) error {
	if supportForm == nil {
		log.Warn("token is nil")
		return nil
	}

	if err := s.db.WithContext(ctx).Create(supportForm).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *SupportFormStore) GetAll(ctx context.Context) ([]*model.SupportForm, error) {
	supportForms := []*model.SupportForm{}
	err := s.db.WithContext(ctx).
		Model(model.SupportForm{}).
		Order("id DESC").Find(&supportForms).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return supportForms, nil
}

func (s *SupportFormStore) GetById(ctx context.Context, id int) (*model.SupportForm, error) {
	supportForm := &model.SupportForm{}
	if err := s.db.WithContext(ctx).Model(model.SupportForm{}).Where("id = ?", id).First(&supportForm).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return supportForm, nil
}
