package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/tidwall/pretty"
	"go.uber.org/zap"

	"api/consts"
	libhttp "api/lib/http"
	"api/lib/util"
	"api/logger"
	"api/pkg/auth"
	"api/pkg/result"
)

// Log 日志中间件
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		body := "<nil>"

		// 请求包含body
		if r.Body != nil && r.Body != http.NoBody {
			switch util.ContentType(r) {
			case libhttp.ContentForm:
				if err := r.ParseForm(); err != nil {
					result.ErrSystem(result.M(err.Error())).JSON(w, r)
					return
				}

				body = r.Form.Encode()
			case libhttp.MultipartForm:
				if err := r.ParseMultipartForm(consts.MaxFormMemory); err != nil {
					if err != http.ErrNotMultipart {
						result.ErrSystem(result.M(err.Error())).JSON(w, r)
						return
					}
				}

				body = r.Form.Encode()
			case libhttp.ContentJSON:
				// 取出Body
				b, err := io.ReadAll(r.Body)
				if err != nil {
					result.ErrSystem(result.M(err.Error())).JSON(w, r)
					return
				}
				// 关闭原Body
				r.Body.Close()

				body = string(pretty.Ugly(b))

				// 重新赋值Body
				r.Body = io.NopCloser(bytes.NewReader(b))
			}
		}

		next.ServeHTTP(w, r)

		logger.Info(r.Context(), "request info",
			zap.String("method", r.Method),
			zap.String("uri", r.URL.String()),
			zap.String("ip", r.RemoteAddr),
			zap.String("body", body),
			zap.String("identity", auth.GetIdentity(r.Context()).String()),
			zap.String("duration", time.Since(now).String()),
		)
	})
}