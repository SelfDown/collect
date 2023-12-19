

CREATE TABLE `config_detail` (
     `config_detail_id` varchar(100) NOT NULL,
     `group_id` varchar(100) DEFAULT NULL,
     `name` varchar(100) DEFAULT NULL,
     `value` varchar(255) DEFAULT NULL,
     PRIMARY KEY (`config_detail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `config_group` (
    `group_id` varchar(100) NOT NULL,
    `name` varchar(100) DEFAULT NULL,
    `description` varchar(255) DEFAULT NULL,
    `create_user` varchar(100) DEFAULT NULL,
    `create_time` varchar(100) DEFAULT NULL,
    PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;