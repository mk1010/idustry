create table sampling_info_meta(
 id int(20) unsigned auto_increment ,
 sampling_info_id varchar(64) NOT NULL,
 sampling_info_type varchar(32) NOT NULL,
 sampling_period int(20) unsigned NOT NULL,
 upload_period int(20) unsigned NOT NULL,
 create_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
 update_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)  ON UPDATE CURRENT_TIMESTAMP(3),
 delete_time datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000',
 primary key(`id`),
 unique key `uniq_sampling_info`(`sampling_info_id`,`delete_time`) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment='采样信息元数据表';