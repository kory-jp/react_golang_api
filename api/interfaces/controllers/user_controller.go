package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kory-jp/react_golang_api/api/domain"
	"github.com/kory-jp/react_golang_api/api/interfaces/database"
	"github.com/kory-jp/react_golang_api/api/usecase"
)

type UserController struct {
	Interfactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interfactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	// ioutilは入出力関連のユーティリティ関数が定義されている
	// ユーティリティ関数 = Function オブジェクトを使用して、特定の関数の GUID を解析、エンコード、復号化、または返します
	// ReadAll = os.Openなどで開いているファイル(*os.File)からファイルの内容を一度にすべて読み込む。
	// ファイルの内容はバイト型なので、文字列型として使用する場合は string(data) します。
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 空のユーザー構造体をインスタンス
	userType := new(domain.User)
	// json形式のデータをユーザー構造体の型に変換
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		fmt.Println(err)
		return
	}
	user, err := controller.Interfactor.Add(*userType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
	fmt.Fprintln(w, user)
}

func (controller *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := controller.Interfactor.Users()
	if err != nil {
		log.Panicln(err)
		return
	}
	us, err := json.Marshal(users)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Fprintf(w, string(us))
}

func (controller *UserController) Show(w http.ResponseWriter, r *http.Request, id int) {
	user, err := controller.Interfactor.UserById(id)
	if err != nil {
		log.Println(err)
		return
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(jsonUser))
}
