where 1=1
{{ if  .user_name}}
and user_name = {{.user_name}}
{{ end }}
{{ if .user_name_list }}
and user_name in ({{.user_name_list}})
{{ end}}

