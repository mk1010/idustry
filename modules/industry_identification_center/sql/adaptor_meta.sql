create table adaptor_meta
(
 id int(20) unsigned auto_increment,
 adaptor_id varchar(64) not null,
 name varchar(64),
 adaptor_type varchar(32) not null,
 description varchar(128),
 device_id varchar(1024) not null,
 device_config varchar(2048),
 create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP,
 delete_time datetime not null default '0000-00-00 00:00:00',
 primary key(`id`),
 unique key `uniq_adaptor`(`adaptor_id`,`delete_time`) 
)ENGINE=InnoDB default CHARSET=utf8mb4 comment='适配器元数据表';