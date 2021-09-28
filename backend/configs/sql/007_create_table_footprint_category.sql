create table footprint_category
(
	id          bigint default id_generator() not null
		constraint footprint_category_pk primary key,
	cname       varchar                       not null,
	description text,
	path        ltree
);
create unique index footprint_category_cname_uidx on footprint_category (cname);
create index footprint_category_path_uidx on footprint_category using gist (path);

create or replace function before_insert_on_footprint_category() returns TRIGGER
	language plpgsql as
$$
declare
	p ltree := ( select path
	             from footprint_category
	             where cname = ltree2text(new.path)::varchar );
begin
	new.path := case when p is not null then p || new.id::text else text2ltree(new.id::text) end;
	return new;
end
$$;


drop trigger if exists before_insert_on_category on footprint_category;
create trigger before_insert_on_storage_footprint_category
	before insert
	on footprint_category
	for each row
execute procedure before_insert_on_footprint_category();


insert into footprint_category
values (default, 'RootCategory', null, null);
insert into footprint_category
values (default, 'BGA', null, 'RootCategory');
insert into footprint_category
values (default, 'CBGA', null, 'BGA');
insert into footprint_category
values (default, 'FCBGA', null, 'BGA');
insert into footprint_category
values (default, 'PBGA', null, 'BGA');
insert into footprint_category
values (default, 'DIP', null, 'RootCategory');
insert into footprint_category
values (default, 'CERDIP', null, 'DIP');
insert into footprint_category
values (default, 'PDIP', null, 'DIP');

---- create above / drop below ----
drop table footprint_category;
drop function before_insert_on_footprint_category();