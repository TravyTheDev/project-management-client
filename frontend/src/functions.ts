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

  export const debounceFunc = (value: string, func: Function) => {
    const timeoutID: number = window.setTimeout(() => { }, 0)

    for (let id: number = timeoutID; id >= 0; id -= 1) {
        window.clearTimeout(id)
    }

    setTimeout(() => {
        if (value.length < 2) {
            value = ''
        }
        func(value)
    }, 300)
  }