package main

import (
	"fmt"
	"innoobijr/collector/types"
)

func main() {


	// Making query

	conf := &types.Conf{}
	conf.GetConf("fake_conf.yaml")

	ctx := &types.Context{
		Params: conf,
	}

	query1 := new(types.Query)
	query1.Command = `
SELECT * FROM {{.Params.Clickhouse.Name}}
WHERE timestamp >= '{{.Params.Clickhouse.GreaterThanEqual}}'
LIMIT {{.Params.Clickhouse.Limit}}`

	query1.Type = "SQL"

	tpl := types.MakeQueryTemplate("clickhouse_query1", query1.Command)

	// tpl.Render("test", query1.Command)

	finalquery1 := tpl.Bind(ctx)

	fmt.Println(finalquery1)

	query2 := new(types.Query)
	query2.Command = `SELECT COUNT(*) FROM {{.Params.Changelog.Name}}`
	query2.Type = "SQL"
	tpl2 := types.MakeQueryTemplate("changelog_query1", query2.Command)
	finalquery2 := tpl2.Bind(ctx)
	fmt.Println(finalquery2)

	// Adding them to a query set, then doing the same thign I did above (creating a query template, binding to the context, then printing)

	var queries types.QuerySet
	queries = append(queries, *query1)
	queries = append(queries, *query2)

	for _, query := range queries {
		tpl := types.MakeQueryTemplate("query", query.Command)
		boundTpl := tpl.Bind(ctx)
		fmt.Println(boundTpl)
	}


	
}
