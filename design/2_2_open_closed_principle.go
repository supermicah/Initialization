package design

import "fmt"

//开闭原则

// AbstractBanker 抽象的银行业务员
type AbstractBanker interface {
	DoBusiness() // DoBusiness 抽象的处理业务接口
}

// SaveBanker 存款的业务员
type SaveBanker struct {
	//AbstractBanker
}

func (sb *SaveBanker) DoBusiness() {
	fmt.Println("进行了存款")
}

// TransferBanker 转账的业务员
type TransferBanker struct {
	//AbstractBanker
}

func (tb *TransferBanker) DoBusiness() {
	fmt.Println("进行了转账")
}

// PayBanker 支付的业务员
type PayBanker struct {
	//AbstractBanker
}

func (pb *PayBanker) DoBusiness() {
	fmt.Println("进行了支付")
}

// BankerBusiness 实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func BankerBusiness(banker AbstractBanker) {
	//通过接口来向下调用，(多态现象)
	banker.DoBusiness()
}

func do2d2() {
	////进行存款
	//sb := &SaveBanker{}
	//sb.DoBusiness()
	//
	////进行转账
	//tb := &TransferBanker{}
	//tb.DoBusiness()
	//
	////进行支付
	//pb := &PayBanker{}
	//pb.DoBusiness()
	//进行存款
	BankerBusiness(&SaveBanker{})

	//进行存款
	BankerBusiness(&TransferBanker{})

	//进行存款
	BankerBusiness(&PayBanker{})

}
