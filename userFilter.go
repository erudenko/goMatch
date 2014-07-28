package main

// Регистрация:
// Запрос: {"register_user":"USER_ID",
// "providesFilters":[массив_фильтров_которые_предоставляет_пользователь],
// "requestsFilters":[массив_фильтров_которые_задал_пользователь_для_поиска]}
// Ответ: {"type":"system","id":"USER_ID","message":"connected"}

// Фильтры:

// Страны: "countryCodes":[...], например, "ru", "ua" и т.д.
// Пол: "gender":"male/female"
// Языки: "languageCodes":[…], например, "ru", "ua" и т.д.
// Интересы: "interests":[...]

type UserFilter struct {
	ProvideFilters FilterDetails `json:"provideFilters"`
	RequestFilters FilterDetails `json:"requestFilters"`
	UserID         string        `json:"userID"`
}

type FilterDetails struct {
	CountryCodes  []string `json:"countryCodes"`
	Interests     []string `json:"interests"`
	LanguageCodes []string `json:"languageCodes"`
	Gender        string   `json:"gender"`
}
