

proto:



gen:
	studioctl model mysql datasource --url="admin:linkmax@tcp(139.9.178.36:3306)/labelstore_ydb" --table=cl_tpl --dir=model
	studioctl model mysql datasource --url="admin:linkmax@tcp(139.9.178.36:3306)/labelstore_ydb" --table=cl_tpl_attrs --dir=model