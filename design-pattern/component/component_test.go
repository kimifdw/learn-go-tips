package component

import "testing"

func TestRunComponent(t *testing.T) {
	// 初始化订单结算页面 这个大组件
	checkoutPage := &CheckoutPageComponent{}

	// 挂载子组件
	storeComponent := &StoreComponent{}
	skuComponent := &SkuComponent{}
	promotionComponent := &PromotionComponent{}
	err := skuComponent.Mount(
		promotionComponent,
		&AftersaleComponent{},
		//promotionComponent,
	)
	if err != nil {
		panic(err)
	}
	err = storeComponent.Mount(
		skuComponent,
		&ExpressComponent{},
	)
	if err != nil {
		panic(err)
	}
	// 挂载组件
	err = checkoutPage.Mount(
		&AddressComponent{},
		&PayMethodComponent{},
		storeComponent,
		&InvoiceComponent{},
		&CouponComponent{},
		&GiftCardComponent{},
		&OrderComponent{},
	)
	if err != nil {
		panic(err)
	}
	// 移除组件测试
	//checkoutPage.Remove(storeComponent)

	// 开始构建页面组件数据
	checkoutPage.Do(&Context{})
}
