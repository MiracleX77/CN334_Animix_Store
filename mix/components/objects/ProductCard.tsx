'use client'
import React, { useState } from "react"
import Image from "next/image"

import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import Middle from "../layouts/Middle"
import Span from "../layouts/Span"
import { useRouter } from "next/navigation"
import Cookies from "js-cookie"
import AlertDialog from "../interactive/layout/alertdialog"

export interface ProductCard {
    id: string
    img_url: string;
    name: string
    description?: string
    price: number
    new?: boolean
    hot?: boolean
    sale?: number
}

type props = {
    product: ProductCard
}

export function ProductCard({ product }: props) {
    const router = useRouter();
    
    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');
    
    const handleAddToCart = () => {
        console.log("Add to cart");
        let cart = Cookies.get("cart");
        let cartIds: string[] = cart ? JSON.parse(cart) : [];
        cartIds.push(product.id);
        Cookies.set("cart", JSON.stringify(cartIds), { expires: 7 });  
        console.log(`Product ID ${product.id} added to cart`);
        setAlertTitle("Add to Cart");
        setAlertContent("Product added to cart successfully");
        setAlertConfirmText("OK");
        setAlertStatus("success");
        setAlertOnConfirm(() => () => setOpenAlert(false));
        setOpenAlert(true);
    }
    const handleViewDetail = () => {
        console.log("View detail");
        router.push(`/product/${product.id}`);
    }

    return (
        <>
            <Card className="h-[350px] w-[100px] sm:w-[130px] md:w-[170px] 2xl:w-[210px] p-2 pb-0 cursor-pointer shadow-lg rounded-xl ring-4"
            >
                <div onClick={handleViewDetail} className="relative overflow-hidden rounded-xl" >
                    <img src={product.img_url} alt="Product" width={300} height={300} className="rounded-lg h-[200px] sm:h-[240px] md:h-[300px] 2xl:h-[250px] hover:scale-110 animate" />
                </div>
                <div className="flex justify-between mt-2 p-2 gap-4">
                    <div onClick={handleViewDetail} className="w-full">
                        <CardTitle className="text-[15px]">{product.name}</CardTitle>
       
                    </div>
                    <AlertDialog open={openAlert} setOpen={setOpenAlert} title={alertTitle} content={alertContent} status={alertStatus} onConfirm={alertOnConfirm} confirmText={alertConfirmText} cancelBottom={false}/>
                    <div>
                        <Span className="text-lg text-green-600">
                            {product.price} ฿
                        </Span>
                        <Button onClick={handleAddToCart} className="bg-secondary ring-2 px-4 rounded-xl text-white">
                            Add
                        </Button>
                    </div>
                </div>
            </Card>
        </>

    )
}
