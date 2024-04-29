'use client'
import { ProductModel,ProductForCard } from "@/models/dto/product";
import { ProductCard } from "./ProductCard";
import { getProducts,getProductsByCategory } from "@/apis/services/productService";
import { use, useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { AuthProvider } from "@/utils/clientAuthProvider";



type props = {
    type: string;
}

export default function ListProductCard({ type }: props) {
    const router = useRouter();
    const token = AuthProvider.getAccessToken()

    const [products, setProducts] = useState([]);

    useEffect(() => {
        loadProducts();
    }, []);

    const loadProducts = async () => {
        if (type === 'all') {
            const fetchedProducts = await getProducts();
            const product = fetchedProducts.data.data
            setProducts(product);
        } else {
            if (type ==='Manga')
            {
                const fetchedProducts = await getProductsByCategory('3');
                const product = fetchedProducts.data.data
                setProducts(product);
            } else if (type === 'Light Novel')
            {
                const fetchedProducts = await getProductsByCategory('2');
                const product = fetchedProducts.data.data
                setProducts(product);
            } else{
                const fetchedProducts = await getProductsByCategory('1');
                const product = fetchedProducts.data.data
                setProducts(product);
            }
        }
    };

    return (
        <>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-10">
                {products.map((product, index) => (
                    <ProductCard key={index} product={product} />
                ))}
            </div>
        </>
        
    )
}