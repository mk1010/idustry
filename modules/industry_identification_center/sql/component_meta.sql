create table component_meta(
 id int(20) unsigned auto_increment ,
 component_id varchar(64) NOT NULL,
 name varchar(64) NOT NULL,
 component_type varchar(32) NOT NULL,
 description varchar(128) NOT NULL default '',
 config varchar(2048) NOT NULL default '',
 data_item_id varchar(1024) NOT NULL,
 sample_info_id varchar(1024) NOT NULL,
 create_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
 update_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)  ON UPDATE CURRENT_TIMESTAMP(3),
 delete_time datetime(3) NOT NULL default '0001-01-01 00:00:00.000',
 primary key(`id`),
 unique key `uniq_component`(`component_id`,`delete_time`) 
)ENGINE=InnoDB default CHARSET=utf8mb4 comment='组件元数据表';