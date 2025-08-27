export interface IUser {
	name: string
}

export interface IAuthResponse {
	token: string
}

export interface DecodedToken {
	exp: number
	nbf: number
	iat: number
	iss: string
	sub: string
}
