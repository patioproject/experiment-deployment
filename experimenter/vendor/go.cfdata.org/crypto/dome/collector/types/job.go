package types

import (
	"fmt"
)

type Job[A Schema[A], B any, C Schema[C], D any] struct {
	In            Resource[A, B]
	Out           Resource[C, D]
	Context       *Context
	QuerySet      []Query
	SuccessHandle func(string) string
	ErrorHandle   func(error) string
}

func (j *Job[A, B, C, D]) SetQuery() *Job[A, B, C, D] {
	return j
}

func (j *Job[A, B, C, D]) SetCommands(query Query) *Job[A, B, C, D] {
	j.QuerySet = append(j.QuerySet, query)
	return j
}

func (j *Job[A, B, C, D]) SetSuccessHandle() *Job[A, B, C, D] {
	return j
}

func (j *Job[A, B, C, D]) SetErrorHandle() *Job[A, B, C, D] {
	return j
}

func (j *Job[A, B, C, D]) Execute() (any, error) {
	querysetLength := len(j.QuerySet)
	switch querysetLength {
	case 0:
		return nil, fmt.Errorf("CommandSet has len 0")
	default:
		for idx, query := range j.QuerySet {
			template := MakeQueryTemplate(j.In.Name, query.Command)
			result := template.Bind(j.Context)

			/* This doesnt pass through the command error. NEed to parse this string */
			res, err := j.In.Handle.Read(result, j.Context)(query)
			if err != nil {
				return nil, fmt.Errorf("collector[%s]: Query #%d failed to execute", j.In.Name, idx)
			} else {
				j.Context.ResultMutex.Lock()
				j.Context.Result[j.In.Name] = string(res)
				j.Context.ResultMutex.Unlock()
			}
		}
		j.Out.Handle.Write(j.In.Name, j.Context)
		return j.Context.Result, nil
	}
}

type Executable[A any] interface {
	*A
	Execute() (any, error)
}
