import { Loading } from "lib/components/loading";
import { AuthRepository } from "lib/repository/auth";
import { GetServerSideProps, NextPage } from "next";
import { useRouter } from "next/router";
import { AppContext } from "pages/_app";
import { useContext, useEffect } from "react";

interface Props {
    state: string;
    code: string;
    scope: string;
    authuser: string;
    prompt: string;
}

const GoogleCallback: NextPage<Props> = (props) => {
    const router = useRouter();
    const { setUser } = useContext(AppContext);

    const googleCallBackLogin = async () => {
        try {
            const auth = await AuthRepository.googleCallBack(props.code, props.state);
            localStorage.setItem("token", auth.token);
            setUser(auth.user);
            router.push("/");
        } catch (e) {
            if (e instanceof Error) {
                alert(e.message);
            }
            router.push("/login");
        }
    };

    useEffect(() => {
        googleCallBackLogin();
    }, []);

    return <Loading />;
};

export const getServerSideProps: GetServerSideProps<Props> = async ({ query }) => {
    const { state, code, scope, authuser, prompt } = query;

    return {
        props: {
            state: String(state),
            code: String(code),
            scope: String(scope),
            authuser: String(authuser),
            prompt: String(prompt),
        },
    };
};

export default GoogleCallback;
