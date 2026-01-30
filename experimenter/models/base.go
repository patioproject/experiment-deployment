package models

import (
	"bufio"
	"fmt"
	"os"
	"time"

	core "go.cfdata.org/crypto/dome/collector/types"
)

type BaseResource struct {
	Name string
	Data BaseResourceSchema
}

func (br *BaseResource) Init(cfg *core.Config) (bool, error) {
	br.Data = BaseResourceSchema{}
	return true, nil
}

type BaseResourceSchema struct {
}

func (br BaseResourceSchema) Unit() BaseResourceSchema {
	return BaseResourceSchema{}
}

func (bss *BaseResource) Read(res string, ctx *core.Context) func(query core.Query) ([]byte, error) {
	return func(query core.Query) ([]byte, error) {
		cmd := core.CreateCommand(res)
		return core.DoCommandType(query.Type, cmd)()
	}
}

func (bss *BaseResource) Write(filename string, ctx *core.Context) (bool, error) {
	result := ctx.Result[filename]
	data := []byte(result)
	now := time.Now().UnixMilli()
	fmt.Println(result)

	//log.Errorf("%s", string(data))
	f, err := os.Create(fmt.Sprintf("%s/%s-%d.json", ctx.Params.Metadata.ExportDataPath, filename, now))
	check(err)
	w := bufio.NewWriter(f)
	bytesWritten, err := w.Write(data)
	check(err)
	fmt.Printf("wrote %d bytes\n", bytesWritten)
	check(w.Flush())
	return true, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
