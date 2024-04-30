import axios, { AxiosRequestConfig } from 'axios';
import { AuthProvider } from '@/utils/clientAuthProvider';
import { cookies } from 'next/headers';
import Cookies from 'js-cookie';

function getRequestConfig(method: string, url: string, type: string,service: string,data?: any,token?: string): AxiosRequestConfig {
    const access_token = token

    if (service == 'user'){
        url = "http://localhost:5003/v1/" + url;
    }
    if (service == 'product'){
        url = "http://localhost:5002/v1/" + url;
    }
    if (service == 'order'){
        url = "http://localhost:5001/v1/" + url;
    }
    if (service == 'review'){
        url = "http://localhost:5004/v1/" + url;
    }


    if (type =='json'){
        const headers: { [key: string]: string } = {
            'Content-Type': 'application/json',
            Accept: 'application/json',
            'X-Requested-With': 'XMLHttpRequest',
        };
        if (token) {
            headers.Authorization = `Bearer ${access_token}`;
        }
        console.log(headers)
        return {
            method,
            url: url,
            headers,
            data,
        };
        
    } else {
        const headers: { [key: string]: string } = {
            'Content-Type': 'multipart/form-data',
            Accept: 'application/json',
            'X-Requested-With': 'XMLHttpRequest',
        };
        if (token) {
            headers.Authorization = `Bearer ${access_token}`;
        }
        return {
            method,
            url: url,
            headers,
            data,
        };
    }
}

async function sendRequest(config: AxiosRequestConfig) {
    try {
        const response = await axios(config);
        return JSON.stringify({
            status: response.status,
            data: response.data,
            message: response.data.message,
        });
    } catch (error: any) {
        throw {
            status: error.response?.status,
            message: error.response?.data?.message,
        };
    }
}

export async function get(url: string,type: string,service:string, params?: any, token?: string) {
    let queryParams = '';
    if (params) {
        const nonEmptyParams = Object.entries(params)
            .filter(([_, value]) => value !== null && value !== undefined && value !== '')
            .map(([key, value]) => {
                if (Array.isArray(value)) {
                    return value.map((item) => `${key}=${item}`).join('&');
                } else {
                    return `${key}=${value}`;
                }
            });
        queryParams = nonEmptyParams.join('&');
    }
    const fullUrl = queryParams ? `${url}?${queryParams}` : url;
    const config = getRequestConfig('get', fullUrl, type,service,undefined, token);
    return sendRequest(config);
}

export async function post(url: string,type: string,service:string, data?: any, token?: string) {
    if (type == 'json'){
        const config = getRequestConfig('post', url,type,service, data, token);
        return sendRequest(config);
    } else {
        const formData = new FormData();
        for (const key in data) {
            if (data.hasOwnProperty(key)) {
                formData.append(key, data[key]);
            }
        }
        const config = getRequestConfig('post', url,type,service, formData, token);
        return sendRequest(config);
    }
}

// export async function postForm(url: string, data?: any, token: string | null = '') {
//     const formData = new FormData();
//     for (const key in data) {
//         if (data.hasOwnProperty(key)) {
//             formData.append(key, data[key]);
//         }
//     }
//     const config = getRequestConfig('post', url, formData, token);
//     if (!config.headers) {
//         config.headers = {}; // Initialize headers as an empty object if it's undefined
//     }
//     config.headers['Content-Type'] = 'multipart/form-data';
//     return sendRequest(config);
// }

export async function deleteRequest(url: string, type: string,service:string,data?: any, token?: string) {
    const config = getRequestConfig('delete', url,type,service, data, token);
    return sendRequest(config);
}

export async function put(url: string,type: string,service:string, data?: any, token?: string) {
    const config = getRequestConfig('put', url,type,service, data, token);
    return sendRequest(config);
}
