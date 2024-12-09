import axios from "axios";
import { API_URLS } from "../utils/utils";
import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../context/AuthContext";
import { useNavigate } from "react-router-dom";

const Items = () => {
    const { auth } = useContext(AuthContext);
    const [items, setItems] = useState([]);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        async function getData() {
            try {
                const response = await axios.get(
                    API_URLS.BASE_URL + API_URLS.ITEMS,
                    {
                        headers: {
                            Authorization: `${auth}`,
                        },
                    }
                );
                setItems(response.data.items);
            } catch (error) {
                if (error.response && error.response.status === 401) {
                    navigate("/login");
                } else {
                    console.error("Error fetching items:", error);
                    setMessage("Failed to load items.");
                }
            }
        }
        getData();
    }, [auth, navigate]);

    const addToCart = async (itemID) => {
        const itemData = {
            item_id: itemID,
        };

        try {
            const response = await axios.post(
                API_URLS.BASE_URL + API_URLS.CARTS,
                itemData,
                {
                    headers: {
                        Authorization: `${auth}`,
                        "Content-Type": "application/json",
                    },
                }
            );
            if (response.ok) setMessage("Item added to cart successfully!");
        } catch (error) {
            if (error.response && error.response.status === 401) {
                navigate("/login");
            } else {
                console.error("Error adding item to cart:", error);
                setMessage(
                    error.response?.data?.error || "Failed to add item to cart."
                );
            }
        }
    };

    return (
        <div className="flex-1 bg-gray-100 p-4 pt-10">
            <h1 className="text-3xl font-bold text-center text-gray-800 mb-6">
                Available Items
            </h1>
            {message && (
                <p className="text-center text-green-600 mb-4">{message}</p>
            )}
            {items.length === 0 ? (
                <p className="text-center text-gray-500">No items available.</p>
            ) : (
                <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                    {items.map((item) => (
                        <div
                            key={item.ID}
                            className="bg-white shadow-md rounded-lg p-6 flex flex-col items-center justify-between"
                        >
                            <h2 className="text-lg font-semibold text-gray-700 mb-2">
                                {item.Name}
                            </h2>
                            <button
                                className="mt-4 bg-indigo-500 text-white px-4 py-2 rounded-md hover:bg-indigo-600"
                                onClick={() => addToCart(item.ID)}
                            >
                                Add to Cart
                            </button>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

export default Items;
