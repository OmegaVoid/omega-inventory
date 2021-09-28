create table attachment
(
	id            bigint default id_generator() not null
		constraint attachment_pk primary key,
	filename      varchar                       not null,
	original_name varchar,
	mimetype      varchar                       not null,
	description   text,
	is_image      bool,
	extension     varchar                       not null,
	size          int                           not null
);

---- create above / drop below ----

drop table attachment;
