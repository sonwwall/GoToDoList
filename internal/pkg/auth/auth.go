package auth

import "golang.org/x/crypto/bcrypt"

//关于该文件的描述：
//在软件项目中，auth（代表“authentication”，即“认证”）目录通常用于存放与用户认证和授权相关的代码。
//认证是验证用户身份的过程，而授权是确定用户是否有权限执行特定操作的过程。
//以下是auth目录可能包含的一些功能和组件：
//1.密码哈希和验证：
//提供密码哈希和验证功能，确保用户密码在存储和验证时都是安全的。
//2.令牌管理：
//处理身份验证令牌（如JWT）的生成、验证和刷新，用于在无状态的HTTP协议中维护用户会话。
//3.会话管理：
//如果使用基于会话的认证，可能会包含会话创建、管理和销毁的逻辑。
//4.OAuth客户端和服务端逻辑：
//如果应用支持OAuth认证流程，可能会包含OAuth相关的实现。
//5.用户登录和登出：
//实现用户登录和登出的业务逻辑。
//6.权限检查：
//检查用户是否有权限访问特定资源或执行特定操作。
//7.第三方认证集成：
//集成第三方认证服务（如Google、Facebook登录）的代码。
//8.安全措施：
//实现安全措施，如防止暴力破解、密码重置逻辑、双因素认证等。
//9.用户模型和数据库访问：
//可能包含用户模型的定义以及与用户认证相关的数据库访问逻辑。
//10.错误处理：
//处理认证过程中可能出现的错误，如无效的用户名或密码。

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) //14是盐值
	//更详细的解释：
	//14 代表的是 bcrypt 哈希算法的代价因子（cost factor）。
	//代价因子决定了哈希计算的复杂度和所需的时间，
	//从而影响密码哈希的安全性和性能。bcrypt 使用这个因子来增加计算哈希的时间，使得暴力破解更加困难
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
