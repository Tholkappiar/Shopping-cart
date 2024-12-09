import { createContext, useEffect, useState } from "react";

export const AuthContext = createContext();

// eslint-disable-next-line react/prop-types
const AuthProvider = ({ children }) => {
    const [auth, setAuth] = useState(localStorage.getItem("auth") || "");

    useEffect(() => {
        localStorage.removeItem("auth");
        localStorage.setItem("auth", auth);
    }, [auth]);

    console.log("auth :" + auth);

    return (
        <AuthContext.Provider value={{ auth, setAuth }}>
            {children}
        </AuthContext.Provider>
    );
};

export default AuthProvider;
