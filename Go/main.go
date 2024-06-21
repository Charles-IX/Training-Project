package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Q struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Type     string `json:"type"`
	// 1：单选；2:判断；3：多选
	Chapter string `json:"chapter"`
	Key     string `json:"key"`
	Answer  string `json:"answer"`
	Times   string `json:"times"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "charles"
	dbPass := "1145141919"
	dbName := "DrivTest"
	dbHost := "127.0.0.1"
	dbPort := "3306"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("成功连接到MySQL题目服务器。")
	return db
}

func GetAllQuestions() ([]Q, error) {
	db := dbConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM Questions")
	if err != nil {
		return nil, err
	}
	defer selDB.Close()

	var questions []Q
	for selDB.Next() {
		var q Q
		err = selDB.Scan(&q.ID, &q.Question, &q.Type, &q.Chapter, &q.Key, &q.Answer, &q.Times)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func GetWrongQuestions() ([]Q, error) {
	db := dbConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM Questions WHERE Times > 0")
	if err != nil {
		return nil, err
	}
	defer selDB.Close()

	var questions []Q
	for selDB.Next() {
		var q Q
		err = selDB.Scan(&q.ID, &q.Question, &q.Type, &q.Chapter, &q.Key, &q.Answer, &q.Times)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func HandleWrongQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("接收到GET请求")
	questions, err := GetWrongQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseBytes, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}

func GetRandomQuestions() ([]Q, error) {
	db := dbConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM Questions ORDER BY RAND() LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer selDB.Close()

	var questions []Q
	for selDB.Next() {
		var q Q
		err = selDB.Scan(&q.ID, &q.Question, &q.Type, &q.Chapter)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

// AnswerPayload 结构用于解析前端发送的请求体
type AnswerPayload struct {
	ID  int `json:"id"`
	Key int `json:"key"`
}

// AnswerResponse 结构用于向前端发送响应
type AnswerResponse struct {
	ID           int  `json:"id"`
	Correct      bool `json:"correct"`
	Expected_Key int  `json:"expected_key"`
	Terminate    bool `json:"terminate"`
}

// HandleAnswer 接受答案，并构建响应
func HandleAnswer(w http.ResponseWriter, r *http.Request) {
	var Expected_Key int
	var payload AnswerPayload
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 校验答案
	Correct, err, Expected_Key := checkAnswer(payload.ID, payload.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := AnswerResponse{
		ID:           payload.ID,
		Correct:      Correct,
		Expected_Key: Expected_Key,
		Terminate:    false, // 这里依旧需要修改！
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}

// checkAnswer 根据问题ID在数据库中检索正确答案，并与用户提交的答案进行比对检查其是否正确
func checkAnswer(ID int, Key int) (bool, error, int) {
	db := dbConn()
	defer db.Close()

	// 查询数据库中的正确答案
	var Expected_Key int
	err := db.QueryRow("SELECT Key FROM Questions WHERE ID = ?", ID).Scan(&Expected_Key)
	if err != nil {
		return false, err, Expected_Key
	}

	return (Key == Expected_Key), nil, Expected_Key
}

func HandleQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("接收到GET请求……")
		questions, err := GetAllQuestions()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responseBytes, err := json.Marshal(questions)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBytes)
	}
}

func HandleResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("接收到POST请求……")
	var updatedQuestion Q

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedQuestion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("json解码完成。")
	UpdateQuestionsInDatabase(updatedQuestion)
	fmt.Println("数据库已更新。")
	w.WriteHeader(http.StatusNoContent)
	fmt.Println("返回HTTP-204。")
}

func UpdateQuestionsInDatabase(updatedQuestion Q) error {
	// 连接数据库
	db := dbConn()
	defer db.Close() // 确保在函数结束时关闭数据库连接
	updatedQuestion.ID = int(updatedQuestion.ID)
	updatedQuestion.Type = string(updatedQuestion.Type)
	updatedQuestion.Chapter = string(updatedQuestion.Chapter)
	updatedQuestion.Key = string(updatedQuestion.Key)
	updatedQuestion.Times = string(updatedQuestion.Times)

	// 准备 SQL 更新语句
	stmt, err := db.Prepare(`
        UPDATE Questions SET
            Question = ?,
            Type = ?,
            Chapter = ?,
            Key = ?,
            Answer = ?,
            Times = ?
        WHERE ID = ?
    `)
	if err != nil {
		return err
	}
	defer stmt.Close() // 确保在函数结束时关闭 statement

	// 执行 SQL 语句，传入需要更新的值
	_, err = stmt.Exec(
		updatedQuestion.Question,
		updatedQuestion.Type,
		updatedQuestion.Chapter,
		updatedQuestion.Key,
		updatedQuestion.Answer,
		updatedQuestion.Times,
		updatedQuestion.ID,
	)
	if err != nil {
		return err
	}
	println("数据库已更新。")

	return nil
}

func main() {
	fmt.Println("Charles ©2024-2077")
	r := mux.NewRouter()

	// 添加路由
	r.HandleFunc("/answer", HandleAnswer).Methods("POST") // 处理提交答案的请求，限制为POST方法
	r.HandleFunc("/questions", HandleQuestions).Methods("GET", "POST")
	r.HandleFunc("/wrong", HandleWrongQuestions).Methods("GET")
	r.HandleFunc("/response", HandleResponse).Methods("POST")
	fmt.Println("Go 路由设置完毕。")

	// 设置CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// 应用CORS中间件
	fmt.Println("初始化CORS中间件……")
	http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(r))
}
