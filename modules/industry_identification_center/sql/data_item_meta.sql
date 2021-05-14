create table data_item_meta(
 id int(20) unsigned auto_increment ,
 data_item_id varchar(64) NOT NULL,
 name varchar(64) NOT NULL,
 data_item_type varchar(32) NOT NULL,
 description varchar(128) NOT NULL default '',
 items varchar(2048) NOT NULL,
 data_unit varchar(2048) NOT NULL,
 create_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
 update_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)  ON UPDATE CURRENT_TIMESTAMP(3),
 delete_time datetime(3) NOT NULL default '0001-01-01 00:00:00.000',
 primary key(`id`),
 unique key `uniq_data_item_meta`(`data_item_id`,`delete_time`) 
)ENGINE=InnoDB default CHARSET=utf8mb4 comment='数据项元数据表';