package api

import (
	"errors"
	"math"
	"receipt-processor-challenge/internal/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func countRetailerNamePoints(purchase *models.Purchase) int {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

	// One point for every alphanumeric character in the retailer name.
	return len(nonAlphanumericRegex.ReplaceAllString(purchase.Retailer, ""))
}

func countTotalCostPoints(purchase *models.Purchase) (int, error) {
	// 50 points if the total is a round dollar number.
	total, parseErr := strconv.ParseFloat(purchase.Total, 64)
	points := 0
	if total == float64(int(total)) {
		points += 50
	}
	if parseErr != nil {
		return 0, errors.New("Could not convert total cost to float")
	}

	// 25 points if the total is a multiple of 0.25.
	if math.Mod(float64(total), 0.25) == 0 {
		points += 25
	}

	return points, nil
}

func countItemDescPoints(purchase *models.Purchase) (int, error) {
	points := 0

	// if the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer
	for _, item := range purchase.Items {
		price, parseErr := strconv.ParseFloat(item.Price, 64)
		if parseErr != nil {
			return 0, errors.New("Could not convert Item price to float")
		}
		if math.Mod(float64(len(strings.TrimSpace(item.ShortDescription))), 3) == 0 {
			points += int(math.Ceil(price * 0.2))
		}
	}

	return points, nil
}

func countPurchaseDatePoints(purchase *models.Purchase) (int, error) {
	dateLayout := "2006-01-02"
	timeLayout := "15:04"
	maxTime, _ := time.Parse(timeLayout, "16:00")
	minTime, _ := time.Parse(timeLayout, "14:00")
	purchaseDate, parseErr := time.Parse(dateLayout, purchase.PurchaseDate)
	points := 0

	if parseErr != nil {
		return 0, errors.New("Could not parse purchase date")
	}
	// 6 points if the day in the purchase date is odd.
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, parseErr := time.Parse(timeLayout, purchase.PurchaseTime)
	if parseErr != nil {
		return 0, errors.New("Could not parse purchase time")
	}
	if purchaseTime.After(minTime) && purchaseTime.Before(maxTime) {
		points += 10
	}

	return points, nil
}

func CalculatePoints(purchase *models.Purchase) (int, error) {
	points := 0

	// One point for every alphanumeric character in the retailer name.
	points += countRetailerNamePoints(purchase)

	// 50 points if the total is a round dollar amount with no cents.
	// 25 points if the total is a multiple of 0.25.
	totalCostPoints, err := countTotalCostPoints(purchase)
	if err != nil {
		return 0, err
	}
	points += totalCostPoints

	// 5 points for every two items on the receipt.
	points += (len(purchase.Items) / 2) * 5

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	itemDescPoints, err := countItemDescPoints(purchase)
	if err != nil {
		return 0, err
	}
	points += itemDescPoints

	// 6 points if the day in the purchase date is odd.
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseDatePoints, err := countPurchaseDatePoints(purchase)
	if err != nil {
		return 0, err
	}
	points += purchaseDatePoints

	return points, nil
}
