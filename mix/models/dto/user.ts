export interface UserModel {
    user_id: number;
    first_name: string;
    last_name: string;
    username: string;
    email: string;
    created_on: string;
    status: string;
}

export interface UserUpdate {
    first_name: string;
    last_name: string;
    email: string;
}