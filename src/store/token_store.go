package store

import (
	"context"

	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/pairswap-be/src/model"
	"gorm.io/gorm"
)

type TokenStore struct {
	db *gorm.DB
}

func NewTokenStore(db *gorm.DB) *TokenStore {
	return &TokenStore{db: db}
}

func (s *TokenStore) CreateToken(ctx context.Context, token *model.Token) error {
	if token == nil {
		log.Warn("token is nil")
		return nil
	}

	if err := s.db.WithContext(ctx).Create(token).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *TokenStore) GetAll(ctx context.Context) ([]*model.Token, error) {
	tokens := []*model.Token{}
	err := s.db.WithContext(ctx).
		Model(model.Token{}).
		Order("id DESC").Find(&tokens).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return tokens, nil
}

func (s *TokenStore) GetById(ctx context.Context, id int) (*model.Token, error) {
	token := &model.Token{}
	if err := s.db.WithContext(ctx).Model(model.Token{}).Where("id = ?", id).First(&token).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return token, nil
}
