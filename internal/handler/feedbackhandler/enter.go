package feedbackhandler

import (
	"JHE_admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(

		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/message/getMessage",
				Handler: MessageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/message/getNewMessage",
				Handler: NewMessageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/message/readNewMessage",
				Handler: ReadNewMessageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/message/dealMessage",
				Handler: DealNewMessageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/message/deleteMessage",
				Handler: DeleteMessageHandler(serverCtx),
			},
		},
	)
}
