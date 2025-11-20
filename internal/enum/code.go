/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/11/3 10:45
 * @Desc: TODO
 */

package enum

const (
	SuccessActionStatus = "OK"   // 成功状态
	FailActionStatus    = "FAIL" // 失败状态
	SuccessCode         = 0      // 成功
	InvalidParamsCode   = -1     // 无效参数（自定义）
	InvalidResponseCode = -2     // 无效响应（自定义）
	NotSetGroupType     = 340001
	NotSetGroupName     = 340002
	GroupNameTooLong    = 340003
	InvalidGroupType    = 340004
	GroupIntroductionTooLong = 340005
	GroupNotificationTooLong = 340006
)
