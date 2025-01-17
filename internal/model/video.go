package model

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      uint      `json:"author_id"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	Title         string    `json:"title"`
	PublishDate   time.Time `json:"publish_date"`
}

func (v Video) TableName() string { return "douyin_video" }

// 新建视频
func (v Video) Create(db *gorm.DB) error {
	return db.Create(&v).Error
}

// 返回用户发布的所有视频id
func (v Video) ListVideoByUserId(db *gorm.DB) ([]Video, error) {
	var video []Video
	err := db.Select("id, play_url, cover_url, favorite_count, comment_count, title").Where("author_id = ?",
		v.AuthorId).Find(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

// 更新视频点赞数
func (v Video) UpdateFavoriteCnt(db *gorm.DB) error {
	return db.Model(&v).Update("favorite_count", v.FavoriteCount).Error
}

// 更新评论数
func (v Video) UpdateCommentCnt(db *gorm.DB) error {
	return db.Model(&v).Update("comment_count", v.CommentCount).Error
}

// QueryVideoById 根据videoId查询video信息
func (v Video) QueryVideoById(videoID uint, db *gorm.DB) (Video, error) {
	var video Video
	err := db.Where("id = ?", videoID).Find(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

// QueryBatchVdieoById  根据videoId批量查询video信息
func (v Video) QueryBatchVdieoById(videoIdList []uint, db *gorm.DB) ([]Video, error) {
	var videos []Video
	err := db.Select("id, play_url, cover_url, favorite_count, comment_count, title").Where("id IN ?", videoIdList).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// QueryFavorCntById 根据Id查询点赞数量
func (v Video) QueryFavorCntById(db *gorm.DB) (int64, error) {
	var cnt int64
	err := db.Model(&v).Select("favorite_count").Find(&cnt).Error
	return cnt, err
}
