package log_entry

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"logger-service/config"
	jh "logger-service/pkg/json_helper"
	"net/http"
	"sync"
)

type Handler struct {
	LogEntryService Service
	WS              *websocket.Conn
	WSMutex         sync.Mutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	err := jh.ReadJSON(w, r, &requestPayload)

	if err != nil {
		jh.ErrorJSON(w, err)
		return
	}

	event := LogEntry{
		Name:      requestPayload.Name,
		Type:      requestPayload.Type,
		Stamp:     requestPayload.Stamp,
		Signature: requestPayload.Signature,
		ProfileID: requestPayload.ProfileID,
		KeyID:     requestPayload.KeyID,
		Data:      requestPayload.Data,
	}

	err = h.LogEntryService.InsertLogEntry(event)
	if err != nil {
		jh.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = jh.WriteJSON(w, http.StatusOK, event)

	go func() {
		h.WSMutex.Lock()
		defer h.WSMutex.Unlock()
		if h.WS != nil {
			jsonObj := make(map[string]int)

			for _, c := range config.Counts {
				count := h.LogEntryService.GetCountService(c)
				if count > 0 {
					jsonObj[c] = h.LogEntryService.GetCountService(c)
				}
			}

			jsonData, err := json.Marshal(jsonObj)

			if err != nil {
				fmt.Printf("WriteLog :: JSON Marshal error: %v", err.Error())
				return
			}

			fmt.Printf("\njsonData: %v\n", jsonData)

			if err := h.WS.WriteMessage(websocket.TextMessage, jsonData); err != nil {
				fmt.Printf("WriteLog :: WS:: error: %v", err.Error())
				return
			}
		}
	}()
}

func (h *Handler) WSHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	h.WS, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("WSHandler :: WS:: error: %v", err.Error())
		return
	}
}
