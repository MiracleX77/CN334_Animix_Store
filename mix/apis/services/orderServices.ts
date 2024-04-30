import {get,post,deleteRequest,put} from '@/apis/api'
import {parseJson} from '@/utils/parseJson'
import * as orderModel from '@/models/dto/order'

// post
export async function postOrder(params: orderModel.OrderCreate,token:string) {

    try {
        const response:any = await post(`order/`,'form','order', params,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

// post
export async function getOrderByStatus(status:string,token:string) {
    try {
        const response:any = await get(`order/status/${status}`,'json','order', null,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function getOrderById(id:string,token:string) {
    try {
        const response:any = await get(`order/${id}`,'json','order', null,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function putOrder(id :string,params: orderModel.OrderUpdate,token:string) {
    try {
        const response:any = await put(`order/${id}`,'json','order', params,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function putDelivery(id :string,params: orderModel.DeliveryUpdate,token:string) {
    try {
        const response:any = await put(`delivery/${id}`,'json','order', params,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}