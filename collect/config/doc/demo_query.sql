SELECT a.*
FROM collect_doc_demo a
where 1=1
{{ if .collect_doc_id }}
and a.collect_doc_id ={{.collect_doc_id}}
{{ end }}