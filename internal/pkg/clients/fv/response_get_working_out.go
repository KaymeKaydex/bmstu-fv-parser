package fv

type ResponseGetWorkingOut struct {
	Items []struct {
		Id                  string        `json:"id"`
		BranchId            string        `json:"branch_id"`
		CategoryId          interface{}   `json:"category_id"`
		AddressId           string        `json:"address_id"`
		Category            []interface{} `json:"category"`
		Title               string        `json:"title"`
		AddressText         interface{}   `json:"address_text"`
		Description         string        `json:"description"`
		Url                 string        `json:"url"`
		DateBegin           string        `json:"date_begin"`
		DateBeginTime       string        `json:"date_begin_time"`
		DateEnd             string        `json:"date_end"`
		DateRegistrationEnd string        `json:"date_registration_end"`
		CountUserMax        *string       `json:"count_user_max"`
		CountUsersAvailable int           `json:"count_users_available"`
		CountUsers          string        `json:"count_users"`
		Address             struct {
			Id      string `json:"id"`
			Title   string `json:"title"`
			Address string `json:"address"`
		} `json:"address"`
		JoinCurrent        bool `json:"join_current"`
		JoinCurrentApprove bool `json:"join_current_approve"`
		JoinNeedRequest    bool `json:"join_need_request"`
		IsAllowJoin        bool `json:"is_allow_join"`
		IsAllowEventCreate bool `json:"is_allow_event_create"`
		IsAllowEventRemove bool `json:"is_allow_event_remove"`
		IsAllowEventExport bool `json:"is_allow_event_export"`
		IsAllowViewUsers   bool `json:"is_allow_view_users"`
		Users              []struct {
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
			Url    string `json:"url"`
		} `json:"users"`
		UsersCountAll interface{} `json:"users_count_all"`
	} `json:"items"`
	SMsgTitle   string `json:"sMsgTitle"`
	SMsg        string `json:"sMsg"`
	BStateError bool   `json:"bStateError"`
}
