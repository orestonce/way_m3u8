package work

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/orestonce/m3u8d"
	"github.com/orestonce/m3u8d/m3u8dcpp"
	log "github.com/sirupsen/logrus"
	"gom3u8/conf"
	"gom3u8/data"
	"gom3u8/model"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	StateReady = 1
	StateEnd   = 2
	StateError = 3
)

func WorkSave(url string, fileName string, save_dir string) (err error) {
	url = strings.Replace(url, " ", "", -1)
	fileName = strings.Replace(fileName, " ", "", -1)
	save_dir = strings.Replace(save_dir, " ", "", -1)

	if fileName == "" {
		fileName, err = extractFilenameFromURL(url)
		if err != nil {
			return err
		}
	}
	log.Info("url:", url)
	log.Info("fileName:", fileName)
	log.Info("save_dir:", save_dir)
	db := data.GetDbInstance()
	temp := md5.Sum([]byte(url))
	id := hex.EncodeToString(temp[:])
	ok := db.Work_D().Select().Where_ID().Equal(id).MustRun_Exist()
	if ok == false {
		db.Work_D().MustInsert(model.Work_D{
			ID:         id,
			Name:       fileName,
			Url:        url,
			SaveDir:    save_dir,
			State:      StateReady,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
	} else {
		db.Work_D().Update().Where_ID().Equal(id).
			Set_Name(fileName).
			Set_Url(url).
			Set_SaveDir(save_dir).
			Set_State(StateReady).
			Set_UpdateTime(time.Now()).MustRun()
	}
	return
}

func GetNotWorkingWork() *model.Work_D {
	instance, ok := data.GetDbInstance().Work_D().Select().Where_State().Equal(StateReady).OrderBy_UpdateTime().ASC().MustRun_ResultOne2()
	if ok == false {
		return nil
	}
	return &instance
}

func WorkEnd(id string) {
	data.GetDbInstance().Work_D().Update().Where_ID().Equal(id).Set_State(StateEnd).MustRun()
}
func WorkError(id string, err_msg string) {
	log.Warn(id, " download err:", err_msg)
	data.GetDbInstance().Work_D().Update().Where_ID().Equal(id).Set_Info(err_msg).Set_State(StateError).MustRun()
}

func extractFilenameFromURL(urlStr string) (string, error) {
	// 解析URL
	_, err := url.ParseRequestURI(urlStr)
	// 如果没有错误，则认为URL是合法的
	if err != nil {
		return "", err
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	// 使用path/filepath包来获取路径的最后一个元素，这将是文件名
	filename := filepath.Base(parsedURL.Path)
	if len(filename) > 5 {
		filename = filename[:len(filename)-5]
	}
	return filename, nil
}

func DownloadFromCmd(req m3u8d.StartDownload_Req) error {
	req.ProgressBarShow = true
	log.Info("DownloadFromCmd M3u8Url: ", req.M3u8Url)
	errMsg := m3u8dcpp.StartDownload(req)
	if errMsg != "" {
		log.Error(errMsg)
		return errors.New(errMsg)
	}

	resp := m3u8dcpp.WaitDownloadFinish()

	if resp.ErrMsg != "" {
		log.Error(resp.ErrMsg)
		return errors.New(errMsg)
	}
	if resp.IsSkipped {
		log.Warn("已经下载过了: " + resp.SaveFileTo)
		return errors.New(errMsg)
	}
	if resp.SaveFileTo == "" {
		log.Info("下载成功.")
		return errors.New(errMsg)
	}
	log.Info("下载成功, 保存路径", resp.SaveFileTo)
	return nil
}

func Working() {
	workListMaxNu := conf.ConfMap.Init.WorkMax
	var workList []Worker
	for i := 0; i < workListMaxNu; i++ {
		workList = append(workList, NewWorker())
	}
	for {
		//过滤已处理的任务
		for {
			for _, worker := range workList {
				if worker.State {
					readywork := GetNotWorkingWork()
					if readywork == nil {
						time.Sleep(5 * time.Second)
						continue
					}
					worker.State = false
					go worker.Start(readywork)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}
}

type Worker struct {
	State bool
}

func NewWorker() Worker {
	return Worker{
		State: true,
	}
}
func (w Worker) Start(readywork *model.Work_D) {
	workOnce(readywork)
	w.State = true
}

func workOnce(readywork *model.Work_D) {
	log.Info("readywork:", readywork)
	_, err := os.Stat(readywork.SaveDir)
	if err != nil {
		os.MkdirAll(readywork.SaveDir, os.ModePerm)
	}
	req := m3u8d.StartDownload_Req{
		M3u8Url:                  readywork.Url,
		Insecure:                 true,
		SaveDir:                  readywork.SaveDir,
		FileName:                 readywork.Name,
		SkipTsExpr:               "",
		SetProxy:                 "",
		HeaderMap:                nil,
		SkipRemoveTs:             false,
		ProgressBarShow:          false,
		ThreadCount:              8,
		SkipCacheCheck:           false,
		SkipMergeTs:              false,
		Skip_EXT_X_DISCONTINUITY: false,
		DebugLog:                 false,
	}
	err = DownloadFromCmd(req)

	if err != nil {
		WorkError(readywork.ID, err.Error())
		return
	}

	WorkEnd(readywork.ID)
}
