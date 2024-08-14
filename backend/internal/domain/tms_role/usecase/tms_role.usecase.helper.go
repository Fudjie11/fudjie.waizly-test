package usecase

import (
	entity "fudjie.waizly/backend-test/internal/db/auth/entity"
	"github.com/samber/lo"
)

func (m *Module) getNewAndDeletPermission(req []string, eTmsPermission []entity.TmsPermission) ([]string, []string) {
	var (
		newIds    []string
		deleteIds []string
	)

	for _, v := range req {
		exists := lo.ContainsBy(eTmsPermission, func(item entity.TmsPermission) bool {
			return item.PermissionId == v
		})

		if !exists {
			newIds = append(newIds, v)
		}
	}

	for _, v := range eTmsPermission {
		exists := lo.Contains(req, v.PermissionId)

		if !exists {
			deleteIds = append(deleteIds, v.PermissionId)
		}
	}

	return newIds, deleteIds
}
