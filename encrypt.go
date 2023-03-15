package encrypt

// Encrypt
// 对字符串进行双重MD5加密
//
// @str 需要处理的字符串
// @salt 盐值: 随机字符串
// @tail 首次MD5后保留的位数
func Encrypt(str, salt string, tail uint) string {

	str = MD5([]byte(str))

	str = str[:tail] + salt

	return MD5([]byte(str))
}
