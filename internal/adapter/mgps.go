package adapter

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"golang-base-structure/config"
	"golang-base-structure/internal/adapter/model"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

// MGPSAdapter represent for api call to MGPS
type (
	MGPSAdapter interface {
		GetInformation(ctx context.Context, transactionID, data, description string) error
		PostData(ctx context.Context, transactionID, data, description string) error
	}

	mgpsAdapter struct {
		cfg *config.Config
	}
)

// NewMGPSAdapter creates an instance
func NewMGPSAdapter(
	cfg *config.Config,
) MGPSAdapter {
	return &mgpsAdapter{
		cfg: cfg,
	}
}

func (a *mgpsAdapter) GetInformation(ctx context.Context, transactionID, data, description string) (err error) {
	return nil
}

func (a *mgpsAdapter) PostData(ctx context.Context, transactionID, data, description string) (err error) {
	return nil
}

func (a *mgpsAdapter) Post(ctx context.Context, path string, reqMsg *model.MGPSRequestData) (resMsg *model.MGPSResponseData, err error) {
	resMsg = &model.MGPSResponseData{}

	bufferReq, err := xml.Marshal(&reqMsg)
	if err != nil {
		zap.S().Errorw("Failed to marshal request", zap.Error(err))
		return nil, err
	}
	url := fmt.Sprintf("%v%v", a.cfg.MGPSURL, path)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(bufferReq))
	if err != nil {
		zap.S().Errorw("Failed to create request", zap.Error(err))
		return nil, err
	}
	zap.S().Infow(fmt.Sprintf("%v - Request to MGPS: %v", reqMsg.TransactionID, bytes.NewBuffer(bufferReq)))
	httpReq.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		zap.S().Errorw("Failed to do request", zap.Error(err))
		return nil, err
	}
	if httpRes.StatusCode == 200 {
		bufferRes, err := ioutil.ReadAll(httpRes.Body)

		if err != nil {
			zap.S().Errorw("Failed to read response", zap.Error(err))
			return nil, err
		}
		zap.S().Infow(fmt.Sprintf("%v - Response from MGPS: %v", reqMsg.TransactionID, bytes.NewBuffer(bufferRes)))
		defer httpRes.Body.Close()
		err = xml.Unmarshal(bufferRes, &resMsg)
		if err != nil {
			zap.S().Errorw("Failed to unmarshal response", zap.Error(err))
			return nil, err
		}
		return resMsg, nil
	}
	return nil, errors.New(httpRes.Status)
}

func (a *mgpsAdapter) Get(ctx context.Context, transactionID, path string) (resMsg *model.MGPSResponseData, err error) {
	resMsg = &model.MGPSResponseData{}

	if err != nil {
		zap.S().Errorw("Failed to marshal request", zap.Error(err))
		return nil, err
	}
	url := fmt.Sprintf("%v%v", a.cfg.MGPSURL, path)
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zap.S().Errorw("Failed to create request", zap.Error(err))
		return nil, err
	}
	zap.S().Infow(fmt.Sprintf("%v - GET Request to MGPS", transactionID))
	httpReq.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		zap.S().Errorw("Failed to do request", zap.Error(err))
		return nil, err
	}
	if httpRes.StatusCode == 200 {
		bufferRes, err := ioutil.ReadAll(httpRes.Body)

		if err != nil {
			zap.S().Errorw("Failed to read response", zap.Error(err))
			return nil, err
		}
		zap.S().Infow(fmt.Sprintf("%v - Response from MGPS: %v", transactionID, bytes.NewBuffer(bufferRes)))
		defer httpRes.Body.Close()
		err = xml.Unmarshal(bufferRes, &resMsg)
		if err != nil {
			zap.S().Errorw("Failed to unmarshal response", zap.Error(err))
			return nil, err
		}
		return resMsg, nil
	}
	return nil, errors.New(httpRes.Status)
}
