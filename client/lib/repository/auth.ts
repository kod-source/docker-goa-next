import axios from "axios";

import { Auth, GoogleRedirect } from "../model/auth";
import { User } from "../model/user";
import { getEndPoint } from "../token";

export const AuthRepository = {
    login: async (email: string, password: string): Promise<Auth> => {
        const res = await axios.post(`${getEndPoint()}/api/v1/login`, {
            email: email,
            password: password,
        });
        return {
            token: res.data.token,
            user: new User(
                res.data.user.id,
                res.data.user.name,
                res.data.user.email,
                new Date(res.data.user.created_at),
                res.data.user.avatar,
            ),
        };
    },

    signUp: async (
        name: string,
        email: string,
        password: string,
        avatarPath?: string,
    ): Promise<Auth> => {
        const res = await axios.post(`${getEndPoint()}/api/v1/sign_up`, {
            name: name,
            email: email,
            password: password,
            avatar: avatarPath,
        });
        return {
            token: res.data.token,
            user: new User(
                res.data.user.id,
                res.data.user.name,
                res.data.user.email,
                new Date(res.data.user.created_at),
                res.data.user.avatar,
            ),
        };
    },

    googleLogin: async (): Promise<GoogleRedirect> => {
        const res = await axios.get(`${getEndPoint()}/api/v1/google/login`);
        return {
            url: res.data.url,
        };
    },

    googleCallBack: async (code: string, state: string): Promise<Auth> => {
        const res = await axios.post(`${getEndPoint()}/api/v1/google/callback`, {
            code: code,
            state: state,
        });
        return {
            token: res.data.token,
            user: new User(
                res.data.user.id,
                res.data.user.name,
                res.data.user.email,
                new Date(res.data.user.created_at),
                res.data.user.avatar,
            ),
        };
    },
};
