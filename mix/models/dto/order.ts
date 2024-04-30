
export interface OrderCreate {
    address_id : number;
    type : string;
    total : number;
    list_product_id: number[];
    img: File | null;
}

export interface ProductUpdate {
    author_id : number;
    category_id : number;
    publisher_id : number;
    name: string;
    description?: string;
    price: number;
    stock: number;
    img: string;
}
export interface OrderModel {
    id: number;
    user_id: number;
    delivery: DeliveryModel;
    payment: PaymentModel;
    total_price: number;
    status: string;
    created_at: string;
    updated_at: string;
}

export interface DeliveryModel {
    id: number;
    address: AddressModel;
    cost: number | null;
    type: string | null;
    tracking_number: string | null;
    status: string;
    created_at: string;
    updated_at: string;
  }
  
  export interface AddressModel {
    id: number;
    user_id: number;
    address_line: string;
    phone: string;
    name: string;
    sub_district: SubDistrictModel;
    district: DistrictModel;
    province: ProvinceModel;
    default: string;
    status: string;
    created_at: string;
    updated_at: string;
  }
  
  export interface SubDistrictModel {
    id: number;
    name_th: string;
    name_en: string;
    post_code: string;
  }
  
  export interface DistrictModel {
    id: number;
    name_th: string;
    name_en: string;
  }
  
  export interface ProvinceModel {
    id: number;
    name_th: string;
    name_en: string;
  }
  
  export interface PaymentModel {
    id: number;
    type: string;
    total: number;
    proof_payment: string;
    status: string;
    created_at: string;
    updated_at: string;
  }


export interface OrderUpdate {
    status: string;
}
export interface DeliveryUpdate {
    address_id : number;
    type : string;
    tracking_number : string;
}