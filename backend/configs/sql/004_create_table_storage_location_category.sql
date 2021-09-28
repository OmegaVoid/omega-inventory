create table storage_location_category
(
	id          bigint default id_generator() not null
		constraint storage_location_category_pk primary key,
	cname       varchar                       not null,
	description text,
	path        ltree
);

create index storage_location_category_path_idx on storage_location_category using gist (path);

create unique index storage_location_category_cname_uidx on storage_location_category (cname);

create or replace function before_insert_on_storage_location_category() returns TRIGGER
	language plpgsql as
$$
declare
	p ltree := ( select path
	             from storage_location_category
	             where cname = ltree2text(new.path)::varchar );
begin
	new.path := case when p is not null then p || new.id::text else text2ltree(new.id::text) end;
	return new;
end
$$;


drop trigger if exists before_insert_on_category on storage_location_category;
create trigger before_insert_on_storage_location_category
	before insert
	on storage_location_category
	for each row
execute procedure before_insert_on_storage_location_category();

---- create above / drop below ----

drop table storage_location_category;
drop function before_insert_on_storage_location_category();