LIST_DATABASE_URI = "/dbs"
# create or delete database use this uri
DATABASE_URI = "/dbs/%(database_name)s"
LIST_SPACE_URI = "/dbs/%(database_name)s/spaces"

SPACE_URI = "/dbs/%(database_name)s/spaces/%(space_name)s"
UPSERT_DOC_URI = "/document/upsert"
DELETE_DOC_URI = "/document/delete"
SEARCH_DOC_URI = "/document/search"
QUERY_DOC_URI = "/document/query"
INDEX_URI = "/document/index"

AUTH_KEY = "Authorization"

CODE_SUCCESS = 0
CODE_INTERNAL_ERROR = 1
CODE_UNKNOWN_ERROR = 2
CODE_AUTHENTICATION_FAILED = 3
CODE_RECOVER = 4
CODE_TIMEOUT = 5
CODE_PARAM_ERROR = 6
CODE_CONFIG_ERROR = 7

CODE_SPACE_NOT_EXIST = 221

CODE_DATABASE_NOT_EXIST = 200
CODE_DB_EXIST = 201

MSG_NOT_EXIST = "not_exist"

# document 260-279
CODE_DOCUMENT_NOT_EXIST = 260
CODE_PRIMARY_KEY_IS_INVALID = 261

# filter 300-319
CODE_FILTER_OPERATOR_TYPE_ERR = 300
CODE_FILTER_CONDITION_OPERATOR_TYPE_ERR = 301

CODE_ERR_CODE_UPSERT_INVALID_PARAMS = 400

# document query 440-459
CODE_QUERY_ENGINE_ERR = 440
CODE_QUERY_INVALID_PARAMS_LENGTH_OF_DOCUMENT_IDS_BEYOND_500 = 441
CODE_QUERY_INVALID_PARAMS_SHOULD_HAVE_ONE_OF_DOCUMENT_IDS_OR_FILTER = 442
CODE_QUERY_INVALID_PARAMS_SHOULD_NOT_HAVE_VECTOR_FIELD = 443
CODE_QUERY_INVALID_PARAMS_BOTH_DOCUMENT_IDS_AND_FILTER = 444
CODE_QUERY_RESPONSE_PARSE_ERR = 445

# document search 460-479
CODE_SEARCH_INVALID_PARAMS_SHOULD_HAVE_VECTOR_FIELD = 460
CODE_SEARCH_ENGINE_ERR = 461
CODE_SEARCH_RESPONSE_PARSE_ERR = 462
