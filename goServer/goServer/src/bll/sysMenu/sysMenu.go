package sysMenu

import (
	"fmt"
	"sort"

	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

var (
	//菜单缓存数据
	sysMenuMap = make(map[int32]*model.SysMenu)

	// 图标
	strIcon = "<i class=\"%s\"></i>"

	// strUl
	strUl = "<ul class=\"nav nav-second-level collapse\">%s</ul>"

	// 无子元素
	strNode = "<li><a class=\"J_menuItem\" href=\"%s\">%s<span class=\"nav-label\">%s</span></a></li>"

	// 有子元素
	strFNode = "<li><a href=\"#\">%s<span class=\"nav-label\">%s</span><span class=\"fa arrow\"></span></a>%s</li>"
)

func init() {
	initTool.RegisterInitFunc(initSysMenuData, initTool.I_NeedInit)
}

// 初始化数据
func initSysMenuData() error {
	sysMenuList, err := dal.SysMenuDALObj.GetAllList()
	if err != nil {
		return err
	}

	for _, item := range sysMenuList {
		sysMenuMap[item.MenuID] = item
	}

	return nil
}

// 通过用户名获取菜单
func getMenuByUser(sysUser *model.SysUser) []*model.SysMenu {
	resultList := make([]*model.SysMenu, 0, 32)

	// 获取玩家角色
	roleIDList := stringTool.SplitToInt32List(sysUser.RoleIDs)
	for _, roleID := range roleIDList {
		if roleInfo, exists := sysRoleMap[roleID]; exists {
			// 获取玩家菜单
			menuIDList := stringTool.SplitToInt32List(roleInfo.MenuIDS)
			for _, menuID := range menuIDList {
				// 判断菜单是否存在
				if menuInfo, exists := sysMenuMap[menuID]; exists {
					// 判断是否其他角色包含了该菜单
					for _, tempMenuInfo := range resultList {
						if tempMenuInfo.MenuID == menuInfo.MenuID {
							continue
						}
					}

					resultList = append(resultList, menuInfo)
				}
			}
		}
	}

	return resultList
}

// 获取该列表中父类下的子类
func getChildMenu(sysMenus []*model.SysMenu, parentID int32) []*model.SysMenu {
	//返回结果
	result := make([]*model.SysMenu, 0, 32)

	for _, menuInfo := range sysMenus {
		if menuInfo.ParentMenuID == parentID {
			result = append(result, menuInfo)
		}
	}

	//按照顺序升序排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].SortOrder < result[j].SortOrder
	})

	return result
}

// 获取玩家的菜单描述
func getMenuScript(sysUser *model.SysUser) string {
	menus := getMenuByUser(sysUser)
	mainMenus := getChildMenu(menus, 0)

	return menuScript(mainMenus, menus)
}

// 菜单描述
func menuScript(sysMenus []*model.SysMenu, allSysMenus []*model.SysMenu) string {
	result := ""

	for _, menuInfo := range sysMenus {
		//如果没有地址，则认为有子页面，如果有，则认为没有子页面
		if menuInfo.MenuUrl == "" {
			childMenus := getChildMenu(allSysMenus, menuInfo.MenuID)
			sonStr := menuScript(childMenus, allSysMenus)

			iconStr := ""
			if menuInfo.MenuIcon != "" {
				iconStr = fmt.Sprintf(strIcon, menuInfo.MenuIcon)
			}

			ulStr := ""
			if sonStr != "" {
				ulStr = fmt.Sprintf(strUl, sonStr)
			}

			result += fmt.Sprintf(strFNode, iconStr, menuInfo.MenuName, ulStr)
		} else {
			iconStr := ""
			if menuInfo.MenuIcon != "" {
				iconStr = fmt.Sprintf(strIcon, menuInfo.MenuIcon)
			}

			result += fmt.Sprintf(strNode, menuInfo.MenuUrl, iconStr, menuInfo.MenuName)
		}
	}

	return result
}
