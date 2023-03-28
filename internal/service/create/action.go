package create

import (
	"context"
	"errors"
	es_cli "es_service/internal/es"
	"es_service/internal/models"
	"es_service/internal/service/base"
	es_service "github.com/solost23/protopb/gen/go/protos/es"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(ctx context.Context, request *es_service.CreateRequest) (reply *es_service.CreateResponse, err error) {
	index := request.GetIndex()
	documentId := request.GetDocumentId()
	body := request.GetDocument()

	if index == "" || documentId == "" || body == "" {
		return nil, errors.New("任一参数都不可为空")
	}
	createResult, err := es_cli.CreateDocument(ctx, a.GetES(), index, documentId, body)
	if err != nil {
		return nil, err
	}

	// 插入数据
	err = (&models.ESRecord{
		CreatorBase: models.CreatorBase{
			CreatorId: uint(request.GetHeader().GetOperatorUid()),
		},
		DocumentID: createResult.Id,
		Document:   body,
		Index:      createResult.Index,
		Type:       createResult.Type,
		Result:     createResult.Result,
		Status:     createResult.Status,
	}).Insert(a.GetMysqlConnect())
	return reply, nil
}
