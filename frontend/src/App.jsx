import { BrowserRouter, Route, Routes } from "react-router-dom";
import Layout from "./Pages/Layout";
import Items from "./Pages/Items";
import AuthProvider from "./context/AuthContext";
import RequireAuth from "./components/RequireAuth";
import { Cart } from "./Pages/Cart";
import { History } from "./Pages/Histroy";
import Login from "./Pages/Login";
import Signup from "./Pages/Signup";

function App() {
    return (
        <AuthProvider>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Layout />}>
                        <Route path="/login" element={<Login />} />
                        <Route path="/signup" element={<Signup />} />
                        <Route element={<RequireAuth />}>
                            <Route path="/" element={<Items />}></Route>
                            <Route path="/cart" element={<Cart />}></Route>
                            <Route
                                path="/history"
                                element={<History />}
                            ></Route>
                        </Route>
                    </Route>
                </Routes>
            </BrowserRouter>
        </AuthProvider>
    );
}

export default App;
