package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const inf = 1000000000000000

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))
	prices := make([]int, n)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		price, _ := strconv.Atoi(strings.TrimSpace(line))
		prices[i] = price
	}
	expenses, ticketsRemain, ticketsSpend, answerRoad := ticketUsage(n, prices)
	writer.WriteString(strconv.Itoa(expenses) + "\n")
	writer.WriteString(strconv.Itoa(ticketsRemain) + " ")
	writer.WriteString(strconv.Itoa(ticketsSpend) + "\n")
	for i := ticketsSpend - 1; i >= 0; i-- {
		writer.WriteString(strconv.Itoa(answerRoad[i]+1) + "\n")
	}
}

func ticketUsage(n int, prices []int) (int, int, int, []int) {
	if n == 0 {
		return 0, 0, 0, []int{}
	}
	
	// Создаем DP таблицу с запасом по размеру
	maxTickets := n + 1
	dp := make([][]int, n)
	prev := make([][]int, n) // для восстановления пути
	
	for i := range dp {
		dp[i] = make([]int, maxTickets)
		prev[i] = make([]int, maxTickets)
		for j := range dp[i] {
			dp[i][j] = inf
			prev[i][j] = -1
		}
	}

	// Инициализация первого дня
	if prices[0] > 100 {
		dp[0][1] = prices[0]
	} else {
		dp[0][0] = prices[0]
	}

	// Заполнение DP таблицы
	for day := 1; day < n; day++ {
		for tickets := 0; tickets < maxTickets; tickets++ {
			if dp[day-1][tickets] == inf {
				continue
			}
			
			currentPrice := prices[day]
			
			// Вариант 1: Покупаем обед
			newTickets := tickets
			if currentPrice > 100 {
				newTickets++
			}
			if newTickets < maxTickets && dp[day][newTickets] > dp[day-1][tickets]+currentPrice {
				dp[day][newTickets] = dp[day-1][tickets] + currentPrice
				prev[day][newTickets] = tickets
			}
			
			// Вариант 2: Используем купон если есть
			if tickets > 0 && dp[day][tickets-1] > dp[day-1][tickets] {
				dp[day][tickets-1] = dp[day-1][tickets]
				prev[day][tickets-1] = tickets
			}
		}
	}

	// Находим минимальные расходы
	expenses := inf
	ticketsRemain := 0
	for tickets := 0; tickets < maxTickets; tickets++ {
		if dp[n-1][tickets] <= expenses {
			expenses = dp[n-1][tickets]
			ticketsRemain = tickets
		}
	}

	// Восстанавливаем дни использования купонов
	answerRoad := []int{}
	curTickets := ticketsRemain
	for day := n - 1; day > 0; day-- {
		prevTickets := prev[day][curTickets]
		if prevTickets > curTickets {
			// Использовали купон в этот день
			answerRoad = append(answerRoad, day)
		}
		curTickets = prevTickets
	}

	return expenses, ticketsRemain, len(answerRoad), answerRoad
}