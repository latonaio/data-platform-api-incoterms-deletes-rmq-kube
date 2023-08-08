package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-incoterms-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-incoterms-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) Incoterms(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Incoterms {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE incoterms.Incoterms = \"%s\ ", input.Incoterms.Incoterms),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	incoterms.Incoterms
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_incoterms_incoterms_data as incoterms 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToIncoterms(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
