package setting

import "github.com/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting()(*Setting,error){
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil,err
	}
	return &Setting{vp},nil
}
