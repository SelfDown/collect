package model

func addTable() {

	//1 attendance
	tableAttendance := Attendance{}
	modelMap["attendance"] = tableAttendance
	primaryKeyMap["attendance"] = tableAttendance.PrimaryKey()
	//2 attendance_result
	tableAttendanceResult := AttendanceResult{}
	modelMap["attendance_result"] = tableAttendanceResult
	primaryKeyMap["attendance_result"] = tableAttendanceResult.PrimaryKey()
	//3 auth_group
	tableAuthGroup := AuthGroup{}
	modelMap["auth_group"] = tableAuthGroup
	primaryKeyMap["auth_group"] = tableAuthGroup.PrimaryKey()
	//4 auth_group_permissions
	tableAuthGroupPermissions := AuthGroupPermissions{}
	modelMap["auth_group_permissions"] = tableAuthGroupPermissions
	primaryKeyMap["auth_group_permissions"] = tableAuthGroupPermissions.PrimaryKey()
	//5 auth_permission
	tableAuthPermission := AuthPermission{}
	modelMap["auth_permission"] = tableAuthPermission
	primaryKeyMap["auth_permission"] = tableAuthPermission.PrimaryKey()
	//6 auth_user
	tableAuthUser := AuthUser{}
	modelMap["auth_user"] = tableAuthUser
	primaryKeyMap["auth_user"] = tableAuthUser.PrimaryKey()
	//7 auth_user_groups
	tableAuthUserGroups := AuthUserGroups{}
	modelMap["auth_user_groups"] = tableAuthUserGroups
	primaryKeyMap["auth_user_groups"] = tableAuthUserGroups.PrimaryKey()
	//8 auth_user_user_permissions
	tableAuthUserUserPermissions := AuthUserUserPermissions{}
	modelMap["auth_user_user_permissions"] = tableAuthUserUserPermissions
	primaryKeyMap["auth_user_user_permissions"] = tableAuthUserUserPermissions.PrimaryKey()
	//9 collect_doc
	tableCollectDoc := CollectDoc{}
	modelMap["collect_doc"] = tableCollectDoc
	primaryKeyMap["collect_doc"] = tableCollectDoc.PrimaryKey()
	//10 collect_doc_demo
	tableCollectDocDemo := CollectDocDemo{}
	modelMap["collect_doc_demo"] = tableCollectDocDemo
	primaryKeyMap["collect_doc_demo"] = tableCollectDocDemo.PrimaryKey()
	//11 collect_doc_important
	tableCollectDocImportant := CollectDocImportant{}
	modelMap["collect_doc_important"] = tableCollectDocImportant
	primaryKeyMap["collect_doc_important"] = tableCollectDocImportant.PrimaryKey()
	//12 collect_doc_params
	tableCollectDocParams := CollectDocParams{}
	modelMap["collect_doc_params"] = tableCollectDocParams
	primaryKeyMap["collect_doc_params"] = tableCollectDocParams.PrimaryKey()
	//13 collect_doc_result
	tableCollectDocResult := CollectDocResult{}
	modelMap["collect_doc_result"] = tableCollectDocResult
	primaryKeyMap["collect_doc_result"] = tableCollectDocResult.PrimaryKey()
	//14 collect_event
	tableCollectEvent := CollectEvent{}
	modelMap["collect_event"] = tableCollectEvent
	primaryKeyMap["collect_event"] = tableCollectEvent.PrimaryKey()
	//15 config_detail
	tableConfigDetail := ConfigDetail{}
	modelMap["config_detail"] = tableConfigDetail
	primaryKeyMap["config_detail"] = tableConfigDetail.PrimaryKey()
	//16 config_detail_change_history
	tableConfigDetailChangeHistory := ConfigDetailChangeHistory{}
	modelMap["config_detail_change_history"] = tableConfigDetailChangeHistory
	primaryKeyMap["config_detail_change_history"] = tableConfigDetailChangeHistory.PrimaryKey()
	//17 config_group
	tableConfigGroup := ConfigGroup{}
	modelMap["config_group"] = tableConfigGroup
	primaryKeyMap["config_group"] = tableConfigGroup.PrimaryKey()
	//18 django_admin_log
	tableDjangoAdminLog := DjangoAdminLog{}
	modelMap["django_admin_log"] = tableDjangoAdminLog
	primaryKeyMap["django_admin_log"] = tableDjangoAdminLog.PrimaryKey()
	//19 django_content_type
	tableDjangoContentType := DjangoContentType{}
	modelMap["django_content_type"] = tableDjangoContentType
	primaryKeyMap["django_content_type"] = tableDjangoContentType.PrimaryKey()
	//20 django_migrations
	tableDjangoMigrations := DjangoMigrations{}
	modelMap["django_migrations"] = tableDjangoMigrations
	primaryKeyMap["django_migrations"] = tableDjangoMigrations.PrimaryKey()
	//21 django_session
	tableDjangoSession := DjangoSession{}
	modelMap["django_session"] = tableDjangoSession
	primaryKeyMap["django_session"] = tableDjangoSession.PrimaryKey()
	//22 doc_group
	tableDocGroup := DocGroup{}
	modelMap["doc_group"] = tableDocGroup
	primaryKeyMap["doc_group"] = tableDocGroup.PrimaryKey()
	//23 ldap_group
	tableLdapGroup := LdapGroup{}
	modelMap["ldap_group"] = tableLdapGroup
	primaryKeyMap["ldap_group"] = tableLdapGroup.PrimaryKey()
	//24 project
	tableProject := Project{}
	modelMap["project"] = tableProject
	primaryKeyMap["project"] = tableProject.PrimaryKey()
	//25 report_gitcommitinfo
	tableReportGitcommitinfo := ReportGitcommitinfo{}
	modelMap["report_gitcommitinfo"] = tableReportGitcommitinfo
	primaryKeyMap["report_gitcommitinfo"] = tableReportGitcommitinfo.PrimaryKey()
	//26 report_gituser_fix
	tableReportGituserFix := ReportGituserFix{}
	modelMap["report_gituser_fix"] = tableReportGituserFix
	primaryKeyMap["report_gituser_fix"] = tableReportGituserFix.PrimaryKey()
	//27 role
	tableRole := Role{}
	modelMap["role"] = tableRole
	primaryKeyMap["role"] = tableRole.PrimaryKey()
	//28 role_ldap_group
	tableRoleLdapGroup := RoleLdapGroup{}
	modelMap["role_ldap_group"] = tableRoleLdapGroup
	primaryKeyMap["role_ldap_group"] = tableRoleLdapGroup.PrimaryKey()
	//29 role_menu
	tableRoleMenu := RoleMenu{}
	modelMap["role_menu"] = tableRoleMenu
	primaryKeyMap["role_menu"] = tableRoleMenu.PrimaryKey()
	//30 sys_code
	tableSysCode := SysCode{}
	modelMap["sys_code"] = tableSysCode
	primaryKeyMap["sys_code"] = tableSysCode.PrimaryKey()
	//31 sys_menu
	tableSysMenu := SysMenu{}
	modelMap["sys_menu"] = tableSysMenu
	primaryKeyMap["sys_menu"] = tableSysMenu.PrimaryKey()
	//32 sys_param
	tableSysParam := SysParam{}
	modelMap["sys_param"] = tableSysParam
	primaryKeyMap["sys_param"] = tableSysParam.PrimaryKey()
	//33 template_event_log
	tableTemplateEventLog := TemplateEventLog{}
	modelMap["template_event_log"] = tableTemplateEventLog
	primaryKeyMap["template_event_log"] = tableTemplateEventLog.PrimaryKey()
	//34 user_account
	tableUserAccount := UserAccount{}
	modelMap["user_account"] = tableUserAccount
	primaryKeyMap["user_account"] = tableUserAccount.PrimaryKey()
	//35 user_change_history
	tableUserChangeHistory := UserChangeHistory{}
	modelMap["user_change_history"] = tableUserChangeHistory
	primaryKeyMap["user_change_history"] = tableUserChangeHistory.PrimaryKey()
	//36 user_role
	tableUserRole := UserRole{}
	modelMap["user_role"] = tableUserRole
	primaryKeyMap["user_role"] = tableUserRole.PrimaryKey()
}