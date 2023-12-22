INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('01044aa5-6f65-4b59-a4a9-ae7a0191be4c', 'session_add', '添加session', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '      - key: session_add
        fields:
          - key: username
            template: "{{.username}}"
          - key: nick
            template: "{{.user_info.nick}}"
          - key: userid
            template: "{{.user_info.user_id}}"
          - key: user_id
            template: "{{.user_info.user_id}}"', '操作session 添加key ，value', 90, '2023-12-05 09:32:21', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('0722aa5a-78e6-4a29-818c-ba33843bdf81', 'model_update', '表修改行数据', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: doc_delete
    http: true
    module: model_update
    params:
      collect_doc_id_list:
        check:
          template: "{{must .collect_doc_id_list}}"
          err_msg: 文档不能为空
      is_delete:
        default: "1"
    table: collect_doc
    filter:
      collect_doc_id__in: "[collect_doc_id_list]"', '示例中是个假删除，将一批记录is_delete 改为1。
主要解决针对数据表的修改，可以单个修改和批量统一修改，主要看你条件怎么写。
注意* 不能，批量不同记录不同值修改，如果需要可以用model_upsert 批量记录不同值修改', 40, '2023-12-04 10:16:57', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('0b3ebb5e-8ce6-4ebc-ae92-da6b0dfe13a2', 'ignore_data', '忽略数据', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: ignore_data
        foreach: "[user_list]"
        params: "params"
        fields:
          - name: "user_id 为空的数据"
            template: "{{ if .user_id }}false{{else}}true{{end}}"', '过滤数组，忽略数组中数据', 140, '2023-12-05 10:01:46', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('1aef7d2e-9336-47ce-b70b-bed9aff9b6ce', 'current_date_time', '生成时间', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', '    params:
      create_time:
        template: "{{current_date_time}}"', '生成时间年月日时分秒', 40, '2023-12-06 18:01:42', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('1ba89c00-fd10-4756-9414-070bf51505d7', 'param2result', '参数转结果', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    result_handler:
      - key: param2result
        field: "[access_token]"', '参数中字段转结果，一般用于只需要返回一个字段', 70, '2023-12-05 09:21:33', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('212b5c89-67ce-4f06-bc23-24421e866974', 'config.doc_edit', '文档编辑', 'service', 'c09992bf-afc8-4abb-a81b-24ff1060fe0b', '{
   "service": "config.doc_edit",
   "demo": [],
   "params": [],
   "important_list": [],
   "doc": {
      "code": "{\\n   \\"service\\": \\"config.doc_save\\",\\n   \\"doc\\": {\\n      \\"title\\": \\"doc.doc_save\\",\\n      \\"sub_title\\": \\"\\",\\n      \\"code\\": \\"\\",\\n      \\"order_index\\": \\"10\\",\\n      \\"type\\": \\"service\\",\\n      \\"parent_dir\\": \\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\",\\n      \\"code_desc\\": \\"\\",\\n      \\"code_result\\": \\"\\"\\n   },\\n   \\"important_list\\": [],\\n   \\"params\\": [],\\n   \\"demo\\": [],\\n   \\"result\\": []\\n}",
      "code_desc": "文档新增",
      "code_result": "{\\n\\t\\"status\\": 0,\\n\\t\\"count\\": 0,\\n\\t\\"success\\": true,\\n\\t\\"code\\": \\"0\\",\\n\\t\\"msg\\": \\"成功\\",\\n\\t\\"data\\": {}\\n}",
      "collect_doc_id": "768f837a-dd60-49ef-bc64-a558eb48bb70",
      "create_time": "2023-12-07 15:39:13",
      "create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
      "is_delete": "0",
      "order_index": 10,
      "parent_dir": "c09992bf-afc8-4abb-a81b-24ff1060fe0b",
      "sub_title": "文档保存",
      "title": "doc.doc_save",
      "type": "service"
   }
}', '文档编辑。注意文档完全是通过对比字段规则，保存的，不是通过数据加字段直接更新，加完字段modify.json 记得加规则', 20, '2023-12-07 15:44:14', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', '{
	"status": 0,
	"count": 0,
	"success": true,
	"code": "0",
	"msg": "成功",
	"data": {
		"base_modify_list": [],
		"change_list": [],
		"demo_add_list": [],
		"demo_modify_list": [],
		"demo_remove_list": [],
		"important_add_list": [],
		"important_remove_list": [],
		"local_doc_detail": {
			"demo": [],
			"doc": {
				"code": "{\\n   \\"service\\": \\"config.doc_save\\",\\n   \\"doc\\": {\\n      \\"title\\": \\"doc.doc_save\\",\\n      \\"sub_title\\": \\"\\",\\n      \\"code\\": \\"\\",\\n      \\"order_index\\": \\"10\\",\\n      \\"type\\": \\"service\\",\\n      \\"parent_dir\\": \\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\",\\n      \\"code_desc\\": \\"\\",\\n      \\"code_result\\": \\"\\"\\n   },\\n   \\"important_list\\": [],\\n   \\"params\\": [],\\n   \\"demo\\": [],\\n   \\"result\\": []\\n}",
				"code_desc": "文档新增",
				"code_result": "",
				"collect_doc_id": "768f837a-dd60-49ef-bc64-a558eb48bb70",
				"create_time": "2023-12-07 15:39:13",
				"create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
				"is_delete": "0",
				"order_index": 10,
				"parent_dir": "c09992bf-afc8-4abb-a81b-24ff1060fe0b",
				"sub_title": "文档保存",
				"title": "doc.doc_save",
				"type": "service"
			},
			"important_list": [],
			"params": []
		},
		"params_modify_list": [],
		"params_remove_list": []
	}
}');
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('234fe81b-c7e3-46a1-a274-2bbc904e6ab7', '生命周期', '服务描述', 'doc', 'dd336894-53b6-405b-98c7-f327407d7cfa', 'params->handler_params->module->result_handler', '服务可以操作的部分目前是四个，
第一步参数定义
第二步参数处理
第三步模块执行
第四步结果处理

请求拦截和定期任务没有定义生命周期。它属于更上一层的包装，而且基本不怎么操作', 50, '2023-12-06 19:37:11', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('25d3fdcb-162a-4ef2-a4af-3b7edcc6e42a', 'count2map', 'count 转字段', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '', '一般是sql查询的count转字段。
这个用比较少，基本用不着count', 180, '2023-12-05 19:50:40', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'hrm.user_list', '用户列表', 'service', 'f0c5c31f-d835-4aa6-8f39-c9fb6c1adfc8', '{
	"service": "hrm.user_list",
	"page": 1,
	"size": 20
}', '查询用户列表', 10, '2023-12-07 09:26:14', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', '{
	"status": 0,
	"count": 127,
	"success": true,
	"code": "0",
	"msg": "执行成功",
	"data": [
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2023-11-03 17:51:12",
			"create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
			"email": "zhangsan@weigaogroup.com",
			"entry_date": "2023-03-11",
			"is_delete": "0",
			"ladp_user_login_id": "",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2023-11-03 17:51:12",
			"modify_user": "",
			"nick": "张三1",
			"password": "01d7f40760960e7bd9443513f22ab9af",
			"phone": "11111111111",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "a33b7fd0-50db-4d55-a29e-573c947721bd",
			"user_name": "zhangsan",
			"user_status": "regular",
			"user_status_name": "正式",
			"username": "zhangsan",
			"wechat_id": "",
			"work_code": "zhangsan"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-05-05 08:52:37",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "tanjingcheng@weigaogroup.com",
			"entry_date": "2022-05-05",
			"is_delete": "0",
			"ladp_user_login_id": "tanjingcheng",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-05-07 15:17:11",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "谭晶城",
			"password": "05d7a0795cf0ef9dfdcee9a3dfd36fe9",
			"phone": "15116423684",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "6f2b3a9f-7ccd-47f5-a5fd-f7988b760127",
			"user_name": "tanjingcheng",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "tanjingcheng",
			"wechat_id": "",
			"work_code": "00119297"
		},
		{
			"attendance_id": "",
			"create_ldap": "0",
			"create_time": "2022-04-26 09:22:46",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "xieyaoyi@weigaogroup.com",
			"entry_date": "2022-04-26",
			"is_delete": "0",
			"ladp_user_login_id": "xieyaoyi",
			"leave_date": "2022-04-26",
			"leave_reason": "不能接受长期出差。",
			"modify_time": "2022-04-27 09:13:21",
			"modify_user": "-1",
			"nick": "谢耀一",
			"password": "c17d7a3905b97ced817f161b21249212",
			"phone": "18976219111",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "93048d5d-198c-4233-91ff-ee6c24d5bb5f",
			"user_name": "xieyaoyi",
			"user_status": "leave",
			"user_status_name": "离职",
			"username": "xieyaoyi",
			"wechat_id": "",
			"work_code": "00119157"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-25 08:46:52",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "hupeng@weigaogroup.com",
			"entry_date": "2022-04-25",
			"is_delete": "0",
			"ladp_user_login_id": "hupeng",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-27 09:13:50",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "胡鹏",
			"password": "0e5768efeb54bdb001cb41b63e8c639e",
			"phone": "13755123980",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "120edd1e-1f4d-4efe-a19b-4c047b61486b",
			"user_name": "hupeng",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "hupeng",
			"wechat_id": "",
			"work_code": "00119109"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-20 09:33:43",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "hecong@weigaogroup.com",
			"entry_date": "2022-04-20",
			"is_delete": "0",
			"ladp_user_login_id": "hecong",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-22 18:05:13",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "何聪",
			"password": "9f46ad5337e0c8285768f7e93d32018f",
			"phone": "15874772562",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "02125451-accb-4acb-b195-435cf42c7d90",
			"user_name": "hecong",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "hecong",
			"wechat_id": "",
			"work_code": "00118975"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-20 09:00:26",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "fuhui@weigaogroup.com",
			"entry_date": "2022-04-20",
			"is_delete": "0",
			"ladp_user_login_id": "fuhui",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-22 18:05:33",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "付辉",
			"password": "a6f54db57b98f3dd767e851e98d9558b",
			"phone": "18574721894",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "daf3dfd4-24de-4c84-9bfc-679ebd9cd0c2",
			"user_name": "fuhui",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "fuhui",
			"wechat_id": "",
			"work_code": "00118976"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-15 09:03:07",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "xionggang@weigaogroup.com",
			"entry_date": "2022-04-15",
			"is_delete": "0",
			"ladp_user_login_id": "xionggang",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-19 11:21:31",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "熊刚",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "13627414094",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "63826a79-3a61-4d90-865f-a480500c8724",
			"user_name": "xionggang",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "xionggang",
			"wechat_id": "",
			"work_code": "00118907"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-14 11:38:33",
			"create_user": "fc46eedc-f5bb-463f-be37-78f08b362828",
			"email": "whslyy@weigaogroup.com",
			"entry_date": "2022-04-14",
			"is_delete": "0",
			"ladp_user_login_id": "menxinyan",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-14 11:38:33",
			"modify_user": "fc46eedc-f5bb-463f-be37-78f08b362828",
			"nick": "门新烟",
			"password": "1c952fbfe31b87c1306bef5bcdf9f935",
			"phone": "1223",
			"role_names": "普通员工,外部威海市立医院",
			"roles": "common,wbweihai",
			"user_id": "6606c08b-6192-4a0c-8f77-67b9cdb33196",
			"user_name": "menxinyan",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "menxinyan",
			"wechat_id": "",
			"work_code": "91002"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-14 11:37:06",
			"create_user": "fc46eedc-f5bb-463f-be37-78f08b362828",
			"email": "whslyy@weigaogroup.com",
			"entry_date": "2022-04-14",
			"is_delete": "0",
			"ladp_user_login_id": "hexinghui",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-14 11:52:31",
			"modify_user": "fc46eedc-f5bb-463f-be37-78f08b362828",
			"nick": "贺幸辉",
			"password": "38f22b7c3b8c343d4315dab636b4e1d0",
			"phone": "111",
			"role_names": "外部威海市立医院",
			"roles": "wbweihai",
			"user_id": "1c85ba13-9911-42f3-be7c-04fc6b8ffea6",
			"user_name": "hexinghui",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "hexinghui",
			"wechat_id": "",
			"work_code": "99001"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-13 08:51:36",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "kuangshengkun@weigaogroup.com",
			"entry_date": "2022-04-13",
			"is_delete": "0",
			"ladp_user_login_id": "kuangshengkun",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-15 09:04:12",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "匡胜昆",
			"password": "6c6687312b92e143444d6fe19e91c18e",
			"phone": "17773102617",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "44c76358-2925-4601-86cf-93c746176f11",
			"user_name": "kuangshengkun",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "kuangshengkun",
			"wechat_id": "",
			"work_code": "00118860"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-13 08:48:28",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "zhouyu@weigaogroup.com",
			"entry_date": "2022-04-13",
			"is_delete": "0",
			"ladp_user_login_id": "zhouyu",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-15 09:05:45",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "周愉",
			"password": "5b0c7138e1fbe31b726a0744bb9747f5",
			"phone": "13548667796",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "0baa8cd8-c711-47cc-956a-cc5985873b11",
			"user_name": "zhouyu",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "zhouyu",
			"wechat_id": "",
			"work_code": "00118862"
		},
		{
			"attendance_id": "",
			"create_ldap": "0",
			"create_time": "2022-04-08 08:43:36",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "wangqian@weigaogroup.com",
			"entry_date": "2022-04-08",
			"is_delete": "0",
			"ladp_user_login_id": "wangqian",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-12 14:59:52",
			"modify_user": "-1",
			"nick": "王倩",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "18098934021",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "84d60fff-fd05-4375-9611-9aabdedde351",
			"user_name": "wangqian",
			"user_status": "leave",
			"user_status_name": "离职",
			"username": "wangqian",
			"wechat_id": "",
			"work_code": "00118759"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-06 08:30:47",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "baibing@weigaogroup.com",
			"entry_date": "2022-04-06",
			"is_delete": "0",
			"ladp_user_login_id": "baibing",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-06 16:33:32",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "白冰",
			"password": "3231e0617fdde4cb9193bb285b402777",
			"phone": "18108450107",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "73d5b5f6-d363-4096-bf25-a37f90d655c7",
			"user_name": "baibing",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "baibing",
			"wechat_id": "",
			"work_code": "00118721"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-31 08:55:38",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "lvting@weigaogroup.com",
			"entry_date": "2022-03-31",
			"is_delete": "0",
			"ladp_user_login_id": "lvting",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-02 14:50:43",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "吕婷",
			"password": "bfb1d909e7a66c52addc9324cb916da3",
			"phone": "15084920749",
			"role_names": "普通员工,产品",
			"roles": "common,product",
			"user_id": "4d8e7b87-7d5a-4558-b4f6-d0918c4a7cc5",
			"user_name": "lvting",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "lvting",
			"wechat_id": "",
			"work_code": "00118646"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-29 08:38:31",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "xiaochunping@weigaogroup.com",
			"entry_date": "2022-03-29",
			"is_delete": "0",
			"ladp_user_login_id": "xiaochunping",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-29 16:35:06",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "肖春平",
			"password": "3a9b5304df5de24d8259f998232d88cf",
			"phone": "17508418876",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "5ff44669-6691-4027-9dc0-7f4a7c43bb2c",
			"user_name": "xiaochunping",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "xiaochunping",
			"wechat_id": "",
			"work_code": "00118574"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-28 15:04:59",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "guanhongqiang@weigaogroup.com",
			"entry_date": "2022-03-28",
			"is_delete": "0",
			"ladp_user_login_id": "guanhongqiang",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-28 15:09:19",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "官红强",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "15243651253",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "b6165b59-5e28-4f02-b7ee-4516229c37b9",
			"user_name": "guanhongqiang",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "guanhongqiang",
			"wechat_id": "",
			"work_code": "00118553"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-28 13:49:44",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "hejian@weigaogroup.com",
			"entry_date": "2022-03-28",
			"is_delete": "0",
			"ladp_user_login_id": "hejian",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-28 15:08:52",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "贺健",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "15574716071",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "bd03d6a6-658c-4fdd-bdd5-9cee31edbdf5",
			"user_name": "hejian",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "hejian",
			"wechat_id": "",
			"work_code": "00118563"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-28 08:50:34",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "xutianyi@weigaogroup.com",
			"entry_date": "2022-03-28",
			"is_delete": "0",
			"ladp_user_login_id": "xutianyi",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-29 16:34:30",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "徐天仪",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "18374601038",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "312e3f8e-e761-4cf5-8416-a4fa471f60a0",
			"user_name": "xutianyi",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "xutianyi",
			"wechat_id": "",
			"work_code": "00118564"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-28 08:47:55",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "wenqia@weigaogroup.com",
			"entry_date": "2022-03-28",
			"is_delete": "0",
			"ladp_user_login_id": "wenqia",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-28 15:07:00",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "文洽",
			"password": "1b594c30e94c8843e65a2e808e1d1000",
			"phone": "13667353867",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "37ef8407-5801-4f6b-bd45-ddfeab848276",
			"user_name": "wenqia",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "wenqia",
			"wechat_id": "",
			"work_code": "00118562"
		},
		{
			"attendance_id": "",
			"create_ldap": "0",
			"create_time": "2022-03-25 08:55:28",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "chenjinke@weigaogroup.com",
			"entry_date": "2022-03-25",
			"is_delete": "0",
			"ladp_user_login_id": "chenjinke",
			"leave_date": "2022-03-28",
			"leave_reason": "个人原因",
			"modify_time": "2022-03-29 09:58:09",
			"modify_user": "-1",
			"nick": "陈金科",
			"password": "0a113ef6b61820daa5611c870ed8d5ee",
			"phone": "18390189893",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "33826318-f291-4361-aeab-de6e9d80800f",
			"user_name": "chenjinke",
			"user_status": "leave",
			"user_status_name": "离职",
			"username": "chenjinke",
			"wechat_id": "",
			"work_code": "00118505"
		}
	]
}');
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('28a1f850-9e5f-4368-9230-220abf98af15', 'excel2data', 'excel转数据', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    excel_config: "./user_list2excel.json"
    handler_params:
      - key: excel2data
        save_field: user_list', 'excel2data和data2excel 配置文件是一致的。导入导出尽量用同一个配置文件，避免出现字段对不上。
实际就是导入

', 125, '2023-12-05 19:42:46', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('293dffa7-1a0b-4e72-8bd7-0f2a44758495', 'session_remove', 'session删除', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: session_remove
        fields:
          - key: username
          - key: nick
          - key: userid
          - key: user_id
', '删除ssession', 110, '2023-12-05 09:40:38', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('2d6b0e66-7188-4660-8d74-0c70176dedca', 'sub_str', '字符串截取', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', 'template: ''./template/{{current_date_format "20220202"}}/user_{{  replace (sub_str current_date_time -8 0) ":" ""}}_{{sub_str uuid -8 0}}.xlsx''', '-8 0 表示从取最后8位字符', 80, '2023-12-06 18:08:03', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('30f9e473-dead-404f-ad33-f2c00fa4e6ad', 'schedule', '定时器', 'doc', 'efbb8e60-eeab-4326-b742-27cedbcc9083', '  - key: ldap_group_sync
    http: true
    schedule:
      enable: "{{eq (get_key \\"schedule_enable\\") \\"true\\"}}"
      schedule_spec: "@every 30s"
    module: empty', '定时执行任务，支持将任何服务做成定时器，服务访问也不受影响。定时器加载只在程序启动加载一次。访问服务，不会重复加载定时器', 20, '2023-12-05 21:07:39', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('31e7d7f9-c87e-4558-bd76-f673d34e1c84', 'group_by', '数据分组', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '      - key: group_by
        enable: "{{must .params_modify_list}}"
        foreach: "[params_modify_list]"
        children: "children"
        fields:
          - field: "[doc_params_id]"
        save_field: params_modify_list', '数据分组，也利用为去重', 260, '2023-12-05 20:45:49', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('3522bb77-0cef-4a33-9e83-5ff05b2cf88f', 'cache', '缓存处理', 'doc', 'efbb8e60-eeab-4326-b742-27cedbcc9083', '  - key: config_detail_query
    http: true
    module: sql
    cache:
      key: "handler_cache"
      enable: "{{eq (get_key \\"can_cache\\") \\"true\\"}}"
      room: config
      second: 0
      fields:
        - field: "[service]"
        - field: "[group_name]"
    params:
      group_name:
        check:
          template: "{{must .group_name}}"
          err_msg: 分组名称不能空
      config:
        default: {}
    data_file: config_detail_query.sql
    result_handler:
      - key: result2params
        fields:
          - to: "[config_params]"
      - key: arr2dict
        enable: "{{must .config_params}}"
        foreach: "[config_params]"
        field: "[name]"
        value: "[value]"
        save_field: config
      - key: param2result
        field: "[config]"', '对服务进行缓存，一般是查询。
如果进来是，查询是否有缓存，有缓存直接返回。没有缓存运行服务，结束设置到缓存，下次进来就有缓存了', 10, '2023-12-05 21:05:03', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('36248998-48ee-41db-a0f4-2010b056b2a3', 'file2result', '返回文件', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '  - key: user_list_download
    module: empty
    http: true
    excel_config: "./user_list2excel.json"
    params:
      excel_path:
        template: ''./template/{{current_date_format "20220202"}}/user_{{  replace (sub_str current_date_time -8 0) ":" ""}}_{{sub_str uuid -8 0}}.xlsx''
      response_name:
        default: "用户列表.xlsx"
    handler_params:
      - key: service2field
        service:
          service: hrm.user_list
        append_param: true
        save_field: user_list
      - key: data2excel
        path: "[excel_path]"
      - key: file2result
        path: "[excel_path]"
        result_name: "[response_name]"', '示例中只看file2result', 130, '2023-12-05 09:55:29', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('3c768fa7-2c9c-48f4-a109-e740ab17ab6c', '什么是参数处理', '服务描述', 'doc', 'dd336894-53b6-405b-98c7-f327407d7cfa', 'handler_params
  - key：xxx
result_handler:
  - key:xxx
', '    我们一般执行模块前，一般前台传过来数据对象，不完全符合我们的预期。比如前台的编码唯一，一般是后台保证。然后前台提交过来的对象，几乎都不是单表保存，数据可能需要保存到好几张表，我们需要根据参数中部分字段存到特定的表里。我之前遇到一个编辑传的全量数据的接口，需要保存20-30张表，根据字段标志存到不同的表
参数处理，是对前台传过来的参数进行特殊处理。
  1.比如前台传了一个ID数组列表，我们需要转换成一个数组对象，批量存入数据表
  2.比如我们要查询一个接口，来校验编码是否唯一
参数处理包括2个部分，请求参数处理，就是模块执行前进行参数处理。结果处理，就是执行完成，针对结果进行处理。
key:xxx 表示用xxx处理器进行处理
是一个数组，表示可以接入多个处理器', 10, '2023-12-03 10:49:33', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('45d1d393-6758-4c37-8ba9-108493f67b8e', 'service2field', '服务转字段', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: service2field
        service:
          service: hrm.ldap_add
        append_param: true', '    将另外一个服务执行结果作为本服务的一个参数。具体服务运行的什么没有任何限制，可以是增删除改查。
    比如校验这个编码是否唯一，我们需要查询数据库表来判断，对比行记录，需要将之前的记录查询出来
    处理有参数处理的公共字段，还有额外字段，service，append_param,append_item_param
', 20, '2023-12-03 16:04:22', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('4a13b507-5eb3-48a5-a99a-28deff080a47', 'md5', '字符串md5加密', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', 'template: "{{ if .item.password }}{{md5 .item.password}}{{ end }}"', '将字符串加密成md5', 70, '2023-12-06 18:06:43', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('4c952eac-4d1e-47c4-b6f2-ac964167ad65', 'must', '判断是否存在参数', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', 'enable: "{{must .remove_list}}"', '存在参数返回true,boolean 变量为false 也是返回true ，因为它只校验存在', 30, '2023-12-06 17:59:37', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('5881b6e1-8217-475f-9b05-4bc20ffdae8e', 'result2params', '结果转参数', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    result_handler:
      - key: result2params
        fields:
          - to: "[config_params]"', '将运行的结果转参数。执行完成结果后，结果还需要进行处理，比如去重，结合拼接一些字段。
前面是param2result 参数转结构，同样也支持结果转参数', 150, '2023-12-05 19:25:50', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('5b86a474-2af6-4e82-8d25-3635645bc315', 'session_get', '获取session', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: session_get
        fields:
          - key: username
            field: username
          - key: nick
            field: nick
          - key: userid
            field: userid
          - key: user_id
            field: user_id', '获取session', 100, '2023-12-05 09:37:22', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('6506e88b-ac97-4724-b6ce-c11f251ce560', 'is_empty', '判断是空', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', ' template: "{{or (is_empty .local_username.cn) (eq .local_username.cn .user_info.username) }}"', '空返回true', 20, '2023-12-06 17:58:27', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('67193cc3-900e-4c27-a3f6-75ee2dba5688', 'arr2dict', '数组转对象', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '      - key: arr2dict
        name: 如果有children 表示有个二级数组
        enable: "{{must .config}}"
        foreach: "[config]"
        children: "children"
        result_name: "children_config"
        field: "[name]"
        value: "[value]"
        save_field: config', '数据列表里面，套了一个二级数组，将二级数组转对象。是转key /value 对象。
比如参数管理', 240, '2023-12-05 20:31:23', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('6cb9a450-5ab1-4a85-9125-54d8844b609c', 'get_key', '获取配置文件参数', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', ' enable: "{{eq (get_key \\"can_cache\\") \\"true\\"}}"', '获取conf/application.properties的参数', 90, '2023-12-06 18:09:49', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('70bbd0a8-5e8a-4083-9986-e8c0c320ebe0', 'combine_array', '数组结合数组', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '      - key: combine_array
        enable: "{{must .width_doc}}"
        foreach: "[group_list]"
        field: "[doc_group_id]"
        right: "[doc_list]"
        right_field: "[parent_dir]"
        children: "children"', '一般利用2个服务结合，将一个数字拼接到另外一个数组中取，比如参数管理，二级数组', 260, '2023-12-05 20:58:23', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('74973d56-9773-432f-b1d5-3b68ad40998d', 'filter_arr', '过滤数组', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '      - key: filter_arr
        foreach: "[change_list]"
        item: item
        if_template: "{{and (eq .item.operation \\"remove\\") (ne .item.has_group \\"0\\") }}"
        save_field: remove_list', '过滤数组，一般把增、删、改的数据过滤，分别对应操作', 220, '2023-12-05 20:19:38', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('768f837a-dd60-49ef-bc64-a558eb48bb70', 'doc.doc_save', '文档保存', 'service', 'c09992bf-afc8-4abb-a81b-24ff1060fe0b', '{
   "service": "config.doc_save",
   "doc": {
      "title": "doc.doc_save",
      "sub_title": "",
      "code": "",
      "order_index": "10",
      "type": "service",
      "parent_dir": "c09992bf-afc8-4abb-a81b-24ff1060fe0b",
      "code_desc": "",
      "code_result": ""
   },
   "important_list": [],
   "params": [],
   "demo": [],
   "result": []
}', '文档新增', 10, '2023-12-07 15:39:13', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', '');
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('77e10054-791e-4889-a5c8-1fc2b9a6e514', 'handler_cache', '处理缓存', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '      - key: handler_cache
        enable: "{{and (eq (get_key \\"can_cache\\") \\"true\\") (must .set_cache)}}"
        name: 批量设置缓存
        method: BULK_SET_CACHE
        foreach: "[config]"
        item: item
        field: "[item.children_config]"
        room: config
        second: 0
        fields:
          - field: "config.config_detail_query"
          - field: "[item.group_id]"', '缓存设置，参数管理的查询就用了缓存，不用每次都充数据库查。
查询与设置利用拦截器，只要设置个cache 标签，服务访问先查询有没有缓存，没有缓存执行，执行完成设置缓存。如果有缓存就直接返回', 250, '2023-12-05 20:41:12', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('79d0939d-16f6-4c8d-b611-5a55409162ad', 'arr2arrayObj', '数组转对象数组', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: field2array
        enable: "{{must .roles}}"
        field: "[roles]"
        save_field: role_list
      - key: arr2arrayObj
        enable: "{{must .role_list}}"
        foreach: "[role_list]"
        item: item
        fields:
          - field: "role_id"
            template: "{{.item}}"
          - field: user_id
            template: "[user_id]"
          - field: user_role_id
            template: "{{uuid}}"
        save_field: user_role_list', '前端传一个普通数组，转对象数组', 210, '2023-12-05 20:14:38', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('7b95f2ea-a69d-47a7-9857-ab301d5b0bb6', 'hash_sha', 'ldap的hash加密', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', '  {"TYPE": "userpassword","Vals": ["{{hash_sha .add_password}}"]}', 'ldap 加密算法', 110, '2023-12-06 18:12:22', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('7da94650-209f-411e-aa0c-47989dbd4409', 'config.doc_detail', '查询详情', 'service', 'c09992bf-afc8-4abb-a81b-24ff1060fe0b', '{
   "service": "config.doc_detail",
   "collect_doc_id": "212b5c89-67ce-4f06-bc23-24421e866974"
}', '查询文档详情', 5, '2023-12-07 16:07:07', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', '{
	"status": 0,
	"count": 0,
	"success": true,
	"code": "0",
	"msg": "成功",
	"data": {
		"demo": [],
		"doc": {
			"code": "{\\n   \\"service\\": \\"config.doc_edit\\",\\n   \\"demo\\": [],\\n   \\"params\\": [],\\n   \\"important_list\\": [],\\n   \\"doc\\": {\\n      \\"code\\": \\"{\\\\n   \\\\\\"service\\\\\\": \\\\\\"config.doc_save\\\\\\",\\\\n   \\\\\\"doc\\\\\\": {\\\\n      \\\\\\"title\\\\\\": \\\\\\"doc.doc_save\\\\\\",\\\\n      \\\\\\"sub_title\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"code\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"order_index\\\\\\": \\\\\\"10\\\\\\",\\\\n      \\\\\\"type\\\\\\": \\\\\\"service\\\\\\",\\\\n      \\\\\\"parent_dir\\\\\\": \\\\\\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\\\\\",\\\\n      \\\\\\"code_desc\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"code_result\\\\\\": \\\\\\"\\\\\\"\\\\n   },\\\\n   \\\\\\"important_list\\\\\\": [],\\\\n   \\\\\\"params\\\\\\": [],\\\\n   \\\\\\"demo\\\\\\": [],\\\\n   \\\\\\"result\\\\\\": []\\\\n}\\",\\n      \\"code_desc\\": \\"文档新增\\",\\n      \\"code_result\\": \\"{\\\\n\\\\t\\\\\\"status\\\\\\": 0,\\\\n\\\\t\\\\\\"count\\\\\\": 0,\\\\n\\\\t\\\\\\"success\\\\\\": true,\\\\n\\\\t\\\\\\"code\\\\\\": \\\\\\"0\\\\\\",\\\\n\\\\t\\\\\\"msg\\\\\\": \\\\\\"成功\\\\\\",\\\\n\\\\t\\\\\\"data\\\\\\": {}\\\\n}\\",\\n      \\"collect_doc_id\\": \\"768f837a-dd60-49ef-bc64-a558eb48bb70\\",\\n      \\"create_time\\": \\"2023-12-07 15:39:13\\",\\n      \\"create_user\\": \\"739ade44-7e83-48a2-8c60-9a7c1e9f3d0a\\",\\n      \\"is_delete\\": \\"0\\",\\n      \\"order_index\\": 10,\\n      \\"parent_dir\\": \\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\",\\n      \\"sub_title\\": \\"文档保存\\",\\n      \\"title\\": \\"doc.doc_save\\",\\n      \\"type\\": \\"service\\"\\n   }\\n}",
			"code_desc": "文档编辑。注意文档完全是通过对比字段规则，保存的，不是通过数据加字段直接更新，加完字段modify.json 记得加规则",
			"code_result": "{\\n\\t\\"status\\": 0,\\n\\t\\"count\\": 0,\\n\\t\\"success\\": true,\\n\\t\\"code\\": \\"0\\",\\n\\t\\"msg\\": \\"成功\\",\\n\\t\\"data\\": {\\n\\t\\t\\"base_modify_list\\": [],\\n\\t\\t\\"change_list\\": [],\\n\\t\\t\\"demo_add_list\\": [],\\n\\t\\t\\"demo_modify_list\\": [],\\n\\t\\t\\"demo_remove_list\\": [],\\n\\t\\t\\"important_add_list\\": [],\\n\\t\\t\\"important_remove_list\\": [],\\n\\t\\t\\"local_doc_detail\\": {\\n\\t\\t\\t\\"demo\\": [],\\n\\t\\t\\t\\"doc\\": {\\n\\t\\t\\t\\t\\"code\\": \\"{\\\\n   \\\\\\"service\\\\\\": \\\\\\"config.doc_save\\\\\\",\\\\n   \\\\\\"doc\\\\\\": {\\\\n      \\\\\\"title\\\\\\": \\\\\\"doc.doc_save\\\\\\",\\\\n      \\\\\\"sub_title\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"code\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"order_index\\\\\\": \\\\\\"10\\\\\\",\\\\n      \\\\\\"type\\\\\\": \\\\\\"service\\\\\\",\\\\n      \\\\\\"parent_dir\\\\\\": \\\\\\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\\\\\",\\\\n      \\\\\\"code_desc\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"code_result\\\\\\": \\\\\\"\\\\\\"\\\\n   },\\\\n   \\\\\\"important_list\\\\\\": [],\\\\n   \\\\\\"params\\\\\\": [],\\\\n   \\\\\\"demo\\\\\\": [],\\\\n   \\\\\\"result\\\\\\": []\\\\n}\\",\\n\\t\\t\\t\\t\\"code_desc\\": \\"文档新增\\",\\n\\t\\t\\t\\t\\"code_result\\": \\"\\",\\n\\t\\t\\t\\t\\"collect_doc_id\\": \\"768f837a-dd60-49ef-bc64-a558eb48bb70\\",\\n\\t\\t\\t\\t\\"create_time\\": \\"2023-12-07 15:39:13\\",\\n\\t\\t\\t\\t\\"create_user\\": \\"739ade44-7e83-48a2-8c60-9a7c1e9f3d0a\\",\\n\\t\\t\\t\\t\\"is_delete\\": \\"0\\",\\n\\t\\t\\t\\t\\"order_index\\": 10,\\n\\t\\t\\t\\t\\"parent_dir\\": \\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\",\\n\\t\\t\\t\\t\\"sub_title\\": \\"文档保存\\",\\n\\t\\t\\t\\t\\"title\\": \\"doc.doc_save\\",\\n\\t\\t\\t\\t\\"type\\": \\"service\\"\\n\\t\\t\\t},\\n\\t\\t\\t\\"important_list\\": [],\\n\\t\\t\\t\\"params\\": []\\n\\t\\t},\\n\\t\\t\\"params_modify_list\\": [],\\n\\t\\t\\"params_remove_list\\": []\\n\\t}\\n}",
			"collect_doc_id": "212b5c89-67ce-4f06-bc23-24421e866974",
			"create_time": "2023-12-07 15:44:14",
			"create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
			"is_delete": "0",
			"order_index": 20,
			"parent_dir": "c09992bf-afc8-4abb-a81b-24ff1060fe0b",
			"sub_title": "文档编辑",
			"title": "config.doc_edit",
			"type": "service"
		},
		"important_list": [],
		"params": []
	}
}');
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('803b90d9-c58c-4113-b4c0-58782e03142c', 'bulk_upsert', '批量修改多行', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: collect_doc_params_update
    module: bulk_upsert
    table: "collect_doc_params"
    params:
      params:
        check:
          template: "{{must .params}}"
          err_msg: 参数不能为空
    model_field: "[params]"', '批量修改多行记录', 80, '2023-12-04 11:43:11', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('849062cb-1f83-454e-b81d-3d7754f9ac4a', 'update_array', '更新数组', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: update_array
        foreach: "[ldap_group_list]"
        item: item
        fields:
          - field: ldap_group_id
            template: "{{ uuid }}"
          - field: name
            template: "{{.item.after}}"
          - field: has_group
            template: "1"
          - field: last_sync_time
            template: "{{current_date_time}}"', '更新参数中数组里面对象字段。一般用于批量保存，批量修改，批量服务流程化中。
比如批量新增生成唯一的uuid，生成创建人，创建时间', 50, '2023-12-05 08:56:16', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('894496fd-4201-4b40-95d6-24ecd978b719', 'bulk_create', '表批量保存', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: config_detail_bulk_create
    name: 配置批量新增
    module: bulk_create
    log: true
    table: "config_detail"
    model_field: "[detail_list]"
    http: true
    params:
      detail_list:
        check:
          template: "{{must .detail_list}}"
          err_msg: 数据列表不能为空
    handler_params:
      - key: update_array
        foreach: "[detail_list]"
        item: item
        fields:
          - field: config_detail_id
            template: "{{uuid}}"', '    数据库表批量保存，bulk_service 针对服务类型批量操作，如果确定是同一张表，bulk_create创建数据，数据库连接只有一次，能减少数据库的连接次数。
    bulk_service 是针对如何服务多线程跑。
    示例中的handler_params是参数处理模块
  ', 70, '2023-12-04 11:02:42', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('8ec4053f-1a54-4f78-9de6-31db6995692e', 'model_delete', '表数据删除', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: role_ldap_group_delete
    module: model_delete
    params:
      role_id_list:
        check:
          template: "{{must .role_id_list}}"
          err_msg: 角色不能为空
    table: "role_ldap_group"
    filter:
      role_id__in: "[role_id_list]"', '针对表数据库删除', 50, '2023-12-04 10:59:03', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('91d4abd1-14c9-4705-9de7-1518da998427', '什么是服务service', '服务描述', 'doc', 'dd336894-53b6-405b-98c7-f327407d7cfa', 'module：empty', '    我们常常编写一个接口，一般是要求服务器做一件事情，比如查询一个用户列表，更新一个用户信息，删除一条记录。我们管服务器每做一个事情叫做服务。
    所以服务，就表示你要服务器做一件什么事情。比如服务器执行sql，是一个服务，删除数据记录也是一个服务。让服务器发个http请求到其他服务器也是一个服务。服务划分这个主要看个人的主观意愿。
    以前我写python、java查询接口，就是执行一段sql，就必须包一个http接口。现在不需要，然后只需要后台配置一个sql文件，服务配置后，对外就可以暴露一个http接口。
    我做的一件事情，将接口更抽象一次。一般人写的接口是面向业务提供一个接口。而我是面向业务接口，提供一个工具。这个工具可以制造业务接口，制造出来的接口是服务，同时也能完成业务接口做的事情。
    报表也是如此，报表根据不同sql，可以渲染不同的数据，以达到不用写接口。但是报表的数据能力不行。比如修改数据模型、批量修改数据模型，需要处理uuid，创建时间，或者过滤一部分数据，报表不行。它根本mvc的model层。不具备代码能力
    服务配置的目录结构
    |conf
    └－－application.properties
    |collect
    └－－service.yml
    └－－hrm
        └－－service.yml
        └－－user
            └－－index.yml
    1.conf/application.properties配置文件入口，一般配置可以缓存开关，数据库地址。
    2.collect 是总服务配置文件夹，collect/service.yml 是总入口
    3.hrm 是业务模块，下面service.yml该页面模块的入口，一般连接业务的文件夹。一般叶子目录下的index.yml 录入服务，service 节点下key 表示具体服务，index.yml 一般按照功能分类，表名分类，再上一级表示项目分类.', 0, '2023-11-29 09:21:40', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('9215873e-d44d-4a8f-b243-f2c45b1833e5', 'ldap', 'ldap', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: ldap_search
    http: true
    module: ldap
    params:
      search_username:
        check:
          template: "{{must .search_username}}"
          err_msg: 用户名不能为空
    handler_params:
      - key: service2field
        service:
          service: config.conf单独ig_detail_query
          group_name: ldap_config
        save_field: ldap_config
        template: "{{must .ldap_config}}"
        err_msg: ldap配置不存在
    data_file: search.json', '    支持ldap模块的增删改查对对对，示例中config_detail_query 是查询ldap 的配置，我将ldap配置到数据库了。
    本模块不太属于公共模块，只是我工作中需要对接，属于业务模块没有过多介绍，至于参数配置请到https://github.com/go-ldap/ldap', 100, '2023-12-04 19:55:07', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('9b4fbf55-b221-4b05-86f1-2e78132ee552', 'model_save', '表新增行', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: collect_doc_save
    module: model_save
    table: collect_doc
    params:
      create_time:
        template: "{{current_date_time}}"
      create_user:
        template: "{{.session_user_id}}"
      is_delete:
        default: "0"', '对数据库表的新增一行记录', 30, '2023-12-04 09:32:36', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('a721393b-32e3-463f-b679-49c9eef2f8cc', 'replace', '字符串替换', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', 'template: ''./template/{{current_date_format "20220202"}}/user_{{  replace (sub_str current_date_time -8 0) ":" ""}}_{{sub_str uuid -8 0}}.xlsx''', '字符串替换，（原字符串，替换来源，替换目标）', 60, '2023-12-06 18:05:34', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('aa2750ff-6c44-44ac-8c95-393eec4c5861', 'current_date_format ', '时间格式化', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', '    params:
      excel_path:
        template: ''./template/{{current_date_format "20220202"}}/user_{{  replace (sub_str current_date_time -8 0) ":" ""}}_{{sub_str uuid -8 0}}.xlsx''
', '将当前时间格式化，golang 没有 yyyy-mm-dd 这种格式，它需要你给个时间示例，然后生成时间格式', 50, '2023-12-06 18:03:25', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('b7a26057-abd7-45a0-9101-63e01af6ed4c', 'get_modify_data', '比对数据', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    modify_config: doc_modify.json
    handler_params:
      - key: get_modify_data
        save_field: change_list
', '    我经常遇到业务级别，需要记录某个字段是谁改的，改之前是什么，改之后是什么。
   但是利用用数据库的binlog日志展示不合适，比如要记录版本号是谁改的，改之前是什么，改之后是什么，然后还原此条。
    一个两个情况到没有什么，主要很多地方有这样的需求，比如保存一个全量列表，之前的搞法就是直接全部删除，然后全部添加，后面发现效率不行，毕竟数量多了，删除和添加总会要占用一定时间，甚至可能触发数据库锁。本身只改了一点点数据，却触发整个表的删除与新增
    有了这个对比工具，我们可以对比列表的差异部分，然后数据进行，哪些删除，哪些新增，哪些修改，理应如此', 0, '2023-12-03 11:05:50', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('b9eab01d-7236-4d15-9053-0f193a3d2ffb', 'result2map', '结果转字段', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '无', '就是将结果，外层在包一层对象。
这个用比较少，一般可能利用params2result。
', 170, '2023-12-05 19:47:53', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('bd6431c8-caf3-4ae9-96af-03cce615204b', 'file2datajson', '配置文件转json', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '  - key: project_router
    http: true
    handler_params:
      - key: file2datajson
        save_field: data
      - key: param2result
        field: data
    data_file: project_router.json
    module: empty', '后台配置文件转json对象。路由配置，构造批量服务的对象，用了这个', 180, '2023-12-05 19:58:58', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('c00a6f14-98c1-4308-8017-cf35ae300de4', 'http', 'http请求', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: gettoken
    module: http
    http_json: gettoken.json
    success: "{{ if .access_token }}true{{ else }}false{{ end }}"
    result_handler:
      - key: result2params
        fields:
          - from: "[access_token]"
            to: "[access_token]"
      - key: param2result
        field: "[access_token]"', '   我们后台经常需要发http请求到其他服务器，比如获取企业微信用户信息，调用集成第三方服务的接口。前台肯定是不能直接调用，需要后台调用，处理完成之后，前台再处理。其主要参考思路是ajax 配置化发送http请求，一般语言像java、python、甚至go对http配置发请求能力还是比较弱，一般封装一个工具类。
   注意module: http 表示是本服务是给其他服务器发请求，而配置http：true 表示改服务对外暴露http接口，允许外部调用，否则只能内部调用', 95, '2023-12-04 13:56:24', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'update_array_from_array', '更新数组', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '      - key: update_array_from_array
        name: "根据group+name 更新操作，有修改，没有就新增"
        foreach: "[change_list]"
        item: item
        field: "[group_id&name_copy]"
        right: "[local_detail_list]"
        right_field: "[group_id&name]"
        fields:
          - field: operation
            name: 如果操作是删除，还是原来的删除，如果是新增操作，存在这改为修改
            template: "{{ if eq .item.operation \\"add\\"}}modify{{else}}{{.item.operation}}{{end}}"
          - field: config_detail_id
            template: "[right.config_detail_id]"', '从另外一个数组，找到相同的记录，本更新字段。
a数组从b数组里面更新字段,治取交集部分更新', 260, '2023-12-05 20:54:07', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('c13b5daf-13ac-47f5-9069-2a6c16e765a4', 'pinyin', '获取中文的拼音', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', 'template: "{{pinyin .content}}"', '获取中文的拼音', 100, '2023-12-06 18:11:08', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('c667be3c-479d-4a29-be13-9ff9963232de', 'config.doc_group_query', '分组查询', 'service', 'c09992bf-afc8-4abb-a81b-24ff1060fe0b', '{
   "service": "config.doc_group_query",
   "width_doc": true
}', '文档分组查询。全量查出', 1, '2023-12-07 16:20:18', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', '{
	"status": 0,
	"count": 0,
	"success": true,
	"code": "0",
	"msg": "执行成功",
	"data": [{
		"children": [{
				"code": "{\\n   \\"service\\": \\"config.doc_save\\",\\n   \\"doc\\": {\\n      \\"title\\": \\"doc.doc_save\\",\\n      \\"sub_title\\": \\"\\",\\n      \\"code\\": \\"\\",\\n      \\"order_index\\": \\"10\\",\\n      \\"type\\": \\"service\\",\\n      \\"parent_dir\\": \\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\",\\n      \\"code_desc\\": \\"\\",\\n      \\"code_result\\": \\"\\"\\n   },\\n   \\"important_list\\": [],\\n   \\"params\\": [],\\n   \\"demo\\": [],\\n   \\"result\\": []\\n}",
				"code_desc": "文档新增",
				"code_result": "",
				"collect_doc_id": "768f837a-dd60-49ef-bc64-a558eb48bb70",
				"create_time": "2023-12-07 15:39:13",
				"create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
				"is_delete": "0",
				"order_index": 10,
				"parent_dir": "c09992bf-afc8-4abb-a81b-24ff1060fe0b",
				"sub_title": "文档保存",
				"title": "doc.doc_save",
				"type": "service"
			},
			{
				"code": "{\\n   \\"service\\": \\"config.doc_edit\\",\\n   \\"demo\\": [],\\n   \\"params\\": [],\\n   \\"important_list\\": [],\\n   \\"doc\\": {\\n      \\"code\\": \\"{\\\\n   \\\\\\"service\\\\\\": \\\\\\"config.doc_save\\\\\\",\\\\n   \\\\\\"doc\\\\\\": {\\\\n      \\\\\\"title\\\\\\": \\\\\\"doc.doc_save\\\\\\",\\\\n      \\\\\\"sub_title\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"code\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"order_index\\\\\\": \\\\\\"10\\\\\\",\\\\n      \\\\\\"type\\\\\\": \\\\\\"service\\\\\\",\\\\n      \\\\\\"parent_dir\\\\\\": \\\\\\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\\\\\",\\\\n      \\\\\\"code_desc\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"code_result\\\\\\": \\\\\\"\\\\\\"\\\\n   },\\\\n   \\\\\\"important_list\\\\\\": [],\\\\n   \\\\\\"params\\\\\\": [],\\\\n   \\\\\\"demo\\\\\\": [],\\\\n   \\\\\\"result\\\\\\": []\\\\n}\\",\\n      \\"code_desc\\": \\"文档新增\\",\\n      \\"code_result\\": \\"{\\\\n\\\\t\\\\\\"status\\\\\\": 0,\\\\n\\\\t\\\\\\"count\\\\\\": 0,\\\\n\\\\t\\\\\\"success\\\\\\": true,\\\\n\\\\t\\\\\\"code\\\\\\": \\\\\\"0\\\\\\",\\\\n\\\\t\\\\\\"msg\\\\\\": \\\\\\"成功\\\\\\",\\\\n\\\\t\\\\\\"data\\\\\\": {}\\\\n}\\",\\n      \\"collect_doc_id\\": \\"768f837a-dd60-49ef-bc64-a558eb48bb70\\",\\n      \\"create_time\\": \\"2023-12-07 15:39:13\\",\\n      \\"create_user\\": \\"739ade44-7e83-48a2-8c60-9a7c1e9f3d0a\\",\\n      \\"is_delete\\": \\"0\\",\\n      \\"order_index\\": 10,\\n      \\"parent_dir\\": \\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\",\\n      \\"sub_title\\": \\"文档保存\\",\\n      \\"title\\": \\"doc.doc_save\\",\\n      \\"type\\": \\"service\\"\\n   }\\n}",
				"code_desc": "文档编辑。注意文档完全是通过对比字段规则，保存的，不是通过数据加字段直接更新，加完字段modify.json 记得加规则",
				"code_result": "{\\n\\t\\"status\\": 0,\\n\\t\\"count\\": 0,\\n\\t\\"success\\": true,\\n\\t\\"code\\": \\"0\\",\\n\\t\\"msg\\": \\"成功\\",\\n\\t\\"data\\": {\\n\\t\\t\\"base_modify_list\\": [],\\n\\t\\t\\"change_list\\": [],\\n\\t\\t\\"demo_add_list\\": [],\\n\\t\\t\\"demo_modify_list\\": [],\\n\\t\\t\\"demo_remove_list\\": [],\\n\\t\\t\\"important_add_list\\": [],\\n\\t\\t\\"important_remove_list\\": [],\\n\\t\\t\\"local_doc_detail\\": {\\n\\t\\t\\t\\"demo\\": [],\\n\\t\\t\\t\\"doc\\": {\\n\\t\\t\\t\\t\\"code\\": \\"{\\\\n   \\\\\\"service\\\\\\": \\\\\\"config.doc_save\\\\\\",\\\\n   \\\\\\"doc\\\\\\": {\\\\n      \\\\\\"title\\\\\\": \\\\\\"doc.doc_save\\\\\\",\\\\n      \\\\\\"sub_title\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"code\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"order_index\\\\\\": \\\\\\"10\\\\\\",\\\\n      \\\\\\"type\\\\\\": \\\\\\"service\\\\\\",\\\\n      \\\\\\"parent_dir\\\\\\": \\\\\\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\\\\\",\\\\n      \\\\\\"code_desc\\\\\\": \\\\\\"\\\\\\",\\\\n      \\\\\\"code_result\\\\\\": \\\\\\"\\\\\\"\\\\n   },\\\\n   \\\\\\"important_list\\\\\\": [],\\\\n   \\\\\\"params\\\\\\": [],\\\\n   \\\\\\"demo\\\\\\": [],\\\\n   \\\\\\"result\\\\\\": []\\\\n}\\",\\n\\t\\t\\t\\t\\"code_desc\\": \\"文档新增\\",\\n\\t\\t\\t\\t\\"code_result\\": \\"\\",\\n\\t\\t\\t\\t\\"collect_doc_id\\": \\"768f837a-dd60-49ef-bc64-a558eb48bb70\\",\\n\\t\\t\\t\\t\\"create_time\\": \\"2023-12-07 15:39:13\\",\\n\\t\\t\\t\\t\\"create_user\\": \\"739ade44-7e83-48a2-8c60-9a7c1e9f3d0a\\",\\n\\t\\t\\t\\t\\"is_delete\\": \\"0\\",\\n\\t\\t\\t\\t\\"order_index\\": 10,\\n\\t\\t\\t\\t\\"parent_dir\\": \\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\",\\n\\t\\t\\t\\t\\"sub_title\\": \\"文档保存\\",\\n\\t\\t\\t\\t\\"title\\": \\"doc.doc_save\\",\\n\\t\\t\\t\\t\\"type\\": \\"service\\"\\n\\t\\t\\t},\\n\\t\\t\\t\\"important_list\\": [],\\n\\t\\t\\t\\"params\\": []\\n\\t\\t},\\n\\t\\t\\"params_modify_list\\": [],\\n\\t\\t\\"params_remove_list\\": []\\n\\t}\\n}",
				"collect_doc_id": "212b5c89-67ce-4f06-bc23-24421e866974",
				"create_time": "2023-12-07 15:44:14",
				"create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
				"is_delete": "0",
				"order_index": 20,
				"parent_dir": "c09992bf-afc8-4abb-a81b-24ff1060fe0b",
				"sub_title": "文档编辑",
				"title": "config.doc_edit",
				"type": "service"
			},
			{
				"code": "{\\n   \\"service\\": \\"config.doc_detail\\",\\n   \\"collect_doc_id\\": \\"212b5c89-67ce-4f06-bc23-24421e866974\\"\\n}",
				"code_desc": "查询文档详情",
				"code_result": "{\\n\\t\\"status\\": 0,\\n\\t\\"count\\": 0,\\n\\t\\"success\\": true,\\n\\t\\"code\\": \\"0\\",\\n\\t\\"msg\\": \\"成功\\",\\n\\t\\"data\\": {\\n\\t\\t\\"demo\\": [],\\n\\t\\t\\"doc\\": {\\n\\t\\t\\t\\"code\\": \\"{\\\\n   \\\\\\"service\\\\\\": \\\\\\"config.doc_edit\\\\\\",\\\\n   \\\\\\"demo\\\\\\": [],\\\\n   \\\\\\"params\\\\\\": [],\\\\n   \\\\\\"important_list\\\\\\": [],\\\\n   \\\\\\"doc\\\\\\": {\\\\n      \\\\\\"code\\\\\\": \\\\\\"{\\\\\\\\n   \\\\\\\\\\\\\\"service\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"config.doc_save\\\\\\\\\\\\\\",\\\\\\\\n   \\\\\\\\\\\\\\"doc\\\\\\\\\\\\\\": {\\\\\\\\n      \\\\\\\\\\\\\\"title\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"doc.doc_save\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"sub_title\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"code\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"order_index\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"10\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"type\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"service\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"parent_dir\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"code_desc\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"code_result\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"\\\\\\\\\\\\\\"\\\\\\\\n   },\\\\\\\\n   \\\\\\\\\\\\\\"important_list\\\\\\\\\\\\\\": [],\\\\\\\\n   \\\\\\\\\\\\\\"params\\\\\\\\\\\\\\": [],\\\\\\\\n   \\\\\\\\\\\\\\"demo\\\\\\\\\\\\\\": [],\\\\\\\\n   \\\\\\\\\\\\\\"result\\\\\\\\\\\\\\": []\\\\\\\\n}\\\\\\",\\\\n      \\\\\\"code_desc\\\\\\": \\\\\\"文档新增\\\\\\",\\\\n      \\\\\\"code_result\\\\\\": \\\\\\"{\\\\\\\\n\\\\\\\\t\\\\\\\\\\\\\\"status\\\\\\\\\\\\\\": 0,\\\\\\\\n\\\\\\\\t\\\\\\\\\\\\\\"count\\\\\\\\\\\\\\": 0,\\\\\\\\n\\\\\\\\t\\\\\\\\\\\\\\"success\\\\\\\\\\\\\\": true,\\\\\\\\n\\\\\\\\t\\\\\\\\\\\\\\"code\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"0\\\\\\\\\\\\\\",\\\\\\\\n\\\\\\\\t\\\\\\\\\\\\\\"msg\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"成功\\\\\\\\\\\\\\",\\\\\\\\n\\\\\\\\t\\\\\\\\\\\\\\"data\\\\\\\\\\\\\\": {}\\\\\\\\n}\\\\\\",\\\\n      \\\\\\"collect_doc_id\\\\\\": \\\\\\"768f837a-dd60-49ef-bc64-a558eb48bb70\\\\\\",\\\\n      \\\\\\"create_time\\\\\\": \\\\\\"2023-12-07 15:39:13\\\\\\",\\\\n      \\\\\\"create_user\\\\\\": \\\\\\"739ade44-7e83-48a2-8c60-9a7c1e9f3d0a\\\\\\",\\\\n      \\\\\\"is_delete\\\\\\": \\\\\\"0\\\\\\",\\\\n      \\\\\\"order_index\\\\\\": 10,\\\\n      \\\\\\"parent_dir\\\\\\": \\\\\\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\\\\\",\\\\n      \\\\\\"sub_title\\\\\\": \\\\\\"文档保存\\\\\\",\\\\n      \\\\\\"title\\\\\\": \\\\\\"doc.doc_save\\\\\\",\\\\n      \\\\\\"type\\\\\\": \\\\\\"service\\\\\\"\\\\n   }\\\\n}\\",\\n\\t\\t\\t\\"code_desc\\": \\"文档编辑。注意文档完全是通过对比字段规则，保存的，不是通过数据加字段直接更新，加完字段modify.json 记得加规则\\",\\n\\t\\t\\t\\"code_result\\": \\"{\\\\n\\\\t\\\\\\"status\\\\\\": 0,\\\\n\\\\t\\\\\\"count\\\\\\": 0,\\\\n\\\\t\\\\\\"success\\\\\\": true,\\\\n\\\\t\\\\\\"code\\\\\\": \\\\\\"0\\\\\\",\\\\n\\\\t\\\\\\"msg\\\\\\": \\\\\\"成功\\\\\\",\\\\n\\\\t\\\\\\"data\\\\\\": {\\\\n\\\\t\\\\t\\\\\\"base_modify_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\\\"change_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\\\"demo_add_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\\\"demo_modify_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\\\"demo_remove_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\\\"important_add_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\\\"important_remove_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\\\"local_doc_detail\\\\\\": {\\\\n\\\\t\\\\t\\\\t\\\\\\"demo\\\\\\": [],\\\\n\\\\t\\\\t\\\\t\\\\\\"doc\\\\\\": {\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"code\\\\\\": \\\\\\"{\\\\\\\\n   \\\\\\\\\\\\\\"service\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"config.doc_save\\\\\\\\\\\\\\",\\\\\\\\n   \\\\\\\\\\\\\\"doc\\\\\\\\\\\\\\": {\\\\\\\\n      \\\\\\\\\\\\\\"title\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"doc.doc_save\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"sub_title\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"code\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"order_index\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"10\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"type\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"service\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"parent_dir\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"code_desc\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"\\\\\\\\\\\\\\",\\\\\\\\n      \\\\\\\\\\\\\\"code_result\\\\\\\\\\\\\\": \\\\\\\\\\\\\\"\\\\\\\\\\\\\\"\\\\\\\\n   },\\\\\\\\n   \\\\\\\\\\\\\\"important_list\\\\\\\\\\\\\\": [],\\\\\\\\n   \\\\\\\\\\\\\\"params\\\\\\\\\\\\\\": [],\\\\\\\\n   \\\\\\\\\\\\\\"demo\\\\\\\\\\\\\\": [],\\\\\\\\n   \\\\\\\\\\\\\\"result\\\\\\\\\\\\\\": []\\\\\\\\n}\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"code_desc\\\\\\": \\\\\\"文档新增\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"code_result\\\\\\": \\\\\\"\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"collect_doc_id\\\\\\": \\\\\\"768f837a-dd60-49ef-bc64-a558eb48bb70\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"create_time\\\\\\": \\\\\\"2023-12-07 15:39:13\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"create_user\\\\\\": \\\\\\"739ade44-7e83-48a2-8c60-9a7c1e9f3d0a\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"is_delete\\\\\\": \\\\\\"0\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"order_index\\\\\\": 10,\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"parent_dir\\\\\\": \\\\\\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"sub_title\\\\\\": \\\\\\"文档保存\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"title\\\\\\": \\\\\\"doc.doc_save\\\\\\",\\\\n\\\\t\\\\t\\\\t\\\\t\\\\\\"type\\\\\\": \\\\\\"service\\\\\\"\\\\n\\\\t\\\\t\\\\t},\\\\n\\\\t\\\\t\\\\t\\\\\\"important_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\t\\\\\\"params\\\\\\": []\\\\n\\\\t\\\\t},\\\\n\\\\t\\\\t\\\\\\"params_modify_list\\\\\\": [],\\\\n\\\\t\\\\t\\\\\\"params_remove_list\\\\\\": []\\\\n\\\\t}\\\\n}\\",\\n\\t\\t\\t\\"collect_doc_id\\": \\"212b5c89-67ce-4f06-bc23-24421e866974\\",\\n\\t\\t\\t\\"create_time\\": \\"2023-12-07 15:44:14\\",\\n\\t\\t\\t\\"create_user\\": \\"739ade44-7e83-48a2-8c60-9a7c1e9f3d0a\\",\\n\\t\\t\\t\\"is_delete\\": \\"0\\",\\n\\t\\t\\t\\"order_index\\": 20,\\n\\t\\t\\t\\"parent_dir\\": \\"c09992bf-afc8-4abb-a81b-24ff1060fe0b\\",\\n\\t\\t\\t\\"sub_title\\": \\"文档编辑\\",\\n\\t\\t\\t\\"title\\": \\"config.doc_edit\\",\\n\\t\\t\\t\\"type\\": \\"service\\"\\n\\t\\t},\\n\\t\\t\\"important_list\\": [],\\n\\t\\t\\"params\\": []\\n\\t}\\n}",
				"collect_doc_id": "7da94650-209f-411e-aa0c-47989dbd4409",
				"create_time": "2023-12-07 16:07:07",
				"create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
				"is_delete": "0",
				"order_index": 30,
				"parent_dir": "c09992bf-afc8-4abb-a81b-24ff1060fe0b",
				"sub_title": "查询详情",
				"title": "config.doc_detail",
				"type": "service"
			}
		],
		"create_time": "2023-12-07 15:27:04",
		"create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
		"desc": "",
		"doc_group_id": "c09992bf-afc8-4abb-a81b-24ff1060fe0b",
		"is_delete": "0",
		"name": "文档管理",
		"order_index": 70,
		"type": "service"
	}]
}');
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'sql', 'sql查询', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', 'module: sql', '   用于数据库表查询,将sql查询出来的数据作为结果。虽然sql能运行任何增删改查，但是新增、修改、删除，尽量用模型修改model_update,模型删除model_delete，模型新增用model_save。
   用渲染html的引擎，来渲染生成sql。基本上99%的场景都能渲染出来，剩下1%场景，是我没有遇到的，遇到了我会解决它。

', 1, '2023-11-25 10:47:58', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('cd14dfe4-ddfd-431e-aeef-b00db64e32f6', 'uuid', '生成uuid', 'doc', '9501b424-606a-433a-9bc5-f9de8064e9d8', '    params:
      ldap_group_id:
        template: "{{uuid}}"', '生成uuid，一般用于数据库唯一主键生成', 0, '2023-12-06 17:56:48', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('cd89c817-3ba4-48c5-baba-74c50c73dbbc', '什么是参数', '服务描述', 'doc', 'dd336894-53b6-405b-98c7-f327407d7cfa', '  - key: bulk_create_user
    module: bulk_create
    log: true
    http: true
    params:
      user_list:
        check:
          template: "{{must .user_list}}"
          err_msg: 用户列表不能为空', '    params 是一个参数池，参数是可以http 请求传来，也可以自己定义，主要存储一些key，value 字典对象，handler_params 和result_handler 可以对参数进行处理。
请求过来的参数，也是直接设置到params,params贯穿了整个服务。handler_params 也是对参数进行处理，result_handler 是运行完成时候对参数处理。
    参数是存储在service ，一个变量区域。就好比是一个大箱子，每个步骤都可以往箱子里面放东西，只要自己保证名称唯一，取的时候就不会混乱。参数定义，handler_params ，result_handler 都是往这个大箱子放与拿东西。
    params 可以做一下简单字段校验，生成默认值；比如字段不能为空；删除is_delete默认值为0，创建时间create_time默认当前时间。创建人当前用户。
', 30, '2023-12-06 19:16:30', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('ce02406f-f729-47da-a119-35ed01c6c1c3', 'arr2obj', '数组结果转对象', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    result_handler:
      - key: arr2obj
        enable: "[to_obj]"', '仅仅用于服务的结果是一个数组，将数组转对象，方便其他服务模板取值，不用从数组取，直接.xx.xx 属性', 60, '2023-12-05 09:11:35', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('d25ddb2a-2c4e-4d3f-83c3-50ddc4215de8', 'empty', '空模块', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: project_router
    http: true
    handler_params:
      - key: file2datajson
        save_field: data
      - key: param2result
        field: data
    data_file: project_router.json
    module: empty
', '空模块是个非常常用的模块，其主要目的是为了处理参数。
空模块就是主体没有做任何事情，主要在handler_params 运行你服务。比如要运行数据保存的服务，先调主体空服务，空服务在转service2field 调用你的数据保存服务，本质没有任何区别。可能就一些传参区别', 90, '2023-12-04 11:49:33', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('d4ee2257-8a51-40bf-8d48-4b5469858e21', 'data2excel', '数据转excel', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '  - key: user_list_download
    module: empty
    http: true
    excel_config: "./user_list2excel.json"
    params:
      excel_path:
        template: ''./template/{{current_date_format "20220202"}}/user_{{  replace (sub_str current_date_time -8 0) ":" ""}}_{{sub_str uuid -8 0}}.xlsx''
      response_name:
        default: "用户列表.xlsx"
    handler_params:
      - key: service2field
        service:
          service: hrm.user_list
        append_param: true
        save_field: user_list
      - key: data2excel
        path: "[excel_path]"
      - key: file2result
        path: "[excel_path]"
        result_name: "[response_name]"', '数据服务转excel,示例中只看处理器key:data2excel。实际就是导出', 120, '2023-12-05 09:51:16', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('e32693d9-a954-48ac-a1bb-097a0684c8b3', 'params2result', '多参数转结果', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '   handler_params:
      - key: params2result
        fields:
          - from: "{{.userid}}"
            to: userid
          - from: "{{.user_id}}"
            to: user_id
          - from: "{{.nick}}"
            to: nick
          - from: "{{.username}}"
            to: username', '将参数中的多个字段返回', 80, '2023-12-05 09:28:27', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('e3d5dce4-c758-41cc-9ffd-578ca39836e3', 'prevent_duplication', '防重复请求', 'doc', 'efbb8e60-eeab-4326-b742-27cedbcc9083', '  - key: doc_edit
    http: true
    module: empty
    modify_config: doc_modify.json
    log: true
    prevent_duplication:
      name: 后台防止重复请求，second 表示毫秒
      key: prevent_duplication
      enable: "{{eq (get_key \\"prevent_duplication\\") \\"true\\"}}"
      second: 300
      room: doc_dpl
      fields:
        - field: "[service]"
        - field: "[session_user_id]"', '和缓存设置的参数是一样的，只是值设置1。底层也是利用的缓存逻辑.
前台http支持重复请求，后台当然也要能支持。
之前遇到的场景就是，重复点击量2次，导致保存数据混乱，有的数据双份了', 30, '2023-12-06 21:05:07', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('eb55515c-df08-4ac5-a4e8-3d811b834e54', 'bulk_service', '批量服务', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '    module: bulk_service
    params:
      collect_doc_id:
        check:
          template: "{{must .collect_doc_id}}"
          err_msg: 文档不能为空
    data_file: service_transfer.json
    handler_params:
      - key: service2field
        service:
          service: config.get_detail_service
          collect_doc_id: "[collect_doc_id]"
        save_field: service_list
    batch:
      foreach: "[service_list]"
      service:
        service: "[service]"
      append_item_param: true
      save_field: ''result''', '    批量运行服务，多线程运行服务。
    我之前写一个发布计划接口，一个查询然后编辑的接口，结合所有表一起查询与保存。保存还好，毕竟修改数据花个3-5秒还能接受。但是一进来查询得花个3-5秒查询，基本受不了。每个模块都是查询一次，但是是顺序进行的。一个模块花0.2秒，但是架不住模块多啊，乘以20，加上模块自身需要点时间，所以时间越来越大，3-5秒很正常。
   可以利用此服务批量进行查询。将服务一次性查询出来
   比如批量去发http请求
', 20, '2023-12-03 18:09:25', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('ede82dd3-1b4f-4097-9260-c2d9c7acdf21', 'update_field', '更新普通字段', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: update_field
        name: 更新字段
        fields:
          - field: user_info
            template: "[modify_data.user_info]"
          - field: change_list
            template: "[modify_data.change_list]"', '更新params中的字段', 30, '2023-12-04 20:32:08', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('eee7f650-4583-4bf6-b770-ce853eba8c54', 'field2array', '字段转数组', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: field2array
        field: "[ids]"
        enable: "{{must .ids}}"
        save_field: ldap_group_id_list', '以逗号隔开的字符串转数字。比如前台批量删除，传一堆的ID,以逗号隔开，需要处理成数组', 200, '2023-12-05 20:03:55', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('f272ba24-4c3e-4072-be07-abae59c93907', 'prop_arr', '对象数组转数组', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: prop_arr
        foreach: "[detail_list]"
        value: "[config_detail_id]"
        save_field: config_detail_id_list', '过滤对象数组的某个字段转数组，比如传对象数组过来，过滤成简单ID数组，好根据ID删除', 230, '2023-12-05 20:24:09', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'service_flow', '服务流程化', 'doc', 'ae78bab6-b68e-4522-86e0-08d35fe201d0', '  - key: system_login
    module: service_flow
    params:
      username:
        check:
          template: "{{must .username}}"
          err_msg: 用户名不能为空
      password_md5:
        check:
          template: "{{must .password}}"
          err_msg: 密码不能为空
        template: "{{md5 .password}}"
      has_user:
        default: true
    data_json: system_login.json
    result_handler:
      - key: param2result
        field: "[session_user]"', '    像工作流一样的运行多个服务，运行一个节点服务，通过计算，运行下一个服务节点。
    服务流程化思想来源于工作流。接触到loonflow工作流，我一直试想着，我写的代码能不能像一个工作流一样流转，我们只写一小部分节点，通过工作流流转，运行到下个服务。比如我写个新建用户，然后流转到新建角色，如果中途失败，流转到删除用户、删除角色然后返回。', 120, '2023-12-04 20:19:37', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);
INSERT INTO collect_doc (collect_doc_id, title, sub_title, `type`, parent_dir, code, code_desc, order_index, create_time, create_user, is_delete, code_result) VALUES('f7350f20-20fb-47f4-8b19-511c8b5b38a9', 'check_field', '检查普通字段', 'doc', '2d07dfdc-1026-40fb-8124-ddc74b566265', '    handler_params:
      - key: check_field
        fields:
          - field: doc_collect_id
            template: "{{must .doc.collect_doc_id}}"
            err_msg: "文档ID不能为空"
          - field: type
            template: "{{must .doc.type}}"
            err_msg: "文档类型不能为空"
          - field: parent_dir
            template: "{{must .doc.parent_dir}}"
            err_msg: "文档上级目录不能为空"', '    检查字段是否合法，params 中的check只能在一开始检查已有的字段，这个可以支持handler_params中运行其他服务再检查。
    我们经常遇到xx字段不能为空', 40, '2023-12-04 20:40:29', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0', NULL);

INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('0043515a-add1-4661-b27e-98d1ea835ab0', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '添加用户', '{
  require(login.common),
  "method": "add",
  "AddParams": {
    "DN": "cn={{.add_username}},{{.ldap_config.ldap_users}},{{.ldap_config.ldap_base_dn}}",
    "Attributes": [
      {"TYPE": "cn","Vals": ["{{.add_username}}"]},
      {"TYPE": "sn","Vals": ["{{.nick}}"]},
      {"TYPE": "givenName","Vals": ["{{.nick}}"]},
      {{if .email}}
      {"TYPE": "mail","Vals": ["{{.email}}"]},
      {{ end }}
      {{if .phone}}
      {"TYPE": "mobile","Vals": ["{{.phone}}"]},
      {{end}}
      {"TYPE": "objectClass","Vals": ["top", "inetOrgPerson"]},
      {"TYPE": "userpassword","Vals": ["{{hash_sha .add_password}}"]}
    ]
  }

}', 40, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('02ffde1c-6727-49bc-b38a-4d0f35266b35', 'd89f00cd-2ecf-43d8-9fb4-f9bd26c1cf05', '111', '111', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('0418c470-b26d-420a-a5b0-8f1c583c2b57', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '删除用户', '{
  require(login.common),
  "method": "delete",
  "DeleteParams": {
    "DN": "cn={{.remove_username}},{{.ldap_config.ldap_users}},{{.ldap_config.ldap_base_dn}}"
  }

}', 80, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('0491a89b-d596-4175-995c-a1a41d5cfb39', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '文档对比', '{
  "desc": "替换前和替换后，概念调整下，方便对比，从左往右，传过来的数据，是前段的数据，left。已有的后台数据是right",
  "left_save_field": "after",
  "right_save_field": "before",
  "op_field_transfer": {
    "name": "op_name"
  },
  "fields": [
    {"rule": "compare_field_value", "field": "[title]", "name": "标题", "left": "[doc]", "right": "[local_doc_detail.doc]", "operation": "modify_base", "append_right_fields": ["[collect_doc_id]"]},
    {"rule": "compare_field_value", "field": "[sub_title]", "name": "子标题", "left": "[doc]", "right": "[local_doc_detail.doc]", "operation": "modify_base", "append_right_fields": ["[collect_doc_id]"]},
    {"rule": "compare_field_value", "field": "[code]", "name": "代码", "left": "[doc]", "right": "[local_doc_detail.doc]", "operation": "modify_base", "append_right_fields": ["[collect_doc_id]"]},
    {"rule": "compare_field_value", "field": "[code_desc]", "name": "代码描述", "left": "[doc]", "right": "[local_doc_detail.doc]", "operation": "modify_base", "append_right_fields": ["[collect_doc_id]"]},
    {"rule": "compare_field_value", "field": "[parent_dir]", "name": "上级目录", "left": "[doc]", "right": "[local_doc_detail.doc]", "operation": "modify_base", "append_right_fields": ["[collect_doc_id]"]},
    {"rule": "compare_field_value", "field": "[type]", "name": "类型", "left": "[doc]", "right": "[local_doc_detail.doc]", "operation": "modify_base", "append_right_fields": ["[collect_doc_id]"]},
    {"rule": "compare_field_value", "field": "[order_index]", "name": "排序", "left": "[doc]", "right": "[local_doc_detail.doc]", "operation": "modify_base", "append_right_fields": ["[collect_doc_id]"]},
    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[demo]", "name": "示例名称", "left": "[demo]", "right": "[local_doc_detail.demo]", "with_add_remove": true, "left_value_field": "[name]", "right_value_field": "[name]", "operation": "modify", "append_right_fields": ["[*]"],"append_left_fields": ["[*]"]},
    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[demo]", "name": "示例代码", "left": "[demo]", "right": "[local_doc_detail.demo]",  "left_value_field": "[code]", "right_value_field": "[code]", "operation": "modify", "append_left_fields": ["[*]"]},
    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[demo]", "name": "示例排序", "left": "[demo]", "right": "[local_doc_detail.demo]",  "left_value_field": "[order_index]", "right_value_field": "[order_index]", "operation": "modify", "append_left_fields": ["[*]"]},

    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[params]", "name": "参数名称", "left": "[params]", "right": "[local_doc_detail.params]", "with_add_remove": true, "left_value_field": "[name]", "right_value_field": "[name]", "operation": "modify", "append_right_fields": ["[*]"]},
    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[params]", "name": "参数类型", "left": "[params]", "right": "[local_doc_detail.params]",  "left_value_field": "[type]", "right_value_field": "[type]", "operation": "modify", "append_left_fields": ["[*]"]},
    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[params]", "name": "参数是否必须", "left": "[params]", "right": "[local_doc_detail.params]",  "left_value_field": "[must]", "right_value_field": "[must]", "operation": "modify", "append_left_fields": ["[*]"]},
    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[params]", "name": "参数描述", "left": "[params]", "right": "[local_doc_detail.params]",  "left_value_field": "[desc]", "right_value_field": "[desc]", "operation": "modify", "append_left_fields": ["[*]"]},
    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[params]", "name": "参数排序", "left": "[params]", "right": "[local_doc_detail.params]",  "left_value_field": "[order_index]", "right_value_field": "[order_index]", "operation": "modify", "append_left_fields": ["[*]"]},

    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[important_list]", "name": "要点", "left": "[important_list]", "right": "[local_doc_detail.important_list]", "with_add_remove": true, "left_value_field": "[name]", "right_value_field": "[name]", "operation": "modify", "append_right_fields": ["[*]"]},
    {"rule": "array_obj_value", "left_field": "[name]", "right_field": "[name]", "field": "[important_list]", "name": "要点排序", "left": "[important_list]", "right": "[local_doc_detail.important_list]",  "left_value_field": "[order_index]", "right_value_field": "[order_index]", "operation": "modify","append_left_fields": ["[*]"]}
  ]
}', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('054c47ed-a6d3-427d-8d61-23fd6e1a073f', 'bd6431c8-caf3-4ae9-96af-03cce615204b', '批量服务的对象生成', '  - key: get_detail_service
    module: empty
    params:
      collect_doc_id:
        check:
          template: "{{must .collect_doc_id}}"
          err_msg: 文档不能为空
    data_file: doc_detail.json
    handler_params:
      - key: file2datajson
        save_field: service_list
      - key: param2result
        field: service_list', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('09ae924f-be53-436f-9801-867fe0553607', 'e32693d9-a954-48ac-a1bb-097a0684c8b3', '使用[变量]', '
      - key: params2result
        fields:
          - from: "[change_list]"
            to: change_list
          - from: "[local_detail_list]"
            to: local_detail_list
', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('0ef42763-bdc4-4e5d-ad49-cc62b7cc3266', 'bd6431c8-caf3-4ae9-96af-03cce615204b', '批量服务配置文件', '[
  {
    "service": "config.doc_query",
    "collect_doc_id": "{{.collect_doc_id}}",
    "target_field": "doc",
    "to_obj":true
  },
  {
    "service": "config.import_list_query",
    "collect_doc_id": "{{.collect_doc_id}}",
    "target_field": "important_list"
  },
  {
    "service": "config.params_query",
    "collect_doc_id": "{{.collect_doc_id}}",
    "target_field": "params"
  },
  {
    "service": "config.demo_query",
    "collect_doc_id": "{{.collect_doc_id}}",
    "target_field": "demo"
  }
]
', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('10cd019e-a2a0-44dd-a59b-27d57d70c5ca', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', '创建时间默认当前时间', '    params:
      create_time:
        template: "{{current_date_time}}"', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('10d074b1-8cea-4113-bd91-8a43d876b732', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', '创建用户的流程', '{
  "services": [
    {
      "node_key": "start",
      "node_type": "start",
      "name": "开始",
      "node_next": "create_user"
    },
    {
      "node_key": "create_user",
      "name": "创建用户信息",
      "node_type": "node",
      "key": "service2field",
      "append_param": true,
      "service": {
        "service": "hrm.create_user"
      },
      "node_fail": "end",
      "node_next": "{{ if .roles }}create_roles{{ else if eq .create_ldap \\"1\\"}}create_ldap{{ else }}end{{ end }}"
    },
    {
      "node_key": "create_roles",
      "key": "service2field",
      "name": "创建角色",
      "service": {
        "service": "hrm.user_role_add",
        "roles": "[roles]",
        "user_id": "[user_id]"
      },
      "node_type": "node",
      "node_next": "{{ if eq .create_ldap \\"1\\" }} create_ldap {{ else }} end {{ end }}",
      "node_fail": "end"
    },
    {
      "node_key": "create_ldap",
      "node_type": "node",
      "name": "ldap 登陆",
      "key": "service2field",
      "service": {
        "service": "hrm.create_ldap_with_group",
        "add_username": "[username]",
        "add_password": "[password]"
      },
      "append_param": true,
      "node_fail": "end",
      "node_next": "end"
    },
    {
      "node_key": "end",
      "node_type": "end",
      "name": "结束"
    }
  ]
}', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('10d414d1-afc6-47c2-ab35-10ff9f0b17b0', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', '动态控制count是否运行', '  - key: role_query
    http: true
    module: sql
    params:
      page:
        type: int
        default: 1
      size:
        default: 20
        type: int
      start:
        template: " ({{.page}}-1) * {{.size}}"
        type: int
      pagination:
        default: true
      count:
        default: true
    data_file: role_query.sql
    count: "[count]"
    count_file: role_query_count.sql
    pagination: pagination', 70, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('12683ac2-e889-473f-8131-5a8da1007d38', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '搜索用户信息', '{
  require(login.common),
  "method": "search",
  "scope_desc": "ScopeBaseObject   = 0\\n\\tScopeSingleLevel  = 1\\n\\tScopeWholeSubtree = 2",
  "SearchParams": {
    "BaseDN": "{{.ldap_config.ldap_users}},{{.ldap_config.ldap_base_dn}}",
    "Scope":2,
    "DerefAliases": 0,
    "filter": "(&(objectClass=inetOrgPerson)(cn={{.search_username}}))",
    "attributes": [
      "cn",
      "sn",
      "givenName",
      "mail",
      "mobile"
    ]
  }

}', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('174032e1-5d69-453f-ba00-4a381ea74c90', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', '生成批量服务的服务配置', '[
  {
    "service": "config.doc_query",
    "collect_doc_id": "{{.collect_doc_id}}",
    "target_field": "doc",
    "to_obj":true
  },
  {
    "service": "config.import_list_query",
    "collect_doc_id": "{{.collect_doc_id}}",
    "target_field": "important_list"
  },
  {
    "service": "config.params_query",
    "collect_doc_id": "{{.collect_doc_id}}",
    "target_field": "params"
  },
  {
    "service": "config.demo_query",
    "collect_doc_id": "{{.collect_doc_id}}",
    "target_field": "demo"
  }
]
', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('23ebe2de-bc51-47e1-a977-32c290b5810d', '45d1d393-6758-4c37-8ba9-108493f67b8e', '拼接参数所有字段', '    handler_params:
      - key: service2field
        service:
          service: hrm.ldap_add
        append_param: true', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('26215f23-e179-42b3-a521-61dc8b0d8e47', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'base.sql 文件示例,包含简单key、value示例，分页示例', 'select a.*
from role a
where 1=1
{{ if .search }}
and (
a.role_name like {{.search}}
or a.role_code like {{.search}}
)
{{ end }}
order by a.order_index desc
{{ if  .pagination  }}
limit {{.start}} , {{.size}}
{{ end }}', 40, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('28112c48-0039-48d3-9218-3547eddb97bc', 'd89f00cd-2ecf-43d8-9fb4-f9bd26c1cf05', 'ddfs', '', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('316e7202-67a9-478f-9d09-336610c7ff6b', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', 'data_file中service_transfer.json注册文件示例', '{
  "config.doc_query": "config.doc_query",
  "config.import_list_query": "config.import_list_query",
  "config.params_query": "config.params_query",
  "config.demo_query": "config.demo_query"
}', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('3315763d-6b7e-4c9a-b9f4-956481b4e156', 'c00a6f14-98c1-4308-8017-cf35ae300de4', '调用nexus的配置，basic 认证', '{
  "url": "http://172.26.0.19:8081/service/extdirect",
  "method": "POST",
  "header": {
    "content-type": "application/json"
  },
  "basic_auth": {
    "username": "zhangzhi",
    "password": "zhang@888"
  },
  "data": {
    "action": "coreui_Search",
    "method": "read",
    "data": [
      {
        "page": 1,
        "start": 0,
        "limit": 5,
        "sort": [
          {
            "property": "version",
            "direction": "DESC"
          }
        ],
        "filter": [
          {
            "property": "group.raw",
            "value": "com.wghis.frontend"
          },
          {
            "property": "name.raw",
            "value": "mis-ui"
          }
        ]
      }
    ],
    "type": "rpc",
    "tid": 20
  }
}', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('33f97374-6a24-407e-af96-a154cde2e2e8', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', '默认当前登陆用户', '      create_user:
        template: "{{.session_user_id}}"', 40, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('34e99073-7ca1-4bf4-8eef-7dbf96c557f0', '91d4abd1-14c9-4705-9de7-1518da998427', 'collect/hrm/service.yml', '# 项目点路由
service:
  - name: "用户管理"
    path: "user/index.yml"
  - name: "角色管理"
    path: "role/index.yml"


  - name: "ldap"
    path: "ldap/index.yml"

  - name: "login"
    path: "login/index.yml"
  - name: user_flow
    path: "user_flow/index.yml"
  - name: user_role
    path: "user_role/index.yml"
  - name: ldap_group
    path: "ldap_group/index.yml"
  - name: role_ldap_group
    path: "role_ldap_group/index.yml"

  - name: user_change_history
    path: "user_change_history/index.yml"', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('40b6ff9a-5053-4ba4-95d7-d4f979918535', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', '默认值', '      is_delete:
        default: 1', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('4403c418-8dfd-4964-bbb5-5ddb7d3f5417', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '删除分组', '{
  require(login.common),
  "method": "delete",
  "DeleteParams": {
    "DN": "ou={{.name}},{{.ldap_config.ldap_groups}},{{.ldap_config.ldap_base_dn}}"
  }

}', 90, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('45a728d1-d9a1-4e09-a010-fffb76b67d4f', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', '生成批量服务接口示例config.get_detail_service', '  - key: get_detail_service
    module: empty
    params:
      collect_doc_id:
        check:
          template: "{{must .collect_doc_id}}"
          err_msg: 文档不能为空
    data_file: doc_detail.json
    handler_params:
      - key: file2datajson
        save_field: service_list
      - key: param2result
        field: service_list', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('47572bff-026b-427c-929c-32c72457019e', '8ec4053f-1a54-4f78-9de6-31db6995692e', '前台传一个数组对象，过滤生简单数组', '  - key: collect_doc_demo_remove
    module: model_delete
    table: collect_doc_demo
    params:
      demo_list:
        check:
          template: "{{must .demo_list}}"
          err_msg: 示例代码不能为空
    handler_params:
      - key: prop_arr
        foreach: "[demo_list]"
        value: "[doc_demo_id]"
        save_field: doc_demo_id_list
    filter:
      doc_demo_id__in: "[doc_demo_id_list]"', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('4c8a5f13-21b0-44ba-8201-aea6168fc182', 'c00a6f14-98c1-4308-8017-cf35ae300de4', '直接post字符串配置', '{
  "url": " http://172.26.0.237:8012/pushalertmsgzbx",
  "method": "POST",
  "data": "3151502276711|[公司云端]-172.26.0.1-[SVN]|172.26.0.1|Average|2019.12.01 03:41:32||1|[xx]产品环境服务器|[TEST001]-[公司云端]-172.26.0.1-[SVN]：测试测试|OPDATA监控细节"
}', 40, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('5721a245-14ec-4888-b3e5-32e384e4037a', '0722aa5a-78e6-4a29-818c-ba33843bdf81', '动态控制update_fields,模仿jira的批量更新', '  - key: update_user_by_user_id_list
    module: model_update
    http: true
    log: true
    params:
      fields:
        default: ["*"]
      user_id_list:
        check:
          template: "{{must .user_id_list}}"
          err_msg: 用户名不能空
    table: "user_account"
    options: "[fields]"
    filter:
      user_id__in: "[user_id_list]"
    ignore_fields:
      - password', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('5907dd4f-f149-4247-b9df-07364716cff4', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', '简单数组 文件示例', '{{ if .ldap_group_id_list }}
and ldap_group_id in ({{.ldap_group_id_list}})
{{ end }}', 50, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('59102662-029e-4f67-92ce-e9353a15ea0d', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '添加分组', '{
  require(login.common),
  "method": "add",
  "AddParams": {
    "DN": "ou={{.name}},{{.ldap_config.ldap_groups}},{{.ldap_config.ldap_base_dn}}",
    "Attributes": [
      {"TYPE": "cn","Vals": ["{{.name}}"]},
      {"TYPE": "ou","Vals": ["{{.name}}"]},
      {"TYPE": "objectClass","Vals": ["top", "groupOfUniqueNames"]},
      {"TYPE": "uniqueMember","Vals": [""]}
    ]
  }

}', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('5d68c76c-31ad-4ce4-a26e-5ffcf965cf14', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', '用户请求', '{
	"service": "hrm.user_list",
	"page": 1,
	"size": 20
}', 10, '{
	"status": 0,
	"count": 127,
	"success": true,
	"code": "0",
	"msg": "执行成功",
	"data": [
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2023-11-03 17:51:12",
			"create_user": "739ade44-7e83-48a2-8c60-9a7c1e9f3d0a",
			"email": "zhangsan@weigaogroup.com",
			"entry_date": "2023-03-11",
			"is_delete": "0",
			"ladp_user_login_id": "",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2023-11-03 17:51:12",
			"modify_user": "",
			"nick": "张三1",
			"password": "01d7f40760960e7bd9443513f22ab9af",
			"phone": "11111111111",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "a33b7fd0-50db-4d55-a29e-573c947721bd",
			"user_name": "zhangsan",
			"user_status": "regular",
			"user_status_name": "正式",
			"username": "zhangsan",
			"wechat_id": "",
			"work_code": "zhangsan"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-05-05 08:52:37",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "tanjingcheng@weigaogroup.com",
			"entry_date": "2022-05-05",
			"is_delete": "0",
			"ladp_user_login_id": "tanjingcheng",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-05-07 15:17:11",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "谭晶城",
			"password": "05d7a0795cf0ef9dfdcee9a3dfd36fe9",
			"phone": "15116423684",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "6f2b3a9f-7ccd-47f5-a5fd-f7988b760127",
			"user_name": "tanjingcheng",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "tanjingcheng",
			"wechat_id": "",
			"work_code": "00119297"
		},
		{
			"attendance_id": "",
			"create_ldap": "0",
			"create_time": "2022-04-26 09:22:46",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "xieyaoyi@weigaogroup.com",
			"entry_date": "2022-04-26",
			"is_delete": "0",
			"ladp_user_login_id": "xieyaoyi",
			"leave_date": "2022-04-26",
			"leave_reason": "不能接受长期出差。",
			"modify_time": "2022-04-27 09:13:21",
			"modify_user": "-1",
			"nick": "谢耀一",
			"password": "c17d7a3905b97ced817f161b21249212",
			"phone": "18976219111",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "93048d5d-198c-4233-91ff-ee6c24d5bb5f",
			"user_name": "xieyaoyi",
			"user_status": "leave",
			"user_status_name": "离职",
			"username": "xieyaoyi",
			"wechat_id": "",
			"work_code": "00119157"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-25 08:46:52",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "hupeng@weigaogroup.com",
			"entry_date": "2022-04-25",
			"is_delete": "0",
			"ladp_user_login_id": "hupeng",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-27 09:13:50",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "胡鹏",
			"password": "0e5768efeb54bdb001cb41b63e8c639e",
			"phone": "13755123980",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "120edd1e-1f4d-4efe-a19b-4c047b61486b",
			"user_name": "hupeng",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "hupeng",
			"wechat_id": "",
			"work_code": "00119109"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-20 09:33:43",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "hecong@weigaogroup.com",
			"entry_date": "2022-04-20",
			"is_delete": "0",
			"ladp_user_login_id": "hecong",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-22 18:05:13",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "何聪",
			"password": "9f46ad5337e0c8285768f7e93d32018f",
			"phone": "15874772562",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "02125451-accb-4acb-b195-435cf42c7d90",
			"user_name": "hecong",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "hecong",
			"wechat_id": "",
			"work_code": "00118975"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-20 09:00:26",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "fuhui@weigaogroup.com",
			"entry_date": "2022-04-20",
			"is_delete": "0",
			"ladp_user_login_id": "fuhui",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-22 18:05:33",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "付辉",
			"password": "a6f54db57b98f3dd767e851e98d9558b",
			"phone": "18574721894",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "daf3dfd4-24de-4c84-9bfc-679ebd9cd0c2",
			"user_name": "fuhui",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "fuhui",
			"wechat_id": "",
			"work_code": "00118976"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-15 09:03:07",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "xionggang@weigaogroup.com",
			"entry_date": "2022-04-15",
			"is_delete": "0",
			"ladp_user_login_id": "xionggang",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-19 11:21:31",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "熊刚",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "13627414094",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "63826a79-3a61-4d90-865f-a480500c8724",
			"user_name": "xionggang",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "xionggang",
			"wechat_id": "",
			"work_code": "00118907"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-14 11:38:33",
			"create_user": "fc46eedc-f5bb-463f-be37-78f08b362828",
			"email": "whslyy@weigaogroup.com",
			"entry_date": "2022-04-14",
			"is_delete": "0",
			"ladp_user_login_id": "menxinyan",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-14 11:38:33",
			"modify_user": "fc46eedc-f5bb-463f-be37-78f08b362828",
			"nick": "门新烟",
			"password": "1c952fbfe31b87c1306bef5bcdf9f935",
			"phone": "1223",
			"role_names": "普通员工,外部威海市立医院",
			"roles": "common,wbweihai",
			"user_id": "6606c08b-6192-4a0c-8f77-67b9cdb33196",
			"user_name": "menxinyan",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "menxinyan",
			"wechat_id": "",
			"work_code": "91002"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-14 11:37:06",
			"create_user": "fc46eedc-f5bb-463f-be37-78f08b362828",
			"email": "whslyy@weigaogroup.com",
			"entry_date": "2022-04-14",
			"is_delete": "0",
			"ladp_user_login_id": "hexinghui",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-14 11:52:31",
			"modify_user": "fc46eedc-f5bb-463f-be37-78f08b362828",
			"nick": "贺幸辉",
			"password": "38f22b7c3b8c343d4315dab636b4e1d0",
			"phone": "111",
			"role_names": "外部威海市立医院",
			"roles": "wbweihai",
			"user_id": "1c85ba13-9911-42f3-be7c-04fc6b8ffea6",
			"user_name": "hexinghui",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "hexinghui",
			"wechat_id": "",
			"work_code": "99001"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-13 08:51:36",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "kuangshengkun@weigaogroup.com",
			"entry_date": "2022-04-13",
			"is_delete": "0",
			"ladp_user_login_id": "kuangshengkun",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-15 09:04:12",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "匡胜昆",
			"password": "6c6687312b92e143444d6fe19e91c18e",
			"phone": "17773102617",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "44c76358-2925-4601-86cf-93c746176f11",
			"user_name": "kuangshengkun",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "kuangshengkun",
			"wechat_id": "",
			"work_code": "00118860"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-13 08:48:28",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "zhouyu@weigaogroup.com",
			"entry_date": "2022-04-13",
			"is_delete": "0",
			"ladp_user_login_id": "zhouyu",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-15 09:05:45",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "周愉",
			"password": "5b0c7138e1fbe31b726a0744bb9747f5",
			"phone": "13548667796",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "0baa8cd8-c711-47cc-956a-cc5985873b11",
			"user_name": "zhouyu",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "zhouyu",
			"wechat_id": "",
			"work_code": "00118862"
		},
		{
			"attendance_id": "",
			"create_ldap": "0",
			"create_time": "2022-04-08 08:43:36",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "wangqian@weigaogroup.com",
			"entry_date": "2022-04-08",
			"is_delete": "0",
			"ladp_user_login_id": "wangqian",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-12 14:59:52",
			"modify_user": "-1",
			"nick": "王倩",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "18098934021",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "84d60fff-fd05-4375-9611-9aabdedde351",
			"user_name": "wangqian",
			"user_status": "leave",
			"user_status_name": "离职",
			"username": "wangqian",
			"wechat_id": "",
			"work_code": "00118759"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-04-06 08:30:47",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "baibing@weigaogroup.com",
			"entry_date": "2022-04-06",
			"is_delete": "0",
			"ladp_user_login_id": "baibing",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-06 16:33:32",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "白冰",
			"password": "3231e0617fdde4cb9193bb285b402777",
			"phone": "18108450107",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "73d5b5f6-d363-4096-bf25-a37f90d655c7",
			"user_name": "baibing",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "baibing",
			"wechat_id": "",
			"work_code": "00118721"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-31 08:55:38",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "lvting@weigaogroup.com",
			"entry_date": "2022-03-31",
			"is_delete": "0",
			"ladp_user_login_id": "lvting",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-04-02 14:50:43",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "吕婷",
			"password": "bfb1d909e7a66c52addc9324cb916da3",
			"phone": "15084920749",
			"role_names": "普通员工,产品",
			"roles": "common,product",
			"user_id": "4d8e7b87-7d5a-4558-b4f6-d0918c4a7cc5",
			"user_name": "lvting",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "lvting",
			"wechat_id": "",
			"work_code": "00118646"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-29 08:38:31",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "xiaochunping@weigaogroup.com",
			"entry_date": "2022-03-29",
			"is_delete": "0",
			"ladp_user_login_id": "xiaochunping",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-29 16:35:06",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "肖春平",
			"password": "3a9b5304df5de24d8259f998232d88cf",
			"phone": "17508418876",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "5ff44669-6691-4027-9dc0-7f4a7c43bb2c",
			"user_name": "xiaochunping",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "xiaochunping",
			"wechat_id": "",
			"work_code": "00118574"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-28 15:04:59",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "guanhongqiang@weigaogroup.com",
			"entry_date": "2022-03-28",
			"is_delete": "0",
			"ladp_user_login_id": "guanhongqiang",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-28 15:09:19",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "官红强",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "15243651253",
			"role_names": "普通员工,研发",
			"roles": "common,development",
			"user_id": "b6165b59-5e28-4f02-b7ee-4516229c37b9",
			"user_name": "guanhongqiang",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "guanhongqiang",
			"wechat_id": "",
			"work_code": "00118553"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-28 13:49:44",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "hejian@weigaogroup.com",
			"entry_date": "2022-03-28",
			"is_delete": "0",
			"ladp_user_login_id": "hejian",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-28 15:08:52",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "贺健",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "15574716071",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "bd03d6a6-658c-4fdd-bdd5-9cee31edbdf5",
			"user_name": "hejian",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "hejian",
			"wechat_id": "",
			"work_code": "00118563"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-28 08:50:34",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "xutianyi@weigaogroup.com",
			"entry_date": "2022-03-28",
			"is_delete": "0",
			"ladp_user_login_id": "xutianyi",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-29 16:34:30",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "徐天仪",
			"password": "f001d5653db7748091388b6983644a7d",
			"phone": "18374601038",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "312e3f8e-e761-4cf5-8416-a4fa471f60a0",
			"user_name": "xutianyi",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "xutianyi",
			"wechat_id": "",
			"work_code": "00118564"
		},
		{
			"attendance_id": "",
			"create_ldap": "1",
			"create_time": "2022-03-28 08:47:55",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "wenqia@weigaogroup.com",
			"entry_date": "2022-03-28",
			"is_delete": "0",
			"ladp_user_login_id": "wenqia",
			"leave_date": "",
			"leave_reason": "",
			"modify_time": "2022-03-28 15:07:00",
			"modify_user": "49ce12c3-b343-43ba-8478-335508726966",
			"nick": "文洽",
			"password": "1b594c30e94c8843e65a2e808e1d1000",
			"phone": "13667353867",
			"role_names": "普通员工,交付",
			"roles": "common,deliver",
			"user_id": "37ef8407-5801-4f6b-bd45-ddfeab848276",
			"user_name": "wenqia",
			"user_status": "trial",
			"user_status_name": "试用",
			"username": "wenqia",
			"wechat_id": "",
			"work_code": "00118562"
		},
		{
			"attendance_id": "",
			"create_ldap": "0",
			"create_time": "2022-03-25 08:55:28",
			"create_user": "49ce12c3-b343-43ba-8478-335508726966",
			"email": "chenjinke@weigaogroup.com",
			"entry_date": "2022-03-25",
			"is_delete": "0",
			"ladp_user_login_id": "chenjinke",
			"leave_date": "2022-03-28",
			"leave_reason": "个人原因",
			"modify_time": "2022-03-29 09:58:09",
			"modify_user": "-1",
			"nick": "陈金科",
			"password": "0a113ef6b61820daa5611c870ed8d5ee",
			"phone": "18390189893",
			"role_names": "普通员工",
			"roles": "common",
			"user_id": "33826318-f291-4361-aeab-de6e9d80800f",
			"user_name": "chenjinke",
			"user_status": "leave",
			"user_status_name": "离职",
			"username": "chenjinke",
			"wechat_id": "",
			"work_code": "00118505"
		}
	]
}');
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('5f1cad35-7b2e-4daa-b81c-0af34dba0797', 'd4ee2257-8a51-40bf-8d48-4b5469858e21', 'data转excel示例', '{
  "name": "将用户列表转excel",
  "sheets": [
    {
      "sheet_index": 0,
      "title_height": 48,
      "desc": "16*行数",
      "title_style": {
        "font": {
          "family": "宋体",
          "size": 12,
          "color": "#ffffff"
        },
        "fill": {
          "type": "pattern",
          "color": [
            "#3366FF"
          ],
          "pattern": 1
        },
        "alignment": {
          "vertical": "center",
          "WrapText": true
        }
      },
      "name_style": {
        "font": {
          "family": "Arial",
          "size": 10,
          "color": "#ffffff"
        },
        "fill": {
          "type": "pattern",
          "color": [
            "#0066CC"
          ],
          "pattern": 1
        },
        "alignment": {
          "vertical": "center"
        }
      },
      "content_style": {
        "NumFmt": 49
      },
      "title": "用户列表\\r\\n1.ID 是作为系统的唯一标志，如果id存在则编辑数据，空着则新增\\r\\n2.带* 是必须保存",
      "data": "[user_list]",
      "fields": [
        {
          "field": "[userid]",
          "name": "id",
          "width": 10
        },
        {
          "field": "[nick]",
          "name": "*昵称",
          "width": 20
        },
        {
          "field": "[username]",
          "name": "*用户名",
          "width": 20
        },
        {
          "field": "[tel]",
          "name": "电话号码",
          "width": 20
        },
        {
          "field": "[email]",
          "name": "邮箱",
          "width": 20
        },
        {
          "field": "[statu]",
          "template": "{{ if  eq .statu 1}} 正常 {{ else }}删除{{ end }}",
          "name": "状态",
          "width": 10
        },
        {
          "field": "[address]",
          "name": "地址",
          "width": 50
        },
        {
          "field": "[create_time]",
          "name": "创建日期",
          "width": 20
        }
      ]
    }
  ]
}', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('648c2085-769d-498b-ad77-57d74fd52e26', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '修改ou', '{
  require(login.common),
  "method": "modifyDn",
  "ModifyDnParams": {
    "DN": "ou={{.old_name}},{{.ldap_config.ldap_groups}},{{.ldap_config.ldap_base_dn}}",
    "NewRDN": "ou={{.new_name}}",
    "DeleteOldRDN": true
  }
}', 100, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('657afc82-e19a-4564-a623-d9251399a8cc', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '分组添加用户', '{
  require(login.common),
  "method": "modify",
  "ModifyParams": {
    "DN": "ou={{.name}},{{.ldap_config.ldap_groups}},{{.ldap_config.ldap_base_dn}}",
    "Changes": [
      {
        "Desc": "AddAttribute       = 0\\n\\tDeleteAttribute    = 1\\n\\tReplaceAttribute   = 2\\n\\tIncrementAttribute = 3",
        "Operation": 0,
        "Modification": {
          "TYPE": "uniqueMember",
          "Vals": [
            "cn={{.add_username}},{{.ldap_config.ldap_groups}},{{.ldap_config.ldap_base_dn}}"
          ]
        }
      }

    ]
  }
}', 60, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('6f607630-7ff2-49d7-9fa4-af92e6399aca', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'common 公共文件', '{
  "connection": {
    "server": "{{.ldap_config.ldap_addr}}",
    "user": "cn={{.username}},{{if .is_not_admin}}{{.ldap_config.ldap_users}},{{end}}{{.ldap_config.ldap_base_dn}}",
    "password": "{{.password}}"
  }
}', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('6fd5b450-1a4c-4e02-9e3c-6d19300b300d', '0722aa5a-78e6-4a29-818c-ba33843bdf81', '忽略密码字段', '  - key: update_user_by_user_id_list
    module: model_update
    http: true
    log: true
    params:
      fields:
        default: ["*"]
      user_id_list:
        check:
          template: "{{must .user_id_list}}"
          err_msg: 用户名不能空
    table: "user_account"
    options: "[fields]"
    filter:
      user_id__in: "[user_id_list]"
    ignore_fields:
      - password', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('736b72ef-71ae-4cb0-84fa-4359768c79dd', '68ac5f58-6e58-4187-91c1-c7bc48152a31', '11', '11', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('88b2ce52-220e-4a04-9496-ab6e2ffad3de', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', '配置文件示例', '  - key: user_list
    module: sql
    http: true
    params:
      search:
        template: "{{ if .search }}%{{.search}}%{{ end }}"
      page:
        type: int
        default: 1
      size:
        default: 20
        type: int
      start:
        template: " ({{.page}}-1) * {{.size}}"
        exec: true
        type: int
      pagination:
        default: true
      count:
        default: true
    data_file: user_list.sql
    count_file: user_list_count.sql
    pagination: "[pagination]"
    count: "[count]"
', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('8b6d51fd-854c-4b42-8de9-13f367c9f686', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '简单字段对比', '{
      "rule": "compare_field_value",
      "field": "[create_ldap]",
      "name": "创建ldap",
      "right": "[user_info]",
      "operation": "modify",
      "append_right_fields": [
        "[user_id]"
      ]
    }', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('90abd94e-6926-4c15-9a02-a0dd42e4eaa2', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '分组删除用户', '{
  require(login.common),
  "method": "modify",
  "ModifyParams": {
    "DN": "ou={{.ou}},{{.ldap_config.ldap_groups}},{{.ldap_config.ldap_base_dn}}",
    "Changes": [
      {
        "Desc": "AddAttribute       = 0\\n\\tDeleteAttribute    = 1\\n\\tReplaceAttribute   = 2\\n\\tIncrementAttribute = 3",
        "Operation": 1,
        "Modification": {
          "TYPE": "uniqueMember",
          "Vals": [
            "cn={{.remove_username}},{{.ldap_config.ldap_users}},{{.ldap_config.ldap_base_dn}}"
          ]
        }
      }

    ]
  }
}', 70, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('99f3fc97-3fee-4c16-ae75-a5fceadcd77e', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', '不能为空校验', '      doc_group_list:
        check:
          template: "{{must .doc_group_list}}"
          err_msg: 分组不能为空', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('9cbad367-89f8-48a1-aa77-deebc7d1517c', '803b90d9-c58c-4113-b4c0-58782e03142c', '批量修改用户，包含update_fields取动态变量', '  - key: bulk_update_user
    module: bulk_upsert
    params:
      user_list:
        check:
          template: "{{must .user_list}}"
          err_msg: 用户列表不能为空
      fields:
        default: ["*"]
    handler_params:
      - key: update_array
        foreach: "[user_list]"
        item: item
        fields:
          - field: password
            template: "{{ if .item.password }}{{md5 .item.password}}{{ end }}"
    table: "user_account"
    options: "[fields]"
    model_field: "[user_list]"', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('9e03b9c4-bc27-470b-88fc-529aac51ca7f', '67193cc3-900e-4c27-a3f6-75ee2dba5688', '只有一级的数组', '      - key: arr2dict
        enable: "{{must .config_params}}"
        foreach: "[config_params]"
        field: "[name]"
        value: "[value]"
        save_field: config', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('a275e234-da99-44e1-8ac8-662f54b5e2c6', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'sql文件示例,注意require(base.sql) 表示是引入当前目录的base.sql文件', 'select a.*,
(
   select GROUP_CONCAT(name )
   from ldap_group lg
   join role_ldap_group rlg on rlg.ldap_group_id =lg.ldap_group_id
   where rlg.role_id =a.role_id
) as ldap_names,
(
    select GROUP_CONCAT(lg.ldap_group_id)
    from ldap_group lg
    join role_ldap_group rlg   on rlg.ldap_group_id =lg.ldap_group_id
    where rlg.role_id =a.role_id
) as ldap_group_ids
from (require(base.sql)) as a', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('b8354a6a-dbe5-4c50-b342-5f692f6ab30f', '91d4abd1-14c9-4705-9de7-1518da998427', 'collect/service.yml', '# 系统总路由
services:
  - key: ''amis_router''
    name: ''项目路由''
    path: ''amis_router/service.yml''
  - key: ''hrm''
    name: ''人资管理''
    path: ''hrm/service.yml''
  - key: ''system''
    name: ''登录''
    path: ''system/service.yml''
  - key: ''wechat''
    name: ''企业微信''
    path: ''wechat/service.yml''
  - key: ''zabbix''
    name: ''zabbix''
    path: ''zabbix/service.yml''
  - key: ''nexus''
    name: ''nexus''
    path: ''nexus/service.yml''
  - key: ''partner''
    name: ''partner''
    path: ''partner/service.yml''
  - key: ''config''
    name: ''config''
    path: ''config/service.yml''

# 模块处理器
module_handler:
  # 模型修改
  - key: sql
    name: 执行sql 查询
    type: inner
    path: SqlService
  # 模型保存
  - key: model_save
    name: 模型保存
    type: inner
    path: ModelSaveService
  # 模型修改
  - key: model_update
    name: 模型修改
    type: inner
    path: ModelUpdateService
  # 模型删除
  - key: model_delete
    name: 模型删除
    type: inner
    path: ModelDeleteService
  # 批量新增
  - key: bulk_create
    name: 批量新增
    type: inner
    path: BulkCreateService
  # 批量新增或更新
  - key: bulk_upsert
    name: 批量新增
    type: inner
    path: BulkUpsertService
  - key: empty
    name: 空模块
    type: inner
    path: EmptyService
  - key: bulk_service
    name: 批量服务
    type: inner
    path: BulkService
  - key: http
    name: 发送http请求
    type: inner
    path: HttpService

  - key: ldap
    name: 发送ldap请求
    type: inner
    path: LdapService
  - key: service_flow
    name: 服务流程化
    type: inner
    path: ServiceFlowService
# 数据处理
data_handler:
  - key: update_field
    name: 添加参数
    type: inner
    path: UpdateField
  - key: prop_arr
    name: 数组对象转数组
    type: inner
    path: PropArr
  - key: check_field
    name: 检查参数
    type: inner
    path: CheckField
  - key: update_array
    name: 添加参数
    type: inner
    path: UpdateArray
  - key: update_array_from_array
    name: 补充右边的数据
    type: inner
    path: UpdateArrayFromArray
  - key: service2field
    name: 服务转字段
    type: inner
    path: Service2Field
  - key: arr2obj
    name: 数组结果转对象
    type: inner
    path: Arr2Obj
  - key: arr2dict
    name: 参数数组转key/value对象
    type: inner
    path: Arr2Dict
  - key: filter_arr
    name: 数组转对象
    type: inner
    path: FilterArr
  - key: param2result
    name: 参数转结果
    type: inner
    path: Param2Result
  - key: params2result
    name: 多个参数转结果
    type: inner
    path: Params2Result
  - key: result2params
    name: 结果转参数
    type: inner
    path: Result2Params
  - key: result2map
    name: 结果转map
    type: inner
    path: Result2Map
  - key: count2map
    name: count转map
    type: inner
    path: Count2Map
  - key: session_add
    name: 添加session
    type: inner
    path: SessionAdd
  - key: session_remove
    name: 删除session
    type: inner
    path: SessionRemove

  - key: session_get
    name: 获取session
    type: inner
    path: SessionGet
  - key: data2excel
    name: 数据转excel
    type: inner
    path: Data2Excel
  - key: excel2data
    name: 数据转excel
    type: inner
    path: Excel2Data
  - key: ignore_data
    name: 忽略数据
    type: inner
    path: IgnoreData
  - key: file2result
    name: 数据转excel
    type: inner
    path: File2Result
  - key: file2datajson
    name: 文件转data_json
    type: inner
    path: File2DataJson
  - key: field2array
    name: 字段转数组
    type: inner
    path: Field2Array
  - key: arr2arrayObj
    name: 字段转数组,简单数组，转数组对象
    type: inner
    path: Arr2arrayObj
  - key: get_modify_data
    name: 获取修改的数据
    type: inner
    path: GetModifyData
  - key: group_by
    name: 分组
    type: inner
    path: GroupBy
  - key: combine_array
    name: 数组添加字段
    type: inner
    path: CombineArray
  # 处理缓存
  - key: handler_cache
    name: 处理缓存
    type: inner
    path: HandlerCache
  # 防止重复请求
  - key: prevent_duplication
    name: 防止重复请求
    type: inner
    path: PreventDuplication

# 加载启动配置，读取文件内容
load_startup_plugin:
  - key: load_data_file
    name: 将文件路径转换成文件内容
    method: LoadDataFile
    fields:
      - from: DataFile
        to: FileData
        name: 将data_file 转换成文件数据
      - from: CountFile
        to: CountFileData
        name: 将count_file 转换成文件数据
      - name: 将excel配置地址转文件内容
        from: ExcelConfig
        to: ExcelConfigContent
      - name: 将修改内容转内容
        from: ModifyConfig
        to: ModifyConfigContent
      - name: 将http配置地址转文件内容
        from: HttpJson
        to: HttpJsonContent
      - name: 将DataJson配置地址转文件内容
        from: DataJson
        to: DataJsonContent
  - key: load_excel_config
    method: LoadExcelConfig
    name: 转换excel配置
  - key: load_modify_config
    method: LoadModifyConfig
    name: 将修改内容转配置
  - key: load_http_config
    method: LoadHttpJson
    name: 转换http配置
  - key: load_data_json_config
    method: LoadDataJson
    name: 转换http配置
  - key: load_schedule
    method: LoadSchedule
    name: 加载定时模板是否启动
  - key: load_cache_config
    method: LoadCacheConfig
    name: 加载定时模板是否启动
  - key: load_prevent_duplication_config
    method: LoadPreventDuplicationConfig
    name: 加载prevent_duplication
  - key: load_router_all_enable
    method: LoadRouterAllEnable
    name: 加载模板能用,处理handler_req_param 是否能用
    fields:
      - rule: array_field
        field: BeforePlugin
        fields:
          - from: Enable
            to: EnableTpl
      - rule: array_field
        field: AfterPlugin
        fields:
          - from: Enable
            to: EnableTpl

  - key: load_data_file_tpl
    name: 将文件转换成模板
    method: LoadDataFileTpl
    fields:

      - from: FileData
        to: FileDataTpl
        name: 将file_data 转换成模板
        rule: simple_field

      - from: CountFileData
        to: CountFileDataTpl
        name: 将count_file转换成模板
        rule: simple_field

      - from: Success
        to: SuccessTpl
        name: 将success转换成模板
        rule: simple_field

      - rule: map_field
        name: 将字典中参数转字转成模板
        field: Params
        fields:
          - from: Template
            to: TemplateTpl
        third_field: Check
        third_fields:
          - from: Template
            to: TemplateTpl
          - from: ErrMsg
            to: ErrMsgTpl
      - rule: array_field
        name: 将data_json 中的services 转参数
        field: DataJsonConfig.Services
        fields:
          - from: NodeNext
            name: next 转 模板
            to: NodeNextTpl
          - from: NodeFail
            name: fail 转 模板
            to: NodeFailTpl
          - from: Enable
            name: 编译是否启用
            to: EnableTpl
          - from: Template
            name: 编辑检查数据模板
            to: TemplateTpl
          - from: ErrMsg
            name: 编译错误提示模板
            to: ErrMsgTpl
          - from: IfTemplate
            name: 编译 if 判断模板
            to: IfTemplateTpl
          - from: Value
            name: 编译取值模板
            to: ValueTpl
        third_array_field: Fields
        third_array_fields:
          - from: Template
            to: TemplateTpl
          - from: From
            to: FromTpl
          - from: Field
            to: FieldTpl
          - from: ErrMsg
            to: ErrMsgTpl
      - rule: array_field
        name: 将字典中参数转字转成模板
        field: HandlerParams
        fields:
          - from: Enable
            name: 编译是否启用
            to: EnableTpl
          - from: Template
            name: 编辑检查数据模板
            to: TemplateTpl
          - from: ErrMsg
            name: 编译错误提示模板
            to: ErrMsgTpl
          - from: IfTemplate
            name: 编译 if 判断模板
            to: IfTemplateTpl
          - from: Value
            name: 编译取值模板
            to: ValueTpl
        third_array_field: Fields
        third_array_fields:
          - from: Template
            to: TemplateTpl
          - from: From
            to: FromTpl
          - from: Field
            to: FieldTpl
          - from: ErrMsg
            to: ErrMsgTpl

      - rule: array_field
        name: 将字典中参数转字转成模板
        field: ResultHandler
        fields:
          - from: Enable
            name: 编译是否启用
            to: EnableTpl
          - from: Template
            name: 编辑检查数据模板
            to: TemplateTpl
          - from: ErrMsg
            name: 编译错误提示模板
            to: ErrMsgTpl
          - from: IfTemplate
            name: 编译 if 判断模板
            to: IfTemplateTpl
          - from: Value
            name: 编译取值模板
            to: ValueTpl
        third_array_field: Fields
        third_array_fields:
          - from: Template
            to: TemplateTpl
          - from: Field
            to: FieldTpl
          - from: ErrMsg
            to: ErrMsgTpl
          - from: From
            to: FromTpl



# 文件处理插件
file_content_plugin:
  - key: require
    name: 引入其他文件
    method: Require
    reg: require[(](.*?)[)]

before_plugin:
  # 处理请求参数
  - key: handler_req_param
    method: HandlerReqParam
    enable: "true"

  # 处理重复请求
  - key: prevent_duplication
    method: PreventDuplication
    enable: "true"
  # 处理缓存
  - key: handler_cache
    method: HandlerCache
    enable: "true"
  # 处理请求参数
  - key: handler_params
    method: HandlerParams
    enable: "true"
after_plugin:
  # 结果处理
  - key: result_handler
    method: ResultHandler
    enable: "true"
  # 处理缓存
  - key: handler_cache
    method: HandlerCache
    enable: "true"', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('bc52c631-170b-439d-92f5-8f6c3072901c', 'e4afe3cd-34eb-4837-b08f-9da741a7f3f8', '11', '11', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('be47c211-cbaa-4702-b211-c0a4f95828aa', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '用户修改对比', '{
  "desc": "替换前和替换后，概念调整下，方便对比，从左往右，传过来的数据，是前段的数据，left。已有的后台数据是right",
  "left_save_field": "after",
  "right_save_field": "before",
  "fields": [
    {
      "rule": "compare_field_value",
      "field": "[nick]",
      "name": "用户昵称",
      "right": "[user_info]",
      "operation": "modify",
      "append_right_fields": [
        "[user_id]"
      ]
    },
    {
      "rule": "compare_field_value",
      "field": "[create_ldap]",
      "name": "创建ldap",
      "right": "[user_info]",
      "operation": "modify",
      "append_right_fields": [
        "[user_id]"
      ]
    },
    {
      "belong": "设置belong 将fields 二层层级去掉，或者在field支持点，xx.xx",
      "rule": "compare_field_value",
      "field": "[user_status]",
      "name": "用户状态",
      "right": "[user_info]",
      "operation": "modify",
      "append_right_fields": [
        "[user_id]"
      ],
      "value_list_field": "current_value_list",
      "target_transfer_key": "[sys_code]",
      "target_transfer_value": "[sys_code_text]",
      "service": {
        "service": "system.get_sys_code",
        "sys_code_type": "user_job_status",
        "sys_code_list": "[current_value_list]"
      }

    },
    {
      "rule": "simple_array_value",
      "field": "[roles]",
      "name": "用户角色",
      "right": "[user_info]",
      "operation": "modify",
      "save_original": true,
      "append_right_fields": [
        "[user_id]"
      ],
      "value_list_field": "current_value_list",
      "target_transfer_key": "[role_code]",
      "target_transfer_value": "[role_name]",
      "service": {
        "service": "hrm.role_query",
        "role_code_list": "[current_value_list]"
      }

    },
    {
      "enable": "{{ eq .create_ldap \\"1\\"}}",
      "rule": "array_obj_value",
      "left_field": "[name]",
      "right_field": "[name]",
      "field": "[ldap_group]",
      "desc": "field匹配规则,value field取值",
      "name": "ldap分组",
      "right": "[right_ldap_group]",
      "left": "[left_ldap_group]",
      "left_value_field": "[name]",
      "right_value_field": "[name]",
      "operation": "modify",
      "with_add_remove": true,
      "save_original": true,
      "append_right_fields": [
        "[user_id]"
      ]
    }
  ]
}', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('bead3efc-8227-4929-9b72-541eefb2d098', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'system_login.json 登陆的流程', '{
  "finish": {
    "key": "service2field",
    "append_param": true,
    "service": {
      "service": "hrm.system_login_finish"
    }
  },
  "services": [
    {
      "node_key": "start",
      "node_type": "start",
      "name": "开始",
      "node_next": "get_user"
    },
    {
      "node_key": "get_user",
      "name": "获取用户信息",
      "node_type": "node",
      "key": "service2field",
      "service": {
        "service": "hrm.user_list",
        "username": "[username]",
        "to_obj": true
      },
      "save_field": "user_info",
      "node_next": "{{ if .user_info }}system_login_check{{ else }}ldap_login{{ end }}"
    },
    {
      "node_key": "system_login_check",
      "key": "check_field",
      "name": "检查名称",
      "node_type": "node",
      "fields": [
        {
          "field": "password",
          "template": "{{eq .password_md5 .user_info.password}}",
          "err_msg": "用户名【{{.username}}】 密码错误！！！"
        }
      ],
      "node_next": "system_login",
      "node_fail": "end"
    },
    {
      "node_key": "ldap_login",
      "node_type": "node",
      "name": "ldap 登陆",
      "key": "service2field",
      "append_param": true,
      "service": {
        "service": "hrm.ldap_login"
      },
      "node_fail": "end",
      "node_next": "get_user_info"
    },
    {
      "node_key": "get_user_info",
      "node_type": "node",
      "name": "获取ldap的用户信息",
      "key": "service2field",
      "service": {
        "service": "hrm.ldap_search",
        "search_username": "[username]",
        "to_obj": true
      },
      "save_field": "userInfo",
      "node_fail": "end",
      "node_next": "create_user"
    },
    {
      "node_key": "create_user",
      "node_type": "node",
      "name": "创建用户",
      "key": "service2field",
      "service": {
        "service": "hrm.create_user",
        "password": "[password]",
        "username": "[username]",
        "email": "[userInfo.mail]",
        "nick": "[userInfo.sn]",
        "phone": "[userInfo.mobile]"
      },
      "node_fail": "end",
      "node_next": "system_login"
    },
    {
      "node_key": "system_login",
      "node_type": "node",
      "name": "系统登陆",
      "key": "service2field",
      "service": {
        "service": "system.login",
        "password": "[password]",
        "username": "[username]"
      },
      "save_field": "session_user",
      "node_fail": "end",
      "node_next": "end"
    },
    {
      "node_key": "end",
      "node_type": "end",
      "name": "结束"
    }
  ]
}', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('c061172d-6ed6-4010-92e5-22e4a84e2af4', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'count 文件示例', 'select count(1) as count
from (require(''./base.sql'')) a', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('c41e2883-0b74-432f-9a05-101139415c4c', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '数组对象对比', '{
	"rule": "array_obj_value",
	"left_field": "[name]",
	"right_field": "[name]",
	"field": "[params]",
	"name": "参数名称",
	"left": "[params]",
	"right": "[local_doc_detail.params]",
	"with_add_remove": true,
	"left_value_field": "[name]",
	"right_value_field": "[name]",
	"operation": "modify",
	"append_right_fields": ["[*]"]
}', 50, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('c5633f18-2d71-4e8a-bbb8-c91c499a7fa6', 'ce02406f-f729-47da-a119-35ed01c6c1c3', 'service2field,调用其他服务，获取对象结果。这里user_info 是对象', '  - key: get_user_modify_data
    http: true
    log: true
    module: empty
    modify_config: user_modify.json
    params:
      user_id:
        check:
          template: "{{must .user_id}}"
          err_msg: 用户ID不能为空
      right_ldap_group:
        default: []
    handler_params:
      - key: service2field
        service:
          service: hrm.user_list
          user_id: "[user_id]"
          count: false
          to_obj: true
        save_field: user_info
      - key: service2field
        enable: "{{must .user_info.roles}}"
        service:
          service: hrm.ldap_group_query
          roles: "[user_info.roles]"
        save_field: right_ldap_group', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('ce83d038-c887-4e2e-9ca8-308218c055a5', '28a1f850-9e5f-4368-9230-220abf98af15', 'excel转数据', '{
  "name": "将用户列表转excel",
  "sheets": [
    {
      "sheet_index": 0,
      "title_height": 48,
      "desc": "16*行数",
      "title_style": {
        "font": {
          "family": "宋体",
          "size": 12,
          "color": "#ffffff"
        },
        "fill": {
          "type": "pattern",
          "color": [
            "#3366FF"
          ],
          "pattern": 1
        },
        "alignment": {
          "vertical": "center",
          "WrapText": true
        }
      },
      "name_style": {
        "font": {
          "family": "Arial",
          "size": 10,
          "color": "#ffffff"
        },
        "fill": {
          "type": "pattern",
          "color": [
            "#0066CC"
          ],
          "pattern": 1
        },
        "alignment": {
          "vertical": "center"
        }
      },
      "content_style": {
        "NumFmt": 49
      },
      "title": "用户列表\\r\\n1.ID 是作为系统的唯一标志，如果id存在则编辑数据，空着则新增\\r\\n2.带* 是必须保存",
      "data": "[user_list]",
      "fields": [
        {
          "field": "[userid]",
          "name": "id",
          "width": 10
        },
        {
          "field": "[nick]",
          "name": "*昵称",
          "width": 20
        },
        {
          "field": "[username]",
          "name": "*用户名",
          "width": 20
        },
        {
          "field": "[tel]",
          "name": "电话号码",
          "width": 20
        },
        {
          "field": "[email]",
          "name": "邮箱",
          "width": 20
        },
        {
          "field": "[statu]",
          "template": "{{ if  eq .statu 1}} 正常 {{ else }}删除{{ end }}",
          "name": "状态",
          "width": 10
        },
        {
          "field": "[address]",
          "name": "地址",
          "width": 50
        },
        {
          "field": "[create_time]",
          "name": "创建日期",
          "width": 20
        }
      ]
    }
  ]
}', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('d2206577-cbb8-4199-8b2b-c407653121e9', '45d1d393-6758-4c37-8ba9-108493f67b8e', '拼接某个字段参数，所有字段', '      - key: service2field
        name: 保存文档
        enable: "{{must .base_modify_list}}"
        service:
          service: config.collect_doc_update
        item: "[doc]"
        append_item_param: true', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('d3f22eac-0daa-4686-b9a2-7f6ff9ee1899', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '简单数组对比', '{
      "rule": "simple_array_value",
      "field": "[roles]",
      "name": "用户角色",
      "right": "[user_info]",
      "operation": "modify",
      "save_original": true,
      "append_right_fields": [
        "[user_id]"
      ],
      "value_list_field": "current_value_list",
      "target_transfer_key": "[role_code]",
      "target_transfer_value": "[role_name]",
      "service": {
        "service": "hrm.role_query",
        "role_code_list": "[current_value_list]"
      }

    }', 40, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('dcfb5437-af4d-44d7-91c5-f9dc1f8fbb49', 'bd6431c8-caf3-4ae9-96af-03cce615204b', '路由生成', '{
  "pages": [
    {
      "children": [
        {
          "label": "系统管理",
          "url": "/system_manage",
          "schema": {
            "type": "page",
            "title": "欢迎进入系统管理"
          },
          "children": [
            {
              "label": "用户管理",
              "url": "user_manage",
              "schemaApi": {
                "method": "post",
                "url": "/template_data/data",
                "data": {
                  "service": "amis_router.user_manage"
                },
                "adaptor": "return {\\n    ...payload,\\n    status: payload.success === true? 0 : 1\\n}"
              }
            },
            {
              "label": "角色管理",
              "url": "role_manage",
              "schemaApi": {
                "method": "post",
                "url": "/template_data/data",
                "data": {
                  "service": "amis_router.role_manage"
                },
                "adaptor": "return {\\n    ...payload,\\n    status: payload.success === true? 0 : 1\\n}"
              }
            } ,
            {
              "label": "ldap分组管理",
              "url": "ldap_group_manage",
              "schemaApi": {
                "method": "post",
                "url": "/template_data/data",
                "data": {
                  "service": "amis_router.ldap_group_manage"
                },
                "adaptor": "return {\\n    ...payload,\\n    status: payload.success === true? 0 : 1\\n}"
              }
            },
            {
              "label": "系统参数管理",
              "url": "config_manage",
              "schemaApi": {
                "method": "post",
                "url": "/template_data/data",
                "data": {
                  "service": "amis_router.config_manage"
                },
                "adaptor": "return {\\n    ...payload,\\n    status: payload.success === true? 0 : 1\\n}"
              }
            }

          ]
        }
      ]
    }
  ]
}
', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('dda58a0f-b687-41d7-b5f0-1fad0fed726f', 'c00a6f14-98c1-4308-8017-cf35ae300de4', '配置文件示例，获取企业微信token', '{
  "url": "https://qyapi.weixin.qq.com/cgi-bin/gettoken",
  "method": "GET",
  "header": {},
  "result_json": true,
  "data": {
    "corpid": "{{get_key `corpid`}}",
    "corpsecret": "{{get_key `corpsecret`}}"
  }
}', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('e07bda70-ec90-42bb-9549-f0e6feb19015', '91d4abd1-14c9-4705-9de7-1518da998427', 'collect/hrm/user/index.yml用户管理的index.yml', 'service:
  - key: user_list_import
    module: empty
    http: true
    excel_config: "./user_list2excel.json"
    handler_params:
      - key: excel2data
        save_field: user_list
      - key: ignore_data
        foreach: "[user_list]"
        params: "params"
        fields:
          - name: "user_id 为空的数据"
            template: "{{ if .user_id }}false{{else}}true{{end}}"
      - key: service2field
        enable: "{{gt (len .user_list) 0 }}"
        service:
          service: hrm.bulk_update_user
          user_list: "[user_list]"
      - key: params2result
        fields:
          - from: "[user_list]"
            to: user_list
  - key: user_list_download
    module: empty
    http: true
    excel_config: "./user_list2excel.json"
    params:
      excel_path:
        template: ''./template/{{current_date_format "20220202"}}/user_{{  replace (sub_str current_date_time -8 0) ":" ""}}_{{sub_str uuid -8 0}}.xlsx''
      response_name:
        default: "用户列表.xlsx"
    handler_params:
      - key: service2field
        service:
          service: hrm.user_list
        append_param: true
        save_field: user_list
      - key: data2excel
        path: "[excel_path]"
      - key: file2result
        path: "[excel_path]"
        result_name: "[response_name]"

  - key: empty_test
    module: empty
    http: true
    handler_params:
      - key: service2field
        service:
          service: hrm.user_list
          username: "[username]"
        save_field: user_info
        template: "{{gt (len .user_info) 0}}"
        err_msg: "用户名【{{.username}}】已经存在"
    result_handler:
      - key: service2field
        service:
          service: hrm.user_list
          username: "[username]"
        save_field: user_info
        template: "{{eq (len .user_info) 0}}"
        err_msg: "用户名【{{.username}}】已经存在"
  - key: bulk_update_user
    module: bulk_upsert
    log: true
    http: true
    params:
      user_list:
        check:
          template: "{{must .user_list}}"
          err_msg: 用户列表不能为空
      fields:
        default: ["*"]
    handler_params:
      - key: update_array
        foreach: "[user_list]"
        item: item
        fields:
          - field: password
            template: "{{ if .item.password }}{{md5 .item.password}}{{ end }}"
    table: "user_account"
    options: "[fields]"
    model_field: "[user_list]"
#    update_fields:
#      - address
    ignore_fields:
      - password
  - key: bulk_create_user
    module: bulk_create
    log: true
    http: true
    params:
      user_list:
        check:
          template: "{{must .user_list}}"
          err_msg: 用户列表不能为空
    handler_params:
      - key: update_array
        foreach: "[user_list]"
        item: item
        fields:
          - field: user_id
            template: "{{uuid}}"
          - field: password
            template: "{{ if .item.password }}{{md5 .item.password}}{{ end }}"
    table: "user_account"
    model_field: "[user_list]"
  - key: get_user_list_by_user_id_list
    http: true
    module: sql
    log: true
    params:
      user_id_list:
        check:
          template: "{{must .user_id_list}}"
          err_msg: 用户ID 不能为空
    data_file: get_user_list_by_user_id_list.sql

  - key: delete_user_by_user_id_list
    module: model_update
    http: true
    log: true
    params:
      user_id_list:
        check:
          template: "{{or (must .user_id_list) (must .ids)}}"
          err_msg: 用户名不能空
      is_delete:
        default: "1"
    table: "user_account"
    update_fields:
      - is_delete
    handler_params:
      - key: field2array
        field: "[ids]"
        enable: "{{must .ids}}"
        save_field: user_id_list
      - key: service2field
        enable: "{{must .user_id_list}}"
        service:
          service: hrm.get_user_list_by_user_id_list
          user_id_list: "[user_id_list]"
        save_field: user_info_list
      - key: service2field
        enable: "{{must .user_info_list}}"
        service:
          service: hrm.remote_user_ldap_delete_bulk
          user_info_list: "[user_info_list]"
        save_field: deleteResult
    filter:
      user_id__in: "[user_id_list]"
    result_handler:
      - key: param2result
        enable: "{{must .user_info_list}}"
        field: deleteResult
  - key: update_user_all
    module: model_update
    http: true
    log: true
    table: "user_account"
    filter:
      user_id__isnull: false
    update_fields:
      - address
      - comments
      - wechat_user_id
  - key: update_user_by_user_id_list
    module: model_update
    http: true
    log: true
    params:
      fields:
        default: ["*"]
      user_id_list:
        check:
          template: "{{must .user_id_list}}"
          err_msg: 用户名不能空
    table: "user_account"
    options: "[fields]"
    filter:
      user_id__in: "[user_id_list]"
    ignore_fields:
      - password
  - key: update_user_by_username_nick
    module: model_update
    http: true
    log: true
    params:
      username:
        check:
          template: "{{must .username}}"
          err_msg: 用户名不能空
      nick:
        check:
          template: "{{must .nick}}"
          err_msg: 昵称不能空
    table: "user_account"
    filter:
      username: "[username]"
      nick: "[username]"
    update_fields:
      - address
  - key: update_user_by_user_id
    module: model_update
    http: true
    log: true
    params:
      user_id:
        check:
          template: "{{must .user_id}}"
          err_msg: 用户ID不能为空
      username:
        check:
          template: "{{must .username}}"
          err_msg: 用户名不能空
      nick:
        check:
          template: "{{must .nick}}"
          err_msg: 昵称不能空
    handler_params:
      - key: service2field
        service:
          service: hrm.user_list
          exclude: "[user_id]"
          username: "[username]"
        save_field: user_info
        template:  "{{ if .user_info  }}false{{ else  }}true{{end}}"
        err_msg: "用户名 {{.username}} 已经存在【{{len .user_info}}】次"
    table: "user_account"
    filter:
      user_id: "[user_id]"



  - key: update_user_by_username
    module: model_update
    http: true
    log: true
    params:
      username:
        check:
          template: "{{must .username}}"
          err_msg: 用户名不能空
      password:
        template: "{{ if .password }}{{md5 .password}}{{ end }}"
    table: "user_account"
    filter:
      user_name: "[username]"
#    update_fields:
#      - address
  - key: get_pinyin
    module: empty
    http: true
    params:
      content:
        template: "{{pinyin .content}}"
      result_name:
        default: "username"
    result_handler:
      - key: params2result
        fields:
          - from: "[content]"
            to: "[result_name]"
          - from : "{{.content}}@weigaogroup.com"
            to: email

  - key: create_user
    module: model_save
    http: true
    log: true
    params:
      password:
        template: "{{ if .password }}{{md5 .password}}{{ end }}"
        default: "123456"
      is_delete:
        default: "0"
      user_id:
        check:
          template: "{{must .user_id}}"
          err_msg: 用户ID 不能为空
#        template: "{{uuid}}"
      user_name:
        template: "{{.username}}"
      username:
        check:
          template: "{{must .username}}"
          err_msg: 用户名不能为空
      user_status:
        default: "trial"
      role_id:
        default: "xxxx"
      avatar:
        template: session_user_id
      create_time:
        template: "{{current_date_time}}"
      create_user:
        template: "{{.session_user_id}}"
      modify_time:
        template: "{{current_date_time}}"
    handler_params:
      - key: service2field
        service:
          service: hrm.user_list
          username: "[username]"
        save_field: user_info
        template: "{{le (len .user_info) 0}}"
        err_msg: "用户名【{{.username}}】已经存在"
    table: "user_account"
    ignore_fields:
      - avatar
#    update_fields:
#      - role_id

  - key: user_list
    module: sql
    http: true
    params:
      search:
        template: "{{ if .search }}%{{.search}}%{{ end }}"
      nick:
        default: ""
      page:
        type: int
        default: 1
      size:
        default: 20
        type: int
      start:
        template: " ({{.page}}-1) * {{.size}}"
        exec: true
        type: int
      pagination:
        default: true
      to_obj:
        default: false
      count:
        default: true
    data_file: user_list.sql
    count_file: user_list_count.sql
    pagination: pagination
    count: "[count]"
    result_handler:
      - key: arr2obj
        enable: "[to_obj]"
#  #   module: service_flow
#  #   http: true
#  #   data_json: install_agent.json
#', 30, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('e1966c8e-8326-4207-9b31-339c8ef4fd9f', 'c00a6f14-98c1-4308-8017-cf35ae300de4', '调用zabbix 的配置', '{
  "url": "http://172.26.0.150/zabbix/api_jsonrpc.php",
  "method": "POST",
  "header": {
    "content-type": "application/json"
  },
  "data": {
    "jsonrpc": "2.0",
    "method": "hostgroup.get",
    "params": {
      "output": "extend",
      "filter": {
        "name": [
          "[LMYY]医院"
        ]
      }

    },
    "auth": "0f8bd590815bf5bf2c74e90f995594f1",
    "id": 1
  }
}', 20, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('e7984820-d8e1-40e7-b8b3-692a86c681b7', '5881b6e1-8217-475f-9b05-4bc20ffdae8e', '取对象的某个字段转参数', '    result_handler:
      - key: result2params
        fields:
          - from: "[access_token]"
            to: "[access_token]"', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('ecd86882-5822-454b-b7dc-f6525d4c1bb6', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '分组编辑', '{
  require(login.common),
  "method": "modify",
  "ModifyParams": {
    "DN": "ou={{.old_name}},{{.ldap_config.ldap_groups}},{{.ldap_config.ldap_base_dn}}",
    "Changes": [
      {
        "Desc": "AddAttribute       = 0\\n\\tDeleteAttribute    = 1\\n\\tReplaceAttribute   = 2\\n\\tIncrementAttribute = 3",
        "Operation": 2,
        "Modification": {
          "TYPE": "cn",
          "Vals": [
            "{{.new_name}}"
          ]
        }
      }

    ]
  }
}', 50, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('f2d8ac2b-ab8a-4dff-9a8d-d66ba5792dff', '0722aa5a-78e6-4a29-818c-ba33843bdf81', '只改地址示例', '  - key: update_user_by_username_nick
    module: model_update
    http: true
    log: true
    params:
      username:
        check:
          template: "{{must .username}}"
          err_msg: 用户名不能空
      nick:
        check:
          template: "{{must .nick}}"
          err_msg: 昵称不能空
    table: "user_account"
    filter:
      username: "[username]"
      nick: "[username]"
    update_fields:
      - address', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('f7a837b5-a01e-4f66-9531-e0bec6057594', '45d1d393-6758-4c37-8ba9-108493f67b8e', '保存服务，拼接参数中字段', '      - key: service2field
        name: 保存修改列
        service:
          service: config.config_detail_change_history_bulk_save
          change_list: "[change_list]"', 10, NULL);
INSERT INTO collect_doc_demo (doc_demo_id, collect_doc_id, name, code, order_index, code_result) VALUES('f9c3163c-0eaa-4e90-ac7d-63c3332b7f7d', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', '数组对象 文件示例', '{{ if .config_detail_list}}
and (g.group_id,g.name) in({{ range $index,$item := .config_detail_list }} {{if $index}},{{end}} ({{$item.group_id}},{{$item.name_copy}}){{ end}})
{{end}}', 60, NULL);



INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('01e26553-44cb-483f-92ed-1d45dba5f589', '91d4abd1-14c9-4705-9de7-1518da998427', 'collect/service.yml 定义了所有模块、处理器。就是为了写配置的时候方便看一眼，有印象', 30);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('03b05fa1-bfc3-4996-84da-ec815b0f33af', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', '支持2个sql，列表sql和统计sql。可以单独控制是否运行统计count', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('03b9dbd6-2914-4233-9383-862c826dcc43', '9b4fbf55-b221-4b05-86f1-2e78132ee552', '可以利用params里面的template生成uuid、创建时间、创建人、设置默认时间', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('15611e60-138f-4a45-a2d4-f592d9927963', '77e10054-791e-4889-a5c8-1fc2b9a6e514', '主要利用room+fields唯一定位缓存。作为key', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('1609d80a-5e76-4aa0-9b9d-43c66f8604d4', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', '服务流程化', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('184ff2db-6aa6-40ee-9a15-ad61d3425e62', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', '优先级template >default >check', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('1fcf10ab-eb50-4afb-93a4-3b3e5699799d', '5b86a474-2af6-4e82-8d25-3635645bc315', '获取session', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('251a0124-06e7-49fc-a27d-a13f9ffc4d1a', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '支持操作名称修改', 50);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('25a248d2-a3ca-4984-976c-f14339c81f19', 'd4ee2257-8a51-40bf-8d48-4b5469858e21', '可以将任何服务的结果转excel ，不过一般都是查询', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('261eba7f-a8c6-46c2-9419-9da9dec99ae9', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'params 下面一级是变量名称，比如示例中create_time 是表里面create_time 字段，下面生成规则，与校验规则', 30);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('2a696412-969b-41f7-9a48-fdcfe42ba308', 'd25ddb2a-2c4e-4d3f-83c3-50ddc4215de8', '在handler_params可以通过enable模板动态判断是否需要运行某个服务', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('2b1710f0-b11f-402a-9882-d36ab9c1fd69', '01044aa5-6f65-4b59-a4a9-ae7a0191be4c', '添加session', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('304a21d1-cd65-438d-8b95-a13095e0d2f7', '293dffa7-1a0b-4e72-8bd7-0f2a44758495', '删除session', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('3271e78b-1607-4bd2-8073-a231aba9baf9', '3c768fa7-2c9c-48f4-a109-e740ab17ab6c', 'handler_params 针对请求参数处理，模块执行前的处理。', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('35aba943-9d28-408c-9f47-fdda285dbb78', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '支持结果数据转换', 40);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('36c1b768-7afd-4d2d-a592-355ee490aa53', 'e4afe3cd-34eb-4837-b08f-9da741a7f3f8', '1', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('378acccc-5986-4738-9aef-8b8ba3da16e0', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', '支持配置控制分页', 40);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('3da809f8-4650-49ae-898c-a73214f67c36', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', '数组更新数组', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('44619c38-3e75-4cb5-81dc-3320a5aef6e3', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', '支持golang的模板渲染成sql；支持3种数据类型变量 ，1.简单类型key/value;2.数组类型,字符串、数字数组;3.数组对象类型', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('45c271e9-bc82-4941-b423-49073c1d020f', '849062cb-1f83-454e-b81d-3d7754f9ac4a', '批量更新数组', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('48b4f4c4-4119-4349-a280-9f00c5eaf052', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', '定义服务的参数，与参数的简单校验', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('49d729d4-fffa-4633-9bb6-9f735ee5812a', 'd89f00cd-2ecf-43d8-9fb4-f9bd26c1cf05', '11', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('4cd2edc0-061e-4adb-a14d-e261b030c98b', '894496fd-4201-4b40-95d6-24ecd978b719', '批量新增数据行', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('52bcd8a1-806e-4c11-a7f0-5d2b947d496f', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '支持require 引入公共文件', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('55b1ba30-61f7-4997-82c9-cbbc698bb965', '77e10054-791e-4889-a5c8-1fc2b9a6e514', '缓存的增删改查', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('5997d7ad-93a2-4b38-a0bb-ace0818c8c26', 'd25ddb2a-2c4e-4d3f-83c3-50ddc4215de8', '主要为了处理参数，一般搭配handler_params', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('5d9ada12-42f7-4fe6-9a2d-2949a0807eac', '894496fd-4201-4b40-95d6-24ecd978b719', '利用handler_params 中update_array 可以对没行记录字段进行调整', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('6306fcd1-6dc1-41c9-ab11-2e78bb8c1e2e', '91d4abd1-14c9-4705-9de7-1518da998427', '二级服务分类，service 一般 xxx.xx。第一级表示项目，第二级表示具体服务.', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('643f6d31-ded1-499d-a8eb-9d23aacea905', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', '多线程调用服务', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('6ec87163-4360-442d-a561-bb119d295628', '3522bb77-0cef-4a33-9e83-5ff05b2cf88f', '缓存设置，根据缓存room+fields 唯一定位缓存', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('742b34e5-4d64-4499-918e-142500d9ae95', '803b90d9-c58c-4113-b4c0-58782e03142c', '批量修改多行记录，支持修改不同记录不同值', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('7636ee28-a01d-4f13-861e-496608813073', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'sql文件支持require引入公共文件', 30);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('76a1df99-b3cd-4754-8ce0-cabeaa79d3e6', 'ce02406f-f729-47da-a119-35ed01c6c1c3', '数组结果转对象结果', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('77db476d-83b9-41a7-96ed-4b56df915d2f', 'f7350f20-20fb-47f4-8b19-511c8b5b38a9', '检查服务字段是否合法', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('78901e95-baed-47c1-bd75-8b318be23e27', '74973d56-9773-432f-b1d5-3b68ad40998d', '从数组中过滤数组', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('79a13c62-4d8d-4bcb-abf8-cb11206e2254', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '数据对象比对新增修改删除', 30);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('7f3cec0b-9aba-45bc-ad2a-5bd962b8892e', 'c00a6f14-98c1-4308-8017-cf35ae300de4', '模板渲染生成http配置', 30);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('908ca256-e4fe-435c-8b44-15604453a098', 'f272ba24-4c3e-4072-be07-abae59c93907', '对象数组转简单数组', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('9d8231c8-c130-4c5a-a028-5e9e1f079b52', '3c768fa7-2c9c-48f4-a109-e740ab17ab6c', 'handler_params和result_handler 可以通过service2field调用内部任何服务', 30);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('9dfc6db8-ef38-4a94-844e-df6923b7569a', '79d0939d-16f6-4c8d-b611-5a55409162ad', '简单数组，转对象数组', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('a106b176-e5f8-4aa2-bf85-dd08cef0ff19', '9215873e-d44d-4a8f-b243-f2c45b1833e5', '支持ldap增删改查', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('a1271b4a-1b6f-4ab2-acba-8c937250a34c', '3c768fa7-2c9c-48f4-a109-e740ab17ab6c', 'result_handler 针对结果执行后进行处理，模块执行后的处理。', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('a56aea23-9741-45f4-b3a1-aa296828160d', '0722aa5a-78e6-4a29-818c-ba33843bdf81', '支持批量和单行记录修改，主要取决你过滤条件', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('a5dfd039-33dd-476f-b2ca-482f4ba030d7', 'eee7f650-4583-4bf6-b770-ce853eba8c54', '将字符串转成数组', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('b0bdab5e-c5b5-4b8e-a383-790b53a505b2', '45d1d393-6758-4c37-8ba9-108493f67b8e', '可以调用其他服务，作为本服务的入参', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('b2a2d715-83fe-4ab4-a083-2516cfbcc618', '45d1d393-6758-4c37-8ba9-108493f67b8e', '可以结合template 来校验，校验服务是否正常', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('b7013ca8-12eb-4e68-aa48-e737d02907a4', '31e7d7f9-c87e-4558-bd76-f673d34e1c84', '数据分组，也可以利用为去重', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('b80e02a9-565a-4155-8d17-25b19378fa9c', '0722aa5a-78e6-4a29-818c-ba33843bdf81', '支持控制只改部分字段', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('b83d3560-ee8b-49bd-8a83-dd2a0b854d36', 'ede82dd3-1b4f-4097-9260-c2d9c7acdf21', '更新参数中的字段', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('b85792ef-ede4-4814-a0fb-c15e9f9558f6', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '简单字段比对修改', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('ba33ed33-a209-48d8-b995-f9256dfe419e', '5881b6e1-8217-475f-9b05-4bc20ffdae8e', '结果转参数', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('bc47d147-85e7-4460-a27d-9c93906cc521', '9b4fbf55-b221-4b05-86f1-2e78132ee552', '针对数据库表进行新增一行', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('bc9f2271-e043-4ac0-a7f9-391491c4ff21', '36248998-48ee-41db-a0f4-2010b056b2a3', '返回服务器上的文件', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('bda7ebdc-0d2a-4061-ba62-51bdf8803c36', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', '简单数组字段比对新增与删除', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('ca4e447d-8e38-4011-b1aa-fa3497411135', 'c00a6f14-98c1-4308-8017-cf35ae300de4', '后台发http请求', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('cd8a0460-3cdd-432b-a8d5-0a31bbf7d250', '0b3ebb5e-8ce6-4ebc-ae92-da6b0dfe13a2', '符合条件的数据，过滤掉', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('ce37f8dd-fd8d-4574-96de-55e6a39cffbd', '91d4abd1-14c9-4705-9de7-1518da998427', '服务名定义，由入口文件services下的key+叶子目录下key2个拼接组成。比如hrm.user_list,是最上层collect/service.yml的key=hrm,user_list 是hrm/user/index.yml下面的key=user_list服务 。定义的时候就得保证唯一', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('d7551f2f-4a4d-4710-a599-7c941cca749f', 'bd6431c8-caf3-4ae9-96af-03cce615204b', '配置文件转字段', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('e0423cb7-08b7-4366-a57f-a6f3e0cbc616', 'c00a6f14-98c1-4308-8017-cf35ae300de4', '支持require另外一个公共文件', 20);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('e17c6f22-901a-443b-a53e-f90f52c68fba', 'e32693d9-a954-48ac-a1bb-097a0684c8b3', '将参数中的多个字段返回', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('ef0c486b-cd68-4c98-ac9a-0a8b44b14a48', '234fe81b-c7e3-46a1-a274-2bbc904e6ab7', '这些模块只有module 是必须的', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('efd2e567-dbd0-4371-abd0-e15f5710ead4', '1ba89c00-fd10-4756-9414-070bf51505d7', '参数中的某个字段转结果', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('f04a591f-682f-4333-82e6-03e17c428577', 'e3d5dce4-c758-41cc-9ffd-578ca39836e3', '防止重复请求，根据fields 算出key', 10);
INSERT INTO collect_doc_important (doc_important_id, collect_doc_id, name, order_index) VALUES('fe131f28-0d6d-4d98-b46b-72504deddffd', '70bbd0a8-5e8a-4083-9986-e8c0c320ebe0', '将数组拼接到另外一个数组中去，将右边的数据拼接到左边', 10);


INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('0046138d-c0fc-43ac-949e-2c92817f7a70', 'b9eab01d-7236-4d15-9053-0f193a3d2ffb', 'key', 'params2result', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('01cdafe7-d4a4-40bf-a8e5-c19decfeb2b5', 'd4ee2257-8a51-40bf-8d48-4b5469858e21', 'excel_config', '主体配置文件中，excel配置路径', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('02034778-44e3-4e59-9d08-ed47271beda2', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'params.你的field.check.err_msg', '校验规则失败的消息', 'template', '', 90);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('024827c3-add7-484e-9518-d89b63d8e9c3', 'ede82dd3-1b4f-4097-9260-c2d9c7acdf21', 'fields[template]', '字段内容,支持模板生成变量，支持[]取参数中值', 'template', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('029e55de-bb9c-4007-b050-67bda221bf07', '70bbd0a8-5e8a-4083-9986-e8c0c320ebe0', 'field', '左边关键字段', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('02e733f2-c299-4097-9e1a-27cbd857af8e', '0b3ebb5e-8ce6-4ebc-ae92-da6b0dfe13a2', 'fields[template]', '忽略规则，true表示忽略', 'template', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('049c7032-e22b-4bb5-a7e1-4d571021b61f', '1ba89c00-fd10-4756-9414-070bf51505d7', 'key', 'param2result', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('052b0da2-684a-4574-8503-96dde90a4dde', 'f7350f20-20fb-47f4-8b19-511c8b5b38a9', 'fields', '校验字段内容', 'array', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('061d179b-1e1f-416f-85cd-b7fc3dca8412', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[]', '数据规则', 'array', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('068502de-972b-4e8c-bcc4-1fe54d17b818', 'ede82dd3-1b4f-4097-9260-c2d9c7acdf21', 'fields[field]', '字段名称', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('06d4e923-05a0-44d8-bd5e-eacf76dd4a5c', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[right_value_field]', '主要用于数组对象，对比行记录里面其他字段值，右边的取值', 'string', '', 160);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('06de358c-f02f-4a89-a019-adecc9653d86', '67193cc3-900e-4c27-a3f6-75ee2dba5688', 'children', '如果有children 表示有个二级数组，转二级字段里面对象', 'string', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('08d9ac6e-b8a4-4302-9c3c-7fee192148c4', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'module', 'service_flow ,服务流程化', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('099ee21c-edf4-458c-873e-d28a38eeecdc', '0b3ebb5e-8ce6-4ebc-ae92-da6b0dfe13a2', 'fields[name]', '名称', 'string', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('0b03eb3d-a72a-4433-9dc8-f0ba2f0f1def', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[field]', '对比的字段名称', 'string', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('0e684496-7782-42a4-a840-8e640a0c32be', 'd4ee2257-8a51-40bf-8d48-4b5469858e21', 'handler_params[path]', '生成excel 的路径，[你的变量]', 'string', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('0fa6df5b-d381-4424-8c7e-7fe1a733bb6c', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'right_field', '右边匹配关键字段[你的item变量] 支持& 取2个字段', 'string', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('0fe757ef-f66e-4def-ba60-9721ca5b19e0', '0722aa5a-78e6-4a29-818c-ba33843bdf81', 'filter.你的field__isnull', 'key表示  你的字段 进行判读是否为空的操作。主要进行全表修改，比如创建时间为空全部修改', 'string', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('10502fc6-85dc-4ff5-a13f-2a3d9035d2a1', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', 'batch', '批量运行的服务配置', 'json', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('126efadf-dc0e-4df3-9f8d-9c568acc83ab', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.ModifyParams', '修改对应的参数，具体参数看示例', 'json', '', 100);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('137e28ef-d0bb-4945-85e1-1c34d892ca2a', '01044aa5-6f65-4b59-a4a9-ae7a0191be4c', 'fields', 'session 操作的内容', 'array', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('13fec02d-b34e-4783-a324-4c653121afc7', 'eee7f650-4583-4bf6-b770-ce853eba8c54', 'key', 'field2array', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('15ee5759-35e3-4853-85a4-e9ab0581489b', '0722aa5a-78e6-4a29-818c-ba33843bdf81', 'options', '支持update_fields动态取变量，[你参数变量]', 'string', '', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('16f410ec-3a39-44c9-883e-a64ae438262f', '803b90d9-c58c-4113-b4c0-58782e03142c', 'options', '支持update_fields 取动态变量，[你参数变量]', 'string', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('16f5a682-cd34-429b-9ce6-3754fe58855f', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'params.你的field', '请求自定义字段', 'json', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('18bcc490-233d-4689-8097-2ab181b3053a', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.services[name]', '节点名称', 'string', '是', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('1a36534d-32aa-48c1-9ce3-cc96326ceba3', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[append_right_fields]', '拼接右边的字段，*表示所有字段', 'array', '', 110);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('1fa01655-469e-4ee9-b9e6-53ece832b7da', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'to_obj', '是否转对象', 'boolean', '', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('201a22d1-f004-48bd-8dd9-4d975a4d8bd3', '849062cb-1f83-454e-b81d-3d7754f9ac4a', 'fields[field]', '字段名', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('22a9d9dd-1c93-48b3-8153-7b94a9c1298d', '894496fd-4201-4b40-95d6-24ecd978b719', 'table', '数据库表名', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('2458d109-e506-4651-ae06-5c98511f4a5d', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'page', '第几页', 'number', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('24619995-82ae-4133-a1c4-8feaa7e090e3', 'f272ba24-4c3e-4072-be07-abae59c93907', 'foreach', '数组循环对象，[你的变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('25f6bd4b-03fd-4ae9-8df2-fe02d9eb6425', '5b86a474-2af6-4e82-8d25-3635645bc315', 'fields[field]', '获取session 的字段', 'string', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('284d931a-fb3e-4e7a-8366-16d10ca9d71f', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[left]', '左边取对象字段，如果没有就从参数中取', 'string', '', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('2a8dedba-d94f-49e4-a7ed-0c76fdf4b143', '74973d56-9773-432f-b1d5-3b68ad40998d', 'if_template', '过滤条件，true 保留下来', 'template', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('2ac36f73-d7f8-4523-bd83-a2ce2cb9ef72', '67193cc3-900e-4c27-a3f6-75ee2dba5688', 'field', '字段名[你的变量]', 'string', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('2b8c4728-bd15-41e0-a660-2a5f4bc2f33c', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.services', '服务流程的节点，', 'array', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('2c51a75a-2055-4018-87ae-e8914c2dd834', '3522bb77-0cef-4a33-9e83-5ff05b2cf88f', 'cache.fields', '缓存的字段规则', 'array', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('2cb699e8-f0b7-4271-a7c1-efa4b691aabe', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', 'module', 'bulk_service', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('2f37eeca-9aa2-4da2-905e-663627eff653', '803b90d9-c58c-4113-b4c0-58782e03142c', 'module', 'bulk_upsert', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('2f6a7634-14f9-4653-a8ce-03fc28da4f35', 'e3d5dce4-c758-41cc-9ffd-578ca39836e3', 'prevent_duplication.key', 'prevent_duplication', 'string', '', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3268d4f5-2666-482f-877b-ee9eee1dfbee', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.connection', 'ldap 连接信息', 'json', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('326e84d4-9b4f-42db-be0f-8924cb069e3c', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', 'params[你字段]', '字段定义规则', 'json', '', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('34253791-de45-4177-8a46-087f25a38369', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'params.你的field.check.template', '校验规则模板，true为正常', 'template', '', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('347640f3-fa53-4aa3-a048-c0018f3e2de9', 'e3d5dce4-c758-41cc-9ffd-578ca39836e3', 'fields', '字段标志', 'array', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3566d235-3f63-4380-83fd-946c090daee6', '79d0939d-16f6-4c8d-b611-5a55409162ad', 'item', 'item 的名称', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('36870f56-8351-41be-a192-82d7aa361f87', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'second', '缓存存的时间', 'string', '是', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('36d4fed2-7581-4336-b1b8-788c4e1bc8c3', '91d4abd1-14c9-4705-9de7-1518da998427', 'module', ' module 表示模块，表示你将要运行的主体模块。module是服务核心关键字', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('37443fbc-fc58-485e-acff-778665c7689f', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.ModifyDnParams', '修改ou对应的参数', 'json', '', 120);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3788440b-8737-49ee-a249-8639be60d6b6', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'right', '右边数组', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3942ae40-a36d-4daa-a9d2-97c2030c48b9', '36248998-48ee-41db-a0f4-2010b056b2a3', 'result_name', '返回文件的名称', 'string', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3b3345b5-5e7d-4a4b-8534-930f3330af99', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'nick', '昵称', 'string', '', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3d25ef14-b51f-48a3-aace-c2a8019f1ab3', '45d1d393-6758-4c37-8ba9-108493f67b8e', 'item', '如果拼接某个字段的参数，字段的取值，[你的变量]', 'string', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3de12bd8-2155-4be3-8fcd-ff7580a39732', '234fe81b-c7e3-46a1-a274-2bbc904e6ab7', 'module', '运行的模块', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3e0e3ffc-1054-4632-90ea-1f1120f39834', '79d0939d-16f6-4c8d-b611-5a55409162ad', 'foreach', '循环数组对象，[你的变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3f07899b-b6c2-4f2f-b320-ffb41f90dce2', 'e4afe3cd-34eb-4837-b08f-9da741a7f3f8', '1', '', '', '', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4016cf75-ce77-454a-b32a-a7433fe1928f', '30f9e473-dead-404f-ad33-f2c00fa4e6ad', 'schedule_spec', '定时执行，具体执行参数https://github.com/robfig/cron', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('40a9b434-8d50-4f3c-9069-fb02679e90c4', '45d1d393-6758-4c37-8ba9-108493f67b8e', 'service', 'json下面service 调用目标服务，如何需要引入参数变量,[你的变量]', 'json', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('42d3175a-bcc2-42bb-9c39-aa4c807dea8b', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'item', '左边对象名称取值', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('42d97868-8279-4da1-928a-3158cf0bb731', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', 'params[你字段].check.err_msg', '当检查不通过的提示信息', 'template', '是', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('431d5597-4e46-422c-925c-bf3674605521', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'fields', '字段取值', 'array', '是', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4330025b-98ee-410d-a3a1-9a41199945e6', '25d3fdcb-162a-4ef2-a4af-3b7edcc6e42a', 'field', '结果里面的字段', 'template', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('478c6549-b317-4681-b47a-e315e61bbff0', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'handler_params[get_modify_data]', '处理器中key', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4aa61ce8-295d-4894-be0a-458ef7fcf987', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.finish', '和handler_params 通样配置，无论流程成功与失败都会运行', 'json', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4cd3302a-72bd-4aa7-8514-b43aa2d04c6c', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.services[node_fail]', '失败后运行的下个状态', 'template', '', 110);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4d570669-bd86-4c77-8704-72f404467dcc', 'e32693d9-a954-48ac-a1bb-097a0684c8b3', 'fields', '返回参数的内容', 'array', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4d8d2865-1bb3-46d6-af82-01a1e141ad97', '0722aa5a-78e6-4a29-818c-ba33843bdf81', 'table', '数据库表名', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4e528264-0083-4fc8-a599-90a9ee597ff0', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[target_transfer_value]', '转换后的取值字段，根据编码换值', 'string', '', 210);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4e939fb2-af7e-45d0-959f-7b47ac6f3251', 'ede82dd3-1b4f-4097-9260-c2d9c7acdf21', 'key', 'update_field', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('4ed574cf-ea23-40ea-b9ad-21e272a400ba', '0722aa5a-78e6-4a29-818c-ba33843bdf81', 'update_fields', '只更新哪些字段', 'array', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('516c94ab-807f-45a6-b9f7-654afa0e0890', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'module', 'http', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('523c774a-9ac1-4198-98f6-b9dd0c9c77b9', 'd25ddb2a-2c4e-4d3f-83c3-50ddc4215de8', 'module', 'empty', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('52a1871b-16d6-4ac0-9b2c-bec2fd0397c7', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', 'params[你字段].check', '字段校验', 'json', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('532f50c0-6ca9-43a7-88be-dcb667020d39', '0b3ebb5e-8ce6-4ebc-ae92-da6b0dfe13a2', 'foreach', '循环哪个数组变量,[你的变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5655cdab-3afc-4f08-afa1-4ebb15a49b2a', '3c768fa7-2c9c-48f4-a109-e740ab17ab6c', 'save_field', '如果有结果，结果存储在哪个字段', 'string', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('572382dd-0ae1-4e9f-86eb-3888f542ff6b', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.services[node_next]', '运行完成后，下个节点的状态', 'template', '是', 90);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('573ee2f4-6606-4e48-8c9e-8b477690a838', 'ce02406f-f729-47da-a119-35ed01c6c1c3', 'key', 'arr2obj', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('58114cec-7180-479c-aa27-dd849ed29379', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json.basic_auth.password', '登陆密码', 'string', '', 110);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5899f5fb-bcfb-4100-b39b-f6578d03ed8d', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[rule]', '规则名称。compare_field_value简单字段，simple_array_value简单数组，array_obj_value数组对象', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5a4142c3-b101-486a-9158-223d62f13e5c', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', 'batch.service', '常规服务json，如果service字典是动态变量，则需要在data_file 中注册，data_file 就是map', 'json', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5ae7e6fa-5802-495f-b60a-6427f5ab88c5', '0722aa5a-78e6-4a29-818c-ba33843bdf81', 'ignore_fields', '忽略哪些字段，比如用户修改表单里面传了密码，但是要求密码不能改', 'array', '', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5bb165f0-cd06-4cff-a908-e43a5d824593', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', 'params[你字段].check.template', '运算结果true 表示正常，false 表示检查不通过', 'template', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5c2ebd6e-219f-41eb-a26f-ffa125dd3031', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json.header', '请求头', 'json', '', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5e891d52-cd7c-40f5-81d7-b1f3def6cdac', '803b90d9-c58c-4113-b4c0-58782e03142c', 'ignore_fields', '忽略哪些字段', 'array', '', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5eef089c-439d-4dcd-8c39-8ee48d298257', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[left_field]', '当左右2边数据对比字段不对等到时候，左边定位数据需要字段a，右边要取字段b。主要用于定位数据，左边的字段', 'string', '', 130);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5f2eb17d-eccf-4015-8d15-99e46c61bb99', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', 'batch.append_item_param', '是否拼接循环变量里面参数', 'string', '', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('5f6b972c-8af8-406d-b11a-9dc1df2dad44', 'd4ee2257-8a51-40bf-8d48-4b5469858e21', 'handler_params[key]', 'data2excel', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('601627b7-e4a1-48b2-b11f-2a7b90d1312b', '3522bb77-0cef-4a33-9e83-5ff05b2cf88f', 'cache', '缓存模块,和参数处理一样的字段', 'json', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('606a3ec6-1cd6-4470-a0d0-e8119a735c5d', '8ec4053f-1a54-4f78-9de6-31db6995692e', 'module', 'module: model_delete', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6218982d-28a6-4bfd-9ce4-c10742c76461', '74973d56-9773-432f-b1d5-3b68ad40998d', 'key', 'filter_arr', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('622511c1-c55f-49ba-80dc-0a0c31a7c4aa', '3c768fa7-2c9c-48f4-a109-e740ab17ab6c', 'key', '运行哪个参数处理器，核心字段，根据这个参数决定运行哪个代码', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('62f27178-66d4-4383-91cd-341446cd764d', 'e3d5dce4-c758-41cc-9ffd-578ca39836e3', 'fields[field]', '字段取值', 'template', '是', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('630b3b4a-a779-4dbf-bc8c-e649212c2eba', '0b3ebb5e-8ce6-4ebc-ae92-da6b0dfe13a2', 'fields', '忽略数据规则', 'array', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('646d4c1d-8937-4b66-914d-6459dce35d96', 'f7350f20-20fb-47f4-8b19-511c8b5b38a9', 'fields[template]', '校验模板，true表示正常', 'template', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('65213015-da48-40f5-aadf-c0c6a71c5230', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json.basic_auth.username', '登陆账号', 'string', '', 100);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('689a67fc-de74-4a52-b36e-18eaacef92b3', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'params.你的field.template', '生成数据的模板', 'template', '', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('691ea915-9aa8-4c4f-8287-d961ae73966e', '3c768fa7-2c9c-48f4-a109-e740ab17ab6c', 'enable', '判断这个参数是否可以用，可以用 [你的变量]，一般会配置个enable 变量，则 [enable],比模板快些', 'template', '', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6a0a9c31-e522-4b1e-b912-c91bb2fea231', '8ec4053f-1a54-4f78-9de6-31db6995692e', 'table', '数据库的表名', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6b5338ae-bec0-4a48-bd2a-e533e49b6a71', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'table', '数据库表名', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6c7d005f-74e8-426e-84e4-948d2b93da51', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.op_field_transfer', '操作的转换字典，change_list 有些固定字段，比如name表示名称，如何有冲突可以在此处修改', 'json', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6cda28c3-c820-42fb-8d50-7a7af5606090', '7da94650-209f-411e-aa0c-47989dbd4409', 'collect_doc_id', '文档ID', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6e258782-9f61-4b4c-a14f-76bcd29754d8', '849062cb-1f83-454e-b81d-3d7754f9ac4a', 'key', 'update_array', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6e789735-03a8-4ffd-a853-439cbb6ad349', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'user_id', '用户ID', 'string', '', 90);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6e8d7080-22bd-4828-8b93-6fe6604c37b7', '45d1d393-6758-4c37-8ba9-108493f67b8e', 'append_item_param', '默认拼接某个字段的参数', 'boolean', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6f89dd73-e683-4663-b6be-95f8a388a78f', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.services[node_type]', 'start 开始，node节点，end结束', 'string', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6f9ec18c-6ac4-43fd-8f04-471138f6325a', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'field', 'item的取值', 'string', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('6fc7239b-d3f6-460c-aafa-7dc6b99660ac', 'e3d5dce4-c758-41cc-9ffd-578ca39836e3', 'enable', '必须有值，true校验', 'template', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('70fd5c9d-f536-4030-882e-b4c7a3bf4a7c', 'c667be3c-479d-4a29-be13-9ff9963232de', 'width_doc', '是否带出文档', 'boolean', '', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('7535f8ce-8cd9-40ef-847d-a9f63bb1e137', '79d0939d-16f6-4c8d-b611-5a55409162ad', 'key', 'arr2arrayObj', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('75b842ef-01c2-44d8-878e-305d24792001', '5881b6e1-8217-475f-9b05-4bc20ffdae8e', 'fields[from]', '如果没有配置值from,将整个结果转字段，如果配置from，这将结果里面字段转参数。[你的结果变量]', 'string', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('7634e0f5-0a9b-4f36-bf14-0252bacd6b3b', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'user_status', '用户状态，关联码表user_job_status', 'string', '', 110);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('76350615-9b97-43a3-9c47-86a893df6704', 'f7350f20-20fb-47f4-8b19-511c8b5b38a9', 'key', 'check_field 检查params中的字段', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('77a462a8-e3f7-4b25-a374-1fed7c70f5b7', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'key', 'update_array_from_array', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('78bd758d-2868-4880-996c-a5d5dad4e5ac', '3c768fa7-2c9c-48f4-a109-e740ab17ab6c', 'template', '判断参数处理器的结果是否正常，true 表示正常，false表示异常', 'template', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('79643bf3-547b-498a-a281-26e5a0ba4e6d', 'e32693d9-a954-48ac-a1bb-097a0684c8b3', 'fields[to]', '目标字段', 'string', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('79b68f06-880a-46ca-98d0-9db1257daefb', '5881b6e1-8217-475f-9b05-4bc20ffdae8e', 'fields', '需要转参数的字段内容', 'array', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('7b6aaba8-3f4f-4bcf-9ab3-2e7269694883', 'f272ba24-4c3e-4072-be07-abae59c93907', 'key', 'prop_arr', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('7c1454a3-915d-4db9-ac0f-07009655aa4b', 'e3d5dce4-c758-41cc-9ffd-578ca39836e3', 'prevent_duplication', '防止重复请求', 'json', '', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('7c7d2fcc-47b7-4a2d-9d2b-8a3c9ce7fb4c', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[save_original]', '是否保留原始值，取值为value，主要用户转换，看下面transfer 和service', 'boolean', '', 180);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('7e8a6b42-1d26-4840-886a-e7e62325e31e', '0b3ebb5e-8ce6-4ebc-ae92-da6b0dfe13a2', 'params', '参数params 定义为哪个变量。注意这里没有定义item', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8038a4f8-4ea5-41d7-a9ea-decee4dfc205', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.connection.password', '登陆的密码', 'string', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('81b77ca2-4d76-4c77-8c0b-eee269bcd6d0', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'fields[field]', '字段的key，[你的变量]', 'string', '是', 90);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('82aae936-f596-432e-b58f-0cfb58775324', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', 'params', 'param 是map ,下面一级表示你的变量，比如示例用user_list ，表示请求服务有个user_list 变量', 'json', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('835930f2-d6fd-4a0c-b1d9-33df71ace1d1', '67193cc3-900e-4c27-a3f6-75ee2dba5688', 'foreach', '循环的数组', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('83ad1fdc-4567-4865-bdb0-f00fbc21a78b', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file', 'ldap 配置路径地址', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('84f189ac-4211-4d75-acc1-1599dbcf7d16', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.connection.user', '登陆的用户信息', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('85330246-e93d-43da-99ae-33eb01552ccc', '31e7d7f9-c87e-4558-bd76-f673d34e1c84', 'foreach', '循环的数组，[你的变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('888c8894-bbec-4509-8467-451fe3a07e80', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', 'params[你字段].default', '默认值，类型不限制', '', '', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8a4c6ba2-6a0b-4a15-9713-a8d46b07ee81', '25d3fdcb-162a-4ef2-a4af-3b7edcc6e42a', 'key', 'count2map', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8b8f4678-aee1-4f7a-9a6d-9c6daf7f5757', '5881b6e1-8217-475f-9b05-4bc20ffdae8e', 'key', 'result2params', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8bdb94de-d0bb-4966-b817-f54d04b0e6eb', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.SearchParams', '搜索对应的参数，具体参数看示例', 'json', '', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8c2a8e65-34d2-46bf-a3e7-df9dfa5700db', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'success', '判断结果是否成功，变量直接返回数据的json字段。true 表示成功', 'template', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8c62149a-3694-412f-bb3f-5584ef8b9a62', 'f7350f20-20fb-47f4-8b19-511c8b5b38a9', 'fields[field]', '字段名称', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8d1b0f59-2fb3-4526-8841-f302256314e7', '234fe81b-c7e3-46a1-a274-2bbc904e6ab7', 'handler_params', '参数处理', 'array', '', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8e22e54d-d5c5-46e1-92d0-947f385b0de9', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[operation]', '仅仅对简单字段修改有效，对操作名称进行重新调整，一般是add，modify，remove', 'string', '', 100);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8e34e99b-07cc-4e75-b5ce-fa6abd74d1ee', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.services[key]', '剩下的字段和handler_params一致，字段取决与具体【参数处理】，开始和结束不需要中间都要key', 'string', '是', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8f053568-1cd9-4ee6-afeb-334174230d32', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'params.你的field.check', '检查数据类型', 'array', '', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8f6c5822-a018-46f5-8a17-24a9ff3a6c43', 'e3d5dce4-c758-41cc-9ffd-578ca39836e3', 'room', '缓存的命名空间', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('91502811-6119-467f-b5fc-4839465206a1', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json.basic_auth', 'basic账号密码认证', 'json', '', 90);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('91fe4231-6ef6-4113-a74f-336d66cd7afd', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json.data', '请求参数，支持模板转义', 'json', '', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9223e89e-0e4a-492c-ab57-6e5075e444a4', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'size', '分页大小', 'number', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('924beff1-5290-4001-ba39-c46377cd2a44', '5b86a474-2af6-4e82-8d25-3635645bc315', 'key', 'session_get', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('92647962-1233-4ae7-94e0-5d8632f1155d', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[left_value_field]', '主要用于数组对象，对比行记录里面其他字段值，左边的取值', 'string', '', 150);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('935ccf73-811e-4102-bf51-53f8e2abfe1e', '3522bb77-0cef-4a33-9e83-5ff05b2cf88f', 'cache.fields[field]', '字段取值', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('95db6e52-9289-4188-9e23-9e69315cc4bb', '0722aa5a-78e6-4a29-818c-ba33843bdf81', 'filter', 'fitler.你的字段__操作符，如果没有__表示进准匹配。操作符号支持__in、__isnull。filter 主要作用就是定位数据。必须有个filter 误配置进行全表更新', 'json', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9694624e-f3e0-4f0d-b1b8-84ba8e822fda', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'search', '模糊匹配昵称+登陆名', 'string', '', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('96d16568-e4c1-4503-9307-706c4a542600', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[target_transfer_key]', '目标服务的取值关键字段，定位行数据，根据编码定位行', 'string', '', 200);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('97285713-37d8-49a7-baa1-ba253111d041', '79d0939d-16f6-4c8d-b611-5a55409162ad', 'fields[field]', '字段名称', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('98b51242-8f17-4a2f-bbf4-b7fbc83aad0b', '31e7d7f9-c87e-4558-bd76-f673d34e1c84', 'children', '分组的数据', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('99f5c36e-2522-41d6-8876-6ca70fe59cc3', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'foreach', '循环数组的变量，批量设置缓存用', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9cf605b3-c05d-4ef7-8c81-1e9e49506964', '01044aa5-6f65-4b59-a4a9-ae7a0191be4c', 'fields[template]', 'session 取值', 'template', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9dae11fd-94ed-4c35-9c06-a58bddae4e58', '45d1d393-6758-4c37-8ba9-108493f67b8e', 'append_param', '是否拼接全部参数,默认不拼接', 'boolean', '', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9e7d0588-c495-4125-92bb-323a1af3a2b9', '293dffa7-1a0b-4e72-8bd7-0f2a44758495', 'fields', '删除session的内容', 'array', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9e943ed8-4d4f-4dd3-ad7d-1e538429a66d', '31e7d7f9-c87e-4558-bd76-f673d34e1c84', 'fields[field]', '具体去重的字段', 'template', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9ebaf6e9-247d-4f41-98bf-abdac777007c', '01044aa5-6f65-4b59-a4a9-ae7a0191be4c', 'key', 'session_add', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9f4f2685-1dd0-4c57-8b21-74dceeaf2b35', '1ba89c00-fd10-4756-9414-070bf51505d7', 'field', '[你参数变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9f8c5c77-a781-4760-b509-0c34bb33a9d6', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', 'batch.foreach', '循环哪个列表', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('a461a8ca-8f6d-48b5-9b3c-a414d4c7fac5', '70bbd0a8-5e8a-4083-9986-e8c0c320ebe0', 'right', '右边的数组', 'string', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('a542abf7-3491-4596-8141-aadf8bfe7443', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'exclude', '排除用户', 'string', '', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('a97e9dc2-e1aa-4093-9e7e-d668d37c0f08', 'eee7f650-4583-4bf6-b770-ce853eba8c54', 'field', '参数来源字段,[你的变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('aac01e3d-d670-407a-864e-10997f1888b3', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'params', 'params 是通用模块，下面key是对应字段', 'json', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ac591ca5-3b26-401c-bac6-8518d11ed618', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'params.你的field.type', '转换的数据类型，支持int，bool，int，int32，int64，bigint，float，time.time，time，sql.nulltime', 'string', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ad7a58b9-6a14-4c1a-852d-703fafd7fad9', '8ec4053f-1a54-4f78-9de6-31db6995692e', 'filter', '参考model_update 的fitler', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('adaa64fe-6b10-43b2-9210-2fa11f0350f6', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[right_field]', '当左右2边数据对比字段不对等到时候，左边定位数据需要字段a，右边要取字段b。主要用于定位数据，右边的字段', 'string', '', 140);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b04e63e3-0fb7-4a6b-aed2-77a80cdce550', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.AddParams', '添加对应的参数，具体参数看示例', 'json', '', 90);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b0b1db7a-d0ba-483a-a013-32d3b64a468f', '849062cb-1f83-454e-b81d-3d7754f9ac4a', 'fields', '字段内容', 'array', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b4a2110a-eb2a-4dfa-946e-9958fdb7138d', '5b86a474-2af6-4e82-8d25-3635645bc315', 'fields', '获取session 数据内容', 'array', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b53f4483-54f5-4192-bc97-a07103bddc48', '849062cb-1f83-454e-b81d-3d7754f9ac4a', 'foreach', '循环数组的对象，[你数组变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b654adc6-6efa-4690-80ca-4bcfc7762a3b', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'item', '循环item 名称', 'string', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b754652f-019e-4bc8-b081-929eff0827bb', '894496fd-4201-4b40-95d6-24ecd978b719', 'model_field', '取哪个列表字段。[你参数变量]', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b781cfb3-432d-4d77-a9b0-565682b0ae17', '36248998-48ee-41db-a0f4-2010b056b2a3', 'path', '文件路径，[你的变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b7aeed30-a819-4a93-b147-6f5e6e2f4fc1', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'room', '缓存的空间', 'string', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b7b62a99-e858-4a6f-8c87-24d230bf7151', '36248998-48ee-41db-a0f4-2010b056b2a3', 'key', 'file2result', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b8e9fb8e-0751-4441-b921-62935d435818', 'e3d5dce4-c758-41cc-9ffd-578ca39836e3', 'second', '这共用缓存的字段，实际单位是毫秒', 'number', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b959607b-376e-476e-a324-8857394fa85d', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'fields[field]', '字段名称', 'template', '是', 80);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('b9b05a30-5fa6-43a0-9b5e-6f354f3abc83', '5b86a474-2af6-4e82-8d25-3635645bc315', 'fields[key]', '获取session 的key', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('bb089c1f-b442-40a5-855d-ab42cb63488f', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', 'batch.save_field', '单个服务运行的结果保存到哪个字段', 'string', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('bb5ca993-7624-446c-a1eb-2e720bd2eb9e', '849062cb-1f83-454e-b81d-3d7754f9ac4a', 'item', '循环时候模板变量名', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('bbd27c80-7c64-4646-be92-710d733a5102', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'file_data', 'sql 文件的相对路径地址', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('bc2aec30-11ff-47b4-8751-83322e9b9e5e', '803b90d9-c58c-4113-b4c0-58782e03142c', 'update_fields', '更新哪些字段，不填所有字段', 'array', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('be4bf3f1-9bfe-4b2d-8345-6e7c9add15ec', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.DeleteParams', '删除对应的参数，具体参数看示例', 'json', '', 110);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('bebec83d-d2a3-4747-be70-d781972d87be', '234fe81b-c7e3-46a1-a274-2bbc904e6ab7', 'result_handler', '和参数处理是一样，所以result_handler的文档是参数处理是一致的', 'array', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('c0efc366-adb8-4260-9a0a-9bafff103ee7', '5881b6e1-8217-475f-9b05-4bc20ffdae8e', 'fields[to]', '转到参数的哪个字段,[你的变量]', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('c1942a06-de61-43dd-8022-0cc67ecaaf09', '67193cc3-900e-4c27-a3f6-75ee2dba5688', 'value', '字段值[你的变量]', 'string', '', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('c2d6a8fc-b1e6-470a-b36a-219ba1e88ef0', '3c768fa7-2c9c-48f4-a109-e740ab17ab6c', 'err_msg', '错误提示信息', 'template', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('c3771ca4-4652-4f33-94d9-55545ff57993', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'key', 'handler_cache', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('c52a48c8-ad0f-4d8f-8551-822a6c1a2053', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[append_left_fields]', '拼接左边的字段，一般数组对象修改，左右2边都有的情况，优先左边的字段，配置op_field_transfer,右边已经存在字段护理', 'array', '', 120);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('c73964b6-5460-4ca6-a264-9d5e838d6f28', '803b90d9-c58c-4113-b4c0-58782e03142c', 'table', '修改哪个表', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('c86777ce-2c3f-4fd6-a09d-26cf35ce2cd8', 'd89f00cd-2ecf-43d8-9fb4-f9bd26c1cf05', '11', '11', 'string', '', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('caf189f1-8779-4587-b27c-cd25d8313acb', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'module', 'sql,conf/application.conf 中配置数据连接，启动程序则连接', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('cc46f71a-c75c-42b7-a9ce-a0ca557653df', '849062cb-1f83-454e-b81d-3d7754f9ac4a', 'fields[fields[field]]', '字段渲染的模板内容，item是当前行数据，可以直接取params中的变量', 'template', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('cc841f73-54a9-4db3-94e1-09eddaa7395b', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'field', '左边匹配关键字段[你的item变量] 支持& 取2个字段', 'string', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ccea8e8c-faa9-4e4d-b57a-dfcfd21b30b2', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'username', '用户名', 'string', '', 100);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('cd0d05f2-71a2-48b4-82d1-1e8d97167760', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.services[ignore_error]', '是否忽略错误', 'boolean', '', 100);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('cd1ca175-7711-4db8-943b-0e410eaaa036', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config', '规则路径，在主体中', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('cd7f834b-619f-4480-ba3c-6d51e4809ed5', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', 'params[你字段].type', '支持转数据类型，比如bool,int', 'string', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('cdbb5c20-d3fd-4c61-81e6-50d5f5f3358b', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'pagination', '是否分页，默认分页', 'boolean', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ce06c82f-f817-4f29-ad0f-e46cf3ececf9', '234fe81b-c7e3-46a1-a274-2bbc904e6ab7', 'params', '参数定义', 'json', '', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ceed476e-10a2-4f6f-8350-5221168b9f3e', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'module', 'ldap,运行ldap模块', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('cf642c9d-c130-4525-b194-4a994958b339', '31e7d7f9-c87e-4558-bd76-f673d34e1c84', 'key', 'group_by', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d1b4953a-4e55-45a9-8ac7-5408b6ac7457', '70bbd0a8-5e8a-4083-9986-e8c0c320ebe0', 'right_field', '右边的关键字段', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d1bc2c21-fb04-4859-8424-6771875a206a', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'pagination', '是否运行分页,一般[pagination]。[]包起的部分是参数字段，在count_sql 中pagination为false', 'string', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d2658294-6182-42c3-9e7c-ea18ae7f5986', 'bd6431c8-caf3-4ae9-96af-03cce615204b', 'data_file', '配置文件地址', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d2cbe928-7149-4822-acba-44e15927d46c', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[name]', '对比字段的中文名称', 'string', '是', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d436eb80-5e5d-4ff4-8f1c-c62d4f1ae06e', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[value_list_field]', '将左右2边的值取出来，从另外一个目标服务查询转换一下，为下面service提供取值列表字段，一般current_value_list', 'string', '', 190);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d51c4bb2-23a2-4d3a-b32a-7db31f83a400', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[service]', '转换调用的服务，如果current_value_list 有冲突，可以改value_list_field', 'json', '', 220);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d5c44120-3d4f-48b2-8ee7-843668a93df2', '31e7d7f9-c87e-4558-bd76-f673d34e1c84', 'fields', '分组的字段', 'array', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d7f3be77-8425-413a-9e24-463954dceb80', 'bd6431c8-caf3-4ae9-96af-03cce615204b', 'handler_params[key]', 'file2datajson', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d91b4859-db42-4bcf-b80c-650f94f3452a', '894496fd-4201-4b40-95d6-24ecd978b719', 'module', 'module: bulk_create', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('d98c1c96-eacc-4083-840e-3e1d0eb19128', '77e10054-791e-4889-a5c8-1fc2b9a6e514', 'method', 'BULK_SET_CACHE 批量设置缓存，GET 获取缓存，SET设置缓存', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('db4ccdcf-b043-4f0c-8416-0dc0d8931c33', '70bbd0a8-5e8a-4083-9986-e8c0c320ebe0', 'key', 'combine_array', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('db8060cd-821e-4f90-aadd-57346f2e863a', '0722aa5a-78e6-4a29-818c-ba33843bdf81', 'filter.你的field__in', 'key表示  你的字段 进行in 操作，值表示 ：[你的参数变量]，取参数中的哪个值', 'string', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('db943402-9431-454a-8ce4-5d2498623ae9', '74973d56-9773-432f-b1d5-3b68ad40998d', 'foreach', '循环数组对象取值，[你的变量]', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('dc49a66f-0a68-497d-ba24-20d7e2b4b68b', 'e32693d9-a954-48ac-a1bb-097a0684c8b3', 'key', 'params2result', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('dc78a839-5906-4eea-a032-11068c576d6b', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.method', '执行方法。search:搜索，add:添加；modify:修改;delete:删除;modifyDn修改dn信息', 'string', '是', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('dd0bc3fc-0a32-4bcc-9f2e-455652d2a7e7', '67193cc3-900e-4c27-a3f6-75ee2dba5688', 'key', 'arr2dict', 'string', '', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('dd0f5b47-92de-4f37-9111-fdd6a94cdd69', 'e32693d9-a954-48ac-a1bb-097a0684c8b3', 'fields[from]', '来源哪个字段，支持[你的变量]', 'template', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ddc8ba73-2890-4591-87b9-daf9e01b45bf', '70bbd0a8-5e8a-4083-9986-e8c0c320ebe0', 'foreach', '左边的数组', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('de397537-1b06-4b63-a222-e77b4fd33a7e', '0b3ebb5e-8ce6-4ebc-ae92-da6b0dfe13a2', 'key', 'ignore_data', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('df630648-8f5f-46b6-adcc-89b20004762f', '3522bb77-0cef-4a33-9e83-5ff05b2cf88f', 'cache.second', '缓存的时效', 'number', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e1282c12-ddb8-428c-9f37-9d884ffee3a1', '67193cc3-900e-4c27-a3f6-75ee2dba5688', 'result_name', '转换的结果对象，二级数组有效', 'string', '', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e3b501c7-0c7d-47e9-acb7-f98deadfe9a2', '79d0939d-16f6-4c8d-b611-5a55409162ad', 'fields', '生成的对象字段内容', 'array', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e582d45a-9d99-4935-9da6-90135ca92052', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[with_add_remove]', '主要用于数组对象，是生成添加修改记录，一个字段数组对象中只要有一个', 'string', '', 170);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e5d7266e-a546-4383-9996-1bc02fc22187', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json', '服务配置地址', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e6697036-1e7b-4086-9b6e-48cc6110abf0', 'cd89c817-3ba4-48c5-baba-74c50c73dbbc', 'params[你字段].template', '模板生产数据', 'template', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e7ad393b-4057-4191-ace5-f8f86119ea4d', '79d0939d-16f6-4c8d-b611-5a55409162ad', 'fields[template]', '字段值，支持[你的参数]', 'string', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e9a03b85-4562-4a19-89a6-5dc50630451e', 'f4a4f3ed-8051-4754-92e2-2a5883fd9f98', 'data_json.services[node_key]', '流程的关键字，流程必须包含start,end', 'string', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e9bcbb8c-b246-4eea-a181-16f1fd11f41d', '01044aa5-6f65-4b59-a4a9-ae7a0191be4c', 'fields[key]', 'session 字段名称', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('eb12abb6-44c4-4807-83e2-5c8745fe5de3', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'foreach', '左边循环的数组，待更新的数组', 'array', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ebb6115c-926e-4afd-8b97-a9c3ede79de0', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json.result_json', '结果转json对象', 'string', '', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('edc1065c-d51f-44b7-a2ea-352bede720b3', '3522bb77-0cef-4a33-9e83-5ff05b2cf88f', 'cache.room', '缓存命名空间', 'string', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('eddd0c98-1df5-4597-982a-a9124b6b80ab', '272b32d6-f193-442f-b8ca-cc42f0f36f7e', 'count', '是否运行count， 默认运行', 'boolean', '', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ee14f301-96ad-4225-a846-70d8143d06f4', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'count_file', 'count sql文件的相对路径地址，没有就不运行，用于统计列表行数', 'string', '', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('eed9efb1-acb5-4d11-92ba-14ad08ca69d0', 'b9eab01d-7236-4d15-9053-0f193a3d2ffb', 'field', '结果字段的key', 'template', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ef0bfb0d-9b90-49a5-8e66-9fc6de45fa6c', '74973d56-9773-432f-b1d5-3b68ad40998d', 'item', 'item 变量取值', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('ef2dbadc-e93c-41c9-8fb9-7029bb5dd28f', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json.url', '请求地址', 'string', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f0218bb4-3034-448b-909b-0be9b1c55582', '803b90d9-c58c-4113-b4c0-58782e03142c', 'model_field', '数据列表取哪个参数,[你的参数]', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f10198d6-dcd8-4cac-9d12-3c83123da187', 'f272ba24-4c3e-4072-be07-abae59c93907', 'value', 'value 取值，[你的item的变量]', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f25dade7-9119-49ef-8e11-3cdd12c34b9e', '9215873e-d44d-4a8f-b243-f2c45b1833e5', 'data_file.connection.server', '服务器地址', 'string', '是', 40);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f597f33d-7967-4c85-b2ee-97ce99da0075', 'eb55515c-df08-4ac5-a4e8-3d811b834e54', 'data_file', '需要转换的服务文件，仅单service是动态变量时候有效', 'string', '', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f6523fe5-9aac-4efc-8060-477157e91ab4', '70bbd0a8-5e8a-4083-9986-e8c0c320ebe0', 'children', '生成的字段名称', 'string', '是', 60);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f7022fb1-e347-455c-9ac5-1d5653f08cf2', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'fields', '更新字段的内容', 'array', '是', 70);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f7b2fe78-bd61-4439-8dfa-13d82d5e0999', 'f7350f20-20fb-47f4-8b19-511c8b5b38a9', 'fields[err_msg]', '失败之后的错误提示', 'template', '是', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f8a84c68-321e-443e-be24-31f4e670f969', 'c0f45c6a-688e-46b7-ba0f-ac544f27c3b5', 'fields[template]', '支持[你的变量]，左边的用item,右边的用right', 'template', '是', 90);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f8eee343-1b60-4a41-b8c6-e25a6a189121', '293dffa7-1a0b-4e72-8bd7-0f2a44758495', 'fields[key]', 'key 名称', 'string', '是', 30);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f8f59df0-fe4a-491a-b681-d8aa33925a71', 'ede82dd3-1b4f-4097-9260-c2d9c7acdf21', 'fields', '更新的字段内容', 'array', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('f9518d71-bebe-44fa-af9f-86d99a2b2c2f', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json', 'http配置文件路径，整个文件对象模板渲染', 'json', '是', 20);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('fad00675-71a3-4717-9dab-cf9d1633be2e', 'c00a6f14-98c1-4308-8017-cf35ae300de4', 'http_json.method', '请求方法，GET,POST', 'string', '', 50);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('fc75172e-b834-4cf2-9549-6569d73a7c79', '293dffa7-1a0b-4e72-8bd7-0f2a44758495', 'key', 'session_remove', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('fd3e413e-1546-4c53-9eb8-09819c75cbe8', 'b7a26057-abd7-45a0-9101-63e01af6ed4c', 'modify_config.fields[right]', '右边取对象字段', 'string', '是', 90);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('fe8ed667-8ecf-4a02-9f60-af7c620bff82', '9b4fbf55-b221-4b05-86f1-2e78132ee552', 'module', 'module: model_save 执行新增模块', 'string', '是', 10);
INSERT INTO collect_doc_params (doc_params_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('fea39634-e9ff-4e1c-9030-9405be8ea173', 'ccc4f5b6-2e98-4692-9e6b-333c1ab404e0', 'count', '参数控制是否运行统计sql,一般配置[count]，并且有count参数。一般count_file 值默认运行,[count]的值为false，就不运行', 'string', '', 50);


INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('100e2fe1-6cac-4d2c-9c8d-e8f944c8b74a', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.parent_dir', '上级文件夹', 'string', '', 40);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('12a01a66-723e-46a9-bce4-61eddf22a523', '7da94650-209f-411e-aa0c-47989dbd4409', 'important_list', '要点', 'array', '', 80);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3a6cb6e2-fd0a-43d5-8864-c945340319f7', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.sub_title', '子标题', 'string', '', 60);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('41136fe9-9e1b-44a2-8cd9-7e6b4c7d42ac', 'c667be3c-479d-4a29-be13-9ff9963232de', 'children[code]', '文档编码', 'string', '', 30);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('54776748-42bb-4728-b46d-68e4e36608fd', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.code_desc', '代码描述', 'string', '', 30);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('70bd0195-7340-4453-b2a8-6192120c15fb', 'c667be3c-479d-4a29-be13-9ff9963232de', 'children', '文档', 'array', '', 20);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('75d8c922-c75e-4479-9b74-3a7e9d9879bc', 'c667be3c-479d-4a29-be13-9ff9963232de', 'name', '文档分组名称', 'string', '', 10);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8fccce9f-3748-44ec-9ec6-0d42a883e010', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.code', '示例代码', 'string', '', 20);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9da0ad32-b18a-42da-b0a4-874262ee0410', '7da94650-209f-411e-aa0c-47989dbd4409', 'params', '参数', 'array', '', 90);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('a33267a3-15a5-410d-a5f3-f20f2347e12e', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc', '文档', 'json', '是', 10);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('a8f7230f-6d2e-411e-9d2a-b667359ae0f0', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.title', '标题', 'string', '', 50);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e6cab251-b6ba-41b0-a6eb-c5b5170a2343', '7da94650-209f-411e-aa0c-47989dbd4409', 'result', '结果描述', 'array', '', 110);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('edc33693-4b09-4cb6-9912-550680af387f', '7da94650-209f-411e-aa0c-47989dbd4409', 'demo', '示例', 'array', '', 100);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('effa3fd6-bcf8-4bff-9f5b-fdc983473993', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.type', '类型，文档doc，服务service', 'string', '', 70);



INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('100e2fe1-6cac-4d2c-9c8d-e8f944c8b74a', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.parent_dir', '上级文件夹', 'string', '', 40);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('12a01a66-723e-46a9-bce4-61eddf22a523', '7da94650-209f-411e-aa0c-47989dbd4409', 'important_list', '要点', 'array', '', 80);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('3a6cb6e2-fd0a-43d5-8864-c945340319f7', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.sub_title', '子标题', 'string', '', 60);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('41136fe9-9e1b-44a2-8cd9-7e6b4c7d42ac', 'c667be3c-479d-4a29-be13-9ff9963232de', 'children[code]', '文档编码', 'string', '', 30);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('54776748-42bb-4728-b46d-68e4e36608fd', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.code_desc', '代码描述', 'string', '', 30);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('70bd0195-7340-4453-b2a8-6192120c15fb', 'c667be3c-479d-4a29-be13-9ff9963232de', 'children', '文档', 'array', '', 20);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('75d8c922-c75e-4479-9b74-3a7e9d9879bc', 'c667be3c-479d-4a29-be13-9ff9963232de', 'name', '文档分组名称', 'string', '', 10);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('8fccce9f-3748-44ec-9ec6-0d42a883e010', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.code', '示例代码', 'string', '', 20);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('9da0ad32-b18a-42da-b0a4-874262ee0410', '7da94650-209f-411e-aa0c-47989dbd4409', 'params', '参数', 'array', '', 90);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('a33267a3-15a5-410d-a5f3-f20f2347e12e', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc', '文档', 'json', '是', 10);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('a8f7230f-6d2e-411e-9d2a-b667359ae0f0', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.title', '标题', 'string', '', 50);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('e6cab251-b6ba-41b0-a6eb-c5b5170a2343', '7da94650-209f-411e-aa0c-47989dbd4409', 'result', '结果描述', 'array', '', 110);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('edc33693-4b09-4cb6-9912-550680af387f', '7da94650-209f-411e-aa0c-47989dbd4409', 'demo', '示例', 'array', '', 100);
INSERT INTO collect_doc_result (doc_result_id, collect_doc_id, name, `desc`, `type`, must, order_index) VALUES('effa3fd6-bcf8-4bff-9f5b-fdc983473993', '7da94650-209f-411e-aa0c-47989dbd4409', 'doc.type', '类型，文档doc，服务service', 'string', '', 70);

INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('1eed4d51-74bd-4bdf-8049-f48fedafcb41', 'ss', 'doc', 'ss', 10, '2023-11-23 08:48:18', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '1');
INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('2d07dfdc-1026-40fb-8124-ddc74b566265', '参数处理', 'doc', '', 40, '2023-11-23 09:01:07', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0');
INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('76ee2f0a-07e3-4202-a62a-2d18c5712d2b', '如何编写服务', 'service', '', 1, '2023-11-22 20:26:48', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '1');
INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('9501b424-606a-433a-9bc5-f9de8064e9d8', '模板函数', 'doc', NULL, 20, '2023-12-06 17:53:58', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0');
INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('ae78bab6-b68e-4522-86e0-08d35fe201d0', '模块处理', 'doc', '', 30, '2023-11-23 09:01:07', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0');
INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('c09992bf-afc8-4abb-a81b-24ff1060fe0b', '文档管理', 'service', NULL, 70, '2023-12-07 15:27:04', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0');
INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('dd336894-53b6-405b-98c7-f327407d7cfa', '如何编写服务', 'doc', '', 10, '2023-11-23 09:01:07', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0');
INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('efbb8e60-eeab-4326-b742-27cedbcc9083', '拦截器', 'doc', NULL, 50, '2023-11-23 09:12:20', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0');
INSERT INTO doc_group (doc_group_id, name, `type`, `desc`, order_index, create_time, create_user, is_delete) VALUES('f0c5c31f-d835-4aa6-8f39-c9fb6c1adfc8', '人资管理', 'service', '', 60, '2023-11-23 09:01:38', '739ade44-7e83-48a2-8c60-9a7c1e9f3d0a', '0');


