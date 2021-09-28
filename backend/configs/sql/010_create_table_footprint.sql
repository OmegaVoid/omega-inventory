create table footprint
(
	id          bigint default id_generator() not null
		constraint footprint_pk primary key,
	fname       varchar                       not null,
	category    bigint
		constraint footprint_footprint_category_id_fk references footprint_category on update cascade on delete cascade,
	description text,
	attachments bigint[],
	image       bigint
);

---- create above / drop below ----

drop table footprint;