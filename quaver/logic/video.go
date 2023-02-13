package logic

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"quaver/dao/mysql"
	"quaver/models"
	"strings"
	"time"
)

// Feed 视频流接口
func Feed(latestTime string, currentUserID ...int64) (videoList *[]models.Video, err error) {
	//if len(latestTime) != 0 {
	//	latest, err := strconv.ParseInt(latestTime, 10, 64)
	//	if err != nil {
	//		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
	//		return nil, err
	//	}
	//	latestTime = time.Unix(latest, 0).Format("2006-01-02 15:04:05")
	//} else {
	//	latestTime = time.Now().Format("2006-01-02 15:04:05") // 本地当前时间
	//}
	// 前端app传来的时间戳太大了，这里就直接用实时的时间，不用上面的代码了，postman测试的时候可以用上面的代码
	latestTime = time.Now().Format("2006-01-02 15:04:05")
	if currentUserID != nil {
		return mysql.Feed(latestTime, currentUserID[0])
	}
	return mysql.Feed(latestTime)
}

// PublishList 发布列表
func PublishList(currentUserID, userID int64) (publishList *[]models.Video, err error) {
	if currentUserID == userID {
		return mysql.PublishList(currentUserID)
	}
	return mysql.PublishList(userID, currentUserID)
}

// Publish 发布视频
func Publish(currentUserID int64, title, filePath, finalName string) (err error) {
	imgNames := strings.Split(finalName, ".")
	// 生成封面
	if _, err = getSnapshot(filePath, "./public/"+imgNames[0], 1); err != nil {
		return
	}
	video := &models.Video{
		UserID:        currentUserID,
		Title:         title,
		CreateTime:    time.Now().Format("2006-01-02 15:04:05"), // 本地当前时间
		PlayUrl:       "/public/" + finalName,
		CoverUrl:      "/public/" + imgNames[0] + ".jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	// 存入数据库
	return mysql.Publish(video)
}

// getSnapshot 生成封面
func getSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {

	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".jpg")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}
