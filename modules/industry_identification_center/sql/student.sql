create table `student`(
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(30) NOT NULL,
    `age` tinyint UNSIGNED NOT NULL,
    `region` varchar(30) NOT NULL,
    `phone_number` varchar(30) default '',
    primary key (`id`)
)ENGINE=InnoDB default CHARSET=utf8mb4 comment='学生表';