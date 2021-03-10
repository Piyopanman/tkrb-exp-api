use touken_api;

SET NAMES utf8mb4;

insert into `toushu` (`toushu_id`, `toushu`) values (1,"特");
insert into `toushu` (`toushu_id`, `toushu`) values (2,"短刀");
insert into `toushu` (`toushu_id`, `toushu`) values (3,"脇差");
insert into `toushu` (`toushu_id`, `toushu`) values (4,"打刀");
insert into `toushu` (`toushu_id`, `toushu`) values (5,"太刀");
insert into `toushu` (`toushu_id`, `toushu`) values (6,"大太刀");
insert into `toushu` (`toushu_id`, `toushu`) values (7,"槍");
insert into `toushu` (`toushu_id`, `toushu`) values (8,"薙刀");

insert into `touken` (`touken_id`,`toushu_id`, `touken`) values(176, 1, "源清麿");
insert into `touken` (`touken_id`,`toushu_id`, `touken`) values(174, 1, "水心子正秀");

insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(101,1, 1, 0);
insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(102,1, 2, 100);
insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(103,1, 3, 300);
insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(104,1, 4, 700);
insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(105,1, 5, 1300);
insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(106,1, 6, 2100);
insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(107,1, 7, 3100);
insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(108,1, 8, 4300);
insert into `exp` (`id`,`toushu_id`, `level`, `sum_exp`) values(109,1, 9, 5700);