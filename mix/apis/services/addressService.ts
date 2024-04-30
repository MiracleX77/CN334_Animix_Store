import {get,post,deleteRequest,put} from '@/apis/api'
import {parseJson} from '@/utils/parseJson'
import * as addressModel from '@/models/dto/address'

// post
export async function getAddressAll(token: string) {

    try {
        const response:any = await get(`address/`,'json','user', null, token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

// post
export async function postAddress(params: addressModel.AddressCreate,token:string) {

    try {
        const response:any = await post(`address/`,'json','user', params,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function getProvince(token: string) {

    try {
        const response:any = await get(`address/province`,'json','user', null, token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function getDistrict(token: string,province_id:number) {

    try {
        const response:any = await get(`address/district/${province_id}`,'json','user', null, token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function getSubDistrict(token: string,district_id:number) {

    try {
        const response:any = await get(`address/subdistrict/${district_id}`,'json','user', null, token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}
