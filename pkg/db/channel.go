package db

import (
	"TeleChan/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableChannel - database table 'channels'
	TableChannel = "channels"
	// TableUserChannel - database table 'users_channels'
	TableUserChannel = "users_channels"
)

// FindChannelByName queries the user by 'name'
func (s *Handler) FindChannelByName(name string) (*model.Channel, error) {
	u := model.Channel{}

	// query the user rows by email
	res := s.db.Table(TableChannel).Where("name = ?", name).First(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &u, nil
}

// FindChannelByID queries the user by unique key 'channel_id'
func (s *Handler) FindChannelByUK(channelID string) (*model.Channel, error) {
	u := model.Channel{}

	// query the user rows by user id
	res := s.db.Table(TableChannel).Where("channel_id = ?", channelID).First(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &u, nil
}

// FindChannelsByUser queries the user's owned channels by 'user_id'
func (s *Handler) FindChannelsByUser(userID string) ([]model.Channel, error) {
	u := []model.Channel{}

	// query the channel rows by user id
	res := s.db.Table(TableChannel).Where("user_id = ?", userID).Find(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return u, nil
		}
		return nil, res.Error
	}

	return u, nil
}

// CreateChannel insert the channel into database
func (s *Handler) CreateChannel(u *model.Channel) error {
	return s.db.Table(TableChannel).Create(u).Error
}

// CreateUserChannel insert the user's joining channel into database
func (s *Handler) CreateUserChannel(u *model.UserChannel) error {
	return s.db.Table(TableUserChannel).Create(u).Error
}

// FindUserChannelByUK queries the user
func (s *Handler) FindUserChannelByUK(userID string, channelID string) (*model.UserChannel, error) {
	u := model.UserChannel{}

	// query the user rows by user id
	res := s.db.Table(TableUserChannel).
		Where("user_id = ?", userID).
		Where("channel_id = ?", channelID).
		First(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &u, nil
}

// FindUserChannelsByUser queries the user's owned and joined channels by 'user_id'
func (s *Handler) FindUserChannelsByUser(userID string) ([]model.UserChannel, error) {
	u := []model.UserChannel{}

	// query the user rows by user id
	res := s.db.Table(TableUserChannel).Where("user_id = ?", userID).Find(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return u, nil
		}
		return nil, res.Error
	}

	return u, nil
}
