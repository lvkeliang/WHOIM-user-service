namespace go user

struct User {
    1: string id,
    2: string username,
    3: string email,
    4: string status
}

service UserService {
    // 注册接口，传入用户名、密码和邮箱，返回成功或失败
    bool Register(1: string username, 2: string password, 3: string email)

    // 登录接口，传入用户名和密码，返回 JWT 令牌
    string Login(1: string username, 2: string password)

    // 获取用户信息，传入用户名，返回用户信息
    User GetUserInfo(1: string username)

    // 设置用户状态为在线
    bool SetUserOnline(1: string username)

    // 设置用户状态为离线
    bool SetUserOffline(1: string username)
}
