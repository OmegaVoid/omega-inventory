create sequence global_id_sequence;

create or replace function id_generator(out result bigint) as
$$
declare
	our_epoch  bigint := 1288834974657;
	seq_id     bigint;
	now_millis bigint;
	-- the id of this DB shard, must be set for each
	-- schema shard you have - you could pass this as a parameter too
	shard_id   int    := 1;
begin
	select nextval('global_id_sequence') % 1024 into seq_id;

	select floor(extract(epoch from clock_timestamp()) * 1000) into now_millis;
	result := (now_millis - our_epoch) << 23;
	result := result | (shard_id << 10);
	result := result | (seq_id);
end;
$$ language PLPGSQL;

alter function id_generator() owner to postgres;

select id_generator();

create extension ltree;