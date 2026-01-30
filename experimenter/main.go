package main

import (
	"fmt"
	"sync"

	// "time"

	"github.com/common-nighthawk/go-figure"
	log "github.com/sirupsen/logrus"
	types "go.cfdata.org/crypto/dome/changelog/models"
	core "go.cfdata.org/crypto/dome/collector/types"
)

/*
 This code does the following things is the prescribed order:
 (1) Grab configurations: this inlcude both secrets which are save a environmental variables and configuration parameters which can be found in the collector.yaml
 (2) Create the Query Jobs for both for this resource. You can find more information about how QuerySets are design in the definition files
 (3) Schedule the querys: Initiatie the cron manager and submit the cron jobs.
 (4) Start HTTP Server: Start teh server and expose endpoints to understand the health of submitted jobs and to change and modify the cron jobs.
*/

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	/*****************************************************************************************************
	Step 0: Configuration
	******************************************************************************************************/
	readConfig := types.ReadConfig{}
	osEnv := core.OsEnv{}
	cfg, err := readConfig.Read(osEnv)
	params := core.Conf{}
	params.GetConf("./collector.yaml")

	banner(cfg)

	if err != nil {
		log.Fatalf("Error reading config: %s", err.Error())
	}

	cfg.TCPPort = params.Metadata.Port
	ctx := core.Context{Params: &params, Config: cfg, Result: make(map[string]string), ResultMutex: sync.RWMutex{}}

	/*****************************************************************************************************
	Step 1: Create all the resources and their sockets
	******************************************************************************************************/

	base := core.MakeResource[types.BaseResourceSchema, types.BaseResource]("baseresource", cfg, params, types.BaseResourceSchema{}, &types.BaseResource{})
	base.Initialize()

	// base2 := core.MakeResource[types.BaseResourceSchema, types.BaseResource]("baseresource", cfg, params, types.BaseResourceSchema{}, &types.BaseResource{})
	// base2.Initialize()

	/*****************************************************************************************************
	Step 2: Create all the resource queries that are to be schedule and their corresponsding Jobs
	******************************************************************************************************/
	/* This is a stub resource mostly use to write to local file system */
	baseQueryset := core.QuerySet{{Type: "Output", Command: `ls`}}
	baseJob := core.Job[types.BaseResourceSchema, types.BaseResource, types.BaseResourceSchema, types.BaseResource]{
		In:       base,
		Out:      base,
		Context:  &ctx,
		QuerySet: baseQueryset[:],
	}

	scamper := core.MakeResource[types.ScamperResourceSchema, types.ScamperResource]("scamperresource", cfg, params, types.ScamperResourceSchema{}, &types.ScamperResource{})
	scamper.Initialize()

	querySet := core.QuerySet{{Type: "Output", Command: `ping 8.8.8.8 8.8.4.4`}}
	job_1 := core.Job[types.ScamperResourceSchema, types.ScamperResource, types.ScamperResourceSchema, types.ScamperResource]{
		In:       scamper,
		Out:      scamper,
		Context:  &ctx,
		QuerySet: querySet[:],
	}

	job_1.Execute()

	/* This create the queryset for changelog and passess that conditions that we want met for issuing queries. */
	// changelogQueryset := core.QuerySet{
	// 	core.Query{Type: "Output",
	// 		Condition: func(ctx *core.Context, i interface{}) bool {
	// 			schema, ok := i.(models.ChangelogSchema)
	// 			if ok {
	// 				length := len(schema.RecordList.Records)
	// 				if length == 0 {
	// 					return false
	// 				}
	// 				time, err := time.Parse(ctx.Params.Changelog.Format, ctx.Params.Changelog.GreaterThanEqual)
	// 				if err != nil {
	// 					fmt.Println("Failed")
	// 					return false
	// 				}
	// 				if time.Compare(schema.RecordList.Records[length-1].Source.Timestamp) == -1 {
	// 					ctx.Params.Changelog.LessThan = schema.RecordList.Records[length-1].Source.Timestamp.Format(ctx.Params.Changelog.Format)
	// 					log.WithFields(log.Fields{"service": "changelog", "NextWindow": schema.RecordList.Records[length-1].Source.Timestamp.Format(ctx.Params.Changelog.Format)}).Info("Query condition Date=" + ctx.Params.Changelog.GreaterThanEqual + " has not been met.\n\tIssuing the next window starting at \033[0;33m==>\033[0m")
	// 					return true
	// 				} else {
	// 					log.WithFields(log.Fields{"service": "changelog", "NextWindow": schema.RecordList.Records[length-1].Source.Timestamp.Format(ctx.Params.Changelog.Format)}).Info("The query condition Date=\033[0;33m" + ctx.Params.Changelog.GreaterThanEqual + "\033[0m has been met! Collecting results")
	// 					return false
	// 				}
	// 			}
	// 			return false
	// 		},
	// 		Command: `cloudflared access curl {{ .Params.Changelog.Endpoint }} -H "Content-Type: application/json" --request POST --data-binary '{"from": 0, "size": {{.Params.Changelog.Size }}, "sort": [{"@timestamp": {"order": "desc"}}], "query": {"bool": {"must": [{"query_string": {"query": "{{.Params.Changelog.Service}}", "default_field": "event_source"}}, {"range": {"@timestamp": {"gte": "{{.Params.Changelog.GreaterThanEqual}}", "lt": "{{.Params.Changelog.LessThan}}"}}}]}}}'`},
	// }

	baseJob.Execute()

	/*****************************************************************************************************
	Step 3: Create the Cron and schedule Jobs
	******************************************************************************************************/
	//c := cron.New()

	//  "*/30 * * * *" : Schedule the login run every 30 minutes*/
	//c.AddFunc("*/30 * * * *", func() {
	//	baseJob.Execute()
	//	log.WithFields(log.Fields{"service": baseJob.In.Name, "command": baseJob.QuerySet[0]}).Info("Executed a Query task.")
	//})

	//  "* * * * *" :  Schedule the changelogs data extraction to run every minute */
	//c.AddFunc("*/5 * * * *", func() {
	//	changelogsJob.Execute()
	//	//fmt.Println(changelogsJob.Context.Result)
	//	log.WithFields(log.Fields{"service": changelogsJob.In.Name, "command": changelogsJob.QuerySet[0]}).Info("Executed a Query task.")
	//})

	//  "* * * * *" :  Schedule the changelogs data extraction to run every minute*/
	//c.AddFunc("* * * * *", func() {
	//	clickhouseJob.Execute()
	//fmt.Println(clickhouseJob.Context.Result)
	//	log.WithFields(log.Fields{"service": clickhouseJob.In.Name, "command": clickhouseJob.QuerySet[1]}).Info("Executed a Query task.")
	//})

	//c.Start()

	//jsonFile, err := os.ReadFile("/Users/innocent/internship/dome/research-design/notebooks/data/test.json")
	// if we os.Open returns an error then handle it
	//if err != nil {
	//	fmt.Println(err)
	//}
	//set := types.ChangelogSchema{}
	//json.Unmarshal(jsonFile, &set)
	//fmt.Println(set)
	//test := fmt.Sprintf("%q", response.RecordList.Records[0].Source.Message)

	//l := lexer.New(strings.ToLower(test))
	//p := parser.New(l)
	//program := p.ParseProgram()

	/*set := types.Message{}

	//fmt.Printf("1. [%s] ", program.Statements[0].TokenLiteral())
	fmt.Printf("2. [%s] | ", program.Statements[1].TokenLiteral())
	for _, stmt := range program.Statements[2].(*ast.BlockStatement).Statements {
		row := types.MessageEntry{
			Dst:    stmt.TokenLiteral(),
			Action: program.Statements[1].TokenLiteral(),
			Plans:  map[string]float64{},
		}
		tmp := stmt.(*ast.AssignmentStatement)
		for _, assign := range tmp.Statements {
			//fmt.Printf("4. [%s]", assign.TokenLiteral())
			exprTemp := assign.(*ast.ExpressionStatement)
			value, err := strconv.ParseFloat(strings.Replace(exprTemp.Expression.(*ast.AttributeLiteral).String(), "%", "", -1), 32)
			if err != nil {
				break
			}
			row.AddPlanShare(assign.TokenLiteral(), value)
		}
		set.AddRow(row)
	}

	fmt.Println(set)*/
	//checkParserErrors(t, p)

	//j := clickhouseJob.Context.Result[clickhouseJob.In.Name]
	/*****************************************************************************************************
	Step 4: Configure and start the HTTP Server
	******************************************************************************************************/

	/*r := mux.NewRouter()
	r.HandleFunc("/test", handlers.Test)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.TCPPort),
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		Handler:        r,
	}

	log.Fatal(s.ListenAndServe())*/
}

func banner(cfg *core.Config) {
	myFigure := figure.NewColorFigure("Collector", "univers", "purple", true)
	myFigure.Print()

	fmt.Println("\n*****************************************************")
	fmt.Printf(" 🦉 PATIO DeX Collector")
	fmt.Println("\n*****************************************************")
	fmt.Printf("\n\t\033[0;33mVersion:\033[0m  \033[1m%s\033[0m\t\t\t\t\033[0;33mMode:\033[0m  %s\n"+
		"\t\033[0;33mAuthor:\033[0m  %s\t"+
		"\t\033[0;33mUser:\033[0m  %s\n\n"+
		"This is a DeX collector that schedules and executes queries to two Resources:\n"+
		"\t1. Clickhouse (Challenge Page)\n"+
		"\t2. Changelogs (Traffic Manager)\n"+
		"\t3. Cloudflare Access : logins to cf_access in order to allow issuance of new token\n", "0.0.1", "Dev", "Innocent Obi Jr", cfg.Secrets.User)
	fmt.Println("*****************************************************")
}
