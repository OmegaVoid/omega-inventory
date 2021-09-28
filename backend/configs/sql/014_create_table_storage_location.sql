create table storage_location
(
	id       bigint default id_generator() not null
		constraint storage_location_pk primary key,
	lname    varchar                       not null,
	category bigint                        not null
		constraint storage_location_category_id_fk references storage_location_category on update cascade on delete cascade,
	image    bigint
		constraint storage_location_image_id_fk references attachment on update cascade on delete cascade

);

---- create above / drop below ----

drop table storage_location;