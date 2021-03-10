drop database if exists `touken_api`;
CREATE SCHEMA IF NOT EXISTS `touken_api` DEFAULT CHARACTER SET utf8mb4 ;
USE `touken_api` ;

SET CHARSET utf8mb4;

-- table "toushu"
create table if not exists `touken_api`.`toushu`(
    `toushuID` int  comment '刀種ID',
    `toushu` varchar(64) comment '刀種名',
    primary key (`toushuID`)
)
comment = '刀種テーブル';

-- table "touken"
create table if not exists `touken_api`.`touken`(
    `toukenID` int comment '刀剣ID',
    `toushuID` int comment '刀種ID',
    `touken` varchar(64) comment '刀剣名',
    primary key (`toukenID`),
    CONSTRAINT `fk_touken_toushu`
        foreign key (`toushuID`)
        references `touken_api`.`toushu` (`toushuID`)
        on delete no action
        on update no action
)
comment = '刀剣テーブル';

-- table "exp"
create table if not exists `touken_api`.`exp`(
    `ID` int comment '主キー',
    `toushuID` int comment '刀種ID',
    `level` int comment 'レベル',
    `sum_exp` int comment '累計経験値',
    primary key(`ID`),
        CONSTRAINT `fk_exp_toushu`
        foreign key (`toushuID`)
        references `touken_api`.`toushu` (`toushuID`)
        on delete no action
        on update no action
)
COMMENT = '経験値テーブル';

