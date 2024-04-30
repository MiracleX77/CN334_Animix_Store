export interface ReviewModel {
    id: number;
    user_id: number;
    product_id: number;
    title: string;
    content: string;
    rating: number;
    polarity: string;
    status: string;
    created_at: string;
    updated_at: string;
}

export interface ReviewCreate {
    product_id: number;
    title: string;
    content: string;
    rating: number;
}

