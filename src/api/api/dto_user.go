package api

func CreateUser(user User) (err error){
  user.Password = GetMd5FromString(user.Password)
  result := getGormDb().Create(&user)
  return result.Error
}

func GetUserByName(name string) (user *User){
  getGormDb().First(&user, "name = ?", name)
  return
}

func GetUserByToken(token string) (user *User){
  getGormDb().First(&user, "token = ?", token)
  if user.Name == "" { user = nil }
  return
}

func UpdateUserToken(username, token string) (err error){
  var user User
  err = getGormDb().First(&user, "name = ?", username).Error
  if err != nil {
    return err
  }

  err = getGormDb().Model(&user).Update("token", token).Error
  if err != nil {
    return err
  }

  return nil
}

