export interface Login {
    username: string;
    password: string;
}

export interface Register {
    first_name: string;
    last_name: string;
    email: string;
    username: string;
    password: string;
}

export interface LoginResponse {
    message: string;
}