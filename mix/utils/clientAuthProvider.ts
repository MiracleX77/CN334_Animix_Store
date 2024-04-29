import Cookies from 'js-cookie'

interface AuthProvider {
    login: (access_token : string,role:string) => void
    logout: () => void
    getAccessToken: () => string | undefined
    setAccessToken: (access_token: string) => void
}

export const AuthProvider : AuthProvider = {

    login: (access_token : string,role:string) => {
        Cookies.set('access_token', access_token, {expires: new Date(Date.now() + 60 * 60 * 1000), sameSite: 'lax', secure: true})
        Cookies.set('role', role, {secure: true})
    },
    logout: () => {
        Cookies.remove('access_token')
    },
    getAccessToken: () => {
        return Cookies.get('access_token')
    },
    setAccessToken: (access_token: string) => {
        Cookies.set('access_token', access_token, {expires: new Date(Date.now() + 15 * 60 * 1000), sameSite: 'lax', secure: true})
    },
}