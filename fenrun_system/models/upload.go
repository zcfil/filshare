package models

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"mime/multipart"
	"strings"
	orm "xAdmin/database"
	log "xAdmin/logrus"
	"xAdmin/pkg/eth"
	"xAdmin/utils"
)

type Addr struct {
	Id    int64  `json:"id"`
	Addr  string `json:"addr"`
	State int    `json:"state"`
}

//导入数据
func (a *Addr) UploadAddress(file multipart.File, Size int64) error {
	buf := make([]byte, Size)
	n, _ := file.Read(buf)

	xf, _ := xlsx.OpenBinary(buf[:n])
	sql := `insert into addr(addr)values`
	for _, sheet := range xf.Sheets {
		for _, row := range sheet.Rows {
			for i, cell := range row.Cells {
				if i == 0 {
					addr := cell.String()
					if eth.ValidAddress(addr) {
						return errors.New("地址" + addr + "不合法！")
					}
					sql += `('` + addr + `'),`
				}
			}
		}
	}
	if err := orm.Eloquent.Exec(sql[:len(sql)-1]).Error; err != nil {
		str := strings.Split(err.Error(), "'")
		return errors.New("已存在地址:" + str[1])
	}
	return nil
}

// 写入数据库图片

func (a *Addr) UploadFile(url string, userid int, adminid int, size string, filename string) error {
	sqlx := `INSERT INTO  kdsystem_upload ( url_id, users_id, admin_id,create_time,file_size,is_del,filename)  VALUES ("%s","%d","%d","%d",%s,%d,"%s")`
	sqlx = fmt.Sprintf(sqlx, url, userid, adminid, utils.TimeUNix(), size, 1, filename)
	log.Info(sqlx)
	return orm.Eloquent.Debug().Exec(sqlx).Error
}

// 查询当前用户的上传的文件数量

func (a *Addr) GetUploadFile(username string) (ret interface{}, err error) {
	type FileList struct {
		Url_Id      string `gorm:"column:url" json:"url"`
		Create_Time string `gorm:"column:datetime" json:"datetime"`
		Admin_Id    string `gorm:"column:username" json:"username"`
		FileName    string `gorm:"column:filename" json:"filename"`
	}

	sql := ` SELECT k.url_id                                            AS url,
                   (SELECT  FROM_UNIXTIME(k.create_time, '%Y-%m-%d %H:%i:%s') ) AS datetime,
                    u.username                                          AS username,
					k.filename                                          AS filename
             FROM (SELECT id FROM customer WHERE name = '` + username + `' ) AS a
                      INNER JOIN kdsystem_upload AS k
                      LEFT JOIN sys_user AS u ON k.admin_id = u.user_id
             WHERE k.users_id = a.id  AND  k.is_del =1;

`
	data := make([]FileList, 0)

	if err = orm.Eloquent.Debug().Raw(sql).Scan(&data).Error; err != nil {
		return
	}

	return data, nil
}

func MaxLimit(userId int) (ret int, err error) {
	sqlx := `SELECT  COUNT(users_id) AS total  FROM  kdsystem_upload WHERE  users_id= %d AND is_del=1`
	sqlx = fmt.Sprintf(sqlx, userId)

	type Total struct {
		Total int `gorm:"column:total"json:"total"`
	}
	data := make([]Total, 0)

	if err = orm.Eloquent.Debug().Raw(sqlx).Scan(&data).Error; err != nil {
		return
	}
	ret = data[0].Total
	return ret, err
}

func (a *Addr) DelStatusFile(filename string) (err error) {
	sqlx := `UPDATE kdsystem_upload SET is_del=0 WHERE filename ='` + filename + `'`
	err = orm.Eloquent.Debug().Exec(sqlx).Error
	return
}
