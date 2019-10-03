package recsys

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func intArrayToString(itemIds []int64) string {
	var itemText []string
	for i := range itemIds {
		number := itemIds[i]
		text := strconv.FormatInt(number, 10)
		itemText = append(itemText, text)
	}
	itemsList := strings.Join(itemText, ", ")
	return itemsList
}

func stringArrayToIntArray(itemIds []string) ([]int64, error) {
	var itemInts []int64
	for i := range itemIds {
		text := itemIds[i]
		number, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return nil, nil
		}
		itemInts = append(itemInts, number)
	}
	return itemInts, nil
}

type RecommendResult struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Predicts []string `json:"predicts"`
}

func SequenceRequest(sequence []int64, itemIds []int64) ([]int64, error) {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return nil, err
	}
	host, _ := os.LookupEnv("SEQUENCE_HOST")
	port, _ := os.LookupEnv("SEQUENCE_PORT")
	token, _ := os.LookupEnv("SEQUENCE_TOKEN")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf(`
	{
		"service_name": "SEQUENCE",
		"service_token": "%s", 
		"data": {
			"request": "rank_items", 
			"inputs": {
				"sequence": [%s], 
				"selected_items": [%s],
				"filter_items": []
			}
		}
	}`, token, intArrayToString(sequence), intArrayToString(itemIds))
	fmt.Fprintf(conn, str)
	resultStr, err := bufio.NewReader(conn).ReadString('\n')
	resultStr = strings.Replace(resultStr, "'", "\"", -1)
	var result = RecommendResult{}
	err = json.Unmarshal([]byte(resultStr), &result)
	if err != nil {
		return nil, err
	}
	finalResult, err := stringArrayToIntArray(result.Predicts)
	return finalResult, err
}

func FactorizationRequest(userId int64, itemIds []int64) ([]int64, error) {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return nil, err
	}
	host, _ := os.LookupEnv("FACTORIZATION_HOST")
	port, _ := os.LookupEnv("FACTORIZATION_PORT")
	token, _ := os.LookupEnv("FACTORIZATION_TOKEN")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf(`
	{
		"service_name": "FACTORIZATION",
		"service_token": "%s", 
		"data": {
			"request": "rank_items", 
			"inputs": {
				"user_id": %d, 
				"selected_items": [%s],
				"filter_items": []
			}
		}
	}`, token, userId, intArrayToString(itemIds))
	fmt.Fprintf(conn, str)
	resultStr, err := bufio.NewReader(conn).ReadString('\n')
	resultStr = strings.Replace(resultStr, "'", "\"", -1)
	var result = RecommendResult{}
	err = json.Unmarshal([]byte(resultStr), &result)
	if err != nil {
		return nil, err
	}
	finalResult, err := stringArrayToIntArray(result.Predicts)
	return finalResult, err

}
