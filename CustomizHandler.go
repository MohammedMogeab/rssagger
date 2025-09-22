package main

import (
    "net/http"
    "context"
)

// HandlerHealthz returns 200 and includes a DB check
func (apicfing *apiconfig) HandlerHealthz(w http.ResponseWriter, r *http.Request) {
    // Best-effort ping; if it fails, return 500 with error
    if err := apicfing.dbConn.PingContext(context.Background()); err != nil {
        respondWithError(w, http.StatusInternalServerError, "database unreachable")
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"status":"ok"})
}
