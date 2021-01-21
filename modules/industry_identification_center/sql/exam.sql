create table `exam`(
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
    `student_id`int UNSIGNED NOT NULL,
    `grade` tinyint UNSIGNED NOT NULL,
    `subject` varchar(30) NOT NULL,
    primary key (`id`),
    unique key `uniq_exam_grade`(`student_id`,`subject`)
)ENGINE=InnoDB default CHARSET=utf8mb4 comment='考试成绩表';