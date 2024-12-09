import { Link } from "react-router-dom";
import useAuth from "../hooks/useAuth";

const Header = () => {
    const { auth } = useAuth();

    return (
        <header className="bg-white dark:bg-indigo-800 border-b border-gray-200 dark:border-gray-700">
            <div className="container mx-auto px-4 py-4">
                <div className="flex flex-col sm:flex-row items-center justify-between space-y-3 sm:space-y-0">
                    <Link
                        to="/"
                        className="text-xl font-semibold text-gray-800 dark:text-white"
                    >
                        Shopping Cart
                    </Link>
                    <nav className="flex flex-wrap justify-center space-x-4">
                        {auth ? (
                            <>
                                <Link
                                    to="/cart"
                                    className="text-sm font-medium text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white"
                                >
                                    Checkout
                                </Link>
                                <Link
                                    to="/history"
                                    className="text-sm font-medium text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white"
                                >
                                    Order History
                                </Link>
                            </>
                        ) : (
                            <>
                                <Link
                                    to="/login"
                                    className="text-sm font-medium text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white"
                                >
                                    Log in
                                </Link>
                                <Link
                                    to="/signup"
                                    className="text-sm font-medium text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white"
                                >
                                    Sign up
                                </Link>
                            </>
                        )}
                    </nav>
                </div>
            </div>
        </header>
    );
};

export default Header;
