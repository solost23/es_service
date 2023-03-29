package delete

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

func (a *Action) Deal(ctx context.Context, request *es_service.DeleteRequest) (reply *es_service.DeleteResponse, err error) {
	index := request.GetIndex()
	documentId := request.GetDocumentId()

	if index == "" || documentId == "" {
		return nil, errors.New("任一参数都不可为空")
	}
	deleteResult, err := es_cli.DeleteDocument(ctx, a.GetES(), index, documentId)
	if err != nil {
		return nil, err
	}

	// 插入数据
	err = (&models.ESRecord{
		CreatorBase: models.CreatorBase{
			CreatorId: uint(request.GetHeader().GetOperatorUid()),
		},
		DocumentID: deleteResult.Id,
		Document:   "",
		Index:      deleteResult.Index,
		Type:       deleteResult.Type,
		Result:     deleteResult.Result,
		Status:     deleteResult.Status,
	}).Insert(a.GetMysqlConnect())
	if err != nil {
		a.GetSl().Errorf("Document Delete Record Create Failed: %s \n", deleteResult.Id)
	}
	return
}
