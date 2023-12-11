-- collect_doc definition

CREATE TABLE `collect_doc` (
                               `collect_doc_id` varchar(50) NOT NULL,
                               `title` varchar(200) DEFAULT NULL,
                               `sub_title` varchar(200) DEFAULT NULL,
                               `type` varchar(200) DEFAULT NULL,
                               `parent_dir` varchar(200) DEFAULT NULL,
                               `code` varchar(2550) DEFAULT NULL,
                               `code_desc` varchar(2550) DEFAULT NULL,
                               `order_index` int(11) DEFAULT NULL,
                               `create_time` varchar(50) DEFAULT NULL,
                               `create_user` varchar(50) DEFAULT NULL,
                               `is_delete` varchar(50) DEFAULT NULL,
                               `code_result` text,
                               PRIMARY KEY (`collect_doc_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- collect_doc_demo definition

CREATE TABLE `collect_doc_demo` (
                                    `doc_demo_id` varchar(50) NOT NULL,
                                    `collect_doc_id` varchar(50) DEFAULT NULL,
                                    `name` varchar(200) DEFAULT NULL,
                                    `code` text,
                                    `order_index` int(11) DEFAULT NULL,
                                    `code_result` mediumtext,
                                    PRIMARY KEY (`doc_demo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- collect_doc_important definition

CREATE TABLE `collect_doc_important` (
                                         `doc_important_id` varchar(50) NOT NULL,
                                         `collect_doc_id` varchar(50) DEFAULT NULL,
                                         `name` varchar(200) DEFAULT NULL,
                                         `order_index` int(11) DEFAULT NULL,
                                         PRIMARY KEY (`doc_important_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- collect_doc_params definition

CREATE TABLE `collect_doc_params` (
                                      `doc_params_id` varchar(50) NOT NULL,
                                      `collect_doc_id` varchar(50) DEFAULT NULL,
                                      `name` varchar(200) DEFAULT NULL,
                                      `desc` varchar(2550) DEFAULT NULL,
                                      `type` varchar(200) DEFAULT NULL,
                                      `must` varchar(200) DEFAULT NULL,
                                      `order_index` int(11) DEFAULT NULL,
                                      PRIMARY KEY (`doc_params_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- collect_doc_result definition

CREATE TABLE `collect_doc_result` (
                                      `doc_result_id` varchar(50) NOT NULL,
                                      `collect_doc_id` varchar(50) DEFAULT NULL,
                                      `name` varchar(200) DEFAULT NULL,
                                      `desc` varchar(2550) DEFAULT NULL,
                                      `type` varchar(200) DEFAULT NULL,
                                      `must` varchar(200) DEFAULT NULL,
                                      `order_index` int(11) DEFAULT NULL,
                                      PRIMARY KEY (`doc_result_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- doc_group definition

CREATE TABLE `doc_group` (
                             `doc_group_id` varchar(50) NOT NULL,
                             `name` varchar(200) DEFAULT NULL,
                             `type` varchar(200) DEFAULT NULL,
                             `desc` varchar(200) DEFAULT NULL,
                             `order_index` int(11) DEFAULT NULL,
                             `create_time` varchar(50) DEFAULT NULL,
                             `create_user` varchar(50) DEFAULT NULL,
                             `is_delete` varchar(50) DEFAULT NULL,
                             PRIMARY KEY (`doc_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;