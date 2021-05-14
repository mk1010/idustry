create table device_meta(
 id int(20) unsigned auto_increment ,
 device_id varchar(64) NOT NULL,
 name varchar(64) NOT NULL,
 device_type varchar(32) NOT NULL,
 description varchar(128) NOT NULL default '',
 device_group varchar(64) NOT NULL default '',
 component_id varchar(1024) NOT NULL,
 config varchar(2048) NOT NULL default '',
 create_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
 update_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)  ON UPDATE CURRENT_TIMESTAMP(3),
 delete_time datetime(3) NOT NULL default '0001-01-01 00:00:00.000',
 primary key(`id`),
 unique key `uniq_device`(`device_id`,`delete_time`) 
)ENGINE=InnoDB default CHARSET=utf8mb4 comment='设备元数据表';