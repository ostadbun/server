package authService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ostadbun/interface/auth"
	authReposipory "ostadbun/repository/auth"
)

type Auth struct {
	repo authReposipory.AuthRepo
}

func Config(r authReposipory.AuthRepo) Auth {
	return Auth{
		repo: r,
	}
}

func (a Auth) RequestToGoogle(data auth.IAuth) (string, error) {
	var recive auth.IAuthSupabase

	url := "https://jpjukqmcfgpwtlaehajo.supabase.co/auth/v1/user"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		//fmt.Println(err)
		return "", err
	}

	token := fmt.Sprintf("bearer %s", data.Access_token)

	fmt.Println("💀", token)
	req.Header.Add("apikey", "sb_publishable_6otjdsE70iLxucZjoslj6g_mrwagC_c")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		//fmt.Println("🫟", err)
		return "", err
	}

	err = json.Unmarshal(body, &recive)
	if err != nil {
		return "", err
	}
	fmt.Println(recive)
	fmt.Println(string(body))

	return recive.Email, nil

}
