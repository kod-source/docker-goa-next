import "../styles/globals.css";
import type { AppProps } from "next/app";

import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";

import { User } from "../lib/model/user";
import "tailwindcss/tailwind.css";
import { UserRepostiory } from "../lib/repository/user";

export const AppContext = React.createContext(
    {} as {
        user: User | null;
        setUser: React.Dispatch<React.SetStateAction<User | null>>;
    },
);

function MyApp({ Component, pageProps }: AppProps) {
    const router = useRouter();
    const [user, setUser] = useState<User | null>(null);
    const path = router.pathname;
    let isFirst = true;

    const fetchData = async () => {
        if (path === "/login" || path === "/sign_up" || path === "/auth/callback/google") {
            return;
        }
        const token = localStorage.getItem("token");
        if (!token) {
            return router.push("/login");
        }
        try {
            const user = await UserRepostiory.currentUser();
            setUser(user);
        } catch {
            if (isFirst) {
                alert("tokenの認証が切れました。再度ログインしてください。");
                localStorage.removeItem("token");
                router.push("/login");
            }
        }
        isFirst = false;
    };

    useEffect(() => {
        fetchData();
    }, []);
    return (
        <AppContext.Provider value={{ user, setUser }}>
            <Component {...pageProps} />
        </AppContext.Provider>
    );
}

export default MyApp;
