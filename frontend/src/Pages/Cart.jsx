import axios from "axios";
import { useEffect, useState } from "react";
import { API_URLS } from "../utils/utils";
import useAuth from "../hooks/useAuth";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useNavigate } from "react-router-dom";

export const Cart = () => {
    const [activeCarts, setActiveCarts] = useState([]);
    const { auth } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        async function getData() {
            try {
                const { data } = await axios.get(
                    API_URLS.BASE_URL + API_URLS.CARTS,
                    {
                        headers: {
                            Authorization: auth,
                            "Content-Type": "application/json",
                        },
                    }
                );
                console.log(data.cart);
                setActiveCarts(data.cart);
            } catch (error) {
                if (error.response && error.response.status === 401) {
                    navigate("/login");
                } else {
                    console.error("Error adding item to cart:", error);
                }
            }
        }
        getData();
    }, [auth, navigate]);

    async function createOrder() {
        try {
            const response = await axios.post(
                API_URLS.BASE_URL + API_URLS.CHECKOUT,
                {},
                {
                    headers: {
                        Authorization: auth,
                        "Content-Type": "application/json",
                    },
                }
            );

            if (response.status === 200) {
                setActiveCarts([]);
                toast.success("Order successful!", {
                    position: "top-right",
                    autoClose: 3000,
                });
            }
            console.log("Order created successfully:", response.data);
        } catch (error) {
            console.error("Error creating order:", error);
            if (error.response && error.response.status === 401) {
                toast.error("Session expired. Please log in again.", {
                    position: toast.POSITION.TOP_CENTER,
                    autoClose: 3000,
                });
                navigate("/login");
            } else {
                toast.error("Failed to create order. Please try again!", {
                    position: toast.POSITION.TOP_CENTER,
                    autoClose: 3000,
                });
            }
        }
    }

    return (
        <div className="min-h-screen bg-gray-100 p-4">
            <h1 className="text-3xl font-bold text-center text-gray-800 mb-6">
                Your Active Carts
            </h1>
            <div className="flex justify-center my-4">
                <button
                    className="bg-indigo-800 p-2 text-white rounded-md font-semibold hover:opacity-90"
                    onClick={() => createOrder()}
                >
                    Checkout
                </button>
            </div>
            <div className="flex flex-col justify-center my-8">
                {!activeCarts.length ? (
                    <p className="text-center text-gray-500">
                        You have no active carts.
                    </p>
                ) : (
                    activeCarts.map((cart) => (
                        <div key={cart.ID} className="mt-10">
                            <div
                                key={cart.Item.ID}
                                className="bg-white shadow-md rounded-lg p-6 flex items-center justify-evenly"
                            >
                                <h3 className="text-lg font-semibold text-gray-700">
                                    {cart.Item.Name}
                                </h3>
                                <p className="text-gray-400 text-xs">
                                    Added on:{" "}
                                    {new Date(
                                        cart.Item.CreatedAt
                                    ).toLocaleDateString()}
                                </p>
                                <p className="text-gray-500 text-sm">
                                    Price: ₹{cart.Item.Price}
                                </p>
                            </div>
                        </div>
                    ))
                )}
            </div>
            <ToastContainer />
        </div>
    );
};

export default Cart;
