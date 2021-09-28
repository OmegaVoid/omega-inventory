create table part_unit
(
	id         serial  not null
		constraint part_unit_pk primary key,
	uname      varchar not null,
	short_name varchar not null,
	is_default boolean not null--,
		constraint part_unit_is_default_not_false check (is_default != false)

);

create unique index on part_unit using btree (is_default nulls last);

create trigger part_unit_only_one_default
	before insert or update of is_default
	on part_unit
	for each row
	when (new.is_default = true)
execute procedure ensure_only_one_default_trigger();

insert into part_unit
values (default, 'Pieces', 'pcs', true);


create view default_part_unit as
select *
from part_unit
where is_default = true;


---- create above / drop below ----

drop view default_part_unit;
drop table part_unit;
