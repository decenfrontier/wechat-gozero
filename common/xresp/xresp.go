package xresp

import (
	"net/http"

	"github.com/wslynn/wechat-gozero/common/xerr"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"

	// en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

var (
	Validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// handler模板会调用该函数
func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err == nil {
		// 若没有错误产生
		body.Code = xerr.OK
		body.Msg = "OK"
		body.Data = resp
	} else {
		// 若有错误
		causeErr := errors.Cause(err)
		switch v := causeErr.(type) {
		case *xerr.CodeError:
			// 自定义错误
			body.Code = v.GetCode()
			body.Msg = v.GetMsg()
		case validator.ValidationErrors:
			// 参数校验错误
			body.Code = xerr.PARAM_ERROR
			body.Msg = v[0].Translate(trans)
		default:
			gstatus, ok := status.FromError(v)
			if ok { // grpc错误
				body.Code = uint32(gstatus.Code())
				body.Msg = gstatus.Message()
			} else { // 未知错误
				body.Code = xerr.SERVER_ERROR
				body.Msg = "unknown"
			}
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】: %+v ", err)
	}
	httpx.OkJson(w, body)
}

func init() {
	zhT := zh.New() //chinese
	enT := en.New() //english
	uni := ut.New(enT, zhT, enT)

	Validate = validator.New()
	trans, _ = uni.GetTranslator("zh")
	zh_translation.RegisterDefaultTranslations(Validate, trans)
}
