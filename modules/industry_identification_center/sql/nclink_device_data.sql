create table nclink_device_data(
 id bigint(20) unsigned auto_increment ,
 data_id varchar(64) NOT NULL,
 adaptor_id varchar(64) NOT NULL,
 device_id varchar(64) NOT NULL,
 component_id varchar(64) NOT NULL,
data_item_id varchar(64) NOT NULL,
payload text NOT NULL,
adaptor_time bigint(20) NOT NULL,
 create_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
 update_time datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)  ON UPDATE CURRENT_TIMESTAMP(3),
 delete_time datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000',
 primary key(`id`),
key `idx_data`(`data_id`),
 key `idx_adaptor`(`adaptor_id`),
 key `idx_device`(`device_id`), 
 key `idx_component`(`component_id`), 
 key `idx_data_item`(`data_item_id`) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment='nclink运行数据表';