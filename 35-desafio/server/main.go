package main

import (
	"context"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type UsdBrlResponse struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type Quotation struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Bid       string `json:"bid"`
	Timestamp string `json:"timestamp"`
	gorm.Model
}

type QuotationDTO struct {
	Bid string `json:"bid"`
}

func getQuotation() (error, *UsdBrlResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return err, nil
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err, nil
	}

	defer res.Body.Close()
	var response UsdBrlResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err, nil
	}

	return nil, &response
}

func QuotationHandler(w http.ResponseWriter, r *http.Request) {
	// get quotation
	err, quotationResponse := getQuotation()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	// create quotation
	db := r.Context().Value("DB").(*gorm.DB)
	persistQuotation(db, quotationResponse)

	// response part
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// return response enconded in writer
	json.NewEncoder(w).Encode(&QuotationDTO{
		Bid: quotationResponse.Usdbrl.Bid,
	})
}

func persistQuotation(db *gorm.DB, quotationResponse *UsdBrlResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	db.WithContext(ctx).Create(&Quotation{
		Name:      quotationResponse.Usdbrl.Name,
		Code:      quotationResponse.Usdbrl.Code,
		Bid:       quotationResponse.Usdbrl.Bid,
		Timestamp: quotationResponse.Usdbrl.Timestamp,
	})
}

func setHttpMiddleware(next http.Handler, db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "DB", db)))
	})
}

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Quotation{})
	if err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	mux := http.NewServeMux()
	quotationHandler := http.HandlerFunc(QuotationHandler)
	mux.Handle("/cotacao", setHttpMiddleware(quotationHandler, db))
	log.Println(http.ListenAndServe(":8080", mux))
}
