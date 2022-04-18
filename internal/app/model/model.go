package model

import (
	"time"
)

// WorkingOutItem - отработка
type WorkingOutItem struct {
	Id                  int       `json:"id"`             // id отработки
	BranchId            int       `json:"branch_id"`      // Неизвестное поле, но совпадает с id в запросе
	PointsGeneral       int       `json:"points_general"` // Основные баллы (за посещения)
	PointsExtra         int       `json:"points_extra"`   // Доп баллы за отработку
	Title               string    `json:"title"`          // Название
	Description         string    `json:"description"`    // Описание "отработки"
	Url                 string    `json:"url"`            // Ведет на адрес с отработкой. Там можно взят
	DateBegin           time.Time `json:"date_begin"`
	DateEnd             time.Time `json:"date_end"`
	DateRegistrationEnd time.Time `json:"date_registration_end"`
	CountUserMax        *string   `json:"count_user_max"` // Максимальное число записей
	CountUsers          string    `json:"count_users"`    // Число записавшихся
	AddressId           int       `json:"id"`
	AddressTitle        string    `json:"title"`
	Address             string    `json:"address"`
	IsAllowJoin         bool      `json:"is_allow_join"` // Можно ли еще записаться?

	UsersCountAll int `json:"users_count_all"` // Число людей записанных на эту отработку
}

type Address struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
}
