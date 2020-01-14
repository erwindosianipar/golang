package docked

import "api-pelabuhan/model"

type DockedRepo interface {
	NewDocked(dock *model.Dock) error
}
