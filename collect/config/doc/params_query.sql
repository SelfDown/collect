SELECT a.*
FROM collect_doc_params a
where 1=1
{{ if .collect_doc_id }}
and a.collect_doc_id ={{.collect_doc_id}}
{{ end }}