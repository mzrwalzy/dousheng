package backend

import (
	"bytedance-douyin/global"
	r "bytedance-douyin/model/data"
	"bytedance-douyin/utils"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"time"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: vedio
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:33
 */
type VideoApi struct{}

func (api *VideoApi) PostVideo(c *gin.Context) {
	file, err := c.FormFile("data")

	if err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	// todo 生成文件名
	path := global.GVA_CONFIG.File.VideoOutput
	videoName := utils.GenerateFilename(file.Filename)
	fn := path + videoName

	if err := c.SaveUploadedFile(file, fn); err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	reader := utils.ReadFrameAsJpeg(fn)
	img, err := imaging.Decode(reader)
	if err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	imagePath := global.GVA_CONFIG.File.ImageOutput

	// replace .mp4 to .jpg
	imageName := videoName[:len(videoName)-4] + ".jpg"
	if err := imaging.Save(img, imagePath+imageName); err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	r.OkWithMessage(c, "上传成功")
}

// 视频流接口
func (api *VideoApi) VideoFeed(c *gin.Context) {
	videoList := make([]*r.Video, 1, 1)
	user := r.Author{Id: 1, Name: "charon", FollowCount: 0, FollowerCount: 1, IsFollow: false}
	videoList[0] = &r.Video{Id: 1, Author: &user, PlayUrl: "http://220.243.147.162:8080/videos/sss.mp4", CoverUrl: "http://220.243.147.162:8080/images/1.jpg", FavoriteCount: 0, CommentCount: 0, IsFavorite: false}
	data := r.FeedData{NextTime: time.Now().Unix(), VideoList: videoList}
	r.OkWithData(c, data)
}

func (api *VideoApi) VideoList(c *gin.Context) {

}
