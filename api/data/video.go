package data

type FeedData struct {
	NextTime  int64    `json:"next_time"`
	VideoList []*Video `json:"video_list" binding:"required"`
}

type PublishData struct {
	VideoList []*Video `json:"video_list" binding:"required"`
}

type FavoriteData struct {
	VideoList []*Video `json:"video_list" binding:"required"`
}

type Video struct {
	Id            int64   `json:"id"`
	Author        *Author `json:"author"`
	PlayUrl       string  `json:"play_url"`
	CoverUrl      string  `json:"cover_url"`
	FavoriteCount int64   `json:"favorite_count"`
	CommentCount  int64   `json:"comment_count"`
	IsFavorite    bool    `json:"is_favorite"`
}

type Author struct {
	Id            int64  `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	FollowCount   int64  `json:"follow_count" binding:"required"`
	FollowerCount int64  `json:"follower_count" binding:"required"`
	IsFollow      bool   `json:"is_follow" binding:"required"`
}
