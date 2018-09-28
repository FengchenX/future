package wechat

import (
	"testing"
	"encoding/xml"
	"fmt"
)
var transferRespXML = `
<xml>

<return_code><![CDATA[SUCCESS]]></return_code>

<return_msg><![CDATA[]]></return_msg>

<mch_appid><![CDATA[wxec38b8ff840bd989]]></mch_appid>

<mchid><![CDATA[10013274]]></mchid>

<device_info><![CDATA[]]></device_info>

<nonce_str><![CDATA[lxuDzMnRjpcXzxLx0q]]></nonce_str>

<result_code><![CDATA[SUCCESS]]></result_code>

<partner_trade_no><![CDATA[10013574201505191526582441]]></partner_trade_no>

<payment_no><![CDATA[1000018301201505190181489473]]></payment_no>

<payment_time><![CDATA[2015-05-19 15：26：59]]></payment_time>

</xml>
`
func TestTransferToChangeRespXmlDecode(t *testing.T) {
	resp := &TransferToChangeResp{}
	xml.Unmarshal([]byte(transferRespXML),resp)
	fmt.Println(resp)
}

var getTransferInfoXML = `
<xml>

<return_code><![CDATA[SUCCESS]]></return_code>

<return_msg><![CDATA[获取成功]]></return_msg>

<result_code><![CDATA[SUCCESS]]></result_code>

<mch_id>10000098</mch_id>

<appid><![CDATA[wxe062425f740c30d8]]></appid>

<detail_id><![CDATA[1000000000201503283103439304]]></detail_id>

<partner_trade_no><![CDATA[1000005901201407261446939628]]></partner_trade_no>

<status><![CDATA[SUCCESS]]></status>

<payment_amount>650</payment_amount >

<openid ><![CDATA[oxTWIuGaIt6gTKsQRLau2M0yL16E]]></openid>

<transfer_time><![CDATA[2015-04-21 20:00:00]]></transfer_time>

<transfer_name ><![CDATA[测试]]></transfer_name >

<desc><![CDATA[福利测试]]></desc>

</xml>
`
func TestGetTransferToChangeInfoXmlDecode(t *testing.T) {
	resp := &GetTransferToChangeInfoResp{}
	xml.Unmarshal([]byte(getTransferInfoXML),resp)
	fmt.Println(resp)
}

