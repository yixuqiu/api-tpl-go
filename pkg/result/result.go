package result

import (
	"api/lib/util"
	"api/logger"

	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

const CodeOK = 0

// Result API结果
type Result interface {
	JSON(w http.ResponseWriter, r *http.Request)
}

type response struct {
	x util.X
}

func (resp *response) JSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp.x["req_id"] = util.GetReqID(ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp.x); err != nil {
		logger.Err(ctx, "err write response", zap.Error(err))
	}
}

// Option API结果选项
type Option func(r *response)

// M 指定返回的msg
func M(m string) Option {
	return func(r *response) {
		r.x["msg"] = m
	}
}

// V 指定返回的data
func V(v any) Option {
	return func(r *response) {
		r.x["data"] = v
	}
}

// KV 指定返回的自定义key-value
func KV(k string, v any) Option {
	return func(r *response) {
		r.x[k] = v
	}
}

// New 返回一个Result
func New(code int, msg string, options ...Option) Result {
	resp := &response{
		x: util.X{
			"code": code,
			"err":  false,
			"msg":  msg,
		},
	}

	if code != CodeOK {
		resp.x["err"] = true
	}

	for _, f := range options {
		f(resp)
	}

	return resp
}
