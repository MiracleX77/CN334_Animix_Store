export interface ProductModel {
    id: number;
    author: any;
    category: any;
    publisher: any;
    name: string;
    description?: string;
    price: number;
    stock: number;
    img_url: string;
    status: string;
    created_at: string;
    updated_at: string;
}

export interface ProductForCard {
    id: number;
    img_url: string;
    name: string;
    price: number;

}

export interface ProductCreate {
    author_id : number;
    category_id : number;
    publisher_id : number;
    name: string;
    description?: string;
    price: number;
    stock: number;
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
