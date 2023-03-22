package customer

type CustomerController struct {
	CustomerModels CustomerModels
}

func (cc *CustomerController) SetConnCcCustModels(cm CustomerModels) {
	cc.CustomerModels = cm
}
