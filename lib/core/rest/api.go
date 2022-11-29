package rest

import (
	"bot/lib/core/discord"
	"bot/lib/core/discord/verification"
	"bot/lib/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Resp struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

type CreateTaskPayload struct {
	TaskID string `json:"task_id"`
}

type SubmitCaptchaPayload struct {
	Response string `json:"response"`
	TaskID   string `json:"task_id"`
}

func VerificationResult(w http.ResponseWriter, r *http.Request) {
	IP := r.Header.Get("X-Forwarded-For")

	fmt.Println(IP)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func VerificationStart(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "challenge")
	http.ServeFile(w, r, filepath.Join(filesDir, "index.html"))
}

func SendJsFile(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "challenge")

	http.ServeFile(w, r, filepath.Join(filesDir, "challenge.js"))
}

func SubmitResult(w http.ResponseWriter, r *http.Request) {
	var p SubmitCaptchaPayload
	Err := json.NewDecoder(r.Body).Decode(&p)

	if utils.HandleError(Err) {
		json.NewEncoder(w).Encode(Resp{
			Error:  "invalid payload",
			Status: "",
		})

		return
	}

	if _, ok := verification.TaskList[p.TaskID]; !ok {
		json.NewEncoder(w).Encode(Resp{
			Error:  "invalid task",
			Status: "",
		})

		return
	}

	w.Header().Add("content/type", "application/json")

	if p.Response == verification.TaskList[p.TaskID].TaskChallenge.OutResponse {
		json.NewEncoder(w).Encode(map[string]string{
			"error":  "",
			"status": "success",
		})

		verification.TaskList[p.TaskID].EndTask(discord.BotSession)
	} else {
		json.NewEncoder(w).Encode(map[string]string{
			"error":  "",
			"status": "invalid",
		})
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var p CreateTaskPayload
	Err := json.NewDecoder(r.Body).Decode(&p)

	if utils.HandleError(Err) {
		json.NewEncoder(w).Encode(Resp{
			Error:  "invalid payload",
			Status: "",
		})

		return
	}

	if _, ok := verification.TaskList[p.TaskID]; !ok {
		json.NewEncoder(w).Encode(Resp{
			Error:  "invalid task",
			Status: "",
		})

		return
	}

	w.Header().Add("content/type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"error": "",
		"token": verification.TaskList[p.TaskID].TaskChallenge.InToken,
	})
}

func HandleRequests() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))

	//r.Get("/result", VerificationResult)
	r.Get("/verify", VerificationStart)
	r.Get("/challenge.js", SendJsFile)

	r.Post("/submit_captcha", SubmitResult)
	r.Post("/create_task", CreateTask)

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		workDir, _ := os.Getwd()
		filesDir := filepath.Join(workDir, "challenge/asset")
		http.ServeFile(w, r, filesDir+r.URL.Path)
	})

	http.ListenAndServe(":80", r)
}
