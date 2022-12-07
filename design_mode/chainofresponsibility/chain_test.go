package chainofresponsibility

import "testing"

/**
在用GoLand写Go代码的时候，不像之前Java那样，可以直接在一个程序中使用Junit可以直接去跑一个代码块了，
在go里面是有单独的测试库testing可以使用，只是需要遵守规则：

一般测试文件以测试程序名_test.go命名
测试函数需要为func TestXxxx注意第一个X一定要大写，否则不生效
*/

func TestChainOfResponsibility(t *testing.T) {
	boardingProcessor := BuildBoardingProcessorChain()
	passenger := &Passenger{
		name:                  "李四",
		hasBoardingPass:       false,
		hasLuggage:            true,
		isPassIdentityCheck:   false,
		isPassSecurityCheck:   false,
		isCompleteForBoarding: false,
	}
	boardingProcessor.ProcessFor(passenger)
}

// BuildBoardingProcessorChain 构建登机流程处理链
// 可以通过构造不同的处理链，满足多场景
func BuildBoardingProcessorChain() BoardingProcessor {
	completeBoardingNode := &completeBoardingProcessor{}

	securityCheckNode := &securityCheckProcessor{}
	securityCheckNode.SetNextProcessor(completeBoardingNode)

	identityCheckNode := &identityCheckProcessor{}
	identityCheckNode.SetNextProcessor(securityCheckNode)

	luggageCheckInNode := &luggageCheckInProcessor{}
	luggageCheckInNode.SetNextProcessor(identityCheckNode)

	boardingPassNode := &boardingPassProcessor{}
	boardingPassNode.SetNextProcessor(luggageCheckInNode)
	return boardingPassNode
}
