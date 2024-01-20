-- Written by yijian on 2024/01/20
-- 创建信箱表 SQL

DROP TABLE IF EXISTS `t_mooon_mailbox`;
CREATE TABLE `t_mooon_mailbox` (
    `f_id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 ID（信件 ID）',
    `f_recipient` VARCHAR(32) NOT NULL COMMENT '收件人',
    `f_deliver_time` DATETIME NOT NULL COMMENT '投递时间',
    `f_arrival_time` DATETIME NOT NULL COMMENT '到达时间',
    `f_state` TINYINT UNSIGNED NOT NULL COMMENT '信件状态（0 未读，1 已读）',
    `f_letter_body` VARCHAR(4096) NOT NULL COMMENT '信件内容',
    PRIMARY KEY (`f_id`),
    KEY (`f_recipient`,`f_state`),
    KEY (`f_arrival_time`,`f_state`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb3;
