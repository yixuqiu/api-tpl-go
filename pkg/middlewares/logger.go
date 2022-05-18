package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"

	"tplgo/pkg/logger"
	"tplgo/pkg/result"
)

var (
	bufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 2<<10)) // 2KB
		},
	}
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Local()

		var body []byte

		// 取出请求Body
		if r.Body != nil && r.Body != http.NoBody {
			var err error

			body, err = ioutil.ReadAll(r.Body)

			if err != nil {
				result.ErrSystem(result.Err(err)).JSON(w, r)

				return
			}

			// 关闭原Body
			r.Body.Close()

			r.Body = ioutil.NopCloser(bytes.NewReader(body))
		}

		// 存储返回结果
		buf := bufPool.Get().(*bytes.Buffer)
		buf.Reset()

		defer bufPool.Put(buf)

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		ww.Tee(buf)

		next.ServeHTTP(ww, r)

		logger.Info(r.Context(), fmt.Sprintf("[%s] %s", r.Method, r.URL.String()),
			zap.ByteString("params", pretty.Ugly(body)),
			zap.String("response", buf.String()),
			zap.Int("status", ww.Status()),
			zap.String("duration", time.Since(now).String()),
		)
	})
}