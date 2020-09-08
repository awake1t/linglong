package models

type Setting struct {
	*Model
	Ip             string `json:"ip"`
	LoginWord      string `json:"login_word"`
	LoginUrl       string `json:"login_url"`
	MasscanThred   int    `json:"masscan_thred"`
	MasscanDeltime int    `json:"masscan_deltime"`
	MasscanIp      string `json:"masscan_ip"`
	MasscanPort    string `json:"masscan_port"`
	MasscanWhite   string `json:"masscan_white"`
	CreatedTime    string `json:"created_time"`
	UpdatedTime    string `json:"updated_time"`
}

func GetSetting(pageNum int, pageSize int, maps interface{}) (setting []Setting) {
	db.Where(maps).First(&setting)
	return
}

func GetSettingTotal(maps interface{}) (count int) {
	db.Model(&Setting{}).Where(maps).Count(&count)
	return
}

func EditSetting(data interface{}) bool {
	db.Model(&Setting{}).Updates(data)
	return true
}

//根据条件获取全部资产爆破
func GetSettingTitle() (setting []Setting) {
	db.First(&setting)
	return
}
