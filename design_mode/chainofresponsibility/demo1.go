package chainofresponsibility

import "fmt"

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

// BoardingProcessor 登机过程中，各节点统一处理接口
type BoardingProcessor interface {
	SetNextProcessor(processor BoardingProcessor)
	ProcessFor(passenger *Passenger)
}

// baseBoardingProcessor 登机流程处理器基类
type baseBoardingProcessor struct {
	//下一个登机处理流程
	nextProcessor BoardingProcessor
}

//实现接口：需要实现接口中的所有方法，go语言中的实现就是将方法与类型进行一个绑定
// SetNextProcessor基类中统一实现设置下一个处理器方法
func (b *baseBoardingProcessor) SetNextProcessor(processor BoardingProcessor) {
	b.nextProcessor = processor
}

//ProcessFor 基类中统一实现下一个处理器流转
func (b *baseBoardingProcessor) ProcessFor(passenger *Passenger) {
	if b.nextProcessor != nil {
		b.nextProcessor.ProcessFor(passenger)
	}
}

//boardingPassProcessor 办理登机牌处理器
type boardingPassProcessor struct {
	baseBoardingProcessor //引用基类
}

func (b *boardingPassProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("为旅客%s办理登机牌;\n", passenger.name)
		passenger.hasBoardingPass = true
	}
	//
	b.baseBoardingProcessor.ProcessFor(passenger)
}

// luggageCheckInProcessor 托运行李处理器
type luggageCheckInProcessor struct {
	baseBoardingProcessor
}

func (l *luggageCheckInProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客%s未办理登机牌，不能托运行李;\n", passenger.name)
		return
	}
	if passenger.hasLuggage {
		fmt.Printf("为旅客%s办理行李托运;\n", passenger.name)
	}
	l.baseBoardingProcessor.ProcessFor(passenger)
}

// identityCheckProcessor 校验身份处理器
type identityCheckProcessor struct {
	baseBoardingProcessor
}

func (i *identityCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客%s未办理登机牌，不能办理身份校验;\n", passenger.name)
		return
	}
	if !passenger.isPassIdentityCheck {
		fmt.Printf("为旅客%s核实身份信息;\n", passenger.name)
		passenger.isPassIdentityCheck = true
	}
	i.baseBoardingProcessor.ProcessFor(passenger)
}

// securityCheckProcessor 安检处理器
type securityCheckProcessor struct {
	baseBoardingProcessor
}

func (s *securityCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客%s未办理登机牌，不能进行安检;\n", passenger.name)
		return
	}
	if !passenger.isPassSecurityCheck {
		fmt.Printf("为旅客%s进行安检;\n", passenger.name)
		passenger.isPassSecurityCheck = true
	}
	s.baseBoardingProcessor.ProcessFor(passenger)
}

// completeBoardingProcessor 完成登机处理器
type completeBoardingProcessor struct {
	baseBoardingProcessor
}

func (c *completeBoardingProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass ||
		!passenger.isPassIdentityCheck ||
		!passenger.isPassSecurityCheck {
		fmt.Printf("旅客%s登机检查过程未完成，不能登机;\n", passenger.name)
		return
	}
	passenger.isCompleteForBoarding = true
	fmt.Printf("旅客%s成功登机;\n", passenger.name)
}


