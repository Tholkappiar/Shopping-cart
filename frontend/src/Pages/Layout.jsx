import { Outlet } from "react-router-dom";
import Header from "../components/Header";

const Layout = () => {
    return (
        <div className="min-h-dvh flex flex-col">
            <Header />
            <section className="flex flex-col flex-1 container mx-auto">
                <Outlet />
            </section>
        </div>
    );
};

export default Layout;
