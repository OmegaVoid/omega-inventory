package graphql

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/bwmarrin/snowflake"
)

func MarshalSnowflake(i snowflake.ID) graphql.Marshaler {
	return graphql.WriterFunc(
		func(w io.Writer) {
			w.Write(i.Bytes())
		},
	)
}

func UnmarshalSnowflake(v interface{}) (snowflake.ID, error) {
	switch v := v.(type) {
	case string:
		id, err := snowflake.ParseBase2(v)
		if err != nil {
			id, err := snowflake.ParseBase36(v)
			if err != nil {
				id, err := snowflake.ParseBase64(v)
				if err != nil {
					return snowflake.ParseString(v)
				}
				return id, nil
			}
			return id, nil
		}
		return id, nil
	case int:
		return snowflake.ParseInt64(int64(v)), nil
	case int64:
		return snowflake.ParseInt64(v), nil
	case json.Number:
		i, err := strconv.Atoi(string(v))
		if err != nil {
			return -1, err
		}
		return snowflake.ParseInt64(int64(i)), nil
	case []byte:
		id, err := snowflake.ParseBase58(v)
		if err != nil {
			id, err := snowflake.ParseBase32(v)
			if err != nil {
				return snowflake.ParseBytes(v)
			}
			return id, nil
		}
		return id, nil
	case [8]byte:
		return snowflake.ParseIntBytes(v), nil
	default:
		return 0, fmt.Errorf("%T is not a snowflake", v)
	}
}
