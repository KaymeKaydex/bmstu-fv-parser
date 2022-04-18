package fv

import (
	"time"
)

type ResponseGetWorkingOut struct {
	Items []struct {
		Id                  string    `json:"id"`             // id отработки
		BranchId            string    `json:"branch_id"`      // Неизвестное поле, но совпадает с id в запросе
		PointsGeneral       int       `json:"points_general"` // Основные баллы (за посещения)
		PointsExtra         int       `json:"points_extra"`   // Доп баллы за отработку
		Title               string    `json:"title"`          // Название
		Description         string    `json:"description"`    // Описание "отработки"
		Url                 string    `json:"url"`            // Ведет на адрес с отработкой. Там можно взят
		DateBegin           time.Time `json:"date_begin"`
		DateEnd             time.Time `json:"date_end"`
		DateRegistrationEnd time.Time `json:"date_registration_end"`
		CountUserMax        *string   `json:"count_user_max"`        // Максимальное число записей
		CountUsersAvailable int       `json:"count_users_available"` // Сколько еще доступно записей?
		CountUsers          string    `json:"count_users"`           // Число записавшихся
		Address             struct {
			Id      string `json:"id"`
			Title   string `json:"title"`
			Address string `json:"address"`
		} `json:"address"`
		JoinCurrent        bool       `json:"join_current"`          // useless
		JoinCurrentApprove bool       `json:"join_current_approve"`  // useless
		JoinNeedRequest    bool       `json:"join_need_request"`     // useless
		IsAllowJoin        bool       `json:"is_allow_join"`         // Можно ли еще записаться?
		IsAllowEventCreate bool       `json:"is_allow_event_create"` // useless
		IsAllowEventRemove bool       `json:"is_allow_event_remove"` // useless
		IsAllowEventExport bool       `json:"is_allow_event_export"` // useless
		IsAllowViewUsers   bool       `json:"is_allow_view_users"`   // useless
		Users              []struct { // Последние 6 или менее людей записавшихся на отработку
			Name   string `json:"name"`   // Имя пользователя
			Avatar string `json:"avatar"` // Ссылка на аватар пользователя
			Url    string `json:"url"`    // Ссылка на профиль пользователя
		} `json:"users"`
		UsersCountAll int `json:"users_count_all"` // Число людей записанных на эту отработку
	} `json:"items"` // Отработки

	SMsgTitle   string `json:"sMsgTitle"`   // Сообщение об ошибке запроса
	SMsg        string `json:"sMsg"`        // Еше какое-то ошибочное сообщение
	BStateError bool   `json:"bStateError"` // Статус ошибки
}
