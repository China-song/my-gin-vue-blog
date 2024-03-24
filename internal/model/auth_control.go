package model

import "gorm.io/gorm"

// Resource

// AddResource 添加一条权限
func AddResource(db *gorm.DB, name, uri, method string, anonymous bool) (*Resource, error) {
	resource := Resource{
		Name:      name,
		Method:    method,
		Url:       uri,
		Anonymous: anonymous,
	}
	result := db.Save(&resource)
	return &resource, result.Error
}

// GetResourcesByRole 获取角色rid具有的权限
func GetResourcesByRole(db *gorm.DB, rid int) (resources []Resource, err error) {
	var role = Role{Model: Model{ID: rid}}
	result := db.Model(&role).Preload("Resources").First(&role)
	return role.Resources, result.Error
}

// Role

// AddRoleWithResources 创建一个角色 并分配一些权限
func AddRoleWithResources(db *gorm.DB, name, label string, rs []int) (*Role, error) {
	role := Role{
		Name:  name,
		Label: label,
	}

	result := db.Create(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	var roleResource []*RoleResource

	for _, rid := range rs {
		roleResource = append(roleResource, &RoleResource{RoleId: role.ID, ResourceId: rid})
	}
	result = db.Create(roleResource)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

// CheckRoleAuth 判断角色rid是否有权限uri
func CheckRoleAuth(db *gorm.DB, rid int, uri, method string) (bool, error) {
	resources, err := GetResourcesByRole(db, rid)
	if err != nil {
		return false, err
	}

	for _, r := range resources {
		if r.Anonymous || (r.Url == uri && r.Method == method) {
			return true, nil
		}
	}

	return false, nil
}
