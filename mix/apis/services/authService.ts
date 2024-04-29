import {get,post,deleteRequest,put} from '@/apis/api'
import {parseJson} from '@/utils/parseJson'
import * as authModel from '@/models/dto/auth'

// post
export async function postLogin(params: authModel.Login) {

    try {
        const response:any = await post(`auth/login`,'json','user', params)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

// post
export async function postRegister(params: authModel.Register) {

    try {
        const response:any = await post(`auth/register`,'json','user', params)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}