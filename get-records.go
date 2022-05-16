package nicrudns

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func (client *Client) GetRecords(zoneName string) ([]*RR, error) {
	url := fmt.Sprintf(GetRecordsUrlPattern, client.provider.DnsServiceName, zoneName)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, RequestError.Error())
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, ResponseError.Error())
	}

	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(response.Body); err != nil {
		return nil, errors.Wrap(err, BufferReadError.Error())
	}

	apiResponse := &Response{}
	if err := xml.NewDecoder(buf).Decode(&apiResponse); err != nil {
		return nil, errors.Wrap(err, XmlDecodeError.Error())
	}
	if apiResponse.Status != SuccessStatus {
		return nil, errors.Wrap(ApiNonSuccessError, describeError(apiResponse.Errors.Error))
	} else {
		var records []*RR
		for _, zone := range apiResponse.Data.Zone {
			if zone.Name != zoneName {
				continue
			}
			records = append(records, zone.Rr...)
		}
		return records, nil
	}

}
