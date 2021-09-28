create table part
(
	id                bigint default id_generator() not null
		constraint part_pk primary key,
	fname             varchar                       not null,
	category          bigint
		constraint part_category_id_fk references part_category on update cascade on delete cascade,
	description       text,
	attachments       bigint[],
	comment           text,
	storage_locations bigint[],
	unit              bigint
		constraint part_unit_id_fk references part_unit on update cascade on delete cascade,
	stock_level       real,
	min_stock_level   real,
	parameters        bigint[],
	footprint         bigint
		constraint part_footprint_id_fk references footprint on update cascade on delete cascade
);

---- create above / drop below ----

drop table part;