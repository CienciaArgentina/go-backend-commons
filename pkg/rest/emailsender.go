package rest

import (
	"fmt"
	"github.com/CienciaArgentina/go-email-sender/commons"
	"github.com/CienciaArgentina/go-enigma/config"
	 "github.com/go-resty/resty/v2"
)

func EmailSenderApiCall(cfg *config.Microservices, dto *commons.DTO) (bool, error, string) {
	c := resty.New()

	res, err := c.EnableTrace().
		R().
		SetAuthScheme(cfg.Scheme).
		SetBody(dto).Post(fmt.Sprintf("%s%s", cfg.BaseUrl, cfg.EmailSenderEndpoints.SendEmail))


	if err != nil || res.IsError() {
		return false, err, res.String()
	}

	return true, nil, ""
}
