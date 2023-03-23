package collect

func GetFilters() map[string]any {
	type register map[string]any
	return register{
		"uuid":                Uuid,
		"is_empty":            IsEmpty,
		"must":                Must,
		"current_date_time":   CurrentDateTime,
		"current_date_format": CurrentDateFormat,
		"replace":             Replace,
		"md5":                 Md5,
		"sub_str":             SubStr,
		"get_key":             GetKey,
		"pinyin":              Pinyin,
	}

}
