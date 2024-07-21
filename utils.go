package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func calculateTotalPoints(transaction TransactionDetail) int64 {
	var points int64 = 0
	retailer := transaction.Retailer
	purchaseDate := transaction.PurchaseDate
	purchaseTime := transaction.PurchaseTime
	totalSpent := transaction.Total
	allItems := transaction.Items
	// count alphanumeric chars
	points += countAlphaNumeric(retailer)
	// check if the total amount has cents
	if hasNoCents(totalSpent) {
		points += 50
	}
	// check if the total amount spent is the multiple of quarter
	if isMultipleOfQuarter(totalSpent) {
		points += 25
	}
	//5 points for every two items
	points += calculateFiveForEveryTwo(allItems)
	//check trimmed length and calculate
	points += calculatePointsFromItemDescription(allItems)
	//check if the purchase date is odd
	if isOddDate(purchaseDate) {
		points += 6
	}
	//check if the business was done after 2pm and before 4pm
	if isAfterTwoBeforeFour(purchaseTime) {
		points += 10
	}
	fmt.Printf("This is the total points earned: %.d\n", points)
	return points
}

func countAlphaNumeric(retailer string) int64 {
	count := 0
	for _, r := range retailer {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			count += 1
		}
	}
	return int64(count)
}

func hasNoCents(amount string) bool {
	cleanedAmount := strings.ReplaceAll(amount, "$", "")
	cleanedAmount = strings.ReplaceAll(cleanedAmount, ",", "")
	value, err := strconv.ParseFloat(cleanedAmount, 64)
	if err != nil {
		return false
	}
	return value == float64(int(value))
}

func isMultipleOfQuarter(n string) bool {
	floatVal, _ := strconv.ParseFloat(n, 64)
	result := floatVal / 0.25
	return math.Abs(result-math.Round(result)) < 1e-10
}

func calculateFiveForEveryTwo(allItems []Item) int64 {
	lengthOfItems := len(allItems)
	tempPoints := int64(lengthOfItems/2) * 5
	return tempPoints
}

func calculatePointsFromItemDescription(allItems []Item) int64 {
	curPoints := 0
	for _, product := range allItems {
		itemDescription := product.ShortDescription
		trimmedStr := strings.TrimSpace(itemDescription)
		trimmedStrLength := len(trimmedStr)
		if trimmedStrLength%3 == 0 {
			priceFloat, err := strconv.ParseFloat(product.Price, 64)
			if err != nil {
				fmt.Println("Error converting price:", err)
				return 0.0
			}
			temp := math.Ceil(priceFloat * 0.2)
			curPoints += int(temp)
		}
	}
	return int64(curPoints)
}

func isOddDate(stringDate string) bool {
	date, err := time.Parse("2006-01-02", stringDate)
	// fmt.Println("This is what the date looks like : ", stringDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}
	day := date.Day()
	return day%2 != 0
}

func isAfterTwoBeforeFour(stringTime string) bool {
	parsedPurchaseTime, err := time.Parse("15:04", stringTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return false
	}
	boundaryStart := time.Date(1970, time.January, 1, 14, 0, 0, 0, time.UTC)
	boundaryEnd := time.Date(1970, time.January, 1, 16, 0, 0, 0, time.UTC)
	parsedTimeOfDay := time.Date(1970, time.January, 1, parsedPurchaseTime.Hour(), parsedPurchaseTime.Minute(), 0, 0, time.UTC)
	return parsedTimeOfDay.After(boundaryStart) && parsedTimeOfDay.Before(boundaryEnd)
}
