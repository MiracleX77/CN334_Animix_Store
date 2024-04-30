import {get,post,deleteRequest,put} from '@/apis/api'
import {parseJson} from '@/utils/parseJson'
import * as userModel from '@/models/dto/user'

// post
export async function getUser(token: string) {

    try {
        const response:any = await get(`user/`,'json','user', null, token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

// post
export async function putUser(params: userModel.UserUpdate,token:string) {

    try {
        const response:any = await put(`user/`,'json','user', params,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}