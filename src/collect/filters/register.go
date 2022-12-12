package collect

func GetFilters() map[string]any {
	type register map[string]any
	return register{
		"uuid":              Uuid,
		"is_empty":          IsEmpty,
		"must":              Must,
		"current_date_time": CurrentDateTime,
		"md5":               Md5,
	}

}
