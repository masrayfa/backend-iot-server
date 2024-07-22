package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/repository"
	"github.com/masrayfa/internals/service"
)

type ChannelControllerImpl struct {
	db *pgxpool.Pool
	channelService service.ChannelService
	channelRepository repository.ChannelRepository
}

func NewChannelController(channelService service.ChannelService, channelRepository repository.ChannelRepository, db *pgxpool.Pool) ChannelController {
	return &ChannelControllerImpl{
		db: db,
		channelService: channelService,
		channelRepository: channelRepository,
	}
}

func (controller *ChannelControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	channelRequest := web.ChannelCreateRequest{}
	helper.ReadFromRequestBody(request, &channelRequest)

	log.Println("channelRequest: ", channelRequest)

	_, err := controller.channelService.Create(request.Context(), channelRequest)
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ChannelControllerImpl) DownloadCSV(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	limitStr := request.URL.Query().Get("limit")
	limit := int64(0)
	if limitStr != "" {

		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			webErrResponse := web.WebErrResponse{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Mesage: err.Error(),
			}

			helper.WriteToResponseBody(writer, webErrResponse)
			return
		}
	}

	startDateStr := request.URL.Query().Get("start")
	var startDate *time.Time 
	if startDateStr != "" {
		start, err := time.Parse("2006-01-02 15:04:05", startDateStr)
		if err != nil {
			webErrResponse := web.WebErrResponse{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Mesage: err.Error(),
			}

			helper.WriteToResponseBody(writer, webErrResponse)
			return
		}
		startDate = &start
	}

	endDateStr := request.URL.Query().Get("end")
	var endDate *time.Time
	
	if endDateStr != "" {
		end, err := time.Parse("2006-01-02 15:04:05", endDateStr)
		if err != nil {
			webErrResponse := web.WebErrResponse{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Mesage: err.Error(),
			}

			helper.WriteToResponseBody(writer, webErrResponse)
			return
		}

		endDate = &end
	}

	
	feed, err := controller.channelRepository.GetNodeChannelCSV(request.Context(), controller.db, id, limit, startDate, endDate)
	if len(feed) == 0 {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: "No data found",
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}
	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}
		
		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	header := []string{"time"}
	for i := range feed[0].Value {
		header = append(header, fmt.Sprintf("value %d", i+1))
	}

	var records []map[string]interface{}
	for _, d := range feed {
		record := make(map[string]interface{})
		record["time"] = d.Time
		for i, v := range d.Value {
			record[fmt.Sprintf("value %d", i+1)] = v
		}
		records = append(records, record)
	}

	if err != nil {
		webErrResponse := web.WebErrResponse{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Mesage: err.Error(),
		}

		helper.WriteToResponseBody(writer, webErrResponse)
		return
	}

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: records,
	}

	helper.WriteToResponseBody(writer, webResponse)
	
	// generate csv version
	// filePath, err := helper.GenerateCSV(feed)
	// if err != nil {
	// 	http.Error(writer, "Could not generate CSV", http.StatusInternalServerError)
	// 	return
	// }
	// defer os.Remove(filePath)

	// writer.Header().Set("Content-Disposition", "attachment; filename=records.csv")
	// writer.Header().Set("Content-Type", "text/csv")
	// http.ServeFile(writer, request, filePath)
}

// type Record struct {
// 	Name  string
// 	Email string
// 	Phone string
// }

// func generateCSV(records []Record) (string, error) {
// 	file, err := os.CreateTemp("", "records-*.csv")
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// Write header
// 	writer.Write([]string{"Name", "Email", "Phone"})

// 	// Write records
// 	for _, record := range records {
// 		writer.Write([]string{record.Name, record.Email, record.Phone})
// 	}

// 	return file.Name(), nil
// }

// func downloadCSVHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	records := []Record{
// 		{"John Doe", "john@example.com", "123-456-7890"},
// 		{"Jane Smith", "jane@example.com", "987-654-3210"},
// 	}

// 	filePath, err := generateCSV(records)
// 	if err != nil {
// 		http.Error(w, "Could not generate CSV", http.StatusInternalServerError)
// 		return
// 	}
// 	defer os.Remove(filePath)

// 	w.Header().Set("Content-Disposition", "attachment; filename=records.csv")
// 	w.Header().Set("Content-Type", "text/csv")
// 	http.ServeFile(w, r, filePath)
// }