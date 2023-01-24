package errorcode

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG       = 10001
	ERROR_EXIST_TAG_FAIL  = 10002
	ERROR_NOT_EXIST_TAG   = 10003
	ERROR_GET_TAGS_FAIL   = 10004
	ERROR_COUNT_TAG_FAIL  = 10005
	ERROR_ADD_TAG_FAIL    = 10006
	ERROR_EDIT_TAG_FAIL   = 10007
	ERROR_DELETE_TAG_FAIL = 10008
	ERROR_EXPORT_TAG_FAIL = 10009
	ERROR_IMPORT_TAG_FAIL = 10010

	ERROR_NOT_EXIST_ARTICLE        = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012
	ERROR_ADD_ARTICLE_FAIL         = 10013
	ERROR_DELETE_ARTICLE_FAIL      = 10014
	ERROR_EDIT_ARTICLE_FAIL        = 10015
	ERROR_COUNT_ARTICLE_FAIL       = 10016
	ERROR_GET_ARTICLES_FAIL        = 10017
	ERROR_GET_ARTICLE_FAIL         = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "請求參數錯誤",
	ERROR_EXIST_TAG:                 "已存在該標籤名稱",
	ERROR_EXIST_TAG_FAIL:            "獲取已存在標籤丟失",
	ERROR_NOT_EXIST_TAG:             "該標籤不存在",
	ERROR_GET_TAGS_FAIL:             "獲取所有標籤丟失",
	ERROR_COUNT_TAG_FAIL:            "統計標籤失敗",
	ERROR_ADD_TAG_FAIL:              "新增加標籤失敗",
	ERROR_EDIT_TAG_FAIL:             "修改標籤失敗",
	ERROR_DELETE_TAG_FAIL:           "刪除標籤失敗",
	ERROR_EXPORT_TAG_FAIL:           "導出標籤失敗",
	ERROR_IMPORT_TAG_FAIL:           "導入標籤失敗",
	ERROR_NOT_EXIST_ARTICLE:         "該文章不存在",
	ERROR_ADD_ARTICLE_FAIL:          "新增加文章失敗",
	ERROR_DELETE_ARTICLE_FAIL:       "刪除文章失敗",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "檢查文章是否丟失",
	ERROR_EDIT_ARTICLE_FAIL:         "修改文章失敗",
	ERROR_COUNT_ARTICLE_FAIL:        "統計文章失敗",
	ERROR_GET_ARTICLES_FAIL:         "獲得多個文章失敗",
	ERROR_GET_ARTICLE_FAIL:          "獲取單個文章失敗",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "生成文章海報失敗",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "令牌鑑權失敗",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "令牌已超時",
	ERROR_AUTH_TOKEN:                "令牌生成失敗",
	ERROR_AUTH:                      "令牌錯誤",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存圖片丟失",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "檢查圖片丟失",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校試圖片錯誤，圖片格式或大小有問題",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
