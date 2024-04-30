export interface Address {
    id: number;
    user_id: number;
    address_line: string;
    phone: string;
    name: string;
    sub_district: any;
    district: any;
    province: any;
    default: string;
    status: string;
    created_at: string;
    updated_at: string;
}

export interface AddressCreate {
    address_line: string;
    phone: string;
    name: string;
    sub_district_id: any;
    district_id: any;
    province_id: any;
    default: string;
}

export interface Province {
    id: number;
    name_th: string;
    name_en: string;
}

export interface District {
    id: number;
    name_th: string;
    name_en: string;
    province_id: number;
}

export interface SubDistrict {
    id: number;
    name_th: string;
    name_en: string;
    post_code: string;
    district_id: number;
}