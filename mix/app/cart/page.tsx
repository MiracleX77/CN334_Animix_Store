'use client'
import ProductComponent from "@/components/objects/product/Product"
import ProductFooter from "@/components/objects/product/ProductFooter"
import { useParams } from "next/navigation"
import CartTable from "@/components/objects/cart/cartTable"

export default function CartPage() {
    return (
        <>
            <div className="mx-auto w-[80%] shadow-neon mt-5">

                <div className="p-4  rounded-xl">
                    <CartTable />
                </div>

            </div>

        </>
    )
}