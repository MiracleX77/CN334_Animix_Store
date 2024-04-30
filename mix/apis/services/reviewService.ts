import {get,post,deleteRequest,put} from '@/apis/api'
import {parseJson} from '@/utils/parseJson'
import * as reviewModel from '@/models/dto/review'

// get
export async function getReviews() {

    try {
        const response:any = await get(`review/`,'json','review', null, "")
        
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}
export async function getReviewByUserId(id:string,token:string) {

    try {
        
        const response:any = await get(`review/user/${id}`,'json','review', null,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function getReviewsByProductId(id:string,token:string) {
    
        try {
            const response:any = await get(`review/product/${id}`,'json','review', null,"")
            const json:any = parseJson(response)
            return json
        } catch (ex) {
            return ex
        }
    }

// post
export async function postReview(params: reviewModel.ReviewCreate,token:string) {

    try {
        const response:any = await post(`review/`,'json','review', params,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

// delete
export async function deleteReviewById(id: string,token:string) {

    try {
        const response:any = await deleteRequest(`review/${id}`,'json','review', null,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}


