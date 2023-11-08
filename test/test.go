package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mail-service/pkg/db"
	"mail-service/pkg/utils"
	"net/http"
	"time"
)

var (
	count = 0

	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}}
)

type Holder struct {
	TokenAddress string
	Address      string
	Balance      float64
	Share        float64
}

type Tokens struct {
	Tokens []Token
}

type Token struct {
	Address string
}

type Holders struct {
	Holders []Holder `json:"holders"`
}

func (tokens *Tokens) GetAllTokenEthereum() error {
	query := `select address from crypto c where chainname = 'ethereum' limit 1000`

	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}

	for rows.Next() {
		token := Token{}
		err := rows.Scan(&token.Address)
		if err != nil {
			return err
		}
		tokens.Tokens = append(tokens.Tokens, token)
	}
	return nil
}

func (holder *Holder) Insert() error {
	query := `INSERT INTO public.holder
	(tokenaddress, address, balance, "share")
	VALUES($1, $2, $3, $4);
	`

	_, err := db.PSQL.Exec(query, holder.TokenAddress, holder.Address, holder.Balance, holder.Share)

	return err
}

func main() {
	tokens := Tokens{}

	err := tokens.GetAllTokenEthereum()
	if err != nil {
		fmt.Println("tokens.GetAllTokenEthereum", err)
	}

	fmt.Println("len tokens", len(tokens.Tokens))
	start := time.Now()

	limit := 1000

	for _, token := range tokens.Tokens {
		apiKey := "EK-dJiPb-pawKjhw-EjLWN"
		err := CallEthplorer(token.Address, apiKey, limit)
		if err != nil {
			fmt.Println("CallEthplorer(address, apiKey)", err)
		}
	}
	fmt.Println("time diff ", time.Since(start))
}

func CallEthplorer(address string, apiKey string, limit int) error {
	request, err := http.NewRequest("GET", `https://api.ethplorer.io/getTopTokenHolders/`+address+`?apiKey=`+apiKey+`&limit=`+fmt.Sprintf("%d", limit), nil)
	request.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36`)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var data = make(map[string]any)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	holders := Holders{}

	err = utils.Mapping(data, &holders)

	if err != nil {
		fmt.Println("utils.Mapping(data, &holders)", err)
	}
	count += 1

	fmt.Println(count, address, len(holders.Holders))

	// return nil

	for _, holder := range holders.Holders {
		holder.TokenAddress = address
		err := holder.Insert()
		if err != nil {
			fmt.Println("holder.Insert()", err)
		}
	}
	return nil
}
