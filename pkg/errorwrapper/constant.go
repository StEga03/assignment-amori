package errorwrapper

const (
	defaultUserMessageTranslationID = "unknown_error"
)

// Define all error keys here, corresponding with error-list.json / custom-list.json
const (
	// Unclassified Error.
	ErrIDTest   ErrID   = "test_error"
	ErrCodeTest ErrCode = -1

	// Unclassified Error.
	ErrIDUnclassified   ErrID   = "Unclassified"
	ErrCodeUnclassified ErrCode = 99999

	// App Initialize Error.
	ErrIDDatabaseInit ErrID = "database_init"
	ErrIDRedisInit    ErrID = "redis_init"
	ErrIDRepoInit     ErrID = "repo_init"
	ErrIDServiceInit  ErrID = "service_init"
	ErrIDGraceInit    ErrID = "grace_init"

	// Generic Error.
	ErrIDUnmarshal                    ErrID = "unmarshal_err"
	ErrIDMarshal                      ErrID = "marshal_err"
	ErrDataAlreadyExist               ErrID = "data_already_exist"
	ErrDataNotFound                   ErrID = "data_not_found"
	ErrDataIsEmpty                    ErrID = "data_is_empty"
	ErrDataIsNotValid                 ErrID = "data_is_not_valid"
	ErrCasting                        ErrID = "casting_err"
	ErrParsing                        ErrID = "parsing_err"
	ErrContextValueNotFound           ErrID = "context_value_not_found"
	ErrPopulateData                   ErrID = "populate_data_error"
	ErrIDFailedVerifyEncryptedContent ErrID = "failed_verify_encrypted_content"
	ErrIDFailedGetEncryptedResult     ErrID = "failed_to_get_encrypted_result"
	ErrIDFailedDecrypt                ErrID = "failed_to_decrypt"
	ErrIDFailedGetSignature           ErrID = "failed_to_get_signature"
	ErrIDFailedGetFeatureFlag         ErrID = "failed_to_get_feature_flag"
	ErrIDFailedToGenerateID           ErrID = "failed_to_generate_id"
	ErrIDPhoneNumberNotValid          ErrID = "phone_number_not_valid"
	ErrIDTxElementIsEmpty             ErrID = "empty_tx_element"

	// HTTP Client Error.
	ErrIDHTTPClient       ErrID = "http_client_error"
	ErrIDNewHTTPRequest   ErrID = "new_http_request_error"
	ErrIDHTTPResponseCode ErrID = "http_response_code_error"

	// Service Error.
	ErrIDServiceExpectedError   ErrID = "service_expected_error"
	ErrIDServiceUnexpectedError ErrID = "service_unexpected_error"

	// MQ Handler Error.
	ErrIDBackgroundJobImplementorNotFound ErrID = "background_job_implementor_not_found"

	// Validation Error.
	ErrIDValidationNotPassed       ErrID = "validation_not_passed"
	ErrIDIPAddressNotAllowed       ErrID = "ip_address_not_allowed"
	ErrIDHandlerBadRequest         ErrID = "bad_request_error"
	ErrIDUnauthorized              ErrID = "unauthorized"
	ErrIDStatusValidationNotPassed ErrID = "status_validation_not_passed"
	ErrIDRuleValidationNotPassed   ErrID = "rule_validation_not_passed"
	ErrIDPublicKeyNotRegistered    ErrID = "public_key_not_registered"

	// Repo Error.
	ErrIDGetFromService      ErrID = "get_from_service_error"
	ErrIDGetFromDB           ErrID = "get_from_db_error"
	ErrIDFailedGetFromRepo   ErrID = "get_from_repo_error"
	ErrIDUpdateDB            ErrID = "update_db_error"
	ErrIDDeleteDB            ErrID = "delete_db_error"
	ErrIDInsertDB            ErrID = "insert_db_error"
	ErrIDUpsertDB            ErrID = "upsert_db_error"
	ErrIDPrepareDB           ErrID = "prepare_db_error"
	ErrIDHashKeyEmpty        ErrID = "hash_key_empty"
	ErrIDInvalidConfigKey    ErrID = "invalid_config_key"
	ErrIDExecuteRedis        ErrID = "redis_exec_error"
	ErrIDFailedParsingEntity ErrID = "entity_parsing_error"
	ErrIDFailedParseUUID     ErrID = "uuid_parsing_error"
	ErrIDFailedParseTime     ErrID = "time_parsing_error"

	// Database Transaction.
	ErrIDFailedToBegin    ErrID = "failed_to_begin"
	ErrIDFailedToCommit   ErrID = "failed_to_commit"
	ErrIDFailedToRollback ErrID = "failed_to_rollback"

	// Repo Module Error.
	// GET.
	ErrIDFailedGetFromRepoChannel       ErrID = "get_from_repo_channel_error"
	ErrIDFailedGetFromRepoMessage       ErrID = "get_from_repo_message_error"
	ErrIDFailedGetFromRepoMessageInput  ErrID = "get_from_repo_message_input_error"
	ErrIDFailedGetFromRepoMessageSource ErrID = "get_from_repo_message_source_error"
	ErrIDFailedGetFromRepoUser          ErrID = "get_from_repo_user_error"

	// CREATE.
	ErrIDFailedCreateFromRepoChannel       ErrID = "create_from_repo_channel_error"
	ErrIDFailedCreateFromRepoMessage       ErrID = "create_from_repo_message_error"
	ErrIDFailedCreateFromRepoMessageInput  ErrID = "create_from_repo_message_input_error"
	ErrIDFailedCreateFromRepoMessageSource ErrID = "create_from_repo_message_source_error"
	ErrIDFailedCreateFromRepoUser          ErrID = "create_from_repo_user_error"

	// File Operation.
	ErrIDFileExtensionNotValid       ErrID = "file_extension_not_valid"
	ErrIDFileIsNotExist              ErrID = "file_is_not_exist"
	ErrIDUploadFailedToRead          ErrID = "failed_to_read_file"
	ErrIDProcessFileFailedToUpload   ErrID = "failed_to_upload"
	ErrIDProcessFileFailedToDownload ErrID = "failed_to_download"
	ErrIDProcessFileFailedToRead     ErrID = "failed_read_process_file"
	ErrIDPerformDynamicActionFailed  ErrID = "failed_to_execute_custom_action_on_jobs"
	ErrIDProcessFileFailedToWrite    ErrID = "failed_write_process_file"

	// Repo Module Error.
	ErrIDChannelDataIsEmpty       ErrID = "channel_data_is_empty"
	ErrIDMessageDataIsEmpty       ErrID = "message_data_is_empty"
	ErrIDMessageInputDataIsEmpty  ErrID = "message_input_data_is_empty"
	ErrIDMessageSourceDataIsEmpty ErrID = "message_source_data_is_empty"
	ErrIDUserDataIsEmpty          ErrID = "user_data_is_empty"

	// Usecase Module Error.
	// Category.
	ErrIDCategoryTypeInvalid ErrID = "category_type_invalid"
)

// Error code list.
const (
	msgUnclassifiedErr = "unclassified error"
)

// Metadata keys.
const (
	metaKeyErrID       MetaKey = "error_id"
	metaKeyErrCode     MetaKey = "error_code"
	metaKeyHTTPStatus  MetaKey = "http_status"
	metaKeyDevMsg      MetaKey = "dev_message"
	metaKeyErrLine     MetaKey = "error_line"
	metaKeyValidator   MetaKey = "validator"
	metaKeyIsRetryable MetaKey = "is_retryable"
)

// User message params keys.
// TODO: Need to find a way how to make this able to store multiple languages.
const (
	// Field: Field which makes this error triggered.
	// Example: "name", "email", "phone", "id", etc.
	Field UserMsgKey = "Field"

	// DataOperator: What went wrong in this data operation.
	/*
				Current list:
					- get -> mendapatkan
					- insert -> memasukkan
					- update -> mengubah
					- delete -> menghapus
					- parse -> mengurai
					- marshal -> membentuk
					- unmarshal -> mengurai
					- decode -> mengurai
					- process -> memproses
					- find -> menemukan
					- initialize -> menginisialisasi
		  			- rollback -> mengembalikan
		  			- lock -> mengunci
		  			- append -> membubuhkan
					- transition -> mengalihkan
					- create -> membuat
					- decrypt -> dekripsi
	*/
	DataOperator UserMsgKey = "DataOperator"

	// DataName: Name of the data which makes this error triggered.
	// example: "user", "auth", etc.
	DataName UserMsgKey = "DataName"

	// SingleOperator: Unary operation/validation for the data.
	/*
		Current List:
			- required -> dibutuhkan
			- missing -> hilang
			- empty -> kosong
			- invalid -> tidak valid
			- incorrect -> salah
			- not found -> tidak ditemukan
			- unfinished -> belum selesai

		Note: can add "not" for the operator.
	*/
	SingleOperator UserMsgKey = "SingleOperator"

	// DoubleOperator: Binary operation/validation for the data.
	/*
		Current List:
			- greater than -> lebih besar dari
			- less than -> lebih kurang dari
			- equal to -> sama dengan
			- between -> di antara

		Note: can add "not" for the operator.
	*/
	DoubleOperator UserMsgKey = "DoubleOperator"

	// Comparator: Comparator for the data.
	// Example: 5.
	Comparator UserMsgKey = "Comparator"

	// Comparator: Comparator Unit for the data.
	// Example: seconds.
	ComparatorUnit UserMsgKey = "ComparatorUnit"

	// TargetState: Target state which wanted to be processed.
	// Example: RECEIVED
	TargetState UserMsgKey = "TargetState"

	// Usecase: The name of the usecase/handler.
	// Example: Create User.
	Usecase UserMsgKey = "Usecase"

	// Reason: The reason why this error is triggered.
	/*
		Example:
			- context cancelled -> konteks dibatalkan
			- process time is longer than expected -> Lama pemrosesan lebih dari batas waktu yang telah ditentukan
	*/
	Reason UserMsgKey = "Reason"

	// Reserved for PluralCount field.
	// Use if you want to differentiate between singular and plural message.
	// Ref: https://cldr.unicode.org/index/cldr-spec/plural-rules
	PluralCount UserMsgKey = "PluralCount"
)
