
const adminKey = 'acs_admin_token'
const adminUser = 'acs_admin_name'

export const SetToken = (token)=>{
    localStorage.setItem(adminKey,token)
}

export const GetToken = ()=>{
    return localStorage.getItem(adminKey)
}

export const RemoveToken = ()=>{
    localStorage.removeItem(adminKey)
}

export const SetUserName = (username)=>{
    localStorage.setItem(adminUser,username)
}

export const GetUserName = ()=>{
    return localStorage.getItem(adminUser)
}