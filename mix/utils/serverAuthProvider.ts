import { cookies } from 'next/headers';

interface AuthProvider {
    login: (access_token : string, refresh_token : string) => void
    logout: () => void
    getAccessToken: () => string | undefined
    getRefreshToken: () => string | undefined
    setAccessToken: (access_token: string) => void
}

export const AuthProvider : AuthProvider = {

    login: (access_token : string, refresh_token : string) => {
        //15 mins
        cookies().set('access_token', access_token, {expires: new Date(Date.now() + 15 * 60 * 1000), sameSite: 'lax', secure: true})
        //1 day
        cookies().set('refresh_token', refresh_token, {expires: new Date(Date.now() + 24 * 60 * 60 * 1000), sameSite: 'lax', secure: true})
    },
    logout: () => {
        cookies().delete('access_token')
        cookies().delete('refresh_token')
    },
    getAccessToken: () => {
        return cookies().get('access_token')?.value
    },
    setAccessToken: (access_token: string) => {
        cookies().set('access_token', access_token, {expires: new Date(Date.now() + 15 * 60 * 1000), sameSite: 'lax', secure: true})
    },
    getRefreshToken: () => {
        return cookies().get('refresh_token')?.value
    }
}