// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package giftcard 微信礼品卡
package giftcard

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiPageAdd               = "/card/giftcard/page/add"
	apiPageGet               = "/card/giftcard/page/get"
	apiPageUpdate            = "/card/giftcard/page/update"
	apiPageBatchGet          = "/card/giftcard/page/batchget"
	apiMaintainSet           = "/card/giftcard/maintain/set"
	apiPayWhitelistAdd       = "/card/giftcard/pay/whitelist/add"
	apiPaySubmchBind         = "/card/giftcard/pay/submch/bind"
	apiWxaSet                = "/card/giftcard/wxa/set"
	apiOrderGet              = "/card/giftcard/order/get"
	apiOrderBatchGet         = "/card/giftcard/order/batchget"
	apiGeneralCardUpdateUser = "/card/generalcard/updateuser"
	apiOrderRefund           = "/card/giftcard/order/refund"
	apiInvoiceSetBizAttr     = "/card/invoice/setbizattr"
	apiInvoiceGetAuthData    = "/card/invoice/getauthdata"
)

/*
创建-礼品卡货架

开发者可以通过该接口创建一个礼品卡货架并且用于公众号、门店的礼品卡售卖

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/add?access_token=ACCESS_TOKEN
*/
func PageAdd(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPageAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询-礼品卡货架信息

开发者可以查询某个礼品卡货架信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/get?access_token=ACCESS_TOKEN
*/
func PageGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPageGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改-礼品卡货架信息

开发者可以通过该接口更新礼品卡货架信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/update?access_token=ACCESS_TOKEN
*/
func PageUpdate(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPageUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询-礼品卡货架列表

开发者可以通过该接口查询当前商户下所有的礼品卡货架id

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/batchget?access_token=ACCESS_TOKEN
*/
func PageBatchGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPageBatchGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
下架-礼品卡货架

开发者可以通过该接口查询当前商户下所有的礼品卡货架id

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/maintain/set?access_token=ACCESS_TOKEN
*/
func MaintainSet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMaintainSet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
申请微信支付礼品卡权限



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/pay/whitelist/add?access_token=TOKEN
*/
func PayWhitelistAdd(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPayWhitelistAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
绑定商户号到礼品卡小程序



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/pay/submch/bind?access_token=TOKEN
*/
func PaySubmchBind(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPaySubmchBind, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
上传小程序代码



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/wxa/set
*/
func WxaSet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiWxaSet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询-单个礼品卡订单信息

开发者可以通过该接口查询某个订单号对应的订单详情

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/get?access_token=ACCESS_TOKEN
*/
func OrderGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOrderGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量查询礼品卡订单信息

开发者可以通过该接口查询该商户某个时间段内创建的所有订单详情

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/batchget?access_token=ACCESS_TOKEN
*/
func OrderBatchGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOrderBatchGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新用户礼品卡信息

当礼品卡被使用后，开发者可以通过该接口变更某个礼品卡的余额信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/generalcard/updateuser?access_token=TOKEN
*/
func GeneralCardUpdateUser(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGeneralCardUpdateUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
退款

开发者可以通过该接口对某一笔订单操作退款，注意该接口比较隐私，请开发者提高操作该功能的权限等级

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/refund?access_token=ACCESS_TOKEN
*/
func OrderRefund(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOrderRefund, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置支付后开票信息

商户可以通过该接口设置某个商户号发生收款后在支付消息上出现开票授权按钮

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/invoice/setbizattr?action=set_pay_mch&access_token={access_token}
*/
func InvoiceSetBizAttr(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiInvoiceSetBizAttr, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询开票信息

用户完成授权后，商户可以调用该接口查询某一个订单

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/invoice/getauthdata
*/
func InvoiceGetAuthData(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiInvoiceGetAuthData, bytes.NewReader(payload), "application/json;charset=utf-8")
}
