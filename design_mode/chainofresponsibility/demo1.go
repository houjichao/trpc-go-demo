package chainofresponsibility

/**
本示例模拟实现机场登机过程，第一步办理登机牌，第二步如果有行李，就办理托运，
第三步核实身份，第四步安全检查，第五步完成登机；其中行李托运是可选的，其他步骤必选，
必选步骤有任何不满足就终止登机；旅客对象作为请求参数上下文，每个步骤会根据旅客对象状态判断是否处理或流转下一个节点；
*/

//Passenger 旅客结构体
type Passenger struct {
	name                  string //姓名
	hasBoardingPass       bool   //是否办理登机牌
	hasLuggage            bool   //是否有行李需要托运
	isPassIdentityCheck   bool   //是否通过身份校验
	isPassSecurityCheck   bool   //是否通过安检
	isCompleteForBoarding bool   //是否完成登记
}

type BoardingProcessor interface {
	SetNextProcessor(processor BoardingProcessor)
}
