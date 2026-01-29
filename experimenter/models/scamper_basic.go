package models

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	core "go.cfdata.org/crypto/dome/collector/types"
)

type ScamperResource struct {
	Name string
	Data ScamperResourceSchema
	Type string
}

func (sr *ScamperResource) Init(cfg *core.Config) (bool, error) {
	sr.Data = ScamperResourceSchema{}
	sr.Type = "ping"
	return true, nil
}

type ScamperResourceSchema struct {
}

func (sr ScamperResourceSchema) Unit() ScamperResourceSchema {
	return ScamperResourceSchema{}
}

func (sr *ScamperResource) Read(res string, ctx *core.Context) func(query core.Query) ([]byte, error) {
	return func(query core.Query) ([]byte, error) {
		queryparts := strings.Split(res, " ")
		if queryparts[0] != "ping" {
			return nil, fmt.Errorf("commands other than ping not currently supported")
		}
		command := "scamper -c " + queryparts[0] + " -i "
		for _, part := range queryparts[1:] {
			if net.ParseIP(part) == nil {
				return nil, fmt.Errorf("not a valid IPV4 or IPV6 address")
			} else {
				command = command + part + " "
			}
		}
		cmd := core.CreateCommand(command)
		return core.DoCommandType(query.Type, cmd)()
	}
}

func (sr *ScamperResource) Write(filename string, ctx *core.Context) (bool, error) {
	result := ctx.Result[filename]
	data := []byte(result)
	now := time.Now().UnixMilli()
	fmt.Println(result)

	//log.Errorf("%s", string(data))
	f, err := os.Create(fmt.Sprintf("%s/%s-%d.txt", ctx.Params.Metadata.ExportDataPath, filename, now))
	check(err)
	w := bufio.NewWriter(f)
	bytesWritten, err := w.Write(data)
	check(err)
	fmt.Printf("wrote %d bytes\n", bytesWritten)
	check(w.Flush())
	return true, nil
}
