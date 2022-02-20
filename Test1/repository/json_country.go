package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/aryadiahmad4689/test-citcall/entity"
)

type CountryRepository struct {
}

func NewCoutryRepository() Country {
	return &CountryRepository{}
}

// get all
func (country *CountryRepository) GetAll() ([]entity.Country, error) {
	body, _ := country.getJsonCountry("https://citcall.com/test/countries.json")

	countrys := []entity.Country{}
	jsonErr := json.Unmarshal(body, &countrys)
	if jsonErr != nil {
		log.Println(jsonErr)
		return nil, jsonErr
	}
	return countrys, nil
}

// getJsonCountry is a function to get json country
func (country *CountryRepository) getJsonCountry(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body, nil
}
