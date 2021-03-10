use touken_api;

SET NAMES utf8mb4;

insert into `toushu` (`toushuID`, `toushu`) values (1,"特");
insert into `toushu` (`toushuID`, `toushu`) values (2,"短刀");
insert into `toushu` (`toushuID`, `toushu`) values (3,"脇差");
insert into `toushu` (`toushuID`, `toushu`) values (4,"打刀");
insert into `toushu` (`toushuID`, `toushu`) values (5,"太刀");
insert into `toushu` (`toushuID`, `toushu`) values (6,"大太刀");
insert into `toushu` (`toushuID`, `toushu`) values (7,"槍");
insert into `toushu` (`toushuID`, `toushu`) values (8,"薙刀");

insert into `touken` (`toukenID`,`toushuID`, `touken`) values(176, 1, "源清麿");
insert into `touken` (`toukenID`,`toushuID`, `touken`) values(174, 1, "水心子正秀");

insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(101,1, 1, 0);
insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(102,1, 1, 100);
insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(103,1, 1, 300);
insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(104,1, 1, 700);
insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(105,1, 1, 1300);
insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(106,1, 1, 2100);
insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(107,1, 1, 3100);
insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(108,1, 1, 4300);
insert into `exp` (`ID`,`toushuID`, `level`, `sum_exp`) values(109,1, 1, 5700);