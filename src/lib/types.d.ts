export type Password = {
  id: string, 
  name: string, 
  password: string,
  user_id: string, 
  username: string, 
  website: string
}

export type User = {
  email: string, 
  id: string, 
  name: string, 
  password: string, 
  salt: string
}
