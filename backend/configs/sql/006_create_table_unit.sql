create table unit
(
	id       serial  not null
		constraint unit_pk primary key,
	uname    varchar not null,
	symbol   varchar not null,
	prefixes int[]   not null
);

create unique index unit_symbol_uidx on unit (symbol);
create unique index unit_prefix_uidx on unit (uname);

insert into unit (id, uname, symbol, prefixes)
values (1, 'Meter', 'm', '{8,11,12,13,14,15,16}'),
       (2, 'Gram', 'g', '{8,11,14}'),
       (3, 'Second', 's', '{11,14}'),
       (4, 'Kelvin', 'K', '{11}'),
       (5, 'Mol', 'mol', '{11}'),
       (6, 'Candela', 'cd', '{14}'),
       (7, 'Ampere', 'A', '{8,11,14,15,16,17}'),
       (8, 'Ohm', 'Ω', '{5,6,8,11,14,15}'),
       (9, 'Volt', 'V', '{8,11,14}'),
       (10, 'Hertz', 'Hz', '{5,6,8,11,14}'),
       (11, 'Newton', 'N', '{8,11}'),
       (12, 'Pascal', 'Pa', '{8,11,14}'),
       (13, 'Joule', 'J', '{8,11,14,15}'),
       (14, 'Watt', 'W', '{6,7,8,11,14,15}'),
       (15, 'Coulomb', 'C', '{8,11}'),
       (16, 'Farad', 'F', '{11,14,15,16,17}'),
       (17, 'Siemens', 'S', '{11,14}'),
       (18, 'Weber', 'Wb', '{11}'),
       (19, 'Tesla', 'T', '{11}'),
       (20, 'Henry', 'H', '{11,14,15}'),
       (21, 'Celsius', '°C', '{11}'),
       (22, 'Lumen', 'lm', '{11}'),
       (23, 'Lux', 'lx', '{11}'),
       (24, 'Becquerel', 'Bq', '{11}'),
       (25, 'Gray', 'Gy', '{11}'),
       (26, 'Sievert', 'Sv', '{11,14,15}'),
       (27, 'Katal', 'kat', '{11}'),
       (28, 'Ampere Hour', 'Ah', '{8,11,14}');

---- create above / drop below ----

drop table unit;