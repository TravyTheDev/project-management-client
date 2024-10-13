import { GetUser } from "../wailsjs/go/main/App"
import { types } from "../wailsjs/go/models"

export const getUser = async () => {
    let user:types.User
    try {
      const res = await GetUser()
      user = res
      if (user.id === 0) {
        return undefined
      }else {
        return user
      }
    } catch (error) {
      console.log(error)
    }
  }