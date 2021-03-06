package sdk

import (
	"github.com/TiMoChao/feishu-sdk-golang/core/consts"
	"github.com/TiMoChao/feishu-sdk-golang/core/model/vo"
	"github.com/TiMoChao/feishu-sdk-golang/core/util/http"
	"github.com/TiMoChao/feishu-sdk-golang/core/util/json"
	"github.com/TiMoChao/feishu-sdk-golang/core/util/log"
)

//校验应用管理员 https://open.feishu.cn/document/ukTMukTMukTM/uITN1EjLyUTNx4iM1UTM
func (t Tenant) IsUserAdmin(openId string, employeeId string) (*vo.IsUserAdminResp, error){
	queryParams := map[string]interface{}{
	}
	if openId != ""{
		queryParams["open_id"] = openId
	}
	if employeeId != ""{
		queryParams["employee_id"] = employeeId
	}
	respBody, err := http.Get(consts.ApiIsUserAdmin, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.IsUserAdminResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}