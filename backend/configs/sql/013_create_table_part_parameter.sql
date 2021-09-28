create table part_parameter
(
	id            bigint default id_generator() not null
		constraint part_parameter_pk primary key,
	pname         varchar                       not null,
	description   text,
	attachments   bigint[],
	unit          bigint
		constraint part_param_unit_id_fk references unit on update cascade on delete cascade,
	value         real,
	max_value     real,
	min_value     real,
	string_value  text,
	value_type    varchar                       not null,
	min_si_prefix int                           not null
		constraint part_param_min_si_prefix_id_fk references si_prefix on update cascade on delete cascade,
	max_si_prefix int                           not null
		constraint part_param_max_si_prefix_id_fk references si_prefix on update cascade on delete cascade,
	si_prefix     int                           not null
		constraint part_param_si_prefix_id_fk references si_prefix on update cascade on delete cascade

);

---- create above / drop below ----

drop table part_parameter;