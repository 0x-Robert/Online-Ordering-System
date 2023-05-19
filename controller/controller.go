package controller

import (
	"online-ordering-system/model"
)

// request > router > controller > model > controller > response
type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}
