create table part_category
(
	id          bigint default id_generator() not null
		constraint part_category_pk primary key,
	cname       varchar                       not null,
	description text,
	path        ltree
);

create index part_category_path_idx on part_category using gist (path);

create unique index part_category_cname_uidx on part_category (cname);

create or replace function before_insert_on_part_category() returns TRIGGER
	language plpgsql as
$$
declare
	p ltree := ( select path
	             from part_category
	             where cname = ltree2text(new.path)::varchar );
begin
	new.path := case when p is not null then p || new.id::text else text2ltree(new.id::text) end;
	return new;
end
$$;

drop trigger if exists before_insert_on_part_category on part_category;
create trigger before_insert_on_part_category
	before insert
	on part_category
	for each row
execute procedure before_insert_on_part_category();

---- create above / drop below ----

drop table part_category;
drop function before_insert_on_part_category();