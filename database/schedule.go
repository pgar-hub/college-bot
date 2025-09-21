package database

import (
	"database/sql"
	"fmt"
)

func SeedSchedule(db *sql.DB) error {
	// Проверяем, есть ли данные
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM schedule").Scan(&count)
	if err != nil {
		return fmt.Errorf("check schedule count: %w", err)
	}

	if count > 0 {
		fmt.Println("Расписание уже заполнено")
		return nil
	}

	// Если пусто — вставляем данные
	schedule := []struct {
		weekType  string
		dayOfWeek string
		subject   string
		cabinet   string
	}{
		{"числитель", "понедельник", "Информационные технологии, Основы алгоритмизации и программирования, Элементы высшей математики", "523,514,512"},
		{"числитель", "вторник", "Техн. физ. уровня передачи данных, Физра, иностранный язык в проф деятельности", "523,зал,307/309"},
		{"числитель", "среда", "Основы алгоритмизации и программирования, Информационные технологии, Техн. физ. уровня передачи данных", "514,523,523"},
		{"числитель", "четверг", "История России, Архетиктура аппаратных средств, Операционные системы и среды", "622,330,503"},
		{"числитель", "пятница", "Архетиктура аппаратных средств, Арзетиктура аппаратных средств, Элементы высшей математики", "330,330,512"},
		{"числитель", "суббота", "выходной", ""},
		{"числитель", "Воскресенье", "выходной", ""},
		{"знаменатель", "понедельник", "Информационные технологии, Основы алгоритмизации и программирования, Архетиктура аппаратных средств", "523,514,330"},
		{"знаменатель", "вторник", "Техн. физ. уровня передачи данных, Физра, иностранный язык в проф деятельности", "523,зал,307/309"},
		{"знаменатель", "среда", "Основы алгоритмизации и программирования, Операционные системы и среды, Основы алгоритмизации и программирования", "514,503,514"},
		{"знаменатель", "четверг", "Информационные технологии, История России, Операционные системы и среды", "523,622,503"},
		{"знаменатель", "пятница", "Архетиктура аппарытных средств, История России, Элементы высшей математики", "330,622,512"},
		{"знаменатель", "суббота", "выходной", ""},
		{"знаменатель", "Воскресенье", "выходной", ""},
	}

	for _, s := range schedule {
		_, err := db.Exec(
			"INSERT INTO schedule (week_type, day_of_week, subject) VALUES ($1, $2, $3)",
			s.weekType, s.dayOfWeek, s.subject,
		)
		if err != nil {
			return fmt.Errorf("insert schedule: %w", err)
		}
	}

	fmt.Println("Расписание добавлено в таблицу")
	return nil
}

func GetSchedule(db *sql.DB, weekType, dayOfWeek string) (string, error) {
	rows, err := db.Query("SELECT subject, cabinet FROM schedule WHERE week_type=$1 AND day_of_week=$2", weekType, dayOfWeek)
	if err != nil {
		return "", fmt.Errorf("query schedule: %w", err)
	}
	defer rows.Close()

	var subjects []string
	for rows.Next() {
		var subject, cabinet string
		if err := rows.Scan(&subject, &cabinet); err != nil {
			return "", fmt.Errorf("scan schedule: %w", err)
		}
		subjects = append(subjects, fmt.Sprintf("%s (каб. %s)", subject, cabinet))
	}

	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("rows error: %w", err)
	}

	if len(subjects) == 0 {
		return "Пар нет!", nil
	}

	return fmt.Sprintf("Расписание на %s (%s):\n%s", dayOfWeek, weekType, Subjects(subjects)), nil
}
func Subjects(subjects []string) string {
	result := ""
	for i, subject := range subjects {
		result += fmt.Sprintf("%d. %s\n", i+1, subject)
	}
	return result
}
