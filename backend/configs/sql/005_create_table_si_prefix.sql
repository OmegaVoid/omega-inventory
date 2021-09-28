create table si_prefix
(
	id       serial  not null
		constraint si_prefix_pk primary key,
	prefix   varchar not null,
	symbol   varchar not null,
	exponent int     not null,
	base     int     not null
);

create unique index si_prefix_symbol_uidx on si_prefix (symbol);
create unique index si_prefix_prefix_uidx on si_prefix (prefix);

insert into si_prefix (id, prefix, symbol, exponent, base)
values (1, 'yotta', 'Y', 24, 10),
       (2, 'zetta', 'Z', 21, 10),
       (3, 'exa', 'E', 18, 10),
       (4, 'peta', 'P', 15, 10),
       (5, 'tera', 'T', 12, 10),
       (6, 'giga', 'G', 9, 10),
       (7, 'mega', 'M', 6, 10),
       (8, 'kilo', 'k', 3, 10),
       (9, 'hecto', 'h', 2, 10),
       (10, 'deca', 'da', 1, 10),
       (11, '-', '', 0, 10),
       (12, 'deci', 'd', -1, 10),
       (13, 'centi', 'c', -2, 10),
       (14, 'milli', 'm', -3, 10),
       (15, 'micro', 'μ', -6, 10),
       (16, 'nano', 'n', -9, 10),
       (17, 'pico', 'p', -12, 10),
       (18, 'femto', 'f', -15, 10),
       (19, 'atto', 'a', -18, 10),
       (20, 'zepto', 'z', -21, 10),
       (21, 'yocto', 'y', -24, 10),
       (22, 'kibi', 'Ki', 1, 1024),
       (23, 'mebi', 'Mi', 2, 1024),
       (24, 'gibi', 'Gi', 3, 1024),
       (25, 'tebi', 'Ti', 4, 1024),
       (26, 'pebi', 'Pi', 5, 1024),
       (27, 'exbi', 'Ei', 6, 1024),
       (28, 'zebi', 'Zi', 7, 1024),
       (29, 'yobi', 'Yi', 8, 1024);
;

---- create above / drop below ----

drop table si_prefix;