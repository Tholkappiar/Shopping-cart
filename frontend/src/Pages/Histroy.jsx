import axios from "axios";
import { useEffect, useState } from "react";
import { API_URLS } from "../utils/utils";
import useAuth from "../hooks/useAuth";
import { useNavigate } from "react-router-dom";

export const History = () => {
    const [orders, setOrders] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const { auth } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        async function getData() {
            try {
                const { data } = await axios.get(
                    API_URLS.BASE_URL + API_URLS.ORDERS,
                    {
                        headers: {
                            Authorization: auth,
                        },
                    }
                );
                setOrders(data.orders);
            } catch (error) {
                if (error.response && error.response.status === 401) {
                    navigate("/login");
                } else {
                    console.error("Error fetching orders:", error);
                }
            } finally {
                setIsLoading(false);
            }
        }

        getData();
    }, [auth, navigate]);

    return (
        <div className="min-h-screen bg-gray-100 p-4">
            <h1 className="text-3xl font-bold text-center text-gray-800 mb-6">
                Your Order History
            </h1>
            {isLoading ? (
                <div className="space-y-4">
                    {[...Array(3)].map((_, index) => (
                        <div
                            key={index}
                            className="w-full h-24 bg-gray-300 rounded-lg animate-pulse"
                        ></div>
                    ))}
                </div>
            ) : orders && orders.length === 0 ? (
                <p className="text-center text-gray-500">You have no orders.</p>
            ) : (
                <div className="space-y-8">
                    {orders.map((order, index) => (
                        <div
                            key={order.ID + index}
                            className="bg-white shadow-md rounded-lg p-4"
                        >
                            <h2 className="text-xl font-semibold text-gray-700">
                                Order #{order.ID} - {order.Status}
                            </h2>
                            <p className="text-gray-600">
                                Total: ₹{order.Total}
                            </p>
                            <p className="text-gray-500">
                                Created At:{" "}
                                {new Date(order.CreatedAt).toLocaleString()}
                            </p>
                            <div className="space-y-2 mt-4">
                                {order.OrderItems.map((item) => (
                                    <div
                                        key={item.ID}
                                        className="flex justify-between text-gray-600"
                                    >
                                        <span>
                                            {item.Item.Name} x {item.Quantity}
                                        </span>
                                        <span>₹{item.Price}</span>
                                    </div>
                                ))}
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};
