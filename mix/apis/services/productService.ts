import {get,post,deleteRequest,put} from '@/apis/api'
import {parseJson} from '@/utils/parseJson'
import * as productModel from '@/models/dto/product'

// get
export async function getProducts() {

    try {
        const response:any = await get(`product/`,'json','product', null, "")
        
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}
export async function getProductById(id:string,token:string) {

    try {
        
        const response:any = await get(`product/${id}`,'json','product', null,"")
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function getProductsByCategory(category:string) {
    
        try {
            const response:any = await get(`product/category/${category}`,'json','product', null,"")
            const json:any = parseJson(response)
            return json
        } catch (ex) {
            return ex
        }
    }

// post
export async function postProduct(params: productModel.ProductCreate,token:string) {

    try {
        const response:any = await post(`product/`,'formdata','product', params,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

// put
export async function putProductById(id: string, params: productModel.ProductUpdate,token:string) {

    try {
        const response:any = await put(`product/${id}`,'json','product', params,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}
// delete
export async function deleteProductById(id: string,token:string) {

    try {
        const response:any = await deleteRequest(`product/${id}`,'json','product', null,token)
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

export async function getAuthor(token:string) {

    try {
        const response:any = await get(`author/`,'json','product', null,token
        )
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}
export async function getPublisher(token:string) {

    try {
        const response:any = await get(`publisher/`,'json','product', null,token
        )
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}
export async function getCategory(token:string) {

    try {
        const response:any = await get(`category/`,'json','product', null,token
        )
        const json:any = parseJson(response)
        return json
    } catch (ex) {
        return ex
    }
}

