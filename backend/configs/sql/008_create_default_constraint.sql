create or replace function ensure_only_one_default_trigger() returns trigger as
$$
begin
	-- nothing to do if updating the row currently enabled
	if (TG_OP = 'UPDATE' and OLD.is_default = true) then return NEW; end if;

	-- disable the currently enabled row
	execute format('UPDATE %I.%I SET is_default = null WHERE is_default = true;', TG_TABLE_SCHEMA, TG_TABLE_NAME);

	-- enable new row
	NEW.is_default := true;
	return NEW;
end;
$$ language plpgsql;


---- create above / drop below ----

drop function ensure_only_one_default_trigger();