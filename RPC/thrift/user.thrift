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

    // 验证 JWT 令牌，返回用户信息
    User ValidateToken(1: string token)

    // 获取用户信息，传入用户 ID (UUID)，返回用户信息
    User GetUserInfo(1: string id)

    // 设置用户状态为在线，传入用户 ID (UUID)
    bool SetUserOnline(1: string id)

    // 设置用户状态为离线，传入用户 ID (UUID)
    bool SetUserOffline(1: string id)

    // 获取用户状态，传入用户 ID (UUID)，仅返回在线/离线状态
    string GetUserStatus(1: string id)
}
