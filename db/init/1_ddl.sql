drop database if exists `touken_api`;
CREATE SCHEMA IF NOT EXISTS `touken_api` DEFAULT CHARACTER SET utf8mb4 ;
USE `touken_api` ;

SET CHARSET utf8mb4;

-- table "toushu"
create table if not exists `touken_api`.`toushu`(
    `toushu_id` int  comment '刀種ID',
    `toushu` varchar(64) comment '刀種名',
    primary key (`toushu_id`)
)
comment = '刀種テーブル';

-- table "touken"
create table if not exists `touken_api`.`touken`(
    `touken_id` int comment '刀剣ID',
    `toushu_id` int comment '刀種ID',
    `touken` varchar(64) comment '刀剣名',
    primary key (`touken_id`),
    CONSTRAINT `fk_touken_toushu`
        foreign key (`toushu_id`)
        references `touken_api`.`toushu` (`toushu_id`)
        on delete no action
        on update no action
)
comment = '刀剣テーブル';

-- table "exp"
create table if not exists `touken_api`.`exp`(
    `id` int comment '主キー',
    `toushu_id` int comment '刀種ID',
    `level` int comment 'レベル',
    `sum_exp` int comment '累計経験値',
    primary key(`id`),
        CONSTRAINT `fk_exp_toushu`
        foreign key (`toushu_id`)
        references `touken_api`.`toushu` (`toushu_id`)
        on delete no action
        on update no action
)
COMMENT = '経験値テーブル';

